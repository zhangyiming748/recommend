package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"reflect"
	"testing"
)

type Result struct {
	id     string `json:"id"`
	titile string `json:"titile"`
}

//搜索
func TestQuery(t *testing.T) {
	var err error
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.4.219:9200/"))
	if err != nil {
		panic(err)
	}

	var res *elastic.SearchResult
	//取所有
	res, err = client.Search("test_v1").Type("doc").Do(context.Background())
	printResult(res, err)

	//字段相等
	q := elastic.NewQueryStringQuery("source:新英")
	res, err = client.Search("test_v1").Type("doc").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}

	//printResult(res, err)

	if res.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits)

		for _, hit := range res.Hits.Hits {
			fmt.Println(hit)
		}
	} else {
		fmt.Printf("Found no Employee \n")
	}

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("test_v1").Type("doc").Query(q).Do(context.Background())
	printResult(res, err)

	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = client.Search("test_v1").Type("doc").Query(matchPhraseQuery).Do(context.Background())
	printResult(res, err)

	//分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests")
	res, err = client.Search("test_v1").Type("doc").Aggregation("all_interests", aggs).Do(context.Background())
	printResult(res, err)

}

//打印查询到的Employee
func printResult(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	_item := make(map[string]interface{})
	for _, item := range res.Each(reflect.TypeOf(_item)) { //从搜索结果中取数据的方法
		t := item
		fmt.Printf("%#v\n", t)
	}
}
