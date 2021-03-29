package query

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	. "recommend/model"
	. "recommend/util"
	"strings"
	"time"
)

func qNewArticlesFromES(size int) []*Article {
	t := time.Now()
	defer func() {
		if err := recover(); err != nil {
			Errorln(err)
		}
	}()
	var res *elastic.SearchResult
	q := elastic.NewBoolQuery()
	q1 := elastic.NewBoolQuery()
	q1.Must(elastic.NewTermQuery("numstatus", 4), elastic.NewTermsQuery("vc2type.keyword", "A", "I"))
	q.Should(q1)
	q2 := elastic.NewBoolQuery()
	q2.Must(elastic.NewTermQuery("numstatus", 4), elastic.NewTermQuery("vc2type.keyword", "V"), elastic.NewTermsQuery("contentsource", "xingying-ppc", "iqiyi-ppc"))
	q.Should(q2)

	q = q.QueryName("idsquery")

	tt := time.Now()
	fsc := elastic.NewFetchSourceContext(true).Include("id", "vc2type", "datpublist", "m_classify_name", "l_classify_name", "s_classify_name", "tagnames", "recommend_time_limit")
	// 线上版本需要按照发布时间排序, 测试该段代码时注意
	res, _ = esClient.Search(index_name).Query(q).FetchSourceContext(fsc).Size(size).Sort("datpublist", false).Do(context.Background())

	dd := time.Since(tt)
	Debugln("idsquery1 Since:", dd)
	data := make([]*Article, 0)
	for _, item := range res.Hits.Hits {
		fields, _ := json.Marshal(item.Source)
		article, _ := NewArticle(fields)
		data = append(data, &article)
	}
	d := time.Since(t)
	Debugln("idsquery2 Since:", d, len(data))
	return data
}

func qNewArticlesFromES4video(size int) []*Article {
	t := time.Now()
	defer func() {
		if err := recover(); err != nil {
			Errorln(err)
		}
	}()
	var res *elastic.SearchResult
	q := elastic.NewBoolQuery()

	q1 := elastic.NewBoolQuery()
	q1.Must(elastic.NewTermQuery("numstatus", 4), elastic.NewTermQuery("vc2type.keyword", "V"), elastic.NewTermsQuery("contentsource", "xingying-pgc"))
	q1.MustNot(elastic.NewBoolQuery().Should(elastic.NewTermQuery("ssportspayflag", 1), elastic.NewTermQuery("iqiyipayflag", 1)))
	q.Should(q1)
	q2 := elastic.NewBoolQuery()
	q2.Must(elastic.NewTermQuery("numstatus", 4), elastic.NewTermQuery("vc2type.keyword", "V"),
		elastic.NewTermsQuery("contentsource", "iqiyi-pgc", "zhiboba", "dongqiudi", "hupu", "lishipin"),
		elastic.NewExistsQuery("sportname"))
	q2.MustNot(elastic.NewBoolQuery().Should(elastic.NewTermQuery("ssportspayflag", 1), elastic.NewTermQuery("iqiyipayflag", 1)))
	q.Should(q2)
	q = q.QueryName("idsquery4video")

	tt := time.Now()
	fsc := elastic.NewFetchSourceContext(true).Include("id", "vc2type", "datpublist", "m_classify_name", "l_classify_name", "s_classify_name", "tagnames", "recommend_time_limit")
	// 线上版本需要按照发布时间排序, 测试该段代码时注意
	res, _ = esClient.Search(index_name).Query(q).FetchSourceContext(fsc).Size(size).Sort("datpublist", false).Do(context.Background())

	dd := time.Since(tt)
	Debugln("idsquery1 Since:", dd)
	data := make([]*Article, 0)
	for _, item := range res.Hits.Hits {
		fields, _ := json.Marshal(item.Source)
		article, _ := NewArticle(fields)
		data = append(data, &article)
	}
	d := time.Since(t)
	Debugln("idsquery2 Since:", d)
	return data
}

func getDataByIds(ids []string) []*Article {
	var res *elastic.SearchResult
	Debugf("ids:%s", ids)
	data := make([]*Article, 0)
	size := len(ids)
	if size == 0 {
		return data
	}
	iddss := make([]interface{}, 0)
	for _, v := range ids {
		iddss = append(iddss, strings.TrimSpace(v))
	}
	Infoln("wangrfxxx:", iddss)
	q := elastic.NewBoolQuery()
	q.Must(elastic.NewTermsQuery("id", iddss...))

	qq := elastic.NewBoolQuery()
	q1 := elastic.NewBoolQuery()
	q1.Must(elastic.NewTermQuery("numstatus", 4), elastic.NewTermsQuery("vc2type.keyword", "A", "I"))
	qq.Should(q1)
	q2 := elastic.NewBoolQuery()
	q2.Must(elastic.NewTermQuery("numstatus", 4), elastic.NewTermQuery("vc2type.keyword", "V"), elastic.NewTermsQuery("contentsource", "xingying-ppc", "iqiyi-ppc"))
	qq.Should(q2)
	q.Filter(qq)
	q = q.QueryName("getInfoByIds")

	//src, err := q.Source()
	//s, err := json.Marshal(src)
	//if err != nil {
	//	//log.Fatalf("marshaling to JSON failed: %v", err)
	//}
	//got := string(s)
	//Infoln(got)

	tt := time.Now()
	fsc := elastic.NewFetchSourceContext(true).Include("id", "vc2type", "datpublist", "m_classify_name", "l_classify_name", "s_classify_name", "tagnames", "recommend_time_limit")
	res, _ = esClient.Search(index_name).Query(q).FetchSourceContext(fsc).Size(size).Do(context.Background())
	dd := time.Since(tt)
	Debugln("getInfoByIds1 Since:", dd)
	for _, item := range res.Hits.Hits {
		Infoln(item.Id)
		fields, _ := json.Marshal(item.Source)
		article, _ := NewArticle(fields)
		data = append(data, &article)
	}
	return data
}

