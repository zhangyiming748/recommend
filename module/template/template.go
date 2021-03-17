package template

import (
	. "../../model"
	"../../storage"
	. "../../util"
	"fmt"
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
	val, found := storage.FetchCacheContent(cacheKey)
	if !found {
		return cacheGetTemplateByUrl(url, cacheKey)
	}
	return val, nil
}

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
		err = fmt.Errorf("Get template error, invalid channel: " + channel)
	}
	return
}

func GetTemplateRetData(channel string) (retData map[string]interface{}, ok bool) {
	val, ok := storage.FetchCacheContent(channel + RETDATATEMPLATKEY)
	if !ok {
		return cacheGetTemplateRetData(channel)
	} else {
		retData = val.(map[string]interface{})
	}
	return
}

func TemplateFixedItems(channel string) (fm map[string]int) {
	fMap, found := storage.FetchCacheContent(channel + FIXEDITEMKEY)
	if !found {
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
