package config

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	cfg "github.com/livegoplayer/go_helper/config"
	myLogger "github.com/livegoplayer/go_logger/logger"
	myWriter "github.com/livegoplayer/go_logger/logger/writer"
	myMiddleware "github.com/livegoplayer/go_rpc_helper/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	realgrpc "google.golang.org/grpc"
	"net"
	"time"
)

func GetListener() net.Listener {

	grpcListener, err := net.Listen("tcp", ":"+viper.GetString("app_port"))
	if err != nil {
		panic(err.Error())
	}

	return grpcListener
}

func GetAndInitUserRpc() *realgrpc.Server {
	//定义中间件
	ms := make([]realgrpc.UnaryServerInterceptor, 0)
	// 這個錯誤處理會返回錯誤，後面銜接上rpc的返回，所以不需要重寫注意先後順序
	logger := logrus.New()
	logger.Out = myWriter.GetFileWriter(viper.GetString("log.access_log_file_path"), logrus.InfoLevel, 30*24*time.Hour, 24*time.Hour)
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Level = logrus.TraceLevel
	logger.AddHook(myLogger.GetDebugHook())

	ms = append(ms, myMiddleware.GetGrpcAccessLoggerOptions(logger)...)
	ms = append(ms, myMiddleware.GetGrpcErrHandler(cfg.IsDebug()))

	baseServer := realgrpc.NewServer(
		grpc_middleware.WithUnaryServerChain(ms...),
	)

	return baseServer
}
