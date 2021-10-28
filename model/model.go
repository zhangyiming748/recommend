package main

import (
	pb "../rpc/protoFile"
	"encoding/json"
	"fmt"
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
	userId    string `json:userId`
	uuid      string
	itemSize  string
	recAction string
	pageNum   string
}

func (p *Param) SetuserId(s string) {
	p.userId = s
}
func (p *Param) Setuuid(s string) {
	p.uuid = s
}
func (p *Param) SetitemSize(s string) {
	p.itemSize = s
}
func (p *Param) SetrecAction(s string) {
	p.recAction = s
}
func (p *Param) SetpageNum(s string) {
	p.pageNum = s
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
	size       int
	PersonalInfo
	Filtermap map[string]bool
	FixedList map[string]int
	FixedPos  map[int][]string
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

type ArticleList []Article
type Article struct {
	pb.Article
}

/*
func (s *Article) SetId(v string) {
	s.id = v
	return
}
func (s *Article) SetTags(v []string) {
	s.tags = v
	return
}
func (s *Article) SetLargeClass(v string) {
	s.largeclass = v
	return
}
func (s *Article) SetMediumClass(v string) {
	s.mediumclass = v
	return
}
func (s *Article) SetSmallClass(v string) {
	s.smallclass = v
	return
}
func (s *Article) SetDatPublish(v string) {
	s.datpublish = v
	return
}
func (s *Article) SetVc2Type(v string) {
	s.vc2type = v
	return
}
*/
