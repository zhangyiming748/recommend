package main

import (
	"../storage"
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	var r storage.RedisInstance
	var p storage.Prefix
	slice_p := make([]storage.Prefix, 0)
	//init
	p.Prefix = "articlescore"
	p.IsCacheable = true
	p.CacheMin = 300
	slice_p = append(slice_p, p)
	r = storage.RedisInit(slice_p, "192.168.4.7:6379?password=888888&db=0", "true")
	table := make(map[string]float64)
	for i := 1; i <= 20; i++ {
		//title:= "articlescore"
		val := "articlescore" + strconv.Itoa(i)
		log.Printf("要查询的key是%v", val)
		value, _ := r.Get(val)
		table["articlescore"+strconv.Itoa(i)], _ = strconv.ParseFloat(value, 64)
		log.Printf("key对应的值是%v", value)
	}
	for k, v := range table {
		fmt.Printf("%v的分数是%v\n", k, v)
	}
	sortByValue(table)
}
func sortByValue(list map[string]float64) {
	allScore := make([]float64, 0)
	for _, v := range list {
		allScore = append(allScore, v)
		log.Printf("append之后的allScore%v", allScore)
	}
	//sort.Slice(allScore,func(x,y int) bool {
	//	return allScore[x] < allScore[y]
	//})
	sort.Float64s(allScore)
	log.Printf("排序之后的allScore%v", allScore)
	for _, vv := range allScore {
		for k, v := range list {
			if v == float64(vv) {
				fmt.Printf("%v is %v\n", k, vv)
				break
			}
		}
	}
	return
}
