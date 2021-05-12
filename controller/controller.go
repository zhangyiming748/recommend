package controller

import (
	_ "google.golang.org/genproto/googleapis/bigtable/admin/v2"
	"net/http"
	"recommend/api"
	"recommend/model"
	"recommend/model/feedback"
	"recommend/model/history"
	"recommend/model/recommend"
	"recommend/model/related"
	. "recommend/util"
)

/**
  推荐首页
*/
func RecommendAppHomePage(r *http.Request, w http.ResponseWriter) (res api.AppResponse) {
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = api.REQUEST_ERR
			Errorln(err)
		}
	}()

	var param model.Param
	userId := r.FormValue("userId")
	param.SetUserId(userId)
	uuid := r.FormValue("uuid")
	param.SetUuid(uuid)
	recAction := r.FormValue("action")
	param.SetRecAction(recAction)
	param.SetChannel(model.HOMECHANNEL)

	//if false {
	//	//校验token值  md5(userId+uuid+action+t+"ssports")
	//	t := r.FormValue("t")
	//	token := r.FormValue("token")
	//	_token := String2MD5(userId + uuid + recAction + t + "ssports")
	//
	//	if token == "" || token != _token {
	//		res.ResCode = api.APPRESPONSE_CODE_SUCCESS
	//		res.ResStatus = api.PARAM_ILLEGAL
	//		return
	//	}
	//}

	res = recommend.RecommendEntry(&param)

	return
}

func RecommendHotVideoPage(r *http.Request, w http.ResponseWriter) (res api.AppResponse) {
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = api.REQUEST_ERR
			Errorln(err)
		}
	}()

	var param model.Param
	userId := r.FormValue("userId")
	param.SetUserId(userId)
	uuid := r.FormValue("uuid")
	param.SetUuid(uuid)
	recAction := r.FormValue("action")
	param.SetRecAction(recAction)
	param.SetChannel(model.HOTVIDEOCHANNEL)

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

	res = recommend.RecommendEntry(&param)

	return
}

func GetClickHistory(r *http.Request, w http.ResponseWriter) (res api.AppResponse) {
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = api.REQUEST_ERR
			Errorln(err)
		}
	}()
	var param model.Param
	userId := r.FormValue("userId")
	param.SetUserId(userId)
	uuid := r.FormValue("uuid")
	param.SetUuid(uuid)
	pageno := r.FormValue("pageno")
	param.SetPageNum(pageno)
	param.SetChannel("gethistory")
	/*
		//验证MD5 校验token值  md5(userId+uuid+pageno+t+"ssports")
		timestamp := r.FormValue("t")

		token := r.FormValue("token")
		_token := String2MD5(userId + uuid + token + timestamp + "ssports")

		if token == "" || token != _token {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = api.PARAM_ILLEGAL
			return
		}
	*/
	res = history.GetHistory(&param)

	return res
}

func GetClassify(r *http.Request, w http.ResponseWriter) (res api.AppResponse) {
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = api.REQUEST_ERR
			Errorln(err)
		}
	}()
	aritcleId := r.FormValue("articleId")
	uid := r.FormValue("uid")
	uuid := r.FormValue("uuid")
	res = feedback.NegativeFb(aritcleId)
	//uid 唯一id 内容id
	DataLogf("%v\t%v\t%v\t%v\n", "NegativeFb", uid, uuid, aritcleId)
	return
}

func GetRelated(r *http.Request, w http.ResponseWriter) (res api.AppResponse) {
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = api.REQUEST_ERR
			Errorln(err)
		}
	}()
	articleId := r.FormValue("articleId")
	Debugf("request articleId is : %v", articleId)
	res = related.GetRelated(articleId)
	return
}
