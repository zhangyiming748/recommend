package template

import (
	. "recommend/module"
	"recommend/storage"
	. "recommend/util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	//	"strconv"
)

func cacheGetTemplateByUrl(url string, cacheKey string) (val interface{}, err error) {
	resp, err := http.Get(url)
	if err != nil {
		Errorln(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		SendAlarm("GetTemplateByUrl: " + url + " : " + err.Error())
		Errorln(err)
		return
	}

	val = string(body)
	if err == nil && len(val.(string)) > 0 {
		storage.SetCacheContent(cacheKey, val, storage.DEFAULTEXPIRATION)
	}
	return
}

func cacheGetTemplateRetData(channel string) (retData map[string]interface{}, ok bool) {
	var template interface{}
	body, err := GetTemplate(channel)
	if err = json.Unmarshal([]byte(body.(string)), &template); err != nil {
		SendAlarm("UnmarshalTemplate: " + channel + " : " + err.Error())
		Errorln("parsing json file", err.Error())
	}

	ret, ok := template.(map[string]interface{})
	if ok {
		retData, ok = ret["retData"].(map[string]interface{})
		if ok {
			storage.SetCacheContent(channel+RETDATATEMPLATKEY, retData, storage.DEFAULTEXPIRATION)
		}
	}
	return
}

func cachetemplateFixedItems(channel string) (fm map[string]int, err error) {
	val, err := getTopTemplate()
	if err != nil { //若有错误，及早返回，以免污染cache
		return
	}
	fixedMap, err := parseTemplateItems(val, true)
	if err != nil {
		return
	}

	val, err = GetTemplate(channel)
	if err != nil {
		return
	}

	tmpmap, err := parseTemplateItems(val, false)
	if err != nil {
		return
	}

	for k, v := range tmpmap {
		fixedMap[k] = v
	}
	if err == nil && len(fixedMap) > 0 {
		storage.SetCacheContent(channel+FIXEDITEMKEY, fixedMap, storage.DEFAULTEXPIRATION)
	}
	return fixedMap, err
}

func foreverLoadTemplate() {
	loopfunc := func() {
		func() {
			defer PanicRecover(PanicPosition())
			cacheGetTemplateByUrl(durl, TEMPLATEKEY)
			cacheGetTemplateByUrl(topurl, TOPTEMPLATEKEY)
			cacheGetTemplateByUrl(hotvurl, HOTVTEMPLATEKEY)
			cacheGetTemplateRetData(HOTVIDEOCHANNEL)
			cacheGetTemplateRetData(HOMECHANNEL)
			cachetemplateFixedItems(HOTVIDEOCHANNEL)
			cachetemplateFixedItems(HOMECHANNEL)

			TemplateFixedItems(HOTVIDEOCHANNEL)
			TemplateFixedItems(HOMECHANNEL)

			Debugln("foreverLoadTemplate")
		}()
		time.Sleep(30 * time.Second) //must little than storage.Defaultexpiration
	}
	//for set cache
	for {
		loopfunc()
	}
}

func init() {
	go foreverLoadTemplate()
}
