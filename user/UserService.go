package user

import (
	"errors"
	"time"

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

func getUserSession(uid int) (session *util.UserSession, exist bool) {
	key := getUserLoginStatusSessionKey(uid)
	res := sessionHelper.GetRedisKey(key)
	val := res.String()
	if val != "" {
		exist = true
		util.JsonDecode(util.StringToBytes(val), session)
	}

	return
}

func setUserSession(uid int, session *util.UserSession) bool {
	key := getUserLoginStatusSessionKey(uid)
	_ = sessionHelper.SetRedisKey(key, util.JsonEncode(session), time.Hour*24)
	return true
}

func getUserLoginStatusSessionKey(uid int) string {
	key := "user_login_status_session_uid_" + string(uid)
	return key
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

//管理员添加用户逻辑
func (userService *UserService) DelUser(uid int, operationUid int) (success bool, err error) {
	return model.DelUser(uid, operationUid)
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

func (userService *UserService) Login(userName string, password string, host string) (uid int, userSession *util.UserSession, tokenStr string, err error) {
	exists := checkHost(host)
	if !exists {
		err = errors.New("非法域名")
		return
	}

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
	userSession, err = model.GetUserDetailInfo(user.ID)
	if err != nil {
		return
	}

	//设置session
	setUserSession(uid, userSession)

	//生成新的token
	tokenStr, err = util.CreateToken(userSession, host)

	return
}

func (userService *UserService) CheckLoginStatus(token string, host string) (isLogin bool, tokenStr string, err error) {
	//todo 分离出去

	exists := checkHost(host)
	if !exists {
		err = errors.New("非法域名")
		return
	}

	claims, err := util.ParseToken(token, host)
	if err != nil {
		//如果token过期了
		if err.Error() == "jwt过期" || err.Error() == "host错误" {
			userSession := claims.UserSession
			//检查session是否过期
			userSessionNew, exsit := getUserSession(userSession.Uid)
			if exsit {
				// 重新根据当前的session生成token
				// 生成新的token
				isLogin = true
				tokenStr, err = util.CreateToken(userSessionNew, host)
			}
		}

		return
	}

	isLogin = true

	return
}

func checkHost(host string) (exists bool) {
	validHostList := []string{"127.0.0.1"}
	exists, _ = util.InArray(host, validHostList)

	return
}
