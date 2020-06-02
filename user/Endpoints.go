package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// ArithmeticRequest define request struct
type UserRequest struct {
	RequestType requesttype `json:"request_type"`
	RequestBody map[string]interface{}
}

// ArithmeticResponse define response struct
type UserResponse struct {
	Data int   `json:"data"`
	Msg  error `json:"msg"`
	Code int   `json:"error_code"`
}

type requesttype int8

const (
	ADD_USER requesttype = iota
	DEl_USER
	LOGIN
	REGISTER
	CHECK_USER_STATUS
)

// rpc_handler 这里做一个路由
func MakeUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserRequest)

		//第一个add user
		switch req.RequestType {
		//添加user
		case ADD_USER:

		}

		return ArithmeticResponse{Result: res, Error: calError}, nil
	}
}
