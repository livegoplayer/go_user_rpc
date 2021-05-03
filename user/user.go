package user

import (
	"errors"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/authority"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/ret_role_authority"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/ret_user_role"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/role"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/user"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/user_detail"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/user_operation_log"
	"github.com/livegoplayer/go_user_rpc/user/user_grpc"
	"strconv"
	"time"

	myHelper "github.com/livegoplayer/go_helper/utils"
	sessionHelper "github.com/livegoplayer/go_redis_helper"
)

func getUserSession(uid int64) (session myHelper.UserSession, exist bool) {
	key := getUserLoginStatusSessionKey(uid)
	val := sessionHelper.GetCacheData(key)
	if val != "" {
		exist = true
		myHelper.JsonDecode(val, session)
	}

	return
}

func setUserSession(uid int64, session *myHelper.UserSession) bool {
	key := getUserLoginStatusSessionKey(uid)
	sessionHelper.SetCacheData(key, myHelper.JsonEncode(session), time.Hour*24)
	return true
}

func clearUserSession(uid int64) {
	key := getUserLoginStatusSessionKey(uid)
	sessionHelper.RemoveCacheData(key)
}

func getUserLoginStatusSessionKey(uid int64) string {
	key := "user_login_status_session_uid_" + strconv.Itoa(int(uid))
	return key
}

//用户注册逻辑
func Register(userName string, password string) (uid int64, err error) {
	if user.CheckExistByUsername(userName) {
		err = errors.New("该用户名已经存在")
		return
	}

	return AddUser(userName, password, 0)
}

const (
	COMMON_USER int64 = iota + 1
	ADMINISTRATOR
)

//管理员添加用户逻辑
func AddUser(userName string, password string, operationUid int64) (uid int64, err error) {
	res := user.Save(&user.User{
		Username: userName,
		Password: password,
	})
	uid = res.Id
	if res.Id > 0 {
		user_operation_log.Save(&user_operation_log.UserOperationLog{
			Uid:         uid,
			OperatorUid: operationUid,
			Cat:         int64(USER_RET),
			Message:     "添加用户userName=" + userName,
		})

		if !AddUserRole(uid, COMMON_USER, operationUid) {
			err = errors.New("初始化用户权限失败")
		}
	}

	return
}

//管理员删除用户
func DelUser(uid int64, operationUid int64) (success bool, err error) {

	res := user.DeleteById(uid)
	if res == 0 {
		err = errors.New("删除用户失败")
		return
	}

	user_operation_log.Save(&user_operation_log.UserOperationLog{
		Uid:         uid,
		OperatorUid: operationUid,
		Cat:         int64(USER_RET),
		Message:     "删除用户uid=" + myHelper.AsString(uid),
	})

	if !DelUserAllRole(uid, operationUid) {
		err = errors.New("删除用户角色失败")
		return
	}

	success = true

	return
}

type OperationType int64

const (
	UNDEFINED_RET OperationType = iota
	USER_RET
	ROLE_RET
	AUTHORITY_RET
)

func AddUserRole(uid int64, roleId int64, operationUid int64) bool {
	if !ret_user_role.CheckExistByUidAndRoleId(uid, roleId) {
		userRole := &ret_user_role.RetUserRole{
			Uid:    uid,
			RoleId: roleId,
		}
		userRole = ret_user_role.Save(userRole)
		if userRole.Id > 0 {
			user_operation_log.Save(&user_operation_log.UserOperationLog{
				Uid:         uid,
				OperatorUid: operationUid,
				Cat:         int64(ROLE_RET),
				Message:     "添加用户角色role_id=" + myHelper.AsString(roleId) + " Id=" + myHelper.AsString(userRole.Id),
			})
		}

	}

	return true
}

func DelUserRole(uid int64, roleId int64, operationUid int64) bool {
	if ret_user_role.CheckExistByUidAndRoleId(uid, roleId) {
		res := ret_user_role.DeleteByUidAndRoleId(uid, roleId)
		if res > 0 {
			user_operation_log.Save(&user_operation_log.UserOperationLog{
				Uid:         uid,
				OperatorUid: operationUid,
				Cat:         int64(ROLE_RET),
				Message:     "delete用户角色role_id=" + myHelper.AsString(roleId),
			})
		}
	}

	return true
}

func DelUserAllRole(uid int64, operationUid int64) bool {
	res := ret_user_role.DeleteAllByUid(uid)
	if res > 0 {
		user_operation_log.Save(&user_operation_log.UserOperationLog{
			Uid:         uid,
			OperatorUid: operationUid,
			Cat:         int64(ROLE_RET),
			Message:     "清空用户角色",
		})
	}

	return true
}

