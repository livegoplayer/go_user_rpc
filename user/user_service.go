package user

import (
	"context"
	"github.com/livegoplayer/go_user_rpc/user/user_grpc"
)

type UserService struct {
}

//一个聚合的数据结构
type UserSession struct {
	Uid               int64            `json:"uid"`
	UserName          string           `json:"username"`
	UserRoleList      map[int64]string `json:"user_role_list"`
	UserAuthorityList map[int64]string `json:"user_authority_list"`
	AddDatetime       string           `json:"add_datetime"`
	UpdateDatetime    string           `json:"update_datetime"`
}

//用户登录 {}
func (us *UserService) Login(c context.Context, req *user_grpc.LoginRequest) (res *user_grpc.LoginResponse, err error) {
	data := &user_grpc.LoginData{}
	data.Uid, data.UserSession, data.Token, err = Login(req.UserName, req.Password)
	res = &user_grpc.LoginResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	if err != nil {
		data.Error = err.Error()
		res.Msg = err.Error()
		res.ErrorCode = 500
	}
	return
}

//用户注销
func (us *UserService) Logout(c context.Context, req *user_grpc.LogoutRequest) (res *user_grpc.LogoutResponse, err error) {
	data := &user_grpc.LogoutData{}
	data.Success = Logout(req.Uid)
	res = &user_grpc.LogoutResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	return
}

//管理员增加用户
func (us *UserService) AddUser(c context.Context, req *user_grpc.AddUserRequest) (res *user_grpc.AddUserResponse, err error) {
	data := &user_grpc.AddUserData{}
	data.Uid, err = AddUser(req.UserName, req.Password, req.OperationUid)
	res = &user_grpc.AddUserResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	if err != nil {
		data.Error = err.Error()
		res.Msg = err.Error()
		res.ErrorCode = 500
	}
	return
}

//管理员删除用户
func (us *UserService) DelUser(c context.Context, req *user_grpc.DelUserRequest) (res *user_grpc.DelUserResponse, err error) {
	data := &user_grpc.DelUserData{}
	data.Success, err = DelUser(req.Uid, req.OperationUid)
	res = &user_grpc.DelUserResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	return
}

//获取用户列表
func (us *UserService) GetUserList(c context.Context, req *user_grpc.GetUserListRequest) (res *user_grpc.GetUserListResponse, err error) {
	data := &user_grpc.GetUserListData{}
	data.UserList, data.Total = GetUserList(req.Page, req.Size)
	res = &user_grpc.GetUserListResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	return
}

//检查用户登陆状态
func (us *UserService) CheckUserStatus(c context.Context, req *user_grpc.CheckUserStatusRequest) (res *user_grpc.CheckUserStatusResponse, err error) {
	data := &user_grpc.CheckUserStatusData{}
	data.IsLogin, data.UserSession, data.Token, err = CheckUserStatus(req.Token)
	res = &user_grpc.CheckUserStatusResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	if err != nil {
		data.Error = err.Error()
		res.Msg = err.Error()
		res.ErrorCode = 500
	}
	return
}

//注册
func (us *UserService) Register(c context.Context, req *user_grpc.RegisterRequest) (res *user_grpc.RegisterResponse, err error) {
	data := &user_grpc.RegisterData{}
	data.Uid, err = Register(req.Username, req.Password)
	res = &user_grpc.RegisterResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	if err != nil {
		data.Error = err.Error()
		res.Msg = err.Error()
		res.ErrorCode = 500
	}
	return
}

//检查用户操作权限
func (us *UserService) CheckUserAuthority(c context.Context, req *user_grpc.CheckUserAuthorityRequest) (res *user_grpc.CheckUserAuthorityResponse, err error) {
	data := &user_grpc.CheckUserAuthorityData{}
	data.Exist = CheckUserAuthority(req.Uid, req.AuthorityId)
	res = &user_grpc.CheckUserAuthorityResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	return
}

//获取用户操作权限列表
func (us *UserService) GetUserAuthorityList(c context.Context, req *user_grpc.GetUserAuthorityListRequest) (res *user_grpc.GetUserAuthorityListResponse, err error) {
	data := &user_grpc.GetUserAuthorityListData{}
	data.UserAuthorityList = GetUserAuthorityMap(req.Uid)
	res = &user_grpc.GetUserAuthorityListResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	return
}

//获取所有操作权限
func (us *UserService) GetAuthorityList(c context.Context, req *user_grpc.GetAuthorityListRequest) (res *user_grpc.GetAuthorityListResponse, err error) {
	data := &user_grpc.GetAuthorityListData{}
	data.UserAuthorityList = GetAuthorityList()
	res = &user_grpc.GetAuthorityListResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	return
}

//增加用户权限
func (us *UserService) AddUserRole(c context.Context, req *user_grpc.AddUserRoleRequest) (res *user_grpc.AddUserRoleResponse, err error) {
	data := &user_grpc.AddUserRoleData{}
	data.Success = AddUserRole(req.Uid, req.RoleId, req.OperationUid)
	res = &user_grpc.AddUserRoleResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	return
}

//删除用户权限
func (us *UserService) DelUserRole(c context.Context, req *user_grpc.DelUserRoleRequest) (res *user_grpc.DelUserRoleResponse, err error) {
	data := &user_grpc.DelUserRoleData{}
	data.Success = DelUserRole(req.Uid, req.RoleId, req.OperationUid)
	res = &user_grpc.DelUserRoleResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	return
}

//获取所有角色列表
func (us *UserService) GetRoleList(c context.Context, req *user_grpc.GetRoleListRequest) (res *user_grpc.GetRoleListResponse, err error) {
	data := &user_grpc.GetRoleListData{}
	data.UserRoleList = GetRoleList()
	res = &user_grpc.GetRoleListResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	return
}

//获取用户角色列表
func (us *UserService) GetUserRoleList(c context.Context, req *user_grpc.GetUserRoleListRequest) (res *user_grpc.GetUserRoleListResponse, err error) {
	data := &user_grpc.GetUserRoleListData{}
	data.UserRoleList = GetUserRoleList(req.Uid)
	res = &user_grpc.GetUserRoleListResponse{}
	res.ErrorCode = 200
	res.Data = data
	res.Msg = "ok"
	return
}
