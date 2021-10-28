package protoFile

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// 转换数据
func NewArticle(fields []byte) (article Article, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(err)
			err = e.(error)
		}
	}()

	var val map[string]interface{}
	if err = json.Unmarshal(fields, &val); err != nil {
		return
	}
	article.SetId(strconv.Itoa(int(val["id"].(float64))))
	switch val["datpublist"].(type) {
	case string:
		article.SetDatPublis(defaultfunc(val["datpublist"]))
	case float64:
		article.SetDatPublis(strconv.FormatInt(int64(val["datpublist"].(float64)), 10))
	}
	article.SetRecommendTimeLimit(int64(val["recommend_time_limit"].(float64)))
	article.SetVc2Type(defaultfunc(val["vc2type"]))
	article.SetLargeClass(defaultfunc(val["l_classify_name"]))
	article.SetMediumClass(defaultfunc(val["m_classify_name"]))
	article.SetSmallClass(defaultfunc(val["s_classify_name"]))
	article.SetTags(strings.Split(defaultfunc(val["tagnames"]), ","))

	return
}

func defaultfunc(v interface{}) string {
	if v == nil {
		return ""
	}
	if strings.ToLower(v.(string)) == "null" {
		return ""
	}
	return v.(string)
}

func (s *Article) SetId(v string) {
	s.Id = v
	return
}
func (s *Article) SetTags(v []string) {
	s.Tags = v
	return
}
func (s *Article) SetLargeClass(v string) {
	s.Largeclass = v
	return
}
func (s *Article) SetMediumClass(v string) {
	s.Mediumclass = v
	return
}
func (s *Article) SetSmallClass(v string) {
	s.Smallclass = v
	return
}
func (s *Article) SetDatPublis(v string) {
	s.Datpublis = v
	return
}
func (s *Article) SetRecommendTimeLimit(v int64) {
	s.RecommendTimeLimit = v
	return
}
func (s *Article) SetVc2Type(v string) {
	s.Vc2Type = v
	return
}
