package history

import (
	"../../api"
	. "../../model"
	"../../storage"
	. "../../util"

	"../query"
	"../template"
	"strconv"
)

const PAGEITEMNUM = 10 //一页多少内容

func getClickHistoryByUid(uid string) []string {
	//根据用户id查历史点击列表
	r := storage.GetRedis(storage.REDIS_CLICKLIST_KEY)
	clh, _ := r.Lrange(storage.REDIS_CLICKLIST_KEY+":"+uid, 0, -1)

	return clh
}

//获取用户点击列表
func GetHistory(param *Param) (res api.AppResponse) {
	res.ResCode = api.APPRESPONSE_CODE_SUCCESS
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.RESPONSE_CODE_FAIL
			res.ResStatus = api.REQUEST_ERR
			Errorln(err)
		}
	}()
	//根据用户id查历史点击列表
	clh := getClickHistoryByUid(param.GetUserId())
	Infoln("xzhaodebugxxx", len(clh))
	clmap := make(map[string]bool)
	for _, i := range clh {
		clmap[i] = true
	}
	Infoln("xzhaodebugxxx", len(clmap), clmap)

	//验证点击历史ids为合法未下架的Articles
	articles := query.QueryArticleByIds("", clh)
	validmap := make(map[string]bool)
	for _, art := range articles {
		validmap[art.GetId()] = true
	}

	//拿到完整列表,去重
	if len(clh) > 0 {
		clh = dedupe(clh, validmap)
	}
	Infoln("xzhaodebugxxx", len(clh))

	//截取相应的ids，从第0页开始，每页10个内容
	pageno, _ := strconv.Atoi(param.GetPageNum())
	if len(clh) < pageno*PAGEITEMNUM {
		//点击内容已用尽，直接返回
		var retdata RecommendInfo
		retdata.SetSize(0)
		retdata.SetAction(param.GetPageNum())
		res.RetData = retdata
		return
	}
	var newclh []string
	if len(clh) > pageno*PAGEITEMNUM+PAGEITEMNUM {
		newclh = clh[pageno*PAGEITEMNUM : pageno*PAGEITEMNUM+PAGEITEMNUM]
	} else {
		newclh = clh[pageno*PAGEITEMNUM : len(clh)]
	}

	//截取相应的内容,调用template包，获取articleDetail，获取内容并返回
	detailList := template.GetArticleDetailList(newclh)
	list := make([]interface{}, 0)
	for _, val := range detailList {
		list = append(list, val)
	}
	var recommendInfo RecommendInfo
	recommendInfo.SetList(list)
	recommendInfo.SetSize(len(list))
	recommendInfo.SetAction(param.GetPageNum())

	res.RetData = recommendInfo
	return
}

// 保序去重
func dedupe(list []string, validmap map[string]bool) []string {
	stand := make(map[string]bool)
	newList := make([]string, 0)
	for _, j := range list {
		if _, ok := validmap[j]; ok && stand[j] == false {
			newList = append(newList, j)
			stand[j] = true
		}
	}
	return newList
}
