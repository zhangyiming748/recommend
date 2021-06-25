package usercf

import (
	"encoding/json"
	"errors"
	"fmt"
	. "recommend/model"
	"recommend/model/query"
	. "recommend/util"
	"strconv"
	"time"
)

type kv struct {
	Key   string
	Value float64
}

//编码json
func (this kv) MarshalJSON() ([]byte, error) {
	var tmps []string
	tmps = append(tmps, fmt.Sprintf("%s", this.Key))
	tmps = append(tmps, fmt.Sprintf("%f", this.Value))
	return json.Marshal(tmps)
}

//解码JSON
func (this *kv) UnmarshalJSON(data []byte) (err error) {
	if this == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	var tmps []string
	if err = json.Unmarshal(data, &tmps); err != nil {
		this.Key = tmps[0]
		this.Value, err = strconv.ParseFloat(tmps[1], 64)
	}

	return nil
}

const (
	DefaultK = 50
	DefaultN = 50
)
const WeeksSec = 14 * 24 * 3600 //time.Second

//得到相似用户点击列表
func GetUserCfItems(st Stgy, K, N int) ArticleList {
	uid := st.GetUserId()
	if K == 0 {
		K = DefaultK
	}
	if N == 0 {
		N = DefaultK
	}

	rank, item_score := usercf(uid, st.Filtermap, K)
	var data []*Article
	if len(rank) > 0 {
		keys := make([]string, 0)
		for _, kv := range rank {
			keys = append(keys, kv.Key)
			Debugf("%s, %f\n", kv.Key, kv.Value)
		}
		//仅仅获取点击内容中当前频道内容
		data = query.QueryArticleByIds(st.GetChannel(), keys)
	}
	Debugln(data)
	filterdata := make([]*Article, 0)
	for i, _ := range data {
		var art Article
		art = *data[i]
		if t, err := strconv.ParseFloat(art.GetDatpublis(), 64); err == nil {
			if int64(t) < time.Now().Unix()-WeeksSec {
				//过滤发布时间过旧的内容
				continue
			}
		} else {
			continue
			Errorln(err)
		}
		Debugln("debugxxx前瞻类到期时间: ", art.GetRecommendTimeLimit(), "debugxxxtimenowunix", time.Now().Unix())
		if art.GetRecommendTimeLimit() > 0 && art.GetRecommendTimeLimit() < time.Now().Unix() {
			//过滤过期前瞻类内容
			continue
		}
		art.Finalscore = item_score[data[i].GetId()]
		art.Servertag = "usercf"
		filterdata = append(filterdata, &art)
	}
	if N > len(filterdata) {
		N = len(filterdata)
	}
	filterdata = filterdata[:N]

	retdata := make([]Article, 0)
	for _, art := range filterdata {
		retdata = append(retdata, *art)
	}
	Debugln("usercf req end:", st.GetUniqueid(), "datalens:", len(filterdata))
	return retdata
}
