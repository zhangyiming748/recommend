package rank

import (

	"fmt"
	"math"
	"recommend/module"
	"recommend/module/query"
	"recommend/storage"
	"recommend/util"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	PUBLISH            = "publish"
	HISTORY            = "history"
	REALTIME           = "realtime"
	HOTCLICK           = "hotclick"
	NEGFBACk           = "negfback"
	MINS               = -1
	MAXS               = 9999999999999999
	COMPUTE_NFB_MAXLEN = 20
	SIMILAR            = "similar"
)

var weights map[string]map[string]float64

type score struct {
	idx        int
	scores     map[string]float64
	finalscore float64
}

func nfbscore(nfb_list []string, data []*model.Article) (scoremap map[string]float64) {
	scoremap = make(map[string]float64)
	var nl []string
	if len(nfb_list) > COMPUTE_NFB_MAXLEN {
		nl = nfb_list[0:COMPUTE_NFB_MAXLEN]
	} else {
		nl = nfb_list
	}

	NFBArticle := query.QueryArticleForTags(nl)

	nfbscorehelper := func(art *Article) float64 {
		var ns float64 = 0.0
		for _, nart := range NFBArticle {
			for _, ntag := range nart.GetTags() {
				for _, tag := range art.GetTags() {
					if ntag == tag {
						ns += 1
					}
				}
			}
			//负反馈依次衰减0.7倍
			if art.GetSmallclass() == art.GetSmallclass() {
				ns += 1 * 0.7
			}
			if art.GetMediumclass() == nart.GetMediumclass() {
				ns += 1 * 0.7 * 0.7
			}
			//不对顶层品类负反馈
		}
		return ns
	}
	for _, art := range data {
		ns := nfbscorehelper(art)
		scoremap[art.GetId()] = ns
	}
	return

}
func hotclickscore(data []*model.Article) (scoremap map[string]float64) {
	scoremap = make(map[string]float64)
	//https://www.cnblogs.com/zhiji6/p/6509770.html
	r := storage.GetRedis(storage.ARTICLECLICK_KEY)
	keys := make([]string, 0)
	for _, art := range data {
		keys = append(keys, storage.ARTICLECLICK_KEY+":"+art.GetId())
		scoremap[art.GetId()] = 0
	}
	clicks, err := r.Mget(keys)
	if err != nil {
		return
	}
	nowsec := float64(time.Now().Unix())
	for i, art := range data {
		//(p-1)/(t+2)**1.5 1/(0+2)^1.5=0.3535
		p, _ := strconv.Atoi(clicks[i])
		t, err := strconv.ParseFloat(art.GetDatpublis(), 64)
		if err != nil {
			Errorln(err)
			continue
		}
		// convert t to publish hours till now
		t = (nowsec - t) / 3600
		if t < 0 {
			continue
		}
		s := float64(p-1) / math.Pow(t+2, 1.5)
		scoremap[art.GetId()] = s
	}
	return scoremap
}

func tagscore(usertag map[string]float64, art *Article) float64 {
	var hs float64 = 0.0
	for k, v := range usertag {
		if k == art.GetLargeclass() {
			hs += v
			continue
		}
		if k == art.GetMediumclass() {
			hs += v
			continue
		}
		if k == art.GetSmallclass() {
			hs += v
			continue
		}
		for _, x := range art.GetTags() {
			if k == x {
				hs += v
				continue
			}
		}
	}
	return hs
}

func rank(st Stgy, data []*Article, formula string, abtest string) (sortdata []*Article) {
	var history_usertag = st.History_usertag
	var realtime_usertag = st.Realtime_usertag

	//1. init scores "Jan 2, 2006 at 3:04pm (MST)"
	scorelist := make([]score, 0)
	hotclickMap := hotclickscore(data)
	nfbMap := nfbscore(st.Nfblist, data)
	for i, art := range data {
		var s score
		s.idx = i
		iscore := make(map[string]float64)
		//if t, err := time.Parse("2006-01-02 15:04:05", art.GetDatpublis()); err == nil { //1562904054
		if t, err := strconv.ParseFloat(art.GetDatpublis(), 64); err == nil {
			iscore[PUBLISH] = float64(t)
		} else {
			util.Errorln(err)
		}

		iscore[HISTORY] = tagscore(history_usertag, art)
		iscore[REALTIME] = tagscore(realtime_usertag, art)
		if hcl, ok := hotclickMap[art.GetId()]; ok {
			iscore[HOTCLICK] = hcl
		} else {
			iscore[HOTCLICK] = 0
		}
		if nf, ok := nfbMap[art.GetId()]; ok {
			iscore[NEGFBACk] = nf
		} else {
			iscore[NEGFBACk] = 0
		}

		s.scores = iscore
		Debugln(iscore)
		scorelist = append(scorelist, s)
	}

	//2. get max&min for normalized
	maxscore := make(map[string]float64)
	minscore := make(map[string]float64)
	for _, w := range []string{PUBLISH, HISTORY, REALTIME, HOTCLICK, NEGFBACk} {
		maxscore[w] = MINS
		minscore[w] = MAXS
	}
	for _, s := range scorelist {
		for _, w := range []string{PUBLISH, HISTORY, REALTIME, HOTCLICK, NEGFBACk} {
			if _, ok := s.scores[w]; ok {
				if s.scores[w] > maxscore[w] {
					maxscore[w] = s.scores[w]
				}
				if s.scores[w] < minscore[w] {
					minscore[w] = s.scores[w]
				}
			}
		}
	}
	Debugln(maxscore, minscore)
	//3. normalized and sort
	wei := weights[formula+"_"+strings.ToUpper(abtest)]
	for i, s := range scorelist {
		smap := make(map[string]float64)
		for _, w := range []string{PUBLISH, HISTORY, REALTIME, HOTCLICK, NEGFBACk} {
			if _, ok := s.scores[w]; ok && (maxscore[w]-minscore[w]-0) > 1e-6 {
				smap[w] = (s.scores[w] - minscore[w]) / (maxscore[w] - minscore[w])
			} else {
				smap[w] = 0
			}
		}
		for w, _ := range wei {
			if strings.Contains(w, "*") { //only support two feature multi
				feature := 1.0
				for _, nw := range strings.Split(w, "*") {
					feature *= smap[nw]
				}
				s.finalscore += wei[w] * feature
			} else {
				s.finalscore += wei[w] * smap[w]
			}
		}
		Debugln(abtest, smap, wei, s.finalscore)
		scorelist[i].scores = smap
		scorelist[i].finalscore = s.finalscore
	}

	sort.Slice(scorelist, func(i, j int) bool {
		//return i.Value > j.Value // 降序
		return scorelist[i].finalscore > scorelist[j].finalscore
	})

	Debugln(scorelist)
	sortdata = make([]*Article, len(data))
	for i, s := range scorelist {
		var art Article
		art = *data[s.idx]
		sortdata[i] = &art
		sortdata[i].Finalscore = s.finalscore
		sortdata[i].Computescore = s.scores
		sortdata[i].Servertag = formula
		Debugln(sortdata[i])
	}
	return sortdata
}

