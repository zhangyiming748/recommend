package main

import (
	. "../util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	//url = "http://localhost:9090/api/recommend/v1/getChannelList?userId=&uuid=000355BCE1C5320067568762B8DB8721&action=up"
	//http://172.16.85.138:9090/api/recommend/admin
	//12229146
	//url = "http://172.16.85.138:9090/api/recommend/v1/getChannelList?userId=&uuid=000355BCE1C5320067568762B8DB8721&action=home"

	url = "http://172.16.85.138:9090/api/recommend/v1/getChannelList?userId=12229146&uuid=&action=home"
)

func pression() (d time.Duration, failed bool) {
	t := time.Now()
	resp, err := http.Get(url)
	d = time.Since(t)
	if err != nil {
		failed = true
		Errorln(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Errorln(err)
	}

	val := string(body)
	Debugln(val)
	var template interface{}
	if err = json.Unmarshal([]byte(val), &template); err != nil {
		Errorln("parsing json file", err.Error())
	}

	var retData map[string]interface{}
	ret, ok := template.(map[string]interface{})
	if ok {
		retData, ok = ret["retData"].(map[string]interface{})
		if strings.Split(retData["strategy"].(string), "_")[0] == "FAIL" {
			Debugln("strategy", retData["strategy"])
			failed = true

		}
	}
	return
}

var wg sync.WaitGroup

func main() {
	fmt.Println(os.Args)
	cur, _ := strconv.Atoi(os.Args[1])
	loops, _ := strconv.Atoi(os.Args[2])
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < cur*runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			latencies := make([]time.Duration, 0)
			failed := 0
			for j := 0; j < loops; j++ {
				t, f := pression()
				latencies = append(latencies, t)
				if f {
					failed += 1
				}
			}
			var d time.Duration = 0
			for _, v := range latencies {
				d += v
			}
			avg := d / time.Duration(loops)
			Errorln("avglatency: ", avg, " failed: ", failed)
		}()
	}
	wg.Wait()
}
