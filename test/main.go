package main

import (
	"math/rand"
	"strconv"
	"time"
	"fmt"
)


const (
	PERSONALALGO string = "personal"
	RANDOMALGO string = "random"

	ACTION_UP string = "up"
	ACTION_DOWN string = "down"
	ACTION_HOME string = "home"

	RECOMMEND_SIZE_UP = 10
	RECOMMEND_SIZE_DOWN = 10
	RECOMMEND_SIZE_HOME = 20
)

type RecommendList []Article

/*
公共参数 uid也是公共参数, abteststgy, algostgy 公共参数
传入的是过滤的filterMap[articleid] 绝对不出的, 这可能为空
1. new 发布时间两天内, 按发布时间倒序排列, 取50条, 如果不够, 有多少给多少
2. rec 这块需要和用户画像结合, rankserver 发布时间大于2天, 2天前的, 按时间倒序排序, top50 后续这块就top500
3. usercf 取50个相近用户, 取50条, 如果不够, 有多少给多少

返回 RecommendList
3.
*/
type Article struct {
	id string
	tags []string
	largeclass string
	mediumclass string
	smallclass string
	datpublish string
	vc2type string
}
func prepare() []RecommendList {
	var rlist []RecommendList
	for j:=0; j<3; j++ {
		var tmpr RecommendList
		for i:=0; i<50; i++ {
			var a Article
			a.id = strconv.Itoa(rand.Intn(50000))
			tmpr = append(tmpr,a)
		}
		rlist = append(rlist,tmpr)
	}
	return rlist
}
/*
type Stgy struct {

}


type mergeList struct {
	articleList []string
	articleMap  map[string]bool
	size int
	articleTypeNum []int
	classNum []int
}
*/

func merge(recli []RecommendList, chooseProb []float64) int64{
	//获取列表之后, 都会经过该函数, 该函数完成打散逻辑, 该函数返回后 将列表交给渲染函数
	//chooseProb [0.6,0.8,1.0]
	resultmap := make(map[string]bool)
	resList := make([]string,0)
	nowidx := make([]int,len(recli))
	failedMap := make(map[int]bool)
	initLoopCtrl := func() {
		nowidx = make([]int,len(recli))  //循环索引全部恢复为0
		failedMap = make(map[int]bool)   //选取失败map恢复为空
	}
	addItem2Res := func(id string) {
		resList = append(resList, id)
		resultmap[id] = true
		initLoopCtrl()
	}
	randomChoose := false
	var xxx int64 = 0
	for {
		xxx++
		if len(resList) >= 20 {
			break
		}

		//以给定概率从多个推荐列表中选择一个列表
		randf := rand.Float64()
		list_i := 0
		for  {
			if randf > chooseProb[list_i] {
				list_i++
				continue
			}else{
				break
			}
		}
		if list_i > len(recli) {
			continue
		}
		nowlist := recli[list_i]  //当前选择的列表
		nowi := nowidx[list_i]    //当前列表中待选择的内容的索引

		if nowi >= len(nowlist) {
			failedMap[list_i] = false
			if len(failedMap) == len(recli) {
				//已选完全部列表且无可选内容, 则随机选择一个内容, 不做任何打散要求
				randomChoose = true
				initLoopCtrl()
			}
		}


		for ;nowi < len(nowlist);nowi++ {
			itemId := nowlist[nowi].id
			if _, ok := resultmap[itemId]; ok {
				continue
			}else{
				break
			}
		}
		nowidx[list_i] = nowi+1 //假设本次选择失败, 则需将该列表的下一次检查位置写回
		if nowi >= len(nowlist) {
			continue
		}

		if randomChoose {  //各个列表中都无法选出合适的内容, 则无需进行任何检查, 直接选择内容
			addItem2Res(nowlist[nowi].id)
			randomChoose = false
			continue
		}

		//item := nowlist[nowi]

		if 0.03 < rand.Float64() { //例如各种失败情况
			nowidx[list_i] = nowi+1  //需将该列表的下一次检查位置写回索引数组
			continue
		}else{
			addItem2Res(nowlist[nowi].id)
			randomChoose = false
		}

	}
	if len(resList) != 20 {
		fmt.Println("ERROR!", len(resList))
	}
	test := make(map[string]bool)
	for _,i := range resList {
		test[i] = true
	}
	if len(test)!=20 {
		fmt.Println("ERROR no 20 items", len(test))
	}
	return xxx
}

