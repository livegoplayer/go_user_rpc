package user

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/endpoint"
)

//熔断器中间件
func HystrixMiddleware(commandName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var resp interface{}
			hystrix.ConfigureCommand(commandName, hystrix.CommandConfig{
				Timeout:                1000 * 30, //Timeout 【请求超时的时间】
				ErrorPercentThreshold:  1,         //ErrorPercentThreshold【允许出现的错误比例】
				SleepWindow:            10000,     //SleepWindow【熔断开启多久尝试发起一次请求】
				MaxConcurrentRequests:  1000,      //MaxConcurrentRequests【允许的最大并发请求数】
				RequestVolumeThreshold: 5,         //RequestVolumeThreshold 【波动期内的最小请求数，默认波动期10S】
			})

			if err := hystrix.Do(commandName, func() (err error) {
				resp, err = next(ctx, request)
				return err
			}, nil); err != nil {
				return nil, err
			}
			return resp, nil
		}
	}
}
