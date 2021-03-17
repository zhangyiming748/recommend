package template

import (
	. "recommend/module"
	"recommend/storage"
	. "recommend/util"
	"encoding/json"
	"sync"
)

/*
func recoverRecommend(st Stgy, reslist ArticleList) (recommendInfo RecommendInfo) {
	var retData map[string]interface{}
	var list []interface{}

	recommendInfo.SetStrategy(st.GetAlgoStgy() + "bugTemplate" + "_" + st.GetAbtestStgy() + "_" + st.GetUniqueid())
	if st.GetChannel() == HOMECHANNEL {
		recommendInfo.SetChannelId("1731262")
	}
	recommendInfo.SetAction(ACTION_DOWN)

	resids := make([]string, 0)
	for _, art := range reslist {
		resids = append(resids, art.GetId())
	}
	detailList := GetArticleDetailList(resids)

}
*/
func DrawTemplate(st Stgy, reslist ArticleList) (recommendInfo RecommendInfo) {
	//t := time.Now()
	var retData map[string]interface{}
	var list []interface{}

	retData, ok := GetTemplateRetData(st.GetChannel())

	if ok {
		recommendInfo.SetStrategy(st.GetAlgoStgy() + "_" + st.GetAbtestStgy() + "_" + st.GetUniqueid())
		if st.GetChannel() == HOMECHANNEL {
			recommendInfo.SetChannelId(retData["channelId"].(string))
		}
		list, ok = retData["list"].([]interface{})
	}
	//DataLogln("parseTemplate:\t",time.Since(t))

	//t = time.Now()

	resids := make([]string, 0)
	for _, art := range reslist {
		resids = append(resids, art.GetId())
	}
	detailList := GetArticleDetailList(resids)
	//DataLogln("getArticleDetailList:\t",time.Since(t))

	Debugln("recommend result:", resids, st.GetUniqueid())
	action := st.GetAction()
	if action == "" {
		action = ACTION_HOME
	}
	if action != ACTION_HOME {
		if len(detailList) > 0 {
			Debugln("......len(detailList):", len(detailList), st.GetUniqueid(), st.GetSize())
			list = make([]interface{}, 0)
			for i, val := range detailList {
				list = append(list, val)
				Infoln(st.GetUniqueid(), " ", st.GetUserId(), " recpos:", i, " recitem:", val["numarticleid"], " ", val["vc2title"])
				if len(list) > RECOMMEND_SIZE_UP {
					break //可能预取到多于要求数量的内容
				}
			}
		} else { //如果没有推荐结果, 则直接返回模板内容
			Debugln("......len(detailList):", len(detailList), "---", st.GetUniqueid(), "----", st.GetSize())
			if RECOMMEND_SIZE_UP < len(list) {
				list = list[0:RECOMMEND_SIZE_UP]
			}
		}
	} else {
		j := 0
		for i, val := range list {
			if j >= len(detailList) {
				break //如果内容已经被用光, 则无需继续渲染
			}

			v, _ := val.(map[string]interface{})
			if l_type, ok := v["list_type"].(string); ok && l_type == "free" {
				Infoln(st.GetUniqueid(), " ", st.GetUserId(), " recpos:", i, " free_list_type")
				continue
			}
			vv := detailList[j]
			// 由于模板长度大于REC_SIZE_HOME的长度，所以此处实际会多打印几行日志，忽略即可。
			Infoln(st.GetUniqueid(), " ", st.GetUserId(), " recpos:", i, " recitem:", vv["numarticleid"], " ", vv["vc2title"])
			j += 1
			list[i] = vv
		}
		if len(list) > RECOMMEND_SIZE_HOME {
			list = list[0:RECOMMEND_SIZE_HOME] //可能预取到多于要求数量的内容
		}
	}

	//Debugln("returnlist",list)
	recommendInfo.SetAction(action)
	recommendInfo.SetList(list)
	recommendInfo.SetSize(len(list))

	return
}

func GetArticleDetailList(list []string) (ret []map[string]interface{}) {
	defer PanicRecover(PanicPosition())
	ret = make([]map[string]interface{}, len(list))
	r := storage.GetRedis(storage.ARTICLEDETAIL_KEY)
	var wg sync.WaitGroup
	for i, l := range list {
		wg.Add(1)
		go func(i int, l string) {
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					Errorln("getArticleDetailHelper", "pos:", i, " id:", l, " err:", err)
				}
			}()
			detail, err := r.HgetAll(storage.ARTICLEDETAIL_KEY + ":" + l)
			//Debugln("debug....................",detail)
			if len(detail) == 0 {
				Errorln("cannot get articleDetail from redis: ", l, " err: ", err)
				return
			}
			//json,commentNumber||"comment_number":"2",
			jsonStr := detail["json"]

			var tt map[string]interface{}

			if err = json.Unmarshal([]byte(jsonStr), &tt); err != nil {
				Errorln("parsing json file", "pos:", i, " id:", l, " err:", err.Error())
				return
			}
			//if false { // TODO 待测试后才可以打开
			_, ok := tt["content"]
			if ok {
				tt["content"] = ""
			}
			//}
			_, ok = detail["commentNumber"]
			if ok {
				tt["comment_number"] = detail["commentNumber"]
			}
			ret[i] = tt
		}(i, l)

	}
	wg.Wait()
	retnew := make([]map[string]interface{}, 0)
	for _, v := range ret {
		if v != nil {
			retnew = append(retnew, v)
		}
	}
	return retnew
}
