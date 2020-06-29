package user

import (
	"context"
	"reflect"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"

	myHelper "github.com/livegoplayer/go_helper"
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

func loggingMiddleware(logger *logrus.Logger) Middleware {
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
			logger.Info(myHelper.BytesToString(myHelper.JsonEncode(logInfo)))

			return response, err
		}
	}
}
