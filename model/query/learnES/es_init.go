package main

import (
	"fmt"
	"github.com/olivere/elastic"
)

var index_name string
var esClient *elastic.Client

/**
  elastic client
*/
func EsClient() (client *elastic.Client, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			//Errorln("EsClient func err:", err)
		}
	}()

	//es_url := GetVal(RunMode+"_es_args", "es_url")
	urls := "http://127.0.0.1:9200"
	//urls := strings.Split(es_url, ",")
	//Infof("es_url:%v", urls)
	client, err = elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(urls),
	)
	return
}

func initIndex() {
	//index_name = GetVal(RunMode+"_es_args", "index_name")
	index_name = "articles"
}

func init() {
	var err error
	esClient, err = EsClient()
	if err != nil {
		//Errorln("get EsClient error:", err)
	}
	initIndex()

	CacheQueryinit()
}
