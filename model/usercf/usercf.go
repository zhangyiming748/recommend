package usercf

import (
	"encoding/json"
	"recommend/storage"
	. "recommend/util"
	"sort"
	"strconv"
	"sync"
)

//获取相似用户
func getSimiUsers(uid string) (simi_users []kv) {

	simi_users = make([]kv, 0)
	r := storage.GetRedis(storage.REDIS_USERSIMI_KEY)

	jstr, err := r.Get(storage.REDIS_USERSIMI_KEY + ":" + uid)
	//userSimilar:00c7c3b81a464b45850b993b2efab562
	jobj := make([]interface{}, 0)
	if err = json.Unmarshal([]byte(jstr), &jobj); err != nil {
		Errorln("parsing json simi_users", err.Error())
	}
	//jobj是解析的json文件
	for _, v := range jobj {
		vv := v.([]interface{})
		var _kv kv
		_kv.Key = vv[0].(string)
		//strconv.ParseFloat字符串转换浮点数
		_kv.Value, _ = strconv.ParseFloat(vv[1].(string), 64)
		//格式uid:相似分数
		simi_users = append(simi_users, _kv)
	}
	return
}

/*
[
	[
		"ee92c5f7c4c2441cac6370c171b02368",
		"0.082203"
	],
	[
		"f72071c79af84f60b833c7dd6941484b",
		"0.081658"
	],
	[
		"4d5e68da5029410a8e0343f5c4bbf5d1",
		"0.073148"
	],
	[
		"18070764",
		"0.071847"
	],
	[
		"b4fd4a82ced749da882d0d0bc2743201",
		"0.071497"
	],
	[
		"6507",
		"0.070434"
	],
	[
		"a2909a066ae74263985a1fa3e2d7fa9b",
		"0.068991"
	],
	[
		"430637",
		"0.068991"
	]
]
*/
func getUserItems(simi_users []kv, K int) (items [][]string) {
	//获取相似用户的点击
	var wg sync.WaitGroup
	r := storage.GetRedis(storage.REDIS_CLICKLIST_KEY)
	items = make([][]string, K)
	for i, s := range simi_users[0:K] {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer PanicRecover()
			itemids, err := r.Lrange(storage.REDIS_CLICKLIST_KEY+":"+s.Key, 0, 50)
			if err != nil {
				Errorln("getUserItems:", storage.REDIS_CLICKLIST_KEY+":"+s.Key, " error")
			}
			items[i] = itemids
		}(i)
	}
	wg.Wait()
	return
}

//基于用户的协同过滤算法
//找到和目标用户兴趣相似的用户集合
//找到这个集合中的用户喜欢的，且目标用户没有听说过的物品推荐给目标用户
//取50个相近用户, 取50条, 如果不够, 有多少给多少
func usercf(uid string, filterMap map[string]bool, K int) (rank []kv, items_score map[string]float64) {
	simi_users := getSimiUsers(uid)
	//defaultK=50
	if K > len(simi_users) {
		K = len(simi_users)
	}
	if K == 0 {
		return
	}

	items := getUserItems(simi_users, K)
	//item=[][]string
	items_score = make(map[string]float64)
	for i, s := range simi_users[0:K] {
		if len(items[i]) > 0 { //人数大于0
			for _, itemid := range items[i] { //人元素分x2
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
