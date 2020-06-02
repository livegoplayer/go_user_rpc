package model

import (
	"errors"

	"go_user_rpc/dbHelper"
	"go_user_rpc/util"
)

type operationType int

//操作类型
const (
	UNDEFINED_RET operationType = iota
	USER_RET
	ROLE_RET
	AUTHORITY_RET
)

const (
	COMMON_USER   int = 1
	ADMINISTRATOR int = iota
)

type User struct {
	Model
	ID       int    `gorm:"column:uid"` //会被自动认为是主键
	Name     string `gorm:"column:username"`
	Password string
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *User) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "users"
}

type Role struct {
	Model
	ID             int    //会被自动认为是主键
	Name           string `gorm:"column:role_name"`
	Level          int    `gorm:"column:role_level"`
	AddDatetime    int    `gorm:"column:add_datetime"`
	UpdateDatetime int    `gorm:"column:update_datetime"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *Role) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "role"
}

type Authority struct {
	Model
	ID   int    //会被自动认为是主键
	Name string `gorm:"column:authority_name"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *Authority) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "role"
}

type RetRoleAuthority struct {
	ID          int //会被自动认为是主键
	RoleId      int `gorm:"column:role_id"`
	AuthorityId int `gorm:"column:authority_id"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *RetRoleAuthority) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "ret_role_authority"
}

type RetUserRole struct {
	ID     int //会被自动认为是主键
	UserId int `gorm:"column:uid"`
	RoleId int `gorm:"column:role_id"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *RetUserRole) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "ret_user_role"
}

type UserOperationLog struct {
	ID          int           //会被自动认为是主键
	Category    operationType `gorm:"column:cat"`
	Uid         int           `gorm:"column:uid"`
	Message     string        `gorm:"column:message"`
	OperatorUid int           `gorm:"column:operator_uid"`
	AddDatetime int           `gorm:"column:add_datetime"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *UserOperationLog) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "user_opeartion_log"
}

func init() {
	// 需要在init中注册定义的model fs_user
}

//role 相关
func GetRoleList() (roles []*Role) {

	db := dbHelper.GetDB()

	db = db.Limit(10).Take(&roles)

	if db.Error != nil {
		//todo
		return
	}
	return
}

func GetUserRoleList(uid int) (roles []*Role) {
	var userRoles []*RetUserRole
	db := dbHelper.GetDB()

	db = db.Where("uid=?", uid).Limit(10).Take(&userRoles)
	if db.Error != nil {
		//todo
		return
	}

	var roleIdList []int
	for _, userRole := range userRoles {
		roleIdList = append(roleIdList, userRole.RoleId)
	}

	db = db.Where("role_id in (?)", roleIdList).Limit(10).Take(&roles)
	if db.Error != nil {
		//todo
		return
	}
	return
}

func AddUserRole(uid int, roleId int, operationUid int) (success bool) {
	db := dbHelper.GetDB()

	//todo 加入事务操作
	retUserRole := &RetUserRole{UserId: uid, RoleId: roleId}

	db = db.Create(retUserRole)
	if db.Error != nil {
		//todo
		return
	}

	_ = addOperationLog(ROLE_RET, operationUid, uid, "删除用户角色")

	success = true
	return success
}

func DelUserRole(uid int, roleId int, operationUid int) (success bool) {
	db := dbHelper.GetDB()

	//todo 加入事务操作
	retUserRole := &RetUserRole{UserId: uid, RoleId: roleId}

	db.Delete(retUserRole)
	if db.Error != nil {
		//todo
		return
	}

	_ = addOperationLog(ROLE_RET, operationUid, uid, "删除用户角色")

	success = true
	return success
}

func addOperationLog(category operationType, operationUid int, targetUid int, message string) error {
	db := dbHelper.GetDB()

	userOperationLog := &UserOperationLog{Category: category, Uid: targetUid, Message: message, OperatorUid: operationUid}

	db.Create(userOperationLog)
	if db.Error != nil {
		//todo
		return db.Error
	}

	return nil
}

//添加用户操作
func AddNewUser(name string, password string, operationUid int) (uid int, err error) {
	db := dbHelper.GetDB()

	//todo 加入事务操作
	User := &User{Name: name, Password: password}

	db = db.Create(User)
	if db.Error != nil {
		err = db.Error
		return
	}

	var id []int
	db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)
	if db.Error != nil {
		err = db.Error
		return
	}
	uid = int(id[0])

	//给用户添加权限
	db = dbHelper.GetDB()
	_ = addOperationLog(USER_RET, operationUid, uid, "增加用户")

	res := AddUserRole(uid, COMMON_USER, operationUid)
	if !res {
		err = errors.New("初始化用户权限失败")
	}

	return
}

//删除用户操作
func DelUser(uid int, operationUid int) (success bool) {
	db := dbHelper.GetDB()

	success = false
	//todo 加入事务操作
	User := &User{ID: uid}

	db = db.Delete(User)
	if db.Error != nil {
		//todo
		return
	}

	if db.RowsAffected > 0 {
		success = true
	}

	_ = addOperationLog(ROLE_RET, operationUid, uid, "删除用户")

	return success
}

//检查用户名密码
func CheckUserPassword(username string, password string) (isRecordFound bool, user *User, err error) {
	db := dbHelper.GetDB()
	user = &User{}

	db = db.Where("username = ?", username).Where("password = ?", password).First(&user)
	if db.Error != nil {
		return
	}

	isRecordFound = !db.RecordNotFound()

	return
}

//检查用户名是否存在
func CheckUserName(username string) (isRecordFound bool, err error) {
	db := dbHelper.GetDB()
	user := &User{}

	db = db.Where("username = ?", username).Find(&user)
	if db.Error != nil {
		return
	}

	isRecordFound = !db.RecordNotFound()

	return
}

/**
获取用户的权限列表
*/
func GetUserDetailInfo(uid int) (userSession *util.UserSession, err error) {
	db := dbHelper.GetDB()

	var res []map[string]interface{}
	db = db.Table("user").Select("fs_user.uid as uid, fs_user.username as username, fs_user.add_datetime as add_datetime, fs_user.upt_datetime as upt_datetime, fs_role.id as role_id, fs_role.role_name as role_name, fs_authority.id as authority_id, fs_authority.authority_name as authority_name").Joins("LEFT JOIN fs_ret_user_role as a ON fs_ret_user_role.uid = users.id").Joins("LEFT JOIN fs_role as c ON b.role_id = c.id").Joins("LEFT JOIN fs_ret_role_authority as d ON c.id = d.role_id").Joins("LEFT JOIN fs_authority as e ON d.authority_id = e.id").Where("user.uid = ?", uid).Scan(&res)
	if db.Error != nil {
		return
	}

	roleList := make(map[int]string)
	authorityList := make(map[int]string)
	userSession = &util.UserSession{}

	userInitFlg := false
	for _, info := range res {
		roleId := util.Int(info["role_id"])
		roleName := util.String(info["role_name"])
		authorityId := util.Int(info["authority_id"])
		authorityName := util.String(info["authority_name"])

		if !userInitFlg {
			uid := util.Int(info["uid"])
			username := util.String(info["username"])
			addDatetime := util.String(info["add_datetime"])
			uptDatetime := util.String(info["upt_datetime"])
			userSession.Uid = uid
			userSession.UserName = username
			userSession.AddDatetime = addDatetime
			userSession.UpdateDatetime = uptDatetime
			userInitFlg = true
		}

		roleList[roleId] = roleName
		authorityList[authorityId] = authorityName
	}

	userSession.UserRoleList = roleList
	userSession.UserAuthorityList = authorityList

	return
}
