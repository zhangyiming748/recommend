package main

import (
	//"context"
	"github.com/olivere/elastic"
	"log"
	"time"
)

func qFromES() {
	t := time.Now()
	defer func() {
		if err := recover(); err != nil {
			//Errorln(err)
			log.Println(err)
		}
	}()
	q := elastic.NewBoolQuery()
	q1 := elastic.NewBoolQuery()
	q1.Must(elastic.NewTermQuery("age", 4))
	q.Should(q1)


	q = q.QueryName("test")

	d:=time.Since(t)
	log.Println(d)

}
