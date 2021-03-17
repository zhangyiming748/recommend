package recommend

import (
	"log"
	"recommend/api"
	"recommend/module"
	"recommend/util"
	"time"
)

const (
	MAXTIMEOUT = 800
)

func RecommendEntey(param *module.Param) (res api.AppResponse)  {
	res.ResCode = api.APPRESPONSE_CODE_SUCCESS
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.RESPONSE_CODE_FAIL
			res.ResStatus = api.REQUEST_ERR
			log.Println(err)
		}
	}()
	stgy := chooseRecStgy(param)

	var reslist module.ArticleList
	//最多等待400ms
	c := make(chan module.ArticleList)
	go func() {
		defer util.PanicRecover(util.PanicPosition())
		//t := time.Now()
		if stgy.GetAlgoStgy() == module.PERSONALALGO {
			reclist := getPersonalmRecommend(stgy)
			for i, rli := range reclist {
				util.Infoln("debug for usercf", i, "      ", len(rli))
			}
			if num := checkArticleNum(reclist); stgy.GetSize() > num {
				stgy.SetSize(num) //if article num less than stgy.GetSize(), then mergelist function will forever loop
				util.Debugln(stgy.GetUniqueid(), "size:", stgy.GetSize())
			}
			if stgy.GetSize() > 0 {
				util.Debugln(".....before mergeList.....")
				c <- mergeList(stgy, reclist, chooseProb)
			}
		} else {
			var cprob = []float64{1.0} //随机推荐仅获取新内容列表
			reclist := getRandomRecommend(stgy)
			if num := checkArticleNum(reclist); stgy.GetSize() > num {
				stgy.SetSize(num)
				util.Debugln(stgy.GetUniqueid(), "size:", stgy.GetSize())
			}
			if stgy.GetSize() > 0 {
				util.Debugln(".....before mergeList.....")
				c <- mergeList(stgy, reclist, cprob)
			}
		}
		//DataLogln("rec&merge:\t",time.Since(t))
	}()
	util.Debugln(stgy.GetUniqueid(), "size:", stgy.GetSize())
	select {
	case reslist = <-c:
	case <-time.After(MAXTIMEOUT * time.Millisecond):
		stgy.SetAlgoStgy("FAIL")
		util.Warningf("wait req:%v response %dms time out!\n", stgy.GetUniqueid(), MAXTIMEOUT)
	}
	util.Debugln(stgy.GetUniqueid(), "size:", stgy.GetSize())
	return
}
func checkArticleNum(reclist []module.ArticleList) (num int) {
	numMap := make(map[string]bool)
	for _, rli := range reclist {
		for _, art := range rli {
			numMap[art.GetId()] = true
		}
	}
	return len(numMap)
}