package main

import (
	"../rpc"
	"../storage"
	. "../util"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
	//	"strconv"
)

type Template struct {
	resCode       string        `json:"resCode"`
	requsetAction string        `json:"requsetAction"`
	resMessage    string        `json:"resMessage"`
	retData       RecommendInfo `json:"retData"`
}

const (
	durl              = "http://json.ssports.com/json/channel/appHomePage_bottom_1.json"
	topurl            = "http://json.ssports.com/json/channel/appHomePage_top.json"
	TEMPLATEKEY       = "template"
	TOPTEMPLATEKEY    = "toptemplate"
	RETDATATEMPLATKEY = "retdatatemplat"
	FIXEDITEMKEY      = "fixedItem"
	ACTICLETAGSKEY    = "articleTags"
)

func originalGetTemplateByUrl(url string, cacheKey string) (val interface{}, err error) {
	resp, err := http.Get(url)
	if err != nil {
		Errorln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Errorln(err)
	}

	val = string(body)
	if err == nil && len(val.(string)) > 0 {
		storage.SetCacheContent(cacheKey, val, storage.DEFAULTEXPIRATION)
	}
	return
}
func getTemplateByUrl(url string, cacheKey string) (val interface{}, err error) {
	val, found := storage.FetchCacheContent(cacheKey)
	if !found {
		return originalGetTemplateByUrl(url, cacheKey)
	}
	return val, err
}

func getTopTemplate() (val interface{}, err error) {
	return getTemplateByUrl(topurl, TOPTEMPLATEKEY)
}

func GetTemplate() (val interface{}, err error) {
	//获取模板内容
	return getTemplateByUrl(durl, TEMPLATEKEY)
}

func originalGetTemplateRetData() (retData map[string]interface{}, ok bool) {
	var template interface{}
	body, err := GetTemplate()
	if err = json.Unmarshal([]byte(body.(string)), &template); err != nil {
		Errorln("parsing json file", err.Error())
	}

	ret, ok := template.(map[string]interface{})
	if ok {
		retData, ok = ret["retData"].(map[string]interface{})
		if ok {
			storage.SetCacheContent(RETDATATEMPLATKEY, retData, storage.DEFAULTEXPIRATION)
		}
	}
	return
}
func GetTemplateRetData() (retData map[string]interface{}, ok bool) {
	val, ok := storage.FetchCacheContent(RETDATATEMPLATKEY)
	if !ok {
		return originalGetTemplateRetData()
	} else {
		retData = val.(map[string]interface{})
	}
	return
}

func TempArtilceList(list []string) (ret []map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			Errorln(err)
		}
	}()

	ret = make([]map[string]interface{}, len(list))
	r := storage.GetRedis(storage.ARTICLEDETAIL_KEY)
	var wg sync.WaitGroup
	for i, l := range list {
		wg.Add(1)
		go func(i int, l string) {
			defer wg.Done()
			detail, err := r.HgetAll(storage.ARTICLEDETAIL_KEY + ":" + l)
			//Debugln("debug....................",detail)
			if len(detail) == 0 {
				Errorln("cannot get articleDetail from redis: ", l, " err: ", err)
				return
			}
			//json,commentNumber||"comment_number":"2",
			jsonStr := detail["json"]

			var tt map[string]interface{}

			if err = json.Unmarshal([]byte(jsonStr), &tt); err != nil {
				Errorln("parsing json file", err.Error())
			}
			tt["comment_number"] = detail["commentNumber"]

			ret[i] = tt
		}(i, l)

	}
	wg.Wait()
	retnew := make([]map[string]interface{}, 0)
	for _, v := range ret {
		if v != nil {
			retnew = append(retnew, v)
		}
	}
	return retnew
}

