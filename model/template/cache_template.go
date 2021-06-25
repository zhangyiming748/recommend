package template

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	. "recommend/model"
	"recommend/storage"
	. "recommend/util"
	"time"
	//	"strconv"
)

//请求url获取内容添加到缓存
func cacheGetTemplateByUrl(url string, cacheKey string) (val interface{}, err error) {
	//Get请求URL
	resp, err := http.Get(url)
	if err != nil {
		Errorln(err)
	}

	defer resp.Body.Close()
	//读取Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Errorln(err)
	}

	val = string(body)
	//如果body不为空
	if err == nil && len(val.(string)) > 0 {
		//设置缓存内容，三分钟生存期
		storage.SetCacheContent(cacheKey, val, storage.DEFAULTEXPIRATION)
	}
	return
}

//请求url获取返回信息设置到缓存
func cacheGetTemplateRetData(channel string) (retData map[string]interface{}, ok bool) {
	var template interface{}
	//获取指定频道的缓存内容
	body, err := GetTemplate(channel)
	if err = json.Unmarshal([]byte(body.(string)), &template); err != nil {
		Errorln("parsing json file", err.Error())
	}
	//设置retdata的内容到指定频道的缓存内容中
	ret, ok := template.(map[string]interface{})
	if ok {
		retData, ok = ret["retData"].(map[string]interface{})
		if ok {
			storage.SetCacheContent(channel+RETDATATEMPLATKEY, retData, storage.DEFAULTEXPIRATION)
		}
	}
	return
}

//获取retdata设置到缓存
func cachetemplateFixedItems(channel string) (fm map[string]int, err error) {
	//获取顶端模版项目 var=值和位置
	val, err := getTopTemplate()
	//解析模版返回过滤列表
	fixedMap, err := parseTemplateItems(val, true)
	//获取模版内容
	val, err = GetTemplate(channel)
	//解析模版
	tmpmap, err := parseTemplateItems(val, false)

	for k, v := range tmpmap {
		fixedMap[k] = v
	}
	//如果没有错误||fixedmap有元素
	if err == nil && len(fixedMap) > 0 {
		//设置缓存，内容，生存期
		storage.SetCacheContent(channel+FIXEDITEMKEY, fixedMap, storage.DEFAULTEXPIRATION)
	}
	//返回
	return fixedMap, err
}

func foreverLoadTemplate() {
	loopfunc := func() {
		func() {
			//获取并缓存
			defer PanicRecover()
			//缓存模版固定内容
			cacheGetTemplateByUrl(durl, TEMPLATEKEY)
			//缓存模版固定内容
			cacheGetTemplateByUrl(topurl, TOPTEMPLATEKEY)
			//缓存模版固定内容
			cacheGetTemplateByUrl(hotvurl, HOTVTEMPLATEKEY)
			//缓存模版固定内容
			cacheGetTemplateRetData(HOMECHANNEL)
			//缓存模版固定内容
			cachetemplateFixedItems(HOTVIDEOCHANNEL)
			//缓存模版固定内容
			cachetemplateFixedItems(HOMECHANNEL)
			//模板固定项目从HOTVIDEO出
			TemplateFixedItems(HOTVIDEOCHANNEL)
			//模板固定项目从HOME出
			TemplateFixedItems(HOMECHANNEL)

			Debugln("foreverLoadTemplate")
		}()
		//等待30sec
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