func main() {
	var chooseProb = []float64{0.6,0.8,1.0}
	fmt.Println(time.Now())
	var minxxx,maxxxx int64 = 99999999999, 0
	for i:=0 ; i<10000000; i++ {
		recli := prepare()
		xxx := merge(recli,chooseProb)
		if minxxx > xxx {
			minxxx = xxx
		}
		if maxxxx < xxx {
			maxxxx = xxx
		}
	}
	fmt.Println(time.Now())
	fmt.Println(minxxx,maxxxx)
	//在选定概率为0.03的前提下, 最小需要循环173次, 最大循环2387次, 平均需要0.1ms, 可估计最大也为2ms
	//假设该函数是CPU密集型, 那么需要占满一个核1ms时间.
	return
}

func checkValidArticle(ckItem Article, resultlist []string, st Stgy) bool{
	pos := len(resultlist)
	var preItems []string
	var nextItems []string
	if stgy.GetAction() == ACTION_HOME and nowpos+1 in fixedpos {
		//若为home页推荐, 则存在运营位
		nexttags = append(nexttags, "123")
	}else if stgy.GetAction() == ACTION_UP and nowpos == 0 {
		pretags = append(pretags,"123")
	}else{
		pretags = append(pretags,"123")
	}

		if len(resList) >= 20 {
			break
		}
		//over-----------------------------------------------------------------

		//备选的article在列表中的序号
		i := idx[li_idx]%len(rli)
		item := rli[i]

		//检查该article是否已经被选择过
		j := 0
		for _, articleid := range rli {
			if _, ok := merge_map[articleid]; ok {
				i = (i+1)%len(rli)
				j += 1
			}else{
				break
			}
		}
		if j == len(ril) { //本列表已被全部选完
			chooseProb[li_idx] = 0
			continue
		}


		//检查这个article是否有资格选(打散)
		nowtags := make([]string)
		nowpos := 0
		pretags := make([]string)
		nexttags := make([]string)
		if stgy.GetAction() == ACTION_HOME and nowpos+1 in fixedpos {
			//若为home页推荐, 则存在运营位
			nexttags = append(nexttags, "123")
		}else if stgy.GetAction() == ACTION_UP and nowpos == 0 {
			pretags = append(pretags,"123")
		}else{
			pretags = append(pretags,"123")
		}
		if checkValidArticle(nowtags,pretags,nexttags) {
			//如果当前内容的大中小项目 和 前面的全部重合, 则continue掉
			//如果当前两个项目的大中相同, 小项目缺失, 则查看是否有重合的tag, 若有, 则continue掉
			continue
		}

		//若当前已选择大于3篇article, 则考虑探索与利用
		if len(merge_list) > 3 {
			// 若当前的大项目 不符合用户画像大项目词, 或用户画像的大项词中排序为最后的词(大项目词多于2个)
			//则直接选择该项目
			//mergelist append
			//mergemap true
			//idx[li_idx] += 1
		}

		//若当前已选择大于3篇article, 则考虑大项比率和图文比率
		if len(merge_list) > 3 {
			//若大想比率过高, 或图文比率过高
			continue
		}

		initProb := 0.7
		//检查item是否出现在别的两个list中, 如果出现, 则增大选择概率1.2倍
		randf := rand.Float64()
		if randf < initProb {
			//选择将该项加入到列表中
		}

		if len(merge_list) > req_size {
			break
		}
		//if recommendList 中的article已经全部被选完, 也break

		//按照概率选择是否将该文章加入到merge_list中
		//如果mergelist已经满了,则break
		item_id := newly_list[newly_idx]
		fmt.Println(item_id)
		break

}
