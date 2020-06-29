package user

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
)

//使用tcp协议做的grpc
type userGrpcService struct {
	UserServiceServer
	login                grpc.Handler
	checkUserStatus      grpc.Handler
	register             grpc.Handler
	addUser              grpc.Handler
	delUser              grpc.Handler
	checkUserAuthority   grpc.Handler
	getUserAuthorityList grpc.Handler
	getAuthorityList     grpc.Handler
	addUserRole          grpc.Handler
	delUserRole          grpc.Handler
	getUserRoleList      grpc.Handler
	getRoleList          grpc.Handler
}

func MakeGRPCHandler(eps UserEndpoints, opts ...grpc_transport.ServerOption) *userGrpcService {
	return &userGrpcService{
		login: grpc.NewServer(
			eps.Login,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		checkUserStatus: grpc.NewServer(
			eps.CheckLoginStatus,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		register: grpc.NewServer(
			eps.Register,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		addUser: grpc.NewServer(
			eps.AddUser,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		delUser: grpc.NewServer(
			eps.DelUser,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		checkUserAuthority: grpc.NewServer(
			eps.CheckUserAuthority,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		getUserAuthorityList: grpc.NewServer(
			eps.GetUserAuthorityList,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		getAuthorityList: grpc.NewServer(
			eps.GetAuthorityList,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		addUserRole: grpc.NewServer(
			eps.AddUserRole,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		delUserRole: grpc.NewServer(
			eps.DelUserRole,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		getUserRoleList: grpc.NewServer(
			eps.GetUserRoleList,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
		getRoleList: grpc.NewServer(
			eps.GetRoleList,
			decodeGetRequest,
			encodeRpcResponse,
			opts...,
		),
	}
}

//预处理吗，相当于中间件，但因为可以存放定制逻辑，所以保留在这里，以后写法推荐是switch，不必每一个都重写，那样太乱
func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	return r, nil
}

func encodeRpcResponse(_ context.Context, r interface{}) (interface{}, error) {
	return r, nil
}
