package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"recommend/api"
	"recommend/controller"
	"recommend/model/recommend"
	"recommend/storage"
	. "recommend/util"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	url_prefix = "/api/recommend"
	server     Server
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	server = NewMonitorServer()

	initMode()

	router := makeRouters()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	n := negroni.New(negroni.NewRecovery())
	n.Use(c)
	n.UseHandler(router)
	Infoln("addr:", ":"+GetVal(RunMode+"_args", "port"))
	s := &http.Server{
		Addr:           ":" + GetVal(RunMode+"_args", "port"),
		Handler:        n,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	Infoln("server started !!!")
	log.Fatal(s.ListenAndServe())
}

func makeRouters() *mux.Router {

	jsonrender := render.New(render.Options{UnEscapeHTML: false})

	wrapper := func(apphandler func(*http.Request, http.ResponseWriter) api.AppResponse) func(w http.ResponseWriter, req *http.Request) {
		return func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			server.Hits.Add(1)
			server.Qps.Add(1)

			path := req.URL.Path
			arr := strings.Split(path, "/")
			version := arr[3]
			if !strings.Contains(req.RequestURI, version) {
				version = "v1"
			}
			Debugln("version:", version)
			resp := apphandler(req, w)
			resp.RequsetAction = strings.Join(arr[0:5], "/")

			jsonrender.JSON(w, http.StatusOK, resp)

			Duration := time.Since(start)
			DataLogf("%s\t%v\t%s\t%s\t%s\n", start.Format("2006-01-02 15:04:05"), Duration, req.Host, req.Method, req.URL.Path)
			server.SetLatency(Duration)
		}
	}

	router := mux.NewRouter()
	/*
		http://localhost:9090/api/recommend/v1/getChannelList?userId=703007&uuid=fsdaf&action=home
		http://localhost:9797/api/recommend/v1/getHotList?userId=703007&uuid=fsdaf&action=home
		http://localhost:9797/api/recommend/v1/getClickHistory?userId=703007&uuid=fsdaf&pageno=1
		http://localhost:9797/api/recommend/v1/admin?arsche=mun12345rm2019-11-26
	*/

	router.HandleFunc(url_prefix+"/v1/getChannelList", wrapper(controller.RecommendAppHomePage))
	router.HandleFunc(url_prefix+"/v1/getHotList", wrapper(controller.RecommendHotVideoPage))

	router.HandleFunc(url_prefix+"/v1/getClickHistory", wrapper(controller.GetClickHistory))

	router.HandleFunc(url_prefix+"/v1/getClassify", wrapper(controller.GetClassify))

	router.HandleFunc(url_prefix+"/v1/getRelatedRecommend", wrapper(controller.GetRelated))
	/*
		http://localhost:9090/api/recommend/v1/getRelatedRecommend?articleId=14552111
		http://localhost:9090/api/recommend/v1/getRelatedRecommend?articleId=14545060
		http://localhost:9090/api/recommend/v1/getRelatedRecommend?articleId=14552111
		articleDetail:14554878
		14556738足球英超
	*/
	router.HandleFunc(url_prefix+"/admin", func(w http.ResponseWriter, req *http.Request) {
		passwd := fmt.Sprintf("mun12345rm%d%d%d", time.Now().Year(), time.Now().Month(), time.Now().Day())
		Infoln("adminpass:", passwd)
		if req.FormValue("arsche") == passwd {
			jsonrender.JSON(w, http.StatusOK, server)
		}
	})

	return router
}

func initMode() {
	//流量控制
	maxtraffic, _ := strconv.Atoi(GetVal(RunMode+"_args", "maxtraffic"))
	if maxtraffic == 0 {
		maxtraffic = 1024 //default value
	}
	recommend.SetMaxTraffic(uint64(maxtraffic))

	//聚合结果概率控制
	cprob := strings.Split(GetVal(RunMode+"_args", "chooseProb"), ",")
	recommend.SetChooseProb(cprob)

	//abtest流量控制
	recommend.SetABtest(GetVal(RunMode+"_args", "abtest"))

	if RunMode == "product" {
		Setloglevel("INFO")
		Debugln(RunMode)
	} else {
		Setloglevel("DEBUG")
		Debugln(RunMode)
	}

	storage.NewRediss(RunMode)

	return
}
