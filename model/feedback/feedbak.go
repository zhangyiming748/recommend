package feedback

import (
	"../../api"
	. "../../util"
)

type KV struct {
	Key   string
	Value string
}

//反馈NegativeFeedback
func NegativeFb(articleId string) (res api.AppResponse) {
	res.ResCode = api.APPRESPONSE_CODE_SUCCESS
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.RESPONSE_CODE_FAIL
			res.ResStatus = api.REQUEST_ERR
			Errorln(err)
		}
	}()

	baseMap := make([]KV, 0)
	baseMap = append(baseMap, KV{"A1", "不感兴趣"})
	baseMap = append(baseMap, KV{"A2", "内容质量不佳"})
	baseMap = append(baseMap, KV{"A3", "看过这条"})

	newMap := make([]KV, 0)
	/*
		var art Article
		arts := query.QueryArticleFortags("", []string{articleId,})
		if art.Id != "" {
			if len(art.Tags) > 0 && art.Tags[0] != "" {
				newMap = append(newMap, KV{"T1", "不喜欢" + art.Tags[0]})
			}
			if art.Largeclass != "" {
				newMap = append(newMap, KV{"C1", "不喜欢" + art.Largeclass})
			}
			if art.Mediumclass != "" {
				newMap = append(newMap, KV{"C2", "不喜欢" + art.Mediumclass})
			}
			tags := art.Tags[1:]
			for i, tag := range tags {
				newMap = append(newMap, KV{"T" + strconv.Itoa(i+2), "不喜欢" + tag})
			}
		}
	*/
	data := make(map[string]interface{}, 0)
	data["baseList"] = baseMap
	data["newList"] = newMap
	res.RetData = data
	return
}
