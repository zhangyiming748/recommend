package controller

import (
	"log"
	"net/http"
	"recommend/api"
	"recommend/module"
	"recommend/module/recommend"
)


func Exam(r *http.Request, w http.ResponseWriter) (res api.AppResponse) {
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = api.REQUEST_ERR
			log.Println(err)
		}
	}()
	var contant api.AppResponse
	res = contant
	return
}
func RecommendHome(r *http.Request,w http.ResponseWriter)(res api.AppResponse){
	defer func() {
		if err := recover(); err != nil {
			res.ResCode = api.APPRESPONSE_CODE_SUCCESS
			res.ResStatus = api.REQUEST_ERR
			log.Println(err)
		}
	}()
	var param module.Param
	userId := r.FormValue("userId")
	param.SetuserId(userId)
	res = recommend.RecommendEntey(&param)
	return
}