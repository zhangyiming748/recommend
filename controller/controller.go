package controller

import (
	"../api"
	"../model"
	. "../util"
	"net/http"
)

/**
  推荐首页
*/
func RecommendAppHomePage(r *http.Request, w http.ResponseWriter) (res api.AppResponse) {
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = ""
			Errorln(err)
		}
	}()

	var param model.Param
	userId := r.FormValue("userId")
	param.SetuserId(userId)
	uuid := r.FormValue("uuid")
	param.Setuuid(uuid)
	recAction := r.FormValue("action")
	param.SetrecAction(recAction)

	if false {
		//校验token值  md5(userId+uuid+action+t+"ssports")
		t := r.FormValue("t")
		token := r.FormValue("token")
		_token := String2MD5(userId + uuid + recAction + t + "ssports")

		if token == "" || token != _token {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = api.PARAM_ILLEGAL
			return
		}
	}

	res = model.AppHomePage(&param)

	return
}

func GetClassify(r *http.Request, w http.ResponseWriter) (res api.AppResponse) {
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = ""
			Errorln(err)
		}
	}()
	aritcleId := r.FormValue("articleId")
	uid := r.FormValue("uid")
	uuid := r.FormValue("uuid")
	res = model.NegativeFb(aritcleId)
	//uid 唯一id 内容id
	DataLogf("%v\t%v\t%v\t%v\n", "NegativeFb", uid, uuid, aritcleId)
	return
}