func parseTemplateItems(val interface{}, usePos bool) (fixedMap map[string]int, err error) {
	var template interface{}
	if err = json.Unmarshal([]byte(val.(string)), &template); err != nil {
		Errorln("parsing json file", err.Error())
	}
	fixedMap = make(map[string]int)
	var recurFind func(interface{}, int, bool)
	recurFind = func(template interface{}, pos int, usePos bool) {
		switch template.(type) {
		case map[string]interface{}:
			xmap := template.(map[string]interface{})
			list_type, ok := xmap["list_type"]
			if ok {
				if list_type.(string) == "free" {
					usePos = true
				}

			}
			ll, ok := xmap["list"]
			if ok {
				//若还存在嵌套的list
				recurFind(ll, pos, usePos)
			} else {
				//已经不存在嵌套的list, 最底层的map
				if _, ok := xmap["jump_type"]; ok {
					jump_type := strings.ToLower(xmap["jump_type"].(string))
					if strings.Contains("avi", jump_type) {
						fid := xmap["jump_url"].(string)
						fixedMap[fid] = pos
						Debugln(pos, err, xmap["jump_url"].(string), " ", xmap["title"])
					}

				}
				//Debugln(xmap["jump_type"])
			}
		case []interface{}:
			for i, v := range template.([]interface{}) {
				if usePos {
					recurFind(v, pos, usePos)
				} else {
					recurFind(v, i, usePos)
				}
			}
		default:
			Debugln(template)
		}
	}

	var retData map[string]interface{}
	ret, ok := template.(map[string]interface{})
	if ok {
		retData, ok = ret["retData"].(map[string]interface{})
		if ok {
			recurFind(retData, 0, usePos)
		}

	}
	return
}

func originaltemplateFixedItems() (fm map[string]int, err error) {
	val, err := getTopTemplate()
	fixedMap, err := parseTemplateItems(val, true)
	val, err = GetTemplate()
	tmpmap, err := parseTemplateItems(val, false)

	for k, v := range tmpmap {
		fixedMap[k] = v
	}
	if err == nil && len(fixedMap) > 0 {
		storage.SetCacheContent(FIXEDITEMKEY, fixedMap, storage.DEFAULTEXPIRATION)
	}
	return fixedMap, err
}
func templateFixedItems() (fm map[string]int) {
	fMap, found := storage.FetchCacheContent(FIXEDITEMKEY)
	if !found {
		fixedMap, err := originaltemplateFixedItems()
		if err == nil && len(fixedMap) > 0 {
			ids := make([]string, 0)
			for k, _ := range fixedMap {
				ids = append(ids, k)
			}
			go getArticleByids(ids)
		}
		return fixedMap
	} else {
		return fMap.(map[string]int)
	}
}

func getArticleByids(ids []string) (artlist ArticleList) {
	idstr := ""
	Infoln("-------getArticleByids----ids---", ids)
	for _, i := range ids {
		val, found := storage.FetchCacheContent(ACTICLETAGSKEY + ":" + string(i))
		if !found {
			idstr += "," + i
		} else {
			artlist = append(artlist, val.(Article))
		}
	}
	Infoln("-------getArticleByids----idstr---", idstr)
	if len(idstr) == 0 {
		return
	}

	idstr = idstr[1:]
	param := make(map[string]string)
	param["ids"] = idstr
	ret, err := rpc.RankClient(rpc.RANK_GETINFOBYIDS, param)
	if err != nil {
		Errorln("rpc err", err)
	}
	Infoln("-------getArticleByids----ret---", ret)
	for _, v := range ret {
		art := Article{*v}
		artlist = append(artlist, art)
		storage.SetCacheContent(ACTICLETAGSKEY+":"+art.GetId(), art, storage.DEFAULTEXPIRATION)
		Debugln(v)
	}
	Infoln("-------getArticleByids----artlist---", artlist)
	return
}

func foreverLoadTemplate() {
	loopfunc := func() {
		defer func() {
			if err := recover(); err != nil {
				Errorln(err)
			}
		}()
		time.Sleep(30 * time.Second) //must little than storage.Defaultexpiration
		originalGetTemplateByUrl(durl, TEMPLATEKEY)
		originalGetTemplateByUrl(topurl, TOPTEMPLATEKEY)
		originalGetTemplateRetData()
		originaltemplateFixedItems()
		Debugln("foreverLoadTemplate")
	}
	//for set cache
	for {
		loopfunc()
	}
}

func init() {
	templateFixedItems()
	go foreverLoadTemplate()
}
