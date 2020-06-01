package user

import (
	"context"
	"strings"
)

// ArithmeticRequest define request struct
type UserRequest struct {
	RequestType string `json:"request_type"`
	RequestBody map[string]interface{}
}

// ArithmeticResponse define response struct
type UserResponse struct {
	Data int   `json:"data"`
	Msg  error `json:"msg"`
	Code int   `json:"error_code"`
}

// rpc_handler 这里做一个路由
func MakeUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserRequest)

		a = req.A
		b = req.B

		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else if strings.EqualFold(req.RequestType, "Substract") {
			res = svc.Subtract(a, b)
		} else if strings.EqualFold(req.RequestType, "Multiply") {
			res = svc.Multiply(a, b)
		} else if strings.EqualFold(req.RequestType, "Divide") {
			res, calError = svc.Divide(a, b)
		} else {
			return nil, ErrInvalidRequestType
		}

		return ArithmeticResponse{Result: res, Error: calError}, nil
	}
}
