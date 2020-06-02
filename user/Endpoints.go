package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"go_user_rpc/util"
)

// ArithmeticRequest define request struct
type UserRequest struct {
	RequestType requesttype `json:"request_type"`
	RequestBody interface{} `json:"request_body"`
}

// ArithmeticResponse define response struct
type UserResponse struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Code int         `json:"error_code"`
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
		requestBody := req.RequestBody.(map[string]interface{})

		res := UserResponse{
			Msg:  "ok",
			Code: 0,
		}

		//第一个add user
		switch req.RequestType {
		//添加user
		case ADD_USER:
			userName := util.String(requestBody["username"])
			password := util.String(requestBody["password"])
			operationUid := util.Int(requestBody["operation_uid"])
			uid, err := svc.AddUser(userName, password, operationUid)
			if err != nil {
				res.Msg = err.Error()
				res.Code = 1
				break
			}

			res.Data = map[string]int{
				"uid": uid,
			}
			//todo
		case DEl_USER:
			uid := util.Int(requestBody["uid"])
			operationUid := util.Int(requestBody["operation_uid"])
			success, err := svc.DelUser(uid, operationUid)
			if err != nil {
				res.Msg = err.Error()
				res.Code = 1
				break
			}
			res.Data = map[string]bool{
				"success": success,
			}
		case LOGIN:
			userName := util.String(requestBody["username"])
			password := util.String(requestBody["password"])
			host := util.String(requestBody["host"])
			uid, userSession, tokenStr, err := svc.Login(userName, password, host)
			if err != nil {
				res.Msg = err.Error()
				res.Code = 1
				break
			}

			res.Data = map[string]interface{}{
				"userSession": userSession,
				"uid":         uid,
				"tokenStr":    tokenStr,
			}
		case REGISTER:
			userName := util.String(requestBody["username"])
			password := util.String(requestBody["password"])
			uid, err := svc.Register(userName, password)
			if err != nil {

				break
			}

			res.Data = map[string]int{
				"uid": uid,
			}
		case CHECK_USER_STATUS:
			tokenStr := util.String(requestBody["token"])
			host := util.String(requestBody["host"])

			isLogin, tokenStr, err := svc.CheckLoginStatus(tokenStr, host)
			if err != nil {
				res.Msg = err.Error()
				res.Code = 1
				break
			}
			res.Data = map[string]interface{}{
				"isLogin":  isLogin,
				"tokenStr": tokenStr,
			}
		}

		response = res
		return
	}
}
