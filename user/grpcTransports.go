package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	user "github.com/livegoplayer/go_user_rpc/user/grpc"
)

//使用tcp协议做的grpc
type userGrpcService struct {
	login                endpoint.Endpoint
	logout               endpoint.Endpoint
	checkUserStatus      endpoint.Endpoint
	register             endpoint.Endpoint
	addUser              endpoint.Endpoint
	delUser              endpoint.Endpoint
	getUserList          endpoint.Endpoint
	checkUserAuthority   endpoint.Endpoint
	getUserAuthorityList endpoint.Endpoint
	getAuthorityList     endpoint.Endpoint
	addUserRole          endpoint.Endpoint
	delUserRole          endpoint.Endpoint
	getUserRoleList      endpoint.Endpoint
	getRoleList          endpoint.Endpoint
}

func MakeGRPCHandler(eps UserEndpoints) *userGrpcService {
	return &userGrpcService{
		login:                eps.Login,
		checkUserStatus:      eps.CheckLoginStatus,
		register:             eps.Register,
		addUser:              eps.AddUser,
		delUser:              eps.DelUser,
		getUserList:          eps.GetUserList,
		checkUserAuthority:   eps.CheckUserAuthority,
		getUserAuthorityList: eps.GetUserAuthorityList,
		getAuthorityList:     eps.GetAuthorityList,
		addUserRole:          eps.AddUserRole,
		delUserRole:          eps.DelUserRole,
		getUserRoleList:      eps.GetUserRoleList,
		getRoleList:          eps.GetRoleList,
		logout:               eps.Logout,
	}
}

//预处理吗，相当于中间件，但因为可以存放定制逻辑，所以保留在这里，以后写法推荐是switch，不必每一个都重写，那样太乱
func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	return r, nil
}

func encodeRpcResponse(_ context.Context, r interface{}) (interface{}, error) {
	return r, nil
}

//以下实现关于UserServer 相关接口
//以下函数是给endpoints调用的，最后会被安装到对应的成员变量上，所以一定要是grpc.Handler 类型的
func (u *userGrpcService) Login(ctx context.Context, request *user.LoginRequest) (loginResponse *user.LoginResponse, err error) {
	response, err := u.login(ctx, request)
	loginResponse = response.(*user.LoginResponse)

	return
}

func (u *userGrpcService) CheckUserStatus(ctx context.Context, request *user.CheckUserStatusRequest) (checkUserStatusResponse *user.CheckUserStatusResponse, err error) {
	response, err := u.checkUserStatus(ctx, request)
	checkUserStatusResponse = response.(*user.CheckUserStatusResponse)

	return
}

func (u *userGrpcService) Register(ctx context.Context, request *user.RegisterRequest) (registerResponse *user.RegisterResponse, err error) {
	response, err := u.register(ctx, request)
	registerResponse = response.(*user.RegisterResponse)

	return
}

func (u *userGrpcService) AddUser(ctx context.Context, request *user.AddUserRequest) (addUserResponse *user.AddUserResponse, err error) {
	response, err := u.addUser(ctx, request)
	addUserResponse = response.(*user.AddUserResponse)

	return
}

func (u *userGrpcService) DelUser(ctx context.Context, request *user.DelUserRequest) (delUserResponse *user.DelUserResponse, err error) {
	response, err := u.delUser(ctx, request)
	delUserResponse = response.(*user.DelUserResponse)

	return
}

func (u *userGrpcService) CheckUserAuthority(ctx context.Context, request *user.CheckUserAuthorityRequest) (checkUserAuthorityResponse *user.CheckUserAuthorityResponse, err error) {
	response, err := u.checkUserAuthority(ctx, request)
	checkUserAuthorityResponse = response.(*user.CheckUserAuthorityResponse)

	return
}
func (u *userGrpcService) GetUserAuthorityList(ctx context.Context, request *user.GetUserAuthorityListRequest) (getUserAuthorityListResponse *user.GetUserAuthorityListResponse, err error) {
	response, err := u.getUserAuthorityList(ctx, request)
	getUserAuthorityListResponse = response.(*user.GetUserAuthorityListResponse)

	return
}

func (u *userGrpcService) GetAuthorityList(ctx context.Context, request *user.GetAuthorityListRequest) (getAuthorityListResponse *user.GetAuthorityListResponse, err error) {
	response, err := u.getAuthorityList(ctx, request)
	getAuthorityListResponse = response.(*user.GetAuthorityListResponse)

	return
}

func (u *userGrpcService) AddUserRole(ctx context.Context, request *user.AddUserRoleRequest) (addUserRoleResponse *user.AddUserRoleResponse, err error) {
	response, err := u.addUserRole(ctx, request)
	addUserRoleResponse = response.(*user.AddUserRoleResponse)

	return
}

func (u *userGrpcService) DelUserRole(ctx context.Context, request *user.DelUserRoleRequest) (delUserRoleResponse *user.DelUserRoleResponse, err error) {
	response, err := u.delUserRole(ctx, request)
	delUserRoleResponse = response.(*user.DelUserRoleResponse)

	return
}

func (u *userGrpcService) GetUserRoleList(ctx context.Context, request *user.GetUserRoleListRequest) (getUserRoleListResponse *user.GetUserRoleListResponse, err error) {
	response, err := u.getUserRoleList(ctx, request)
	getUserRoleListResponse = response.(*user.GetUserRoleListResponse)

	return
}

func (u *userGrpcService) GetRoleList(ctx context.Context, request *user.GetRoleListRequest) (getRoleListResponse *user.GetRoleListResponse, err error) {
	response, err := u.getRoleList(ctx, request)
	getRoleListResponse = response.(*user.GetRoleListResponse)

	return
}

func (u *userGrpcService) GetUserList(ctx context.Context, request *user.GetUserListRequest) (getUserListResponse *user.GetUserListResponse, err error) {
	response, err := u.getUserList(ctx, request)
	getUserListResponse = response.(*user.GetUserListResponse)

	return
}

func (u *userGrpcService) Logout(ctx context.Context, request *user.LogoutRequest) (logoutResponse *user.LogoutResponse, err error) {
	response, err := u.logout(ctx, request)
	logoutResponse = response.(*user.LogoutResponse)

	return
}
