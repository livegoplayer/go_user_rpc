package grpc

import (
	"github.com/go-kit/kit/transport/grpc"

	"go_user_rpc/user"
	"go_user_rpc/util"
)

//专门的go rpc server代码，为了和基础的service代码解耦，主要是避免输入输出的对象化，那样写很难受

type UserServiceServer struct {
	get grpc.Handler
}

func (userServiceServer *UserServiceServer) Login(request *LoginRequest) (loginResponse LoginResponse) {

	uid, userSession, token, err := user.UserServiceInstance.Login(request.GetUserName(), request.GetPassword(), request.GetHost())
	if err != nil {
		loginResponse.Msg = err.Error()
		loginResponse.ErrorCode = 1
		return
	}
	loginData := &LoginData{}
	loginData.Uid = uid
	loginData.UserSession = utilUserSessionToResponse(userSession)
	loginData.Token = token

	loginResponse.Data = loginData

	return
}

func (userServiceServer *UserServiceServer) CheckUserStatus(checkUserStatusRequest *CheckUserStatusRequest) (checkUserStatusResponse CheckUserStatusResponse) {

	isLogin, tokenStr, err := user.UserServiceInstance.CheckLoginStatus(checkUserStatusRequest.GetToken(), checkUserStatusRequest.GetHost())
	if err != nil {
		checkUserStatusResponse.Msg = err.Error()
		checkUserStatusResponse.ErrorCode = 1
		return
	}
	checkUserStatusData := &CheckUserStatusData{}
	checkUserStatusData.IsLogin = isLogin
	checkUserStatusData.Token = tokenStr

	checkUserStatusResponse.Data = checkUserStatusData

	return
}

func (userServiceServer *UserServiceServer) Register(request *RegisterRequest) (registerResponse RegisterResponse) {
	uid, err := user.UserServiceInstance.Register(request.GetToken(), request.GetHost())
	if err != nil {
		registerResponse.Msg = err.Error()
		registerResponse.ErrorCode = 1
		return
	}
	registerData := &RegisterData{}
	registerData.Uid = uid

	registerResponse.Data = registerData

	return
}

func (userServiceServer *UserServiceServer) AddUser(request *AddUserRequest) (addUserResponse AddUserResponse) {
	uid, err := user.UserServiceInstance.AddUser(request.GetUserName(), request.GetPassword(), request.GetOperationUid())
	if err != nil {
		addUserResponse.Msg = err.Error()
		addUserResponse.ErrorCode = 1
		return
	}
	addUserData := &AddUserData{}
	addUserData.Uid = uid

	addUserResponse.Data = addUserData
	return
}

func (userServiceServer *UserServiceServer) DelUser(request *DelUserRequest) (delUserResponse DelUserResponse) {
	success, err := user.UserServiceInstance.DelUser(request.GetUid(), request.GetOperationUid())
	if err != nil {
		delUserResponse.Msg = err.Error()
		delUserResponse.ErrorCode = 1
		return
	}
	delUserData := &DelUserData{}
	delUserData.Success = success

	delUserResponse.Data = delUserData
	return
}

func (userServiceServer *UserServiceServer) CheckUserAuthority(request *CheckUserAuthorityRequest) (checkUserAuthorityResponse CheckUserAuthorityResponse) {
	exists := user.UserServiceInstance.CheckUserAuthority(request.GetUid(), request.GetAuthorityId())

	checkUserAuthorityData := &CheckUserAuthorityData{}
	checkUserAuthorityData.Exist = exists

	checkUserAuthorityResponse.Data = checkUserAuthorityData
	return
}
func (userServiceServer *UserServiceServer) GetUserAuthorityList(request *GetUserAuthorityListRequest) (getUserAuthorityListResponse GetUserAuthorityListResponse) {
	userAuthorityList := user.UserServiceInstance.GetUserAuthorityList(request.GetUid())

	getUserAuthorityListData := &GetUserAuthorityListData{}
	getUserAuthorityListData.UserAuthorityList = userAuthorityList

	getUserAuthorityListResponse.Data = getUserAuthorityListData
	return
}

func (userServiceServer *UserServiceServer) GetAuthorityList(request *GetAuthorityListRequest) (getAuthorityListResponse GetAuthorityListResponse) {
	authorityList := user.UserServiceInstance.GetAuthorityList()

	getAuthorityListData := &GetAuthorityListData{}
	getAuthorityListData.UserAuthorityList = authorityList

	getAuthorityListResponse.Data = getAuthorityListData
	return
}

func (userServiceServer *UserServiceServer) AddUserRole(request *AddUserRoleRequest) (addUserRoleResponse AddUserRoleResponse) {
	success := user.UserServiceInstance.AddUserRole(request.GetUid(), request.GetRoleId(), request.GetOperationUid())

	addUserRoleData := &AddUserRoleData{}
	addUserRoleData.Success = success

	addUserRoleResponse.Data = addUserRoleData
	return
}
func (userServiceServer *UserServiceServer) DelUserRole(request *DelUserRoleRequest) (delUserRoleResponse DelUserRoleResponse) {
	success := user.UserServiceInstance.AddUserRole(request.GetUid(), request.GetRoleId(), request.GetOperationUid())

	delUserRoleData := &DelUserRoleData{}
	delUserRoleData.Success = success

	delUserRoleResponse.Data = delUserRoleData
	return
}

func (userServiceServer *UserServiceServer) GetUserRoleList(request *GetUserRoleListRequest) (getUserRoleListResponse GetUserRoleListResponse) {
	userRoleList := user.UserServiceInstance.GetUserRoleList(request.GetUid())

	getUserRoleListData := &GetUserRoleListData{}
	getUserRoleListData.UserRoleList = userRoleList

	getUserRoleListResponse.Data = getUserRoleListData
	return
}

func (userServiceServer *UserServiceServer) GetRoleList(request *GetRoleListRequest) (getRoleListResponse GetRoleListResponse) {
	roleList := user.UserServiceInstance.GetRoleList()

	getRoleListData := &GetRoleListData{}
	getRoleListData.UserRoleList = roleList

	getRoleListResponse.Data = getRoleListData
	return
}

func utilUserSessionToResponse(session *util.UserSession) *UserSessions {
	responseUserSession := &UserSessions{
		Uid:               int32(session.Uid),
		Username:          session.UserName,
		UserRoleList:      session.UserRoleList,
		UserAuthorityList: session.UserAuthorityList,
		AddDatetime:       session.AddDatetime,
		UpdateDatetime:    session.UpdateDatetime,
	}

	return responseUserSession
}

func responseToUtilUserSession(session *UserSessions) *util.UserSession {
	responseUserSession := &util.UserSession{
		Uid:               session.Uid,
		UserName:          session.Username,
		UserRoleList:      session.UserRoleList,
		UserAuthorityList: session.UserAuthorityList,
		AddDatetime:       session.AddDatetime,
		UpdateDatetime:    session.UpdateDatetime,
	}

	return responseUserSession
}
