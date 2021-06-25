package query

import (
	. "recommend/model"
	"recommend/storage"
	. "recommend/util"
	"time"
)

func qArticleFromES(channel string, ids []string) (al []*Article) {
	if channel == HOMECHANNEL {
		return getDataByIds(ids)
	} else if channel == HOTVIDEOCHANNEL {
		return getDataByIds4video(ids)
	}
	return
}

func QueryArticleByIds(channel string, ids []string) []*Article {
	//获取id给定的合法article
	var data []*Article
	idsmap := make(map[string]bool)

	// first find in cache, if channel is "", then find all cache
	if channel == "" {
		data = qArticleFromCache(HOMECHANNEL, ids)
		data2 := qArticleFromCache(HOTVIDEOCHANNEL, ids)
		data = append(data, data2...)
	} else {
		data = qArticleFromCache(channel, ids)
	}
	Infoln("lendata", len(data))

	// second find is es
	resids := make([]string, 0)
	for _, item := range data {
		idsmap[item.GetId()] = true
	}
	for _, i := range ids {
		if _, ok := idsmap[i]; !ok {
			resids = append(resids, i)
		}
	}
	Infoln("lenresid", len(resids))

	if channel == "" {
		data1 := qArticleFromES(HOMECHANNEL, resids)
		Infoln("lenresid1: ", len(data1))

		data = append(data, data1...)
		data2 := qArticleFromES(HOTVIDEOCHANNEL, resids)
		Infoln("lenresid2: ", len(data2))

		data = append(data, data2...)
	} else {
		data1 := qArticleFromES(channel, resids)
		data = append(data, data1...)
	}
	return data
}

func QueryArticleForTags(ids []string) []Article {
	//缓存article以供personal， merge， template快速调用, 仅供计算tags使用， 允许下架等不合法article存在
	//注意cache中的article可能已经不合法
	artlist := make([]Article, 0)
	resids := make([]string, 0)
	for _, i := range ids {
		val, found := storage.FetchCacheContent(ACTICLETAGSKEY + ":" + string(i))
		if !found {
			resids = append(resids, i)
		} else {
			artlist = append(artlist, val.(Article))
		}
	}
	if len(resids) == 0 {
		return artlist
	}
	resdata := QueryArticleByIds("", resids)

	for _, art := range resdata {
		var a Article = *art
		storage.SetCacheContent(ACTICLETAGSKEY+":"+art.GetId(), a, storage.MAXEXPIRATION)
		artlist = append(artlist, a)
	}

	return artlist
}

func QueryNewArticles(st Stgy, filterids []string, size int) []*Article {
	data := qFromCache(st.GetChannel())
	newdata := make([]*Article, 0)
	filterMap := make(map[string]bool)
	for _, f := range filterids {
		filterMap[f] = true
	}
	j := 0
	for _, v := range data {
		//Debugln("debugxxx前瞻类到期时间: ",v.GetRecommendTimeLimit(),"debugxxxtimenowunix", time.Now().Unix())
		if _, ok := filterMap[v.GetId()]; ok {
			continue
		} else if v.GetRecommendTimeLimit() > 0 && v.GetRecommendTimeLimit() < time.Now().Unix() {
			//过滤过期前瞻类内容
			continue
		} else {
			newdata = append(newdata, v)
			j += 1
			if j > size {
				break
			}

		}
	}
	return newdata
}

func QueryRecArticles(st Stgy, filterids []string, size int) []*Article {
	data := qFromCache(st.GetChannel())
	// 准备用户画像词, 优先召回符合画像词内容
	var history_usertag = st.History_usertag
	var realtime_usertag = st.Realtime_usertag

	tagmap := make(map[string]bool)
	for k, _ := range history_usertag {
		tagmap[k] = true
	}
	for k, _ := range realtime_usertag {
		tagmap[k] = true
	}

	newdata := make([]*Article, 0)
	filterMap := make(map[string]bool)
	for _, f := range filterids {
		filterMap[f] = true
	}

	j := 0
	for _, v := range data { //优先填充画像词相关内容
		if _, ok := filterMap[v.GetId()]; ok {
			continue
		}
		if v.GetRecommendTimeLimit() > 0 && v.GetRecommendTimeLimit() < time.Now().Unix() {
			//过滤过期前瞻类内容
			continue
		}

		added := false
		for _, x := range v.GetTags() {
			if _, ok := tagmap[x]; ok {
				added = true
			}
		}
		if _, ok := tagmap[v.GetLargeclass()]; ok {
			added = true
		}
		if _, ok := tagmap[v.GetMediumclass()]; ok {
			added = true
		}
		if _, ok := tagmap[v.GetSmallclass()]; ok {
			added = true
		}
		if added {
			newdata = append(newdata, v)
			filterMap[v.GetId()] = true //已经填充过的内容, 不再填充
			j += 1
			if j > size {
				break
			}
		}
	}
	if j <= size { //若画像词相关内容不够数量, 则以其余内容补上
		Infoln("usertag article smaller than size: ", j)
		for _, v := range data {
			if _, ok := filterMap[v.GetId()]; ok {
				continue
			} else if v.GetRecommendTimeLimit() > 0 && v.GetRecommendTimeLimit() < time.Now().Unix() {
				//过滤过期前瞻类内容
				continue
			} else {
				newdata = append(newdata, v)
				j += 1
				if j > size {
					break
				}
			}
		}
	} else {
		Debugln("usertag article bigger than size: ", j)
	}
	return newdata
}
