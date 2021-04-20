package main

import (
	"fmt"
	"github.com/go-kit/kit/sd"
	"github.com/livegoplayer/go_user_rpc/config"
	"net"

	"github.com/go-kit/kit/log"

	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/viper"
	realgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/livegoplayer/go_user_rpc/user"
	userpb "github.com/livegoplayer/go_user_rpc/user/user_grpc"
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

	config.InitDb()

	grpcListener, err := net.Listen("tcp", ":"+viper.GetString("app_port"))
	if err != nil {
		panic(err.Error())
	}

	grpcLogConfig := user.GrpcLoggerConfig{AccessFilePath: viper.GetString("log.access_log_file_path"), AccessFileName: viper.GetString("app_name") + "_" + viper.GetString("log.access_log_file_name"), PrintStd: viper.GetBool("log.print_to_std")}
	//定义中间件
	baseServer := realgrpc.NewServer()

	config.InitLog()
	config.InitRedis()

	var register sd.Registrar
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
		if register != nil {
			register.Deregister()
		}
		baseServer.Stop()
	})
}
