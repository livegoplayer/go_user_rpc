package main

import (
	"fmt"
	"net"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	myHelper "github.com/livegoplayer/go_helper"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/viper"
	realgrpc "google.golang.org/grpc"

	redisHelper "github.com/livegoplayer/go_redis_helper"

	dbHelper "github.com/livegoplayer/go_db_helper"

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

	myHelper.LoadEnv()
	dbHelper.InitDbHelper(&dbHelper.MysqlConfig{Username: viper.GetString("database.username"), Password: viper.GetString("database.password"), Host: viper.GetString("database.host"), Port: int32(viper.GetInt("database.port")), Dbname: viper.GetString("database.dbname")}, viper.GetBool("database.log_mode"), viper.GetInt("database.max_open_connection"), viper.GetInt("database.max_idle_connection"))

	//报错日志
	grpcOpts := []grpctransport.ServerOption{
		//不是处理器，只是一个错误打印器
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	grpcListener, err := net.Listen("tcp", ":"+viper.GetString("app_port"))
	if err != nil {
		panic(err.Error())
	}

	//定义中间件
	endpoints := user.MakeUserEndpoints(&user.UserServiceServer{})
	baseServer := realgrpc.NewServer()

	redisHelper.InitRedisHelper(viper.GetString("redis.host"), viper.GetString("redis.port"), viper.GetString("redis.password"), viper.GetInt("redis.db"), viper.GetString("redis.prefix"), 2*time.Second)

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
