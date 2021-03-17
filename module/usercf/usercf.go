package usercf

import (
	"../../storage"
	. "../../util"
	"encoding/json"
	"sort"
	"strconv"
	"sync"
)

func getSimiUsers(uid string) (simi_users []kv) {
	//获取相似用户
	simi_users = make([]kv, 0)
	r := storage.GetRedis(storage.REDIS_USERSIMI_KEY)

	jstr, err := r.Get(storage.REDIS_USERSIMI_KEY + ":" + uid)
	jobj := make([]interface{}, 0)
	if err = json.Unmarshal([]byte(jstr), &jobj); err != nil {
		Errorln("parsing json simi_users", err.Error())
	}
	for _, v := range jobj {
		vv := v.([]interface{})
		var _kv kv
		_kv.Key = vv[0].(string)
		_kv.Value, _ = strconv.ParseFloat(vv[1].(string), 64)
		simi_users = append(simi_users, _kv)
	}
	return
}

func getUserItems(simi_users []kv, K int) (items [][]string) {
	//获取相似用户的点击
	var wg sync.WaitGroup
	r := storage.GetRedis(storage.REDIS_CLICKLIST_KEY)
	items = make([][]string, K)
	for i, s := range simi_users[0:K] {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer PanicRecover(PanicPosition())
			itemids, err := r.Lrange(storage.REDIS_CLICKLIST_KEY+":"+s.Key, 0, 50)
			if err != nil {
				Errorln("getUserItems:", storage.REDIS_CLICKLIST_KEY+":"+s.Key, " error")
				return
			}
			items[i] = itemids
		}(i)
	}
	wg.Wait()
	return
}

func usercf(uid string, filterMap map[string]bool, K int) (rank []kv, items_score map[string]float64) {
	simi_users := getSimiUsers(uid)

	if K > len(simi_users) {
		K = len(simi_users)
	}
	if K == 0 {
		return
	}

	items := getUserItems(simi_users, K)

	items_score = make(map[string]float64)
	for i, s := range simi_users[0:K] {
		if len(items[i]) > 0 {
			for _, itemid := range items[i] {
				items_score[itemid] += s.Value
			}
		}
	}

	rank = make([]kv, 0)
	for k, v := range items_score {
		//已曝光过滤逻辑
		if _, ok := filterMap[k]; !ok {
			rank = append(rank, kv{k, v})
		}
	}

	sort.Slice(rank, func(i, j int) bool {
		return rank[i].Value > rank[j].Value // 降序
	})
	Debugln(rank)

	return
}
