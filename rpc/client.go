package rpc

import (
	pb "../rpc/protoFile"
	. "../util"
	"./common"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"
	"strings"
)

const (
	// rank server
	RPC_RankServer    = "RankServer"
	RANK_TEST         = "test"
	RANK_GETINFOBYIDS = "getInfoByIds"
	RANK_FORNEW       = "fornew"
	RANK_FORREC       = "forrec"
	RANK_FORDATA      = "fordata"

	//todo orther server
	RPC_UserCfServer = "UserCfServer"
	USERCF_RECOMMEND = "recommend"
)

var conns map[string]*grpc.ClientConn

func RpcInit() {
	conns = make(map[string]*grpc.ClientConn)
	section := RunMode + "_rpc_args"
	server_names := GetVal(section, "server_names")
	Infof("reg rpc server_names:%v", server_names)
	names := strings.Split(server_names, ",")
	for _, n := range names {
		c, err := getConnInternal(n)
		if err != nil {
			conns[n] = nil
			Errorf("[%v] no connected", n)
		} else {
			conns[n] = c
			Infof("[%v] connected", n)
		}
	}
}

func getConnInternal(serviceName string) (conn *grpc.ClientConn, err error) {
	if conns == nil {
		conns = make(map[string]*grpc.ClientConn)
	}
	if len(conns) > 0 {
		conn = conns[serviceName]
		if conn == nil {
			conn, err = getConn(serviceName)
		}
	} else {
		conn, err = getConn(serviceName)
		conns[serviceName] = conn
	}
	if err != nil {
		Errorf("get rpc Conn err:%v", err)
		conn, err = getConn(serviceName)
		if err == nil {
			conns[serviceName] = conn
			Infof("try get rpc Conn %s", serviceName)
		} else {
			Errorf("try get rpc Conn %s:%v", serviceName, err)
		}
	}
	return
}

func getConn(serviceName string) (conn *grpc.ClientConn, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf(serviceName+" GetConn %v", e)
			Errorln(err)
		}
	}()
	section := RunMode + "_rpc_args"
	address := GetVal(section, "address")
	schema, err := common.ConsulResolver(address, serviceName)
	if err != nil {
		log.Fatal("init consul resovler err", err.Error())
		return nil, nil
	}
	addr := fmt.Sprintf("%s:///"+serviceName, schema)
	pool := common.New(func(addr string) (*grpc.ClientConn, error) {
		return grpc.Dial(
			addr,
			grpc.WithInsecure(),
			grpc.WithBalancerName(roundrobin.Name),
		)
	})
	conn, err = pool.GetConn(addr)
	return
}

/**
rank server 入口方法
*/
func RankClient(method string, param map[string]string) (ret []*pb.Article, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf(method+" GetConn %v", e)
			Errorln(err)
		}
	}()
	conn, err := getConnInternal(RPC_RankServer)
	c := pb.NewRankServiceClient(conn)
	Debugf("param:%v", param)
	request := pb.RankRequest{
		Method: method,
		Param:  param,
	}
	resp, err := c.Communication(context.Background(), &request)
	if err != nil {
		Errorf("RankReply error %v", err)
		return
	}
	if resp.Code != pb.RankReply_OK {
		Errorf("RankReply not ok %v", err)
	} else {
		ret = resp.Data
	}
	return
}

func UserCfClient(request pb.RecReqeust) (ret *pb.RecReply, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf(" GetConn %v", e)
			Errorln(err)
		}
	}()
	conn, err := getConnInternal(RPC_UserCfServer)
	c := pb.NewUserCfServiceClient(conn)
	ret, err = c.Recommend(context.Background(), &request)
	if err != nil {
		Errorf("RankReply error %v", ret)
		return
	}
	if ret.Code == pb.RecReply_FAIL {
		Errorf("RankReply not ok %v", ret)
	}
	return
}

/**
rank server 入口方法
*/
func UserCfClientMethod(method string, param map[string]string) (ret []*pb.Article, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf(method+" GetConn %v", e)
			Errorln(err)
		}
	}()
	conn, err := getConnInternal(RPC_UserCfServer)
	c := pb.NewUserCfServiceClient(conn)
	request := pb.UserCfRequest{
		Method: method,
		Param:  param,
	}
	resp, err := c.Communication(context.Background(), &request)
	if err != nil {
		Errorf("RankReply error %v", resp)
		return
	}
	if resp.Code == pb.UserCfReply_FAIL {
		Errorf("RankReply not ok %v", resp)
	} else {
		ret = resp.Data
	}
	return
}
