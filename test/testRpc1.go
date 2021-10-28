package main

import (
	"../model"
	"../rpc"
	. "../util"
)

func main() {
	//	rpc.RpcInit()
	//testRankClient()
	testGETINFOBYIDS1()
	//rankServer()
	//usercfServer()
}

func rankServer1() {
	model.Test_rank_server()
}

func usercfServer1() {
	model.Test_usercf_server()
}

func testRankClient1() {

	param := make(map[string]string, 0)
	param["test1"] = "1111test"
	param["test2"] = "中文测试"
	param["test3"] = "33333333333333"

	ret, _ := rpc.RankClient(rpc.RANK_TEST, param)
	DataLogln(ret)
}

func testGETINFOBYIDS1() (err error) {

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
