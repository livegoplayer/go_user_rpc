package user

import (
	"errors"

	"go_user_rpc/model"
	sessionHelper "go_user_rpc/redisHelper"
	"go_user_rpc/util"
)

//定义一个接口，service包含的方法
type UserServiceInterFace interface {
	Login(userName string, password string) (uid int, token string, err error)
	CheckLoginStatus(uid int) (isLogin bool, err error)
	Register(userName string, password string) (uid int, err error)

	CheckUserAuthority(uid int, authorityId int) bool
	GetUserAuthorityList(uid int) map[int]string
	GetAuthorityList() map[int]string

	AddUserRole(uid int, roleId int, operationUid int) bool
	DelUserRole(uid int, roleId int, operationUid int) bool
	GetRoleList() map[int]string
	GetUserRoleList() map[int]string
}

type UserService struct{}

func getUserSession(uid int) (val int) {
	key := "user_login_status_session_uid_" + string(uid)
	res := sessionHelper.GetRedisKey(key)
	var err error

	val, err = res.Int()
	if err != nil {
		//todo
		return
	}

	return
}

//用户注册逻辑
func (userService *UserService) Register(userName string, password string) (uid int, err error) {
	isRecordFound, err := model.CheckUserName(userName)
	if isRecordFound {
		err = errors.New("该用户名已经存在")
		return
	}

	return model.AddNewUser(userName, password, 0)
}

//管理员添加用户逻辑
func (userService *UserService) AddUser(userName string, password string, operationUid int) (uid int, err error) {
	return model.AddNewUser(userName, password, operationUid)
}

func AddUserRole(uid int, roleId int, operationUid int) bool {
	return model.AddUserRole(uid, roleId, operationUid)
}

//不给用
func DelUserRole(uid int, roleId int, operationUid int) bool {
	return model.DelUserRole(uid, roleId, operationUid)
}

func GetRoleList() []*model.Role {
	roles := model.GetRoleList()
	return roles
}

func GetUserRoleList(uid int) []*model.Role {
	roles := model.GetUserRoleList(uid)
	return roles
}

func (userService *UserService) Login(userName string, password string) (uid int, tokenStr string, err error) {
	//第一步获取用户信息
	isRecordFound, user, err := model.CheckUserPassword(userName, password)
	if err != nil {
		return
	}

	if !isRecordFound {
		err = errors.New("用户名或者密码错误")
		return
	}

	//获取detail信息
	userSession, err := model.GetUserDetailInfo(user.ID)
	if err != nil {
		return
	}

	// 设置新的session给客户端派发token
	tokenStr, err = util.CreateToken(userSession)

	return
}

func (userService *UserService) CheckLoginStatus(token string) (isLogin bool, err error) {
	_, err = util.ParseToken(token)
	if err != nil {
		return
	}

	//todo 可能需要在这里增加过期token续期逻辑
	isLogin = true

	return
}
