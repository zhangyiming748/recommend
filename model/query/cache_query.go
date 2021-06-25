package query

import (
	. "recommend/model"
	"recommend/storage"
	. "recommend/util"
	"time"
)

const (
	CACHEHOMEQUERYRES = "cachehomequeryres"
	CACHEHOTQUERYRES  = "cachehotqueryres"
)

func CacheQueryinit() {
	loopfunc := func(cachekey string) {
		for {
			func() {
				defer PanicRecover(PanicPosition())
				var data []*Article
				if cachekey == CACHEHOMEQUERYRES {
					data = qNewArticlesFromES(5000)
				} else if cachekey == CACHEHOTQUERYRES {
					data = qNewArticlesFromES4video(5000)
				}
				Infoln(cachekey, " fill cache from es nums: ", len(data))
				if len(data) > 0 {
					storage.SetCacheContent(cachekey, data, 60*time.Second)
					//for i, v := range data {
					//	Infoln(i,"----",v.GetDatpublis())
					//}
				}
			}()
			time.Sleep(16 * time.Second)
		}
	}
	go loopfunc(CACHEHOMEQUERYRES)
	go loopfunc(CACHEHOTQUERYRES)
}

func qFromCache(channel string) []*Article {
	var data []*Article
	cachekey := ""
	if channel == "" {
		channel = HOMECHANNEL
	}
	if channel == HOMECHANNEL {
		cachekey = CACHEHOMEQUERYRES
	//} else if channel == HOTVIDEOCHANNEL {
	//	cachekey = CACHEHOTQUERYRES
	}

	for i := 0; i < 5; i++ {
		datain, found := storage.FetchCacheContent(cachekey)
		if !found {
			time.Sleep(5 * time.Millisecond)
		} else {
			data = datain.([]*Article)
			break
		}
	}
	return data
}

func qArticleFromCache(channel string, ids []string) []*Article {
	data := qFromCache(channel)

	idsmap := make(map[string]bool)
	for _, i := range ids {
		idsmap[i] = true
	}

	newdata := make([]*Article, 0)
	for _, item := range data {
		if _, ok := idsmap[item.GetId()]; ok {
			newdata = append(newdata, item)
		}
	}
	return newdata
}
