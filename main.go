package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render" //Render是一个软件包，提供了轻松呈现JSON，XML，文本，二进制数据和HTML模板的功能。该程序包基于Martini 渲染工作。
	"github.com/urfave/negroni"  //Negroni 不是一个框架，它是为了方便使用 net/http 而设计的一个库而已。
	"log"
	"net/http"
	"recommend/api"
	"recommend/controller"
	"time"
	// "github.com/xujiajun/nutsdb"
	"strings"
)

var (
	url_prefix = "/api/recommend"
	//server     Server
)

func main() {

	router := makeRouters()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	n := negroni.New(negroni.NewRecovery())
	n.Use(c)
	n.UseHandler(router)

	s := &http.Server{
		Addr:           ":6400",
		Handler:        n,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	address := strings.Join([]string{"127.0.0.1", s.Addr}, "")
	log.Printf("addr:%s", address)
	log.Println("server started !!!")
	log.Fatal(s.ListenAndServe())
}
func makeRouters() *mux.Router {
	jsonrender := render.New(render.Options{UnEscapeHTML: false})
	wrapper := func(apphandler func(*http.Request, http.ResponseWriter) api.AppResponse) func(w http.ResponseWriter, req *http.Request) {
		return func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			//server.Hits.Add(1)
			//server.Qps.Add(1)
			path := req.URL.Path
			arr := strings.Split(path, "/")
			version := arr[3]
			if !strings.Contains(req.RequestURI, version) {
				version = "v1"
			}
			log.Println("version:", version)
			resp := apphandler(req, w)
			resp.RequsetAction = strings.Join(arr[0:5], "/")
			jsonrender.JSON(w, http.StatusOK, resp)
			Duration := time.Since(start)
			log.Printf("%s\t%v\t%s\t%s\t%s\n", start.Format("2006-01-02 15:04:05"), Duration, req.Host, req.Method, req.URL.Path)
			//server.SetLatency(Duration)
		}
	}
	router := mux.NewRouter()
	router.HandleFunc(url_prefix+"/v1/exam", wrapper(controller.Exam))
	return router
}
