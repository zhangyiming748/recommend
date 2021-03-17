package module

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Recommend struct {
	datcreate             string `json:"datcreate"`
	datpublist            string `json:"datpublist"`
	datpublistime         string `json:"datpublistime"`
	expertnewstime        string `json:"expertnewstime"`
	indexrightmessagetime string `json:"indexrightmessagetime"`
	is_have_logo          string `json:"is_have_logo"`
	logo_position         string `json:"logo_position"`
	logo_url              string `json:"logo_url"`
	numarticleid          string `json:"numarticleid"`
	numclick              string `json:"numclick"`
	numkeywordid          string `json:"numkeywordid"`
	numsort               string `json:"numsort"`
	vc2brieftitle         string `json:"vc2brieftitle"`
	vc2keyword            string `json:"vc2keyword"`
	vc2keywordname        string `json:"vc2keywordname"`
	vc2summary            string `json:"vc2summary"`
	vc2thumbpicurl        string `json:"vc2thumbpicurl"`
	vc2timelen            string `json:"vc2timelen"`
	vc2title              string `json:"vc2title"`
	vc2type               string `json:"vc2type"`
	vc2video              string `json:"vc2video"`
	vc2videourltelecom    string `json:"vc2videourltelecom"`
	vc2videourlunited     string `json:"vc2videourlunited"`
	videotime             string `json:"videotime"`
	comment_number        string `json:"comment_number"`
	iqiyibackpicurl       string `json:"iqiyibackpicurl"`
	iqiyi_pic             string `json:"iqiyi_pic"`
	iqiyibackvideoh5url   string `json:"iqiyibackvideoh5url"`
	iqiyibackvideopcurl   string `json:"iqiyibackvideopcurl"`
	iqiyipayflag          string `json:"iqiyipayflag"`
	is1080pay             string `json:"is1080pay"`
	iscash                string `json:"iscash"`
	qipuid                string `json:"qipuid"`
	ssports_pic           string `json:"ssports_pic"`
	ssportspayflag        string `json:"ssportspayflag"`
	tag_bg_iqiyi          string `json:"tag_bg_iqiyi"`
	tag_bg_ssports        string `json:"tag_bg_ssports"`
	tag_iqiyi             string `json:"tag_iqiyi"`
	tag_ssports           string `json:"tag_ssports"`
	uploadiqiyiflag       string `json:"uploadiqiyiflag"`
	vc2displaymode        string `json:"vc2displaymode"`
	publish_time          string `json:"publish_time"`
	create_time           string `json:"create_time"`
	vc2thumbpicurl_icon   string `json:"vc2thumbpicurl_icon"`
	vc2thumbpicurl2       string `json:"vc2thumbpicurl2"`
	vc2thumbpicurl3       string `json:"vc2thumbpicurl3"`
	vc2topicpicurl        string `json:"vc2topicpicurl"`
	vc2picurl             string `json:"vc2picurl"`
	vc2clickgourl         string `json:"vc2clickgourl"`
	inum                  string `json:"inum"`
	album_id              string `json:"album_id"`
	vc2source             string `json:"vc2source"`
	new_version_action    string `json:"new_version_action"`
	new_vc2type           string `json:"new_vc2type"`
	new_version_type      string `json:"new_version_type"`
	isSingleRight         string `json:"isSingleRight"`
	isPublicRight         string `json:"isPublicRight"`
	isSpecificRight       string `json:"isSpecificRight"`
	display_model         string `json:"display_model"`
	name                  string `json:"name"`
	list_type             string `json:"list_type"`
}

type RecommendInfo struct {
	strategy  string        `json:"strategy"`
	list      []interface{} `json:"list"`
	size      int           `json:"size"`
	channelId string        `json:"channelId"`
	action    string        `json:"action"`
}

func (r *RecommendInfo) SetStrategy(s string) {
	r.strategy = s
}
func (r RecommendInfo) GetStrategy() string {
	return r.strategy
}
func (r *RecommendInfo) SetChannelId(s string) {
	r.channelId = s
}
func (r RecommendInfo) GetChannelId() string {
	return r.channelId
}
func (r *RecommendInfo) SetAction(s string) {
	r.action = s
}
func (r RecommendInfo) GetAction() string {
	return r.action
}
func (r *RecommendInfo) SetList(s []interface{}) {
	r.list = s
}
func (r RecommendInfo) GetList() []interface{} {
	return r.list
}
func (r *RecommendInfo) SetSize(s int) {
	r.size = s
}
func (r RecommendInfo) GetSize() int {
	return r.size
}
func (this RecommendInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"strategy":  this.strategy,
		"list":      this.list,
		"size":      this.size,
		"channelId": this.channelId,
		"action":    this.action,
	})
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
