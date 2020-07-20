package main

import (
	"fmt"
	"net"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	realgrpc "google.golang.org/grpc"

	redisHelper "github.com/livegoplayer/go_redis_helper"

	"github.com/livegoplayer/go_user_rpc/dbHelper"
	"github.com/livegoplayer/go_user_rpc/user"
	userpb "github.com/livegoplayer/go_user_rpc/user/grpc"
)

const rateBucketNum = 20

var (
	logger log.Logger
)

func main() {
	Run()
}

func Run() {
	//initHttpHandler()
	//g oklog是线程组管理工具
	g := &group.Group{}
	initUserRpcHandler(g)
	_ = g.Run()
}

func initUserRpcHandler(g *group.Group) {

	dbHelper.InitDbHelper(&dbHelper.MysqlConfig{Username: "myuser", Password: "myuser", Host: "139.224.132.234", Port: 3306, Dbname: "user"}, true, 100, 20)
	//报错日志
	grpcOpts := []grpctransport.ServerOption{
		//不是处理器，只是一个错误打印器
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	grpcListener, err := net.Listen("tcp", ":8888")
	if err != nil {
		//todo log
	}

	//定义中间件
	endpoints := user.MakeUserEndpoints(&user.UserServiceServer{})
	baseServer := realgrpc.NewServer()

	redisHelper.InitRedisHelper("139.224.132.234", "6379", "myredis", 0, "us_redis_")

	g.Add(func() error {
		//这里是执行函数
		userpb.RegisterUserServer(baseServer, user.MakeGRPCHandler(*endpoints, grpcOpts...))
		fmt.Printf("start..")
		return baseServer.Serve(grpcListener)
	}, func(err error) {
		//这里是遇到错误的中断处理函数
		fmt.Printf(err.Error())
		baseServer.Stop()
	})
}