func getDataByIds4video(ids []string) []*Article {
	var res *elastic.SearchResult
	Debugf("ids:%s", ids)
	data := make([]*Article, 0)
	size := len(ids)
	if size == 0 {
		return data
	}
	iddss := make([]interface{}, 0)
	for _, v := range ids {
		iddss = append(iddss, strings.TrimSpace(v))
	}
	q := elastic.NewBoolQuery()
	q.Must(elastic.NewTermsQuery("id", iddss...))

	qq := elastic.NewBoolQuery()
	q1 := elastic.NewBoolQuery()
	q1.Must(elastic.NewTermQuery("numstatus", 4), elastic.NewTermQuery("vc2type.keyword", "V"), elastic.NewTermQuery("contentsource", "xingying-pgc"))
	qq.Should(q1)
	q2 := elastic.NewBoolQuery()
	q2.Must(elastic.NewTermQuery("numstatus", 4), elastic.NewTermQuery("vc2type.keyword", "V"),
		elastic.NewTermsQuery("contentsource", "iqiyi-pgc", "zhiboba", "dongqiudi", "hupu", "lishipin"),
		elastic.NewExistsQuery("sportname"))
	qq.Should(q2)
	q.Filter(qq)
	q = q.QueryName("getDataByIds4video")

	//src, err := q.Source()
	//s, err := json.Marshal(src)
	//if err != nil {
	//	//log.Fatalf("marshaling to JSON failed: %v", err)
	//}
	//got := string(s)
	//Infoln(got)

	tt := time.Now()
	fsc := elastic.NewFetchSourceContext(true).Include("id", "vc2type", "datpublist", "m_classify_name", "l_classify_name", "s_classify_name", "tagnames", "recommend_time_limit")
	res, _ = esClient.Search(index_name).Query(q).FetchSourceContext(fsc).Size(size).Do(context.Background())
	dd := time.Since(tt)
	Debugln("getInfoByIds1 Since:", dd)
	for _, item := range res.Hits.Hits {
		//Infoln(item.Id)
		fields, _ := json.Marshal(item.Source)
		article, _ := NewArticle(fields)
		data = append(data, &article)
	}
	return data
}

/**
 * 查询elstaicsearch
 */
func test(ids_str string) (err error) {
	t := time.Now()
	var res *elastic.SearchResult
	//ids_str := in.Param["ids"]
	ids := strings.Split(ids_str, ",")
	Debugf("ids:%s", ids)
	q := elastic.NewBoolQuery()
	q.Must(NewIdsQuery().Ids(ids...))
	q.Filter(elastic.NewTermQuery("numstatus", 4), elastic.NewTermsQuery("vc2type.keyword", "A", "V", "I"))
	//q.MinimumNumberShouldMatch(1)
	q = q.QueryName("getInfoByIds")
	//src, err := q.Source()
	//s, err := json.Marshal(src)
	//if err != nil {
	//	log.Fatalf("marshaling to JSON failed: %v", err)
	//}
	//got := string(s)
	//Debugln(got)

	tt := time.Now()
	fsc := elastic.NewFetchSourceContext(true).Include("id", "vc2type", "datpublist", "m_classify_name", "l_classify_name", "s_classify_name", "tagnames", "recommend_time_limit")
	res, _ = esClient.Search(index_name).Query(q).FetchSourceContext(fsc).Do(context.Background())

	esClient.Scroll(index_name).Slice(q)
	dd := time.Since(tt)
	Debugln("getInfoByIds1 Since:", dd)
	data := make([]*Article, 0)
	for _, item := range res.Hits.Hits {
		//Infoln(item.Id)
		fields, _ := json.Marshal(item.Source)
		article, _ := NewArticle(fields)
		data = append(data, &article)
	}
	/*
		reply = &pb.RankReply{
			Code:    pb.RankReply_OK,
			Message: "success",
			Data:    data,
		}
	*/
	d := time.Since(t)
	Debugln("getInfoByIds2 Since:", d)
	return
}
