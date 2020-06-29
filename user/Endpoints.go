package user

import (
	"context"
	"errors"
	"reflect"

	"github.com/go-kit/kit/endpoint"

	mylogger "go_user_rpc/logger"
	user "go_user_rpc/user/grpc"
)

//handler

// ArithmeticRequest define request struct
type UserRequest struct {
	RequestType string            `json:"request_type"`
	RequestBody map[string]string `json:"request_body"`
}

// ArithmeticResponse define response struct
type UserResponse struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Code int32       `json:"error_code"`
}

type Requesttype int8

const (
	ADD_USER int32 = iota
	DEl_USER
	LOGIN
	REGISTER
	CHECK_USER_STATUS
)

type UserEndpoints struct {
	//用户登录
	Login endpoint.Endpoint
	//管理员增加用户
	AddUser endpoint.Endpoint
	//管理员删除用户
	DelUser endpoint.Endpoint
	//检查用户登陆状态
	CheckLoginStatus endpoint.Endpoint
	//注册
	Register endpoint.Endpoint
	//检查用户操作权限
	CheckUserAuthority endpoint.Endpoint
	//获取用户操作权限列表
	GetUserAuthorityList endpoint.Endpoint
	//获取所有操作权限
	GetAuthorityList endpoint.Endpoint
	//增加用户权限
	AddUserRole endpoint.Endpoint
	//删除用户权限
	DelUserRole endpoint.Endpoint
	//获取所有角色列表
	GetRoleList endpoint.Endpoint
	//获取用户角色列表
	GetUserRoleList endpoint.Endpoint
}

// rpc_handler 这里做一个路由
func MakeUserEndpoints(svc *UserServiceServer) *UserEndpoints {
	eps := &UserEndpoints{
		Login: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.LoginRequest)
			res := user.LoginResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.Login(ctx, req)
			return
		},
		AddUser: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.AddUserRequest)
			res := user.AddUserResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.AddUser(ctx, req)
			return
		},
		DelUser: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.DelUserRequest)
			res := user.DelUserResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.DelUser(ctx, req)
			return
		},
		CheckLoginStatus: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.CheckUserStatusRequest)
			res := user.CheckUserStatusResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.CheckUserStatus(ctx, req)
			return
		},
		Register: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.RegisterRequest)
			res := user.RegisterResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.Register(ctx, req)
			return
		},
		CheckUserAuthority: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.CheckUserAuthorityRequest)
			res := user.CheckUserAuthorityResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.CheckUserAuthority(ctx, req)
			return
		},
		GetUserAuthorityList: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.GetUserAuthorityListRequest)
			res := user.GetUserAuthorityListResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.GetUserAuthorityList(ctx, req)
			return
		},
		GetAuthorityList: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.GetAuthorityListRequest)
			res := user.GetAuthorityListResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.GetAuthorityList(ctx, req)
			return
		},
		AddUserRole: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.AddUserRoleRequest)
			res := user.AddUserRoleResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.AddUserRole(ctx, req)
			return
		},
		DelUserRole: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.DelUserRoleRequest)
			res := user.DelUserRoleResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.DelUserRole(ctx, req)
			return
		},
		GetRoleList: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.GetRoleListRequest)
			res := user.GetRoleListResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.GetRoleList(ctx, req)
			return
		},
		GetUserRoleList: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req, ok := request.(*user.GetUserRoleListRequest)
			res := user.GetUserRoleListResponse{}
			if !ok {
				res.Msg = "request body form error !"
				res.ErrorCode = 1
				response = res
				err = errors.New(res.Msg)
				return
			}

			response, err = svc.GetUserRoleList(ctx, req)
			return
		},
	}

	//貌似这种中间件的写法贴别局限性
	////增加中间件
	//for _, m := range mdw["Login"] {
	//	eps.Login = m(eps.Login)
	//}
	//for _, m := range mdw["AddUser"] {
	//	eps.AddUser = m(eps.AddUser)
	//}
	//for _, m := range mdw["DelUser"] {
	//	eps.DelUser = m(eps.DelUser)
	//}
	//for _, m := range mdw["CheckLoginStatus"] {
	//	eps.CheckLoginStatus = m(eps.CheckLoginStatus)
	//}
	//for _, m := range mdw["Register"] {
	//	eps.Register = m(eps.Register)
	//}
	//for _, m := range mdw["CheckUserAuthority"] {
	//	eps.CheckUserAuthority = m(eps.CheckUserAuthority)
	//}
	//for _, m := range mdw["GetUserAuthorityList"] {
	//	eps.GetUserAuthorityList = m(eps.GetUserAuthorityList)
	//}
	//for _, m := range mdw["GetAuthorityList"] {
	//	eps.GetAuthorityList = m(eps.GetAuthorityList)
	//}
	//for _, m := range mdw["AddUserRole"] {
	//	eps.AddUserRole = m(eps.AddUserRole)
	//}
	//for _, m := range mdw["DelUserRole"] {
	//	eps.DelUserRole = m(eps.DelUserRole)
	//}
	//for _, m := range mdw["GetRoleList"] {
	//	eps.GetRoleList = m(eps.GetRoleList)
	//}
	//for _, m := range mdw["GetUserRoleList"] {
	//	eps.GetUserRoleList = m(eps.GetUserRoleList)
	//}

	//全局访问日志中间件
	logger := mylogger.GetLogger()
	s := reflect.ValueOf(eps).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		name := typeOfT.Field(i).Name
		f := s.Field(i).Interface().(endpoint.Endpoint)
		//全局日志中间件
		loggerEP := loggingMiddleware(logger)(f)

		//全局微服务熔断器中间件
		commandName := "my_endpoint"
		breakerMw := HystrixMiddleware(commandName)
		hystrixEP := breakerMw(loggerEP)

		s.FieldByName(name).Set(reflect.ValueOf(hystrixEP))

	}

	return eps
}
