package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	HOMECHANNEL = "homepage"
)

type ArticleList []Article
type Article struct {
	// cp from pb
	Id                 string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tags               []string           `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	Largeclass         string             `protobuf:"bytes,3,opt,name=largeclass,proto3" json:"largeclass,omitempty"`
	Mediumclass        string             `protobuf:"bytes,4,opt,name=mediumclass,proto3" json:"mediumclass,omitempty"`
	Smallclass         string             `protobuf:"bytes,5,opt,name=smallclass,proto3" json:"smallclass,omitempty"`
	Datpublis          string             `protobuf:"bytes,6,opt,name=datpublis,proto3" json:"datpublis,omitempty"`
	Vc2Type            string             `protobuf:"bytes,7,opt,name=vc2type,proto3" json:"vc2type,omitempty"`
	Finalscore         float64            `protobuf:"fixed64,8,opt,name=finalscore,proto3" json:"finalscore,omitempty"`
	Computescore       map[string]float64 `protobuf:"bytes,9,rep,name=computescore,proto3" json:"computescore,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Servertag          string             `protobuf:"bytes,10,opt,name=servertag,proto3" json:"servertag,omitempty"`
	RecommendTimeLimit int64              `protobuf:"varint,11,opt,name=recommendTimeLimit,proto3" json:"recommendTimeLimit,omitempty"`
}

func NewArticle(fields []byte) (article Article, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("NewArticleError: ", err)
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

	rtl, ok := val["recommend_time_limit"].(float64)
	if ok {
		article.SetRecommendTimeLimit(int64(rtl))
	}

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

func (m *Article) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Article) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Article) GetLargeclass() string {
	if m != nil {
		return m.Largeclass
	}
	return ""
}

func (m *Article) GetMediumclass() string {
	if m != nil {
		return m.Mediumclass
	}
	return ""
}

func (m *Article) GetSmallclass() string {
	if m != nil {
		return m.Smallclass
	}
	return ""
}

func (m *Article) GetDatpublis() string {
	if m != nil {
		return m.Datpublis
	}
	return ""
}

func (m *Article) GetVc2Type() string {
	if m != nil {
		return m.Vc2Type
	}
	return ""
}

func (m *Article) GetFinalscore() float64 {
	if m != nil {
		return m.Finalscore
	}
	return 0
}

func (m *Article) GetComputescore() map[string]float64 {
	if m != nil {
		return m.Computescore
	}
	return nil
}

func (m *Article) GetServertag() string {
	if m != nil {
		return m.Servertag
	}
	return ""
}

func (m *Article) GetRecommendTimeLimit() int64 {
	if m != nil {
		return m.RecommendTimeLimit
	}
	return 0
}

type Stgy struct {
	userId     string
	deviceId   string
	abtestStgy string
	algoStgy   string
	action     string
	uniqueid   string
	channel    string
	size       int
	PersonalInfo
	Filtermap map[string]bool
	FixedList map[string]int
	FixedPos  map[int][]string
}

func (s Stgy) GetUniqueid() string {
	return s.uniqueid
}
func (s *Stgy) SetUniqueid(u string) {
	s.uniqueid = u
	return
}
func (s Stgy) GetUserId() string {
	return s.userId
}
func (s *Stgy) SetUserId(u string) {
	s.userId = u
	return
}
func (s Stgy) GetDeviceId() string {
	return s.deviceId
}
func (s *Stgy) SetDeviceId(u string) {
	s.deviceId = u
	return
}
func (s Stgy) GetAlgoStgy() string {
	return s.algoStgy
}
func (s *Stgy) SetAlgoStgy(st string) {
	s.algoStgy = st
	return
}
func (s Stgy) GetAbtestStgy() string {
	return s.abtestStgy
}
func (s *Stgy) SetAbtestStgy(st string) {
	s.abtestStgy = st
	return
}
func (s Stgy) GetAction() string {
	return s.action
}
func (s *Stgy) SetAction(st string) {
	s.action = st
	return
}
func (s Stgy) GetSize() int {
	return s.size
}
func (s *Stgy) SetSize(st int) {
	s.size = st
	return
}
func (s Stgy) GetChannel() string {
	return s.channel
}
func (s *Stgy) SetChannel(st string) {
	s.channel = st
	return
}

type PersonalInfo struct {
	Showlist         []string
	Clicklist        []string
	Quick_showlist   []string
	Nfblist          []string
	History_usertag  map[string]float64
	Realtime_usertag map[string]float64
}

func (p *PersonalInfo) SetHistoryUsertag(hu map[string]string) {
	p.History_usertag = make(map[string]float64)
	for k, v := range hu {
		var val float64
		fmt.Sscanf(v, "%f", &val)
		p.History_usertag[k] = val
	}
	return
}

type Param struct {
	userId    string `json:"userId"`
	uuid      string
	itemSize  string
	recAction string
	pageNum   string
	channel   string
}

func (p *Param) SetUserId(s string) {
	p.userId = s
}
func (p Param) GetUserId() string {
	return p.userId
}
func (p *Param) SetUuid(s string) {
	p.uuid = s
}
func (p Param) GetUuid() string {
	return p.uuid
}
func (p *Param) SetItemSize(s string) {
	p.itemSize = s
}
func (p *Param) SetRecAction(s string) {
	p.recAction = s
}
func (p Param) GetRecAction() string {
	return p.recAction
}
func (p *Param) SetPageNum(s string) {
	p.pageNum = s
}
func (p Param) GetPageNum() string {
	return p.pageNum
}
func (p *Param) SetChannel(s string) {
	p.channel = s
}
func (p Param) GetChannel() string {
	return p.channel
}
