package main

import (
	"../model"
	"../rpc"
	. "../util"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//func main() {
//	//	rpc.RpcInit()
//	//testRankClient()
//	testGETINFOBYIDS()
//	//rankServer()
//	//usercfServer()
//}

func rankServer() {
	//model.Test_rank_server()
}

func usercfServer() {
	model.Test_usercf_server()
}

func testRankClient() {

	param := make(map[string]string, 0)
	param["test1"] = "1111test"
	param["test2"] = "中文测试"
	param["test3"] = "33333333333333"

	ret, _ := rpc.RankClient(rpc.RANK_TEST, param)
	DataLogln(ret)
}

func testGETINFOBYIDS() (err error) {

	param := make(map[string]string, 0)
	//param["ids"] = "557223,561464,1791736,1780955,1781578"
	param["ids"] = "1411440,9028870,9029065,9048396,9046543,9048511,9057611,9070659,9071159,9071193,6000,6008"
	//param["ids"] = "1798957"

	ret, err := rpc.RankClient(rpc.RANK_GETINFOBYIDS, param)
	for _, v := range ret {
		art := model.Article{*v}
		DataLogln(art.Id)
	}
	return
}

func grpcTest() (d time.Duration, failed bool) {
	t := time.Now()
	err := testGETINFOBYIDS()
	if err != nil {
		failed = true
	}
	d = time.Since(t)
	DataLogln("Since", d)
	return
}

var wgg sync.WaitGroup

func main() {

	fmt.Println(os.Args)
	cur, _ := strconv.Atoi(os.Args[1])
	loops, _ := strconv.Atoi(os.Args[2])

	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < cur*runtime.NumCPU(); i++ {
		wgg.Add(1)

		go func() {
			defer wgg.Done()
			latencies := make([]time.Duration, 0)
			failed := 0
			for j := 0; j < loops; j++ {
				t, f := grpcTest()
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
			DataLogln("avglatency: ", avg, " failed: ", failed)
		}()

	}
	wgg.Wait()
}