func GetRoleList() (roleList map[int64]string) {
	roles := role.FetchRoleAll()
	roleList = make(map[int64]string)
	for _, r := range roles {
		roleList[r.Id] = r.RoleName
	}

	return roleList
}

func GetUserRoleList(uid int64) (userRoleList map[int64]string) {
	roles := ret_user_role.FetchByUid(uid)
	userRoleList = make(map[int64]string)
	for _, rs := range roles {
		r := role.GetOneById(rs.RoleId)
		if r == nil {
			continue
		}
		userRoleList[r.Id] = r.RoleName
	}
	return userRoleList
}

func Login(userName string, password string) (uid int64, session *user_grpc.UserSessions, tokenStr string, err error) {
	//第一步获取用户信息
	userSession := &myHelper.UserSession{}
	one := user.GetOneByUsernameAndPassword(userName, password)
	if one == nil {
		err = errors.New("用户名或者密码错误")
		return
	}

	//获取detail信息
	userSession = GetUserDetailInfo(one.Id)

	uid = one.Id
	//设置session
	setUserSession(uid, userSession)

	//生成新的token
	tokenStr, err = myHelper.CreateToken(userSession, "")

	session = InitUserSession(userSession)

	return
}

func GetUserDetailInfo(uid int64) *myHelper.UserSession {
	userDetailList := user_detail.GetUserAuthority(uid)

	roleList := make(map[int64]string)
	authorityList := make(map[int64]string)
	userSession := &myHelper.UserSession{}

	userInitFlg := false
	for _, userDetail := range userDetailList {

		if !userInitFlg {
			userSession.Uid = userDetail.Uid
			userSession.UserName = userDetail.Username
			userSession.AddDatetime = userDetail.AddDatetime
			userSession.UpdateDatetime = userDetail.UptDatetime
			userInitFlg = true
		}

		roleList[userDetail.RoleId] = userDetail.RoleName
		authorityList[userDetail.AuthorityId] = userDetail.AuthorityName
	}

	userSession.UserRoleList = roleList
	userSession.UserAuthorityList = authorityList

	return userSession
}

func InitUserSession(session *myHelper.UserSession) *user_grpc.UserSessions {
	return &user_grpc.UserSessions{
		Uid:               session.Uid,
		Username:          session.UserName,
		UserRoleList:      session.UserRoleList,
		UserAuthorityList: session.UserAuthorityList,
		AddDatetime:       session.AddDatetime,
		UpdateDatetime:    session.UpdateDatetime,
	}
}

func CheckUserStatus(token string) (isLogin bool, session *user_grpc.UserSessions, tokenStr string, err error) {
	claims, err := myHelper.ParseToken(token, "")
	tokenStr = token
	userSession := claims.UserSession
	if err != nil {
		userSession = myHelper.UserSession{}
		return
	}

	//检查session
	_, exsit := getUserSession(userSession.Uid)
	if exsit {
		isLogin = true
	} else {
		userSession = myHelper.UserSession{}
	}

	session = InitUserSession(&userSession)

	return
}

func CheckUserAuthority(uid int64, authorityId int64) (exists bool) {
	ret_role_authority.CountRetRoleAuthorityAll()
	mp := GetUserAuthorityMap(uid)
	if _, ok := mp[authorityId]; ok {
		return true
	}
	return false
}

func Logout(uid int64) (success bool) {
	clearUserSession(uid)
	return true
}

func GetUserAuthorityMap(uid int64) (userAuthorityList map[int64]string) {
	one := user.GetOneById(uid)
	if one == nil {
		return
	}

	roleIds := ret_user_role.FetchByUid(uid).PluckRoleId()
	if len(roleIds) == 0 {
		return
	}

	userAuthorityList = make(map[int64]string, 0)
	authIds := ret_role_authority.FetchByRoleIds(roleIds).PluckAuthorityId()

	list := authority.FetchByIds(authIds)
	for _, v := range list {
		userAuthorityList[v.Id] = v.AuthorityName
	}

	return
}

func GetAuthorityList() (authorityList map[int64]string) {
	authorityList = make(map[int64]string, 0)
	list := authority.FetchAuthorityAll()
	for _, v := range list {
		authorityList[v.Id] = v.AuthorityName
	}
	return
}

func GetUserList(page int64, size int64) (userList []*user_grpc.UserInfo, total int64) {
	userList = InitUserInfo(user.FetchUserAllWithPageSize(int(page), int(size)))
	total = user.CountUserAll()
	return
}

func InitUserInfo(users []user.User) []*user_grpc.UserInfo {
	userList := make([]*user_grpc.UserInfo, 0)
	for _, v := range users {
		userList = append(userList, &user_grpc.UserInfo{
			Id:          v.Id,
			Username:    v.Username,
			AddDataTime: v.AddDatetime,
			UptDatetime: v.UptDatetime,
		})
	}

	return userList
}
