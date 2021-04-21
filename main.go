package main

import (
	"fmt"
	"github.com/livegoplayer/go_user_rpc/config"
	"github.com/livegoplayer/go_user_rpc/user"
	userpb "github.com/livegoplayer/go_user_rpc/user/user_grpc"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc/reflection"
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

	config.InitLog()
	config.InitDb()

	grpcListener := config.GetListener()

	baseServer := config.GetAndInitUserRpc()

	config.InitRedis()

	g.Add(func() error {
		//这里是执行函数
		userpb.RegisterUserServer(baseServer, &user.UserService{})
		// 注册reflection服务 ，可以使用 grpcurl -plaintext localhost:8888 list 测试
		reflection.Register(baseServer)
		fmt.Printf("start..")
		return baseServer.Serve(grpcListener)
	}, func(err error) {
		//这里是遇到错误的中断处理函数
		fmt.Printf(err.Error())
		baseServer.Stop()
	})
}
