package main

import (
	"../rpc"
	"../storage"
	. "../util"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

func chooseRecStgy(param *Param) (stgy Stgy) {
	//0. 获取AB策略
	stgy.SetAbtestStgy(getABtest(param.uuid))
	stgy.uniqueid = fmt.Sprintf("%x", rand.Int63())

	//1. 获取action
	switch strings.ToLower(param.recAction) {
	case ACTION_UP:
		stgy.SetAction(ACTION_UP)
		stgy.SetSize(RECOMMEND_SIZE_UP)
	case ACTION_DOWN:
		stgy.SetAction(ACTION_DOWN)
		stgy.SetSize(RECOMMEND_SIZE_DOWN)
	default:
		stgy.SetAction(ACTION_HOME)
		stgy.SetSize(RECOMMEND_SIZE_HOME)
	}
	//2. 设置固定位
	stgy.FixedList = templateFixedItems()
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

	stgy.SetDeviceId(param.uuid)
	stgy.SetAlgoStgy(PERSONALALGO)

	//3. 流量控制
	if isTrafficJam() {
		//若当前并发请求过多, 则降级为随机策略
		Warningln("to many request! more than ", MaxTraffic)
		stgy.SetAlgoStgy(RANDOMALGO)
		return stgy
	}
	//4. 检查uid是否有效, 若无效yuid, 则通过uuid反查uid
	uid := param.userId
	if uid == "" || uid == "NULL" || uid == "null" {
		uid = getUserIdByUuid(param.uuid)
		if uid == "" {
			// 若为匿名用户, 则降级为随机策略
			stgy.SetAlgoStgy(RANDOMALGO)
			return stgy
		}
	}
	stgy.SetUserId(uid)

	//5. 获取个性化信息
	pinfo := getUserPersonalInfo(uid)

	if len(pinfo.Showlist)+len(pinfo.Clicklist)+len(pinfo.History_usertag) == 0 {
		// 若个性化信息为空, 则降级为随机策略
		stgy.SetAlgoStgy(RANDOMALGO)
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

func getPersonalmRecommend(st Stgy) []ArticleList {
	var wg sync.WaitGroup
	reclist := make([]ArticleList, 3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c := make(chan ArticleList)
		go func() {
			//传入uid, algostgy, abteststgy, filtermap, 拿到按照时间倒序排列的列表, 长度为50
			c <- rank_server(rpc.RANK_FORNEW, st, 100)
		}()
		select {
		case reclist[0] = <-c:
		case <-time.After(300 * time.Millisecond):
			Errorln("wait rank_server_fornew 300ms timeout! ", st.uniqueid)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		c := make(chan ArticleList)
		go func() {
			//传入uid, algostgy, abteststgy, filtermap, 拿到个性化推荐内容的列表, 长度为50
			c <- rank_server(rpc.RANK_FORREC, st, 50)
		}()
		select {
		case reclist[1] = <-c:
		case <-time.After(300 * time.Millisecond):
			Errorln("wait rank_server_forrec 300ms timeout!! ", st.uniqueid)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		c := make(chan ArticleList)
		go func() {
			//传入uid, algostgy, abteststgy, filtermap, 拿到usercf的列表, 长度为50
			c <- usercf_server(rpc.USERCF_RECOMMEND, st, 50)

		}()
		select {
		case reclist[2] = <-c:
		case <-time.After(300 * time.Millisecond):
			Errorln("wait usercf 300ms timeout!! ", st.uniqueid)
		}
	}()
	wg.Wait()
	Debugln(st)
	return reclist

}

const (
	NEWLYLISTKEY = "newlylist"
)

func getRandomRecommend(st Stgy) []ArticleList {
	reclist := make([]ArticleList, 0)
	val_, ok := storage.FetchCacheContent(NEWLYLISTKEY)
	if !ok {
		//获得更大的长度, 新内容, 新内容, 假设长度为200
		val_ = rank_server(rpc.RANK_FORNEW, st, 200)
		if len(val_.(ArticleList)) > 0 {
			storage.SetCacheContent(NEWLYLISTKEY, val_, storage.DEFAULTEXPIRATION)
		}
	}
	val := val_.(ArticleList)
	if len(val) > 50 {
		var rli ArticleList
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
				Debugln(idxmap)
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

func Test_rank_server() {
	var st Stgy
	st.SetUserId("5268cf835abc4b4686c90ff62cab3a14")
	maps := make(map[string]bool)
	//maps["1794717"] = true
	//maps["1794718"] = false
	st.Filtermap = maps
	rank_server(rpc.RANK_FORNEW, st, 50)
}

func rank_server(method string, st Stgy, size int) ArticleList {
	var artlist ArticleList
	param := make(map[string]string, 0)
	// 公共参数传入uid, algostgy, abteststgy, filtermap
	param["uid"] = st.userId
	param["algostgy"] = st.algoStgy
	param["abteststgy"] = st.abtestStgy
	param["uniqueid"] = st.uniqueid
	param["size"] = strconv.Itoa(size)
	filterids := make([]string, 0)
	for id, _ := range st.Filtermap {
		filterids = append(filterids, id)
	}
	param["filterids"] = strings.Join(filterids, ",")

	history_usertag_str, _ := json.Marshal(st.History_usertag)
	param["history_usertag"] = string(history_usertag_str)
	realtime_usertag_str, _ := json.Marshal(st.Realtime_usertag)
	param["realtime_usertag"] = string(realtime_usertag_str)

	if method == rpc.RANK_FORNEW {

	} else if method == rpc.RANK_FORREC {

	}
	ret, err := rpc.RankClient(method, param)
	if err != nil {
		Errorln("rpc err", err, "uniqueid:", st.uniqueid)
	}

	for _, v := range ret {
		Debugln(v)
		art := Article{*v}
		artlist = append(artlist, art)
	}

	return artlist
}

func Test_usercf_server() {
	var st Stgy
	st.SetUserId("5268cf835abc4b4686c90ff62cab3a14")
	maps := make(map[string]bool)
	maps["1794717"] = true
	maps["1794718"] = false
	st.Filtermap = maps
	usercf_server(rpc.USERCF_RECOMMEND, st, 50)
}

func usercf_server(method string, st Stgy, size int) ArticleList {
	var artlist ArticleList

	param := make(map[string]string, 0)
	// 公共参数传入uid, algostgy, abteststgy, filtermap
	param["uid"] = st.userId
	param["algostgy"] = st.algoStgy
	param["abteststgy"] = st.abtestStgy
	param["filterMap"] = Marshal2String(st.Filtermap)
	param["size"] = strconv.Itoa(size)

	ret, _ := rpc.UserCfClientMethod(method, param)
	for _, v := range ret {
		Debugln(v)
		art := Article{*v}
		artlist = append(artlist, art)
	}

	Debugln(artlist)
	return artlist
}
