package user

import (
	"context"
	"reflect"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"

	"go_user_rpc/helper"
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
			requestbody := helper.JsonEncode(request)
			errStr := ""

			response, err := next(ctx, request)
			responseBody := helper.JsonEncode(response)
			if err != nil {
				errStr = err.Error()
				logger.Error(errStr)
			}

			logInfo := LogInfo{
				NowTimeStr:   nowTimeStr,
				nowTimeUnix:  nowTimeUnix,
				RequestName:  typStr,
				RequestBody:  helper.BytesToString(requestbody),
				ResponseBody: helper.BytesToString(responseBody),
				Error:        errStr,
			}

			//定位是访问日志
			logger.Info(helper.BytesToString(helper.JsonEncode(logInfo)))

			return response, err
		}
	}
}
