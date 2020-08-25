package main

import (
	"fmt"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	myHelper "github.com/livegoplayer/go_helper"
	myLogger "github.com/livegoplayer/go_logger"
	redisHelper "github.com/livegoplayer/go_redis_helper"
	"github.com/oklog/oklog/pkg/group"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	realgrpc "google.golang.org/grpc"

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

	grpcListener, err := net.Listen("tcp", ":"+viper.GetString("app_port"))
	if err != nil {
		panic(err.Error())
	}

	grpcLogConfig := user.GrpcLoggerConfig{AccessFilePath: viper.GetString("log.access_log_file_path"), AccessFileName: viper.GetString("app_name") + "_" + viper.GetString("log.access_log_file_name"), PrintStd: viper.GetBool("log.print_to_std")}
	//定义中间件
	endpoints := user.MakeUserEndpoints(&user.UserServiceServer{}, grpcLogConfig)
	baseServer := realgrpc.NewServer()

	//初始化 app_log， 以后使用mylogger.Info() 打印log
	//如果是debug模式的话，直接打印到控制台
	var appLogger *logrus.Logger
	if gin.IsDebugging() {
		appLogger = myLogger.GetConsoleLogger()
	} else {
		appLogger = myLogger.GetMysqlLogger(viper.GetString("log.app_log_mysql_host"), viper.GetString("log.app_log_mysql_port"), viper.GetString("log.app_log_mysql_db_name"), viper.GetString("log.app_log_mysql_table_name"), viper.GetString("log.app_log_mysql_user"), viper.GetString("log.app_log_mysql_pass"))
	}
	myLogger.SetLogger(appLogger)

	redisHelper.InitRedisHelper(viper.GetString("redis.host"), viper.GetString("redis.port"), viper.GetString("redis.password"), viper.GetInt("redis.db"), viper.GetString("redis.prefix"), 2*time.Second)

	//如果开启了服务治理的话
	var register sd.Registrar
	if viper.GetBool("consul.open") {
		consulReg := user.NewConsulRegister(viper.GetString("consul.consul_address"), viper.GetString("consul.service_name"), viper.GetString("consul.service_port"), viper.GetInt("consul.service_ip"), viper.GetStringSlice("consul.tags"))
		consuelServerConfig := user.NewServerConfig(viper.GetString("consul.consul_server_addr"), viper.GetString("consul.token"), viper.GetString("consul.scheme"), viper.GetString("consul.dataCenter"))
		register = consulReg.NewConsulGRPCRegister(consuelServerConfig)
	}

	g.Add(func() error {
		//这里是执行函数
		userpb.RegisterUserServer(baseServer, user.MakeGRPCHandler(*endpoints))
		//附加的健康检查服务
		userpb.RegisterHealthServer(baseServer, &user.HealthImpl{})
		if register != nil {
			register.Register()
		}
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
