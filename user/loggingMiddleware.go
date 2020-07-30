package user

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/go-kit/kit/endpoint"
	myHelper "github.com/livegoplayer/go_helper"
	logger "github.com/livegoplayer/go_logger"
)

//全局中间件日志组件
type Middleware func(endpoint.Endpoint) endpoint.Endpoint

type LogInfo struct {
	NowTimeStr   string
	nowTimeUnix  int64
	RequestName  string
	RequestBody  string
	ResponseBody string
	Error        string
}

type GrpcLoggerConfig struct {
	AccessFilePath string
	AccessFileName string
	PrintStd       bool
}

func loggingMiddleware(grpcLoggerConfig GrpcLoggerConfig) Middleware {
	if grpcLoggerConfig.AccessFilePath == "" {
		grpcLoggerConfig.AccessFilePath = "../log"
	}

	if grpcLoggerConfig.AccessFileName == "" {
		grpcLoggerConfig.AccessFileName = "access.log"
	}

	if !myHelper.Exists(myHelper.PathToCommon(grpcLoggerConfig.AccessFilePath)) {
		err := os.MkdirAll(grpcLoggerConfig.AccessFilePath, os.ModeDir)
		if err != nil {
			panic("创建日志文件目录失败")
		}
	}

	if !myHelper.Exists(myHelper.PathToCommon(grpcLoggerConfig.AccessFilePath + "/" + grpcLoggerConfig.AccessFileName)) {
		file, err := os.Create(grpcLoggerConfig.AccessFilePath + "/" + grpcLoggerConfig.AccessFileName)
		if err != nil {
			panic("创建日志文件失败")
		}
		err = os.Chmod(grpcLoggerConfig.AccessFilePath+"/"+grpcLoggerConfig.AccessFileName, 0777)
		if err != nil {
			panic("修改文件权限失败")
		}
		file.Close()
	}

	accessLogFile, err := os.OpenFile(grpcLoggerConfig.AccessFilePath+"/"+grpcLoggerConfig.AccessFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic("打开日志文件失败" + err.Error())
	}

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		//这里面就是一个endpoint
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			//获取固定格式
			//调用了哪个方法，当前时间等等
			nowTime := time.Now()
			nowTimeStr := nowTime.String()
			nowTimeUnix := nowTime.Unix()
			s := reflect.ValueOf(&request).Elem()
			typeOfT := s.Type()
			typStr := typeOfT.Kind().String()
			requestbody := myHelper.JsonEncode(request)
			errStr := ""

			response, err := next(ctx, request)
			responseBody := myHelper.JsonEncode(response)
			if err != nil {
				errStr = err.Error()
				logger.Error(errStr)
			}

			logInfo := LogInfo{
				NowTimeStr:   nowTimeStr,
				nowTimeUnix:  nowTimeUnix,
				RequestName:  typStr,
				RequestBody:  myHelper.BytesToString(requestbody),
				ResponseBody: myHelper.BytesToString(responseBody),
				Error:        errStr,
			}

			//定位是访问日志
			_, _ = fmt.Fprintf(accessLogFile, MyGRPCAccessLoggerFormatter(logInfo))
			if grpcLoggerConfig.PrintStd {
				_, _ = fmt.Fprintf(os.Stdout, MyGRPCAccessLoggerStringFormatter(logInfo))
			}

			return response, err
		}
	}
}

//专门给MyGRPCLoggerFormatter定制的access log文件输出模式
//直接json encode
func MyGRPCAccessLoggerFormatter(logInfo LogInfo) string {
	return myHelper.BytesToString(myHelper.JsonEncode(logInfo))
}

func MyGRPCAccessLoggerStringFormatter(logInfo LogInfo) string {
	return fmt.Sprintf("[GRPC] %s %s %s %s %s", logInfo.NowTimeStr, logInfo.NowTimeStr, logInfo.RequestName, "\nrequest:"+string(myHelper.JsonEncode(logInfo.RequestBody)), "\nresponse:"+string(myHelper.JsonEncode(logInfo.ResponseBody)))
}
