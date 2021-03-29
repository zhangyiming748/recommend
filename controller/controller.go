package controller

import (
	"net/http"
	"recommend/api"
	"recommend/model"
	"recommend/model/recommend"
	"recommend/util"
)

func Exam(r *http.Request, w http.ResponseWriter) (res api.AppResponse) {

	return res
}
func RecommendAppHomePage(r *http.Request, w http.ResponseWriter) (res api.AppResponse) {
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = api.REQUEST_ERR
			util.Errorln(err)
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
	res = recommend.RecommendEntry(&param)
	return
}
