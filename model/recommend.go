package main

import (
	"../api"
	"../storage"
	. "../util"
	"fmt"
	"time"
)

var chooseProb []float64

func SetChooseProb(cprob []string) {
	for _, ff := range cprob {
		var f float64
		fmt.Sscanf(ff, "%f", &f)
		chooseProb = append(chooseProb, f)
	}
	Infoln(chooseProb)
}

/**
  首页推荐
*/
func AppHomePage(param *Param) (res api.AppResponse) {
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
		//t := time.Now()
		//if stgy.GetAlgoStgy() == PERSONALALGO && stgy.GetAbtestStgy() == "A" {
		if stgy.GetAlgoStgy() == PERSONALALGO {
			reclist := getPersonalmRecommend(stgy)
			for i, rli := range reclist {
				Infoln("debug for usercf", i, "      ", len(rli))
			}
			if num := checkArticleNum(reclist); stgy.size > num {
				stgy.size = num //if article num less than stgy.size, then mergelist function will forever loop
				Debugln(stgy.uniqueid, "size:", stgy.size)
			}
			if stgy.size > 0 {
				Debugln(".....before mergeList.....")
				c <- mergeList(stgy, reclist, chooseProb)
			}
		} else {
			var cprob = []float64{1.0} //随机推荐仅获取新内容列表
			reclist := getRandomRecommend(stgy)
			if num := checkArticleNum(reclist); stgy.size > num {
				stgy.size = num
				Debugln(stgy.uniqueid, "size:", stgy.size)
			}
			if stgy.size > 0 {
				Debugln(".....before mergeList.....")
				c <- mergeList(stgy, reclist, cprob)
			}
		}
		//DataLogln("rec&merge:\t",time.Since(t))
	}()
	Debugln(stgy.uniqueid, "size:", stgy.size)
	select {
	case reslist = <-c:
	case <-time.After(400 * time.Millisecond):
		stgy.SetAlgoStgy("FAIL")
		Errorf("wait req:%v response 400ms time out!\n", stgy.uniqueid)
	}
	Debugln(stgy.uniqueid, "size:", stgy.size)

	//使用reslist渲染结果
	//t := time.Now()
	res.RetData = drawTemplate(stgy, reslist)
	//DataLogln("drawTemplate:\t",time.Since(t))

	//异步处理快速过滤队列
	if stgy.GetAlgoStgy() == PERSONALALGO && stgy.GetAbtestStgy() == "A" {
		go func(u string, reslist ArticleList) {
			//quick_showlist
			r := storage.GetRedis(storage.REDIS_QUICKSHOWLIST_KEY)
			for _, article := range reslist {
				r.LPUSH(storage.REDIS_QUICKSHOWLIST_KEY+":"+u, article.GetId())
			}
		}(stgy.GetUserId(), reslist)
	}
	//异步打印日志
	go func() {
		for i, val := range reslist {
			//算法 AB策略 唯一id uid 设备id 动作 位置 内容id 排序服务 最终得分 算分详情
			ReportLogf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", stgy.GetAlgoStgy(), stgy.GetAbtestStgy(), stgy.uniqueid, stgy.userId, stgy.deviceId, stgy.action, i, val.GetId(), val.GetServertag(), val.GetFinalscore(), val.GetComputescore())
		}
	}()
	return
}

func checkArticleNum(reclist []ArticleList) (num int) {
	numMap := make(map[string]bool)
	for _, rli := range reclist {
		for _, art := range rli {
			numMap[art.GetId()] = true
		}
	}
	return len(numMap)
}

func drawTemplate(st Stgy, reslist ArticleList) (recommendInfo RecommendInfo) {
	//t := time.Now()
	var retData map[string]interface{}
	var list []interface{}
	retData, ok := GetTemplateRetData()
	if ok {
		recommendInfo.strategy = st.GetAlgoStgy() + "_" + st.GetAbtestStgy() + "_" + st.uniqueid
		recommendInfo.channelId = retData["channelId"].(string)
		recommendInfo.action = retData["action"].(string)
		list, ok = retData["list"].([]interface{})
	}
	//DataLogln("parseTemplate:\t",time.Since(t))

	//t = time.Now()

	resids := make([]string, 0)
	for _, art := range reslist {
		resids = append(resids, art.GetId())
	}
	tempList := TempArtilceList(resids)
	//DataLogln("TempArtilceList:\t",time.Since(t))

	Debugln(".......推荐结果id:", resids, st.uniqueid)
	action := st.GetAction()
	if action == "" {
		action = ACTION_HOME
	}
	if action != ACTION_HOME {
		if len(tempList) > 0 {
			Debugln("......len(tempList):", len(tempList), st.uniqueid, st.size)
			list = make([]interface{}, 0)
			for _, val := range tempList {
				list = append(list, val)
			}
		} else { //如果没有推荐结果, 则直接返回模板内容
			Debugln("......len(tempList):", len(tempList), "---", st.uniqueid, "----", st.size)
			if RECOMMEND_SIZE_UP < len(list) {
				list = list[0:RECOMMEND_SIZE_UP]
			}
		}
	} else {
		j := 0
		for i, val := range list {
			if j >= len(tempList) {
				break //如果内容已经被用光, 则无需继续渲染
			}

			v, _ := val.(map[string]interface{})
			if v["list_type"].(string) == "free" {
				continue
			}
			vv := tempList[j]
			j += 1
			list[i] = vv
		}
	}
	//Debugln("returnlist",list)
	recommendInfo.action = action
	recommendInfo.list = list
	recommendInfo.size = len(list)

	return
}
