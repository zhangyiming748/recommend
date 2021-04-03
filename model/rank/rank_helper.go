package rank

import (
	. "recommend/model"
	. "recommend/util"
	"recommend/model/query"
)

const (
	FINDSIZE = 100
	//从缓存中找多少article用来查找相似
)

func Rank_fornew(st Stgy, size int) ArticleList {
	filterids := make([]string, 0)
	for id, _ := range st.Filtermap {
		filterids = append(filterids, id)
	}
	var data []*Article
	//for explore, only get 150
	data = query.QueryNewArticles(st, filterids, 150)

	_data := rank(st, data, NEWLYFORMULA, st.GetAbtestStgy())
	newsize := len(_data)
	if newsize > size {
		newsize = size
	}
	_data = _data[0:newsize]

	retdata := make([]Article, 0)
	for _, art := range _data {
		retdata = append(retdata, *art)
	}
	Debugln("fornew req end:", st.GetUniqueid(), "datalens:", len(_data))
	return retdata
}

func Rank_forrec(st Stgy, size int) ArticleList {
	filterids := make([]string, 0)
	for id, _ := range st.Filtermap {
		filterids = append(filterids, id)
	}

	var data []*Article
	//for explore, only get 150
	data = query.QueryRecArticles(st, filterids, 150)

	_data := rank(st, data, RECFORMULA, st.GetAbtestStgy())
	newsize := len(_data)
	if newsize > size {
		newsize = size
	}
	_data = _data[0:newsize]

	retdata := make([]Article, 0)
	for _, art := range _data {
		retdata = append(retdata, *art)
	}

	Debugln("forrec req end:", st.GetUniqueid(), "datalens:", len(_data))
	return retdata
}
func Rank_forSimiart(article Article, size int) ArticleList {
	filterids := make([]string, 0)
	//过滤列表只过滤当前article
	filterids = append(filterids, article.GetId())
	var data []*Article
	var st Stgy

	//获取新的Article
	data = query.QueryNewArticles(st, filterids, FINDSIZE)

	_data := rankSimilar(article, data, NEWLYFORMULA)
	newsize := len(_data)
	if newsize > size {
		newsize = size
	}
	_data = _data[0:newsize]
	retdata := make([]Article, 0)
	for _, art := range _data {
		retdata = append(retdata, *art)
	}
	//Debugln("for this article req end:", st.GetUniqueid(), "datalens:", len(_data))

	return retdata
}
