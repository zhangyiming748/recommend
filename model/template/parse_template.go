package template

import (
	"encoding/json"
	. "recommend/util"
	"strings"
	//	"strconv"
)

//usePos  自由运营位 freePosition
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