const (
	NEWLYFORMULA = "newly_formula"
	RECFORMULA   = "rec_formula"
)

func WeightsInit() {
	weights = make(map[string]map[string]float64)
	for _, m1 := range []string{NEWLYFORMULA, RECFORMULA} {
		for _, m2 := range []string{"A", "B", "C"} {
			m := m1 + "_" + m2
			formula := GetVal(RunMode+"_rank_args", m)
			if formula == "" {
				continue
			}
			weights[m] = make(map[string]float64)
			fs := strings.Split(formula, ",")
			for _, f := range fs {
				k := strings.Split(f, ":")[0]
				v := strings.Split(f, ":")[1]
				var val float64
				fmt.Sscanf(v, "%f", &val)
				weights[m][k] = val
			}
			Infoln(m, weights[m])
		}
	}
}

func init() {
	WeightsInit()
}

func rankSimilar(article Article, data []*Article, formula string) (sortdata []*Article) {
	small := article.GetSmallclass()
	medium := article.GetMediumclass()
	large := article.GetLargeclass()
	taglist := article.GetTags()
	tags := []string{small, medium, large}
	for _, v := range taglist {
		tags = append(tags, v)
	}

	tagsMap := make(map[string]float64)
	for _, v := range tags {
		tagsMap[v] = 1.0
	}

	//新建分数列表
	//1.
	scorelist := make([]score, 0)

	for i, art := range data {
		var s score
		s.idx = i
		iscore := make(map[string]float64)
		//字符串转换float
		if t, err := strconv.ParseFloat(art.GetDatpublis(), 64); err == nil {
			iscore[PUBLISH] = float64(t)
		} else {
			Errorln(err)
		}
		iscore[SIMILAR] = tagscore(tagsMap, art)
		s.scores = iscore
		//Debugln(iscore)
		scorelist = append(scorelist, s)
	}
	//2.
	maxscore := make(map[string]float64)
	minscore := make(map[string]float64)
	for _, w := range []string{SIMILAR, PUBLISH} {
		maxscore[w] = MAXS
		minscore[w] = MINS
	}
	for _, s := range scorelist {
		for _, w := range []string{SIMILAR, PUBLISH} {
			if _, ok := s.scores[w]; ok {
				if s.scores[w] > maxscore[w] {
					maxscore[w] = s.scores[w]
				}
				if s.scores[w] < minscore[w] {
					minscore[w] = s.scores[w]
				}
			}
		}
	}
	Debugln(maxscore, minscore)

	//3.
	wei := weights[formula+"_A"]
	for i, s := range scorelist {
		smap := make(map[string]float64)
		for _, w := range []string{SIMILAR, PUBLISH} {
			if _, ok := s.scores[w]; ok && (maxscore[w]-minscore[w]-0) > 1e-6 {
				smap[w] = (s.scores[w] - minscore[w]) / (maxscore[w] - minscore[w])
			} else {
				smap[w] = 0
			}
		}
		for w, _ := range wei {
			if strings.Contains(w, "*") { //only support two feature multi
				feature := 1.0
				for _, nw := range strings.Split(w, "*") {
					feature *= smap[nw]
				}
				s.finalscore += wei[w] * feature
			} else {
				s.finalscore += wei[w] * smap[w]
			}
		}
		//Debugln(smap, wei, s.finalscore)
		scorelist[i].scores = smap
		scorelist[i].finalscore = s.finalscore

	}
	sort.Slice(scorelist, func(i, j int) bool {
		//return i.Value > j.Value // 降序
		return scorelist[i].finalscore > scorelist[j].finalscore
	})

	Debugln(scorelist)
	sortdata = make([]*Article, len(data))
	for i, s := range scorelist {
		var art Article
		art = *data[s.idx]
		sortdata[i] = &art
		sortdata[i].Finalscore = s.finalscore
		sortdata[i].Computescore = s.scores
		sortdata[i].Servertag = formula
		Debugf("similar and sort list :%v", sortdata[i])
	}
	return sortdata
}
