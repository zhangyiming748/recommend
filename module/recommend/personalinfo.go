package recommend

import (
	"encoding/json"
	"log"
	"math"
	"recommend/module"
	"recommend/module/query"
	"recommend/storage"
	"recommend/util"
	"sync"
)

func getUserIdByUuid(uuid string) string {
	//根据设备id反查userid
	r := storage.GetRedis(storage.REDIS_UUIDMAP_KEY)
	uid, _ := r.Get(storage.REDIS_UUIDMAP_KEY + ":" + uuid)
	return uid
}
var names = []string{storage.REDIS_SHOWLIST_KEY, storage.REDIS_CLICKLIST_KEY, storage.REDIS_QUICKSHOWLIST_KEY}
var lens = []int{module.SHOWLISTLEN, module.CLICKLISTLEN, module.QUICKSHOWLISTLEN}

func getUserPersonalInfo(user string) (pinfo module.PersonalInfo) {
	var wg sync.WaitGroup
	lists := make([][]string, 3)
	for i, s := range names {
		wg.Add(1)
		go func(i int, s string) {
			defer wg.Done()
			defer util.PanicRecover(util.PanicPosition())
			r := storage.GetRedis(s)
			lists[i], _ = r.Lrange(s+":"+user, 0, lens[i])
		}(i, s)
	}

	var history_usertag map[string]string
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer util.PanicRecover(util.PanicPosition())
		r := storage.GetRedis(storage.REDIS_USERTAGS_KEY)
		history_usertag_str, _ := r.Get(storage.REDIS_USERTAGS_KEY + ":" + user)
		json.Unmarshal([]byte(history_usertag_str), &history_usertag)

		log.Println("getUserPersonalinfo ", user, " usertaglen:", len(history_usertag), ".")
		log.Println("getUserPersonalinfo ", user, " usertaglen:", history_usertag, ".")

	}()
	// "usertags:5899"
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer util.PanicRecover(util.PanicPosition())
		rr := storage.GetRedis(storage.ARTICLENFB_KEY)
		nfb_maps, _ := rr.HgetAll(storage.ARTICLENFB_KEY + ":" + user)
		nfb_list := make([]string, 0)
		for k, _ := range nfb_maps {
			nfb_list = append(nfb_list, k)
		}
		pinfo.Nfblist = nfb_list
	}()

	wg.Wait()

	pinfo.Showlist, pinfo.Clicklist, pinfo.Quick_showlist = lists[0], lists[1], lists[2]
	pinfo.SetHistoryUsertag(history_usertag)

	go func() {
		defer PanicRecover(PanicPosition())
		//异步清理redis
		for i, s := range names {
			r := storage.GetRedis(s)
			si, _ := r.LLEN(s + ":" + user)
			lens := []int{SHOWLISTLEN + 100, CLICKLISTLEN + 100, QUICKSHOWLISTLEN + 10}
			for ; si > lens[i]; si-- {
				r.RPOP(s + ":" + user)
			}

		}
	}()
	return
}

const (
	ComputeClicklen = 5
)

func getUserRealtimeTags(clicklist []string) map[string]float64 {
	// only use latest ComputeClicklen articles to compute realtime usertags
	var cl []string
	if len(clicklist) > ComputeClicklen {
		cl = clicklist[0:ComputeClicklen]
	} else {
		cl = clicklist
	}

	clickArticle := query.QueryArticleForTags(cl)

	//fmt.Println("-----getUserRealtimeTags_clickArticle-----", clickArticle)
	realtime_usertag := make(map[string]float64)
	for i, article_id := range cl {
		for _, art := range clickArticle {
			if article_id != art.GetId() {
				continue
			}
			tags := make([]string, 0)
			if art.GetLargeclass() != "" {
				tags = append(tags, art.GetLargeclass())
			}
			if art.GetMediumclass() != "" {
				tags = append(tags, art.GetMediumclass())
			}
			if art.GetSmallclass() != "" {
				tags = append(tags, art.GetSmallclass())
			}
			if art.GetTags() != nil {
				tags = append(tags, art.GetTags()...)
			}
			for _, tag := range tags {
				realtime_usertag[tag] += 100 * (math.Pow(0.8, float64(i)))
			}
		}
	}

	//fmt.Println("-----getUserRealtimeTags_realtime_usertag-----", realtime_usertag)
	return realtime_usertag
}
