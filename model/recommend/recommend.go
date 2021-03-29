package recommend

import (
	"recommend/api"
	. "recommend/model"
	"recommend/storage"
	"time"
	. "recommend/util"
)

func RecommendEntry(param *Param) (res api.AppResponse) {
	res.ResCode = api.APPRESPONSE_CODE_SUCCESS

	defer func() {
		trafficDone()
		if err := recover(); err != nil {
			res.ResCode = api.RESPONSE_CODE_FAIL
			res.ResStatus = api.REQUEST_ERR
			Errorln(err)
		}
	}()
	trafficBegin()

	stgy := chooseRecStgy(param)

	var reslist ArticleList
	//最多等待400ms
	c := make(chan ArticleList)
	go func() {
		defer PanicRecover(PanicPosition())
		//t := time.Now()
		if stgy.GetAlgoStgy() == PERSONALALGO {
			reclist := getPersonalmRecommend(stgy)
			for i, rli := range reclist {
				Infoln("debug for usercf", i, "      ", len(rli))
			}
			if num := checkArticleNum(reclist); stgy.GetSize() > num {
				stgy.SetSize(num) //if article num less than stgy.GetSize(), then mergelist function will forever loop
				Debugln(stgy.GetUniqueid(), "size:", stgy.GetSize())
			}
			if stgy.GetSize() > 0 {
				Debugln(".....before mergeList.....")
				c <- mergeList(stgy, reclist, chooseProb)
			}
		} else {
			var cprob = []float64{1.0} //随机推荐仅获取新内容列表
			reclist := getRandomRecommend(stgy)
			if num := checkArticleNum(reclist); stgy.GetSize() > num {
				stgy.SetSize(num)
				Debugln(stgy.GetUniqueid(), "size:", stgy.GetSize())
			}
			if stgy.GetSize() > 0 {
				Debugln(".....before mergeList.....")
				c <- mergeList(stgy, reclist, cprob)
			}
		}
		//DataLogln("rec&merge:\t",time.Since(t))
	}()
	Debugln(stgy.GetUniqueid(), "size:", stgy.GetSize())
	select {
	case reslist = <-c:
	case <-time.After(MAXTIMEOUT * time.Millisecond):
		stgy.SetAlgoStgy("FAIL")
		Warningf("wait req:%v response %dms time out!\n", stgy.GetUniqueid(), MAXTIMEOUT)
	}
	Debugln(stgy.GetUniqueid(), "size:", stgy.GetSize())
	//使用reslist渲染结果
	//t := time.Now()
	res.RetData = template.DrawTemplate(stgy, reslist)
	//DataLogln("drawTemplate:\t",time.Since(t))

	//异步处理快速过滤队列
	if stgy.GetAlgoStgy() == PERSONALALGO {
		go func(u string, reslist ArticleList) {
			defer PanicRecover(PanicPosition()) //quick_showlist
			r := storage.GetRedis(storage.REDIS_QUICKSHOWLIST_KEY)
			for _, article := range reslist {
				r.LPUSH(storage.REDIS_QUICKSHOWLIST_KEY+":"+u, article.GetId())
			}
		}(stgy.GetUserId(), reslist)
	}
	//异步打印日志
	go func() {
		defer PanicRecover(PanicPosition())
		for i, val := range reslist {
			//算法 AB策略 唯一id uid 设备id 动作 位置 内容id 排序服务 最终得分 算分详情
			ReportLogf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", stgy.GetAlgoStgy(), stgy.GetAbtestStgy(), stgy.GetUniqueid(), stgy.GetUserId(), stgy.GetDeviceId(), stgy.GetChannel()+stgy.GetAction(), i, val.GetId(), val.GetServertag(), val.GetFinalscore(), val.GetComputescore())
		}
	}()
	return
}
