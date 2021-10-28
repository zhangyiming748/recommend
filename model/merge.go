package main

import (
	"../storage"
	. "../util"
	"math/rand"
	"strings"
	"time"
)

func mergeList(stgy Stgy, recli []ArticleList, chooseProb []float64) (resList []Article) {
	//获取列表之后, 都会经过该函数, 该函数完成打散逻辑, 该函数返回后 将列表交给渲染函数
	resList = make([]Article, 0)
	resultmap := make(map[string]bool)
	nowidx := make([]int, len(recli))
	failedMap := make(map[int]bool)
	randomChoose := false
	initLoopCtrl := func() {
		nowidx = make([]int, len(recli)) //循环索引全部恢复为0
		failedMap = make(map[int]bool)   //选取失败map恢复为空
	}
	addItem2Res := func(nowlist ArticleList, nowi int) {
		resList = append(resList, nowlist[nowi])
		resultmap[nowlist[nowi].GetId()] = true
		randomChoose = false
		initLoopCtrl()
	}
	llloop := 0
	for {
		llloop += 1
		if llloop > 30000 { //最多允许执行三万次循环
			Errorln("So many loops when mergeList, may something error")
			break
		}

		if len(resList) >= stgy.size {
			break
		}
		//以给定概率从多个推荐列表中选择一个列表
		randf := rand.Float64()
		list_i := 0
		for {
			if randf > chooseProb[list_i] {
				list_i++
				continue
			} else {
				break
			}
		}
		if list_i > len(recli) {
			continue
		}
		nowlist := recli[list_i] //当前选择的列表
		nowi := nowidx[list_i]   //当前列表中待选择的内容的索引

		if nowi >= len(nowlist) {
			failedMap[list_i] = false
			if len(failedMap) == len(recli) {
				//已选完全部列表且无可选内容, 则随机选择一个内容, 不做任何打散要求
				randomChoose = true
				initLoopCtrl()
			}
		}
		Debugln(randomChoose, nowi)
		for ; nowi < len(nowlist); nowi++ { //迭代到第一个未被选择的内容
			itemId := nowlist[nowi].GetId()
			if _, ok := resultmap[itemId]; ok {
				continue
			} else {
				break
			}
		}
		if nowi >= len(nowlist) {
			continue
		}
		nowidx[list_i] = nowi + 1 //假设本次选择失败, 则需将该列表的下一次检查位置写回

		Debugln(randomChoose, nowi)

		if randomChoose { //各个列表中都无法选出合适的内容, 则无需进行任何检查, 直接选择内容
			addItem2Res(nowlist, nowi)
			continue //随机选择一个之后,恢复初始循环条件,继续下次循环
		}

		item := nowlist[nowi]

		//if 0.05 < rand.Float64() { //例如各种失败情况
		if checkValidArticle(item, resList, stgy) {
			nowidx[list_i] = nowi + 1 //需将该列表的下一次检查位置写回索引数组
			continue
		} else {
			addItem2Res(nowlist, nowi)
			initLoopCtrl()
		}
	}
	return resList
}

func checkValidArticle(ckItem Article, resList []Article, stgy Stgy) bool {
	fixedPos := stgy.FixedPos
	nowpos, prepos, nextpos := len(resList), len(resList)-1, len(resList)+1
	//7,11 运营位在7和11, 以及顶部有运营位
	var preItems []string
	var nextItems []string
	if stgy.GetAction() == ACTION_HOME {
		if _, ok := fixedPos[nowpos]; ok {
			//例如当前待填充位置为7, 而运营位的位置也为7, 则不需考虑位置6, 只需考虑运营位
			preItems = append(preItems, fixedPos[nowpos]...)
		} else if prepos >= 0 {
			//否则选取结果集中前一个内容
			preItems = append(preItems, resList[prepos].GetId())
		}
		if _, ok := fixedPos[nextpos]; ok {
			//例如当前待选位置为7, 而运营位的位置为8, 则需要考虑运营位内容
			nextItems = append(nextItems, fixedPos[nextpos]...)
		}
	} else if stgy.GetAction() == ACTION_UP && nowpos == 0 {
		if _, ok := fixedPos[0]; ok {
			preItems = append(preItems, fixedPos[0]...)
		}
	} else if prepos >= 0 {
		preItems = append(preItems, resList[prepos].GetId())
	}
	checktags := func(ids []string, ckItem Article) bool {
		status := false
		for _, i := range ids {
			a := getArticlehelper(i, resList)
			if a.GetLargeclass() != ckItem.GetLargeclass() {
				continue
			}
			if a.GetMediumclass() != ckItem.GetMediumclass() {
				continue
			}
			if a.GetSmallclass() != "" && strings.ToLower(a.GetSmallclass()) != "null" && a.GetSmallclass() == ckItem.GetSmallclass() {
				//大中小项全部相同, 则ckItem需被打散
				return true
			}
			if a.GetSmallclass() == "" && strings.ToLower(a.GetSmallclass()) == "null" && ckItem.GetSmallclass() == "" && strings.ToLower(ckItem.GetSmallclass()) == "null" {
				//大中项全部一致, 小项缺失, 此时校验tags, 若存在重复tag, 则ckItem需要被打散
				for _, x := range a.GetTags() {
					for _, y := range ckItem.GetTags() {
						if x == y {
							return true
						}
					}
				}
			}
		}
		return status
	}

	if checktags(preItems, ckItem) {
		return true
	}
	if checktags(nextItems, ckItem) {
		return true
	}

	if nowpos > 3 { //当长度大于3时, 考察比例
		classes := make(map[string]int)
		for _, i := range resList {
			a := getArticlehelper(i.GetId(), resList)
			if _, ok := classes[a.GetLargeclass()]; ok {
				classes[a.GetLargeclass()] = classes[a.GetLargeclass()] + 1
			} else {
				classes[a.GetLargeclass()] = 1
			}
		}
		if _, ok := classes[ckItem.GetLargeclass()]; ok && float64(classes[ckItem.GetLargeclass()])/float64(len(resList)) > CATEGORY_RATIO {
			//如果当前大项词在列表中所占比率已经超过0.8, 则打散
			return true
		}
	}
	//计算当前的图文比率, 若全是同一类内容, 且待选项目与该类比率相同, 则返回失败
	return false
}

func getArticlehelper(id string, resList ArticleList) (art Article) {
	// 内容可能在reslist中
	for _, art := range resList {
		if art.GetId() == id {
			return art
		}
	}
	val, found := storage.FetchCacheContent(ACTICLETAGSKEY + ":" + string(id))
	if !found {
		return //如果没找到则放弃
	}

	return val.(Article)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
