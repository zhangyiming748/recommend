package recommend

import (
	"log"
	"recommend/module"
	"recommend/module/rank"
	"recommend/module/template"
	"recommend/util"

	"fmt"
	"math/rand"
	"recommend/module/usercf"
	"strings"
	"sync"
	"time"
)

func chooseRecStgy(param *module.Param) (stgy module.Stgy) {
	//0. 获取频道和AB策略
	stgy.SetChannel(param.GetChannel())
	stgy.SetAbtestStgy(getABtest(param.GetUuid()))
	stgy.SetUniqueid(fmt.Sprintf("%x", rand.Int63()))
	//1. 获取action
	switch strings.ToLower(param.GetRecAction()) {
	case module.ACTION_UP:
		stgy.SetAction(module.ACTION_UP)
		stgy.SetSize(module.RECOMMEND_SIZE_UP)
	case module.ACTION_DOWN:
		stgy.SetAction(module.ACTION_DOWN)
		stgy.SetSize(module.RECOMMEND_SIZE_DOWN)
	default:
		stgy.SetAction(module.ACTION_HOME)
		stgy.SetSize(module.RECOMMEND_SIZE_HOME)
	}

	//2. 设置固定位
	stgy.FixedList = template.TemplateFixedItems(stgy.GetChannel())
	fixpos := make(map[int][]string)
	for article, pos := range stgy.FixedList {
		if _, ok := fixpos[pos]; ok {
			fixpos[pos] = append(fixpos[pos], article)
		} else {
			fixpos[pos] = make([]string, 0)
		}
	}

	stgy.FixedPos = fixpos
	filtermap := make(map[string]bool)
	for article_id, _ := range stgy.FixedList {
		filtermap[article_id] = true
	}
	stgy.Filtermap = filtermap

	stgy.SetDeviceId(param.GetUuid())
	stgy.SetAlgoStgy(module.PERSONALALGO)

	//3. 流量控制
	if isTrafficJam() {
		//若当前并发请求过多, 则降级为随机策略
		log.Println("to many request! more than ", MaxTraffic)
		stgy.SetAlgoStgy(module.RANDOMALGO)
		return stgy
	}
	//4. 检查uid是否有效, 若无效yuid, 则通过uuid反查uid
	uid := param.GetUserId()
	if uid == "" || uid == "NULL" || uid == "null" {
		uid = getUserIdByUuid(param.GetUuid())
		if uid == "" {
			// 若为匿名用户, 则降级为随机策略
			stgy.SetAlgoStgy(module.RANDOMALGO)
			return stgy
		}
	}
	stgy.SetUserId(uid)

	//5. 获取个性化信息
	pinfo := getUserPersonalInfo(uid)

	if len(pinfo.Showlist)+len(pinfo.Clicklist)+len(pinfo.History_usertag) == 0 {
		// 若个性化信息为空, 则降级为随机策略
		stgy.SetAlgoStgy(module.RANDOMALGO)
		return stgy
	}

	//6. 获取过滤列表
	var lli = [][]string{pinfo.Showlist, pinfo.Clicklist, pinfo.Quick_showlist, pinfo.Nfblist}
	for _, li := range lli {
		for _, article_id := range li {
			filtermap[article_id] = true
		}
	}
	stgy.Filtermap = filtermap

	//7. 获取实时用户画像, 以填充stgy
	pinfo.Realtime_usertag = getUserRealtimeTags(pinfo.Clicklist)
	stgy.PersonalInfo = pinfo

	return stgy
}

func getPersonalmRecommend(st module.Stgy) []module.ArticleList {
	var wg sync.WaitGroup
	reclist := make([]module.ArticleList, 3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer util.PanicRecover(util.PanicPosition())

		c := make(chan module.ArticleList)
		go func() {
			defer util.PanicRecover(util.PanicPosition())
			//传入uid, algostgy, abteststgy, filtermap, 拿到按照时间倒序排列的列表, 长度为50
			c <- rank.Rank_fornew(st, 100)
		}()
		select {
		case reclist[0] = <-c:
		case <-time.After(300 * time.Millisecond):
			log.Println("wait rank_server_fornew 300ms timeout! ", st.GetUniqueid())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer util.PanicRecover(util.PanicPosition())

		c := make(chan module.ArticleList)
		go func() {
			defer util.PanicRecover(util.PanicPosition())
			//传入uid, algostgy, abteststgy, filtermap, 拿到个性化推荐内容的列表, 长度为50
			c <- rank.Rank_forrec(st, 50)
		}()
		select {
		case reclist[1] = <-c:
		case <-time.After(300 * time.Millisecond):
			util.Errorln("wait rank_server_forrec 300ms timeout!! ", st.GetUniqueid())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer util.PanicRecover(util.PanicPosition())
		c := make(chan module.ArticleList)
		go func() {
			defer util.PanicRecover(util.PanicPosition())
			//传入uid, algostgy, abteststgy, filtermap, 拿到usercf的列表, 长度为50
			c <- usercf.GetUserCfItems(st, 50, 50)

		}()
		select {
		case reclist[2] = <-c:
		case <-time.After(300 * time.Millisecond):
			util.Errorln("wait usercf 300ms timeout!! ", st.GetUniqueid())
		}
	}()
	wg.Wait()
	util.Debugln(st)
	return reclist

}

const (
	NEWLYLISTKEY = "newlylist"
)

func getRandomRecommend(st module.Stgy) []module.ArticleList {
	reclist := make([]module.ArticleList, 0)
	//获得更大的长度, 新内容, 新内容, 假设长度为200

	val := rank.Rank_fornew(st, 200)
	if len(val) > 50 {
		var rli module.ArticleList
		// 从200个新内容中随机选取50个内容, 越新的内容, 被选择的概率略大
		for {
			idxmap := make(map[int]bool)
			for i, _ := range val {
				randf := rand.Float64()
				//概率从0.73到0.07
				if randf < float64(float64((20+len(val)-i))/(1.5*float64(len(val)))) {
					idxmap[i] = true
				}
			}
			if len(idxmap) >= 50 {
				util.Debugln(idxmap)
				for i, _ := range idxmap {
					rli = append(rli, val[i])
				}
				break
			}
		}
		reclist = append(reclist, rli)
	} else {
		//若长度不足, 则直接返回
		reclist = append(reclist, val)
	}
	return reclist
}
