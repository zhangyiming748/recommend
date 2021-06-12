package related

import (
	_ "log"
	"recommend/api"
	"recommend/model"
	"recommend/model/query"
	"recommend/model/rank"
	"recommend/model/template"
	. "recommend/util"
	"time"
)

//TO-DO
//完善代码可读性
//加入发布时间的算分
//加入tags的算分
//改代码取消指定频道列表

const (
	SIZE = 50
	//取多少新的article
	SELECTSIZE = 5
	//推荐几个符合条件的article
)

func GetRelated(id string) (res api.AppResponse) {
	t1 := time.Now()
	res.ResCode = api.APPRESPONSE_CODE_SUCCESS
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.RESPONSE_CODE_FAIL
			res.ResStatus = api.REQUEST_ERR
			Errorln(err)
		}
	}()
	//在缓存中查找当前id的Article
	var ids []string
	ids = append(ids, id)
	articles := query.QueryArticleByIds("", ids)
	var article model.Article
	article = *articles[0]

	suitable := rank.Rank_forSimiart(article, SIZE)

	//拿到排序列表后需要返回几个?
	if len(suitable) > SELECTSIZE {
		suitable = suitable[0:SELECTSIZE]
	}
	recids := make([]string, 0)

	for _, v := range suitable {
		recids = append(recids, v.GetId())
		//Debugf("get articleId:%v", v)

	}
	Debugf("符合条件的Top 5 articleId:%v\n", recids)
	q := query.QueryArticleByIds("", recids)
	for _, v := range q {
		Debugf("发布时间为:%v\t大中小标签:%v\t%v\t%v\t%v\n", v.Datpublis, v.GetLargeclass(), v.GetMediumclass(), v.GetSmallclass(), v.GetTags())
	}
	detailList := template.GetArticleDetailList(recids)
	res.RetData = detailList
	t2 := time.Now()
	Debugf("一次请求运行用时%v", t2.Sub(t1))
	return
}
