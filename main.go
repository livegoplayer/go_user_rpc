package main

import (
	"fmt"
	"net"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	realgrpc "google.golang.org/grpc"

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

	//报错日志
	grpcOpts := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	grpcListener, err := net.Listen("tcp", ":8888")
	if err != nil {
		//todo log
	}

	//定义中间件
	endpoints := user.MakeUserEndpoints(&user.UserServiceServer{})
	baseServer := realgrpc.NewServer()

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
