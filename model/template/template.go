package template

import (
	. "recommend/model"
	"recommend/storage"
	. "recommend/util"
)

type Template struct {
	resCode       string        `json:"resCode"`
	requsetAction string        `json:"requsetAction"`
	resMessage    string        `json:"resMessage"`
	retData       RecommendInfo `json:"retData"`
}

const (
	hotvurl           = "http://json.ssports.com/json/channel/appHotlList_1.json"
	durl              = "http://json.ssports.com/json/channel/appHomePage_bottom_1.json"
	topurl            = "http://json.ssports.com/json/channel/appHomePage_top.json"
	HOTVTEMPLATEKEY   = "hotvtemplate"
	TEMPLATEKEY       = "homedtemplate"
	TOPTEMPLATEKEY    = "hometoptemplate"
	RETDATATEMPLATKEY = "retdatatemplat"
	FIXEDITEMKEY      = "fixedItem"
)

func getTemplateByUrl(url string, cacheKey string) (val interface{}, err error) {
	//获取缓存内容
	val, found := storage.FetchCacheContent(cacheKey)
	//如果没找到
	if !found {
		//从url请求返回的body中获取并且加入缓存中
		return cacheGetTemplateByUrl(url, cacheKey)
	}
	return val, err
}

//获取顶端模版内容

func getTopTemplate() (val interface{}, err error) {
	//only homepage recommend has toptemplate
	return getTemplateByUrl(topurl, TOPTEMPLATEKEY)
}

func GetTemplate(channel string) (val interface{}, err error) {
	//获取模板内容
	if channel == HOMECHANNEL {
		return getTemplateByUrl(durl, TEMPLATEKEY)
	} else if channel == HOTVIDEOCHANNEL {
		return getTemplateByUrl(hotvurl, HOTVTEMPLATEKEY)
	} else {
		Errorln("Get template error, invalid channel:", channel)
	}
	return
}

//获取频道模版
func GetTemplateRetData(channel string) (retData map[string]interface{}, ok bool) {
	val, ok := storage.FetchCacheContent(channel + RETDATATEMPLATKEY)
	if !ok {
		return cacheGetTemplateRetData(channel)
	} else {
		retData = val.(map[string]interface{})
	}
	return
}

//模版固定项目
func TemplateFixedItems(channel string) (fm map[string]int) {
	fMap, found := storage.FetchCacheContent(channel + FIXEDITEMKEY)
	//如果没在缓存中找到
	if !found {
		//缓存模版
		fixedMap, err := cachetemplateFixedItems(channel)
		if err == nil && len(fixedMap) > 0 {
			ids := make([]string, 0)
			for k, _ := range fixedMap {
				ids = append(ids, k)
			}
		}
		return fixedMap
	} else {
		return fMap.(map[string]int)
	}
}
