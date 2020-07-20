package model

import (
	"errors"

	myHelper "github.com/livegoplayer/go_helper"

	"github.com/livegoplayer/go_user_rpc/dbHelper"
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
	COMMON_USER   int32 = 1
	ADMINISTRATOR int32 = iota
)

type User struct {
	Model
	ID       int32  `gorm:"column:id"` //会被自动认为是主键
	Name     string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *User) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "user"
}

type Role struct {
	Model
	ID    int32  //会被自动认为是主键
	Name  string `gorm:"column:role_name"`
	Level int32  `gorm:"column:role_level"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *Role) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "role"
}

type Authority struct {
	Model
	ID   int32  //会被自动认为是主键
	Name string `gorm:"column:authority_name"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *Authority) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "role"
}

type RetRoleAuthority struct {
	ID          int32 //会被自动认为是主键
	RoleId      int32 `gorm:"column:role_id"`
	AuthorityId int32 `gorm:"column:authority_id"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *RetRoleAuthority) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "ret_role_authority"
}

type RetUserRole struct {
	ID     int32 //会被自动认为是主键
	UserId int32 `gorm:"column:uid"`
	RoleId int32 `gorm:"column:role_id"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *RetUserRole) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "ret_user_role"
}

type UserOperationLog struct {
	ID          int32         //会被自动认为是主键
	Category    operationType `gorm:"column:cat"`
	Uid         int32         `gorm:"column:uid"`
	Message     string        `gorm:"column:message"`
	OperatorUid int32         `gorm:"column:operator_uid"`
	AddDatetime int32         `gorm:"column:add_datetime;-"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u *UserOperationLog) TableName() string {
	//绑定MYSQL表名为users
	return prefix + "user_opeartion_log"
}

func init() {
	// 需要在init中注册定义的model us_user
}

//role 相关
func GetRoleList() (roles []*Role) {

	db := dbHelper.GetDB()

	db = db.Limit(10).Find(&roles)

	if db.Error != nil {
		//todo
		return
	}
	return
}

func GetUserRoleList(uid int32) (roles []*Role) {
	var userRoles []*RetUserRole
	db := dbHelper.GetDB()

	db = db.Where("uid=?", uid).Limit(10).Find(&userRoles)
	if db.Error != nil {
		//todo
		return
	}

	var roleIdList []int32
	for _, userRole := range userRoles {
		roleIdList = append(roleIdList, userRole.RoleId)
	}

	db = dbHelper.GetDB()
	db = db.Where("id in (?)", roleIdList).Limit(10).Find(&roles)
	if db.Error != nil {
		//todo
		return
	}
	return
}

func AddUserRole(uid int32, roleId int32, operationUid int32) (success bool) {
	db := dbHelper.GetDB()

	//todo 加入事务操作
	retUserRole := &RetUserRole{UserId: uid, RoleId: roleId}

	db = db.Create(retUserRole)
	if db.Error != nil {
		//todo
		return
	}

	_ = addOperationLog(ROLE_RET, operationUid, uid, "添加用户角色")

	success = true
	return success
}

func CheckUserRole(uid int32, roleId int32) bool {
	db := dbHelper.GetDB()

	//todo 加入事务操作
	retUserRole := &RetUserRole{}

	//查出任意一条
	db = db.Where("uid = ? and role_id = ?", uid, roleId).Take(retUserRole)

	if retUserRole.ID != 0 {
		return true
	}

	return false
}

func DelUserRole(uid int32, roleId int32, operationUid int32) (success bool) {
	db := dbHelper.GetDB()

	//todo 加入事务操作
	retUserRole := &RetUserRole{}

	db.Where("uid = ? and role_id = ?", uid, roleId).Delete(retUserRole)
	if db.Error != nil {
		//todo
		return
	}

	_ = addOperationLog(ROLE_RET, operationUid, uid, "删除用户角色")

	success = true
	return success
}

func addOperationLog(category operationType, operationUid int32, targetUid int32, message string) error {
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
func AddNewUser(name string, password string, operationUid int32) (uid int32, err error) {
	db := dbHelper.GetDB()

	//todo 加入事务操作
	User := &User{Name: name, Password: password}

	db = db.Create(User)
	if db.Error != nil {
		err = db.Error
		return
	}

	db = dbHelper.GetDB()
	var id []int
	db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)
	if db.Error != nil {
		err = db.Error
		return
	}
	uid = int32(id[0])

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
func DelUser(uid int32, operationUid int32) (success bool, err error) {
	db := dbHelper.GetDB()

	success = false
	//todo 加入事务操作
	User := &User{}

	db = db.Where("id=?", uid).Delete(User)
	if db.Error != nil {
		err = db.Error
		return
	}

	if db.RowsAffected == 0 {
		err = errors.New("删除用户失败")
		return
	}

	//删除用户相关的role
	retUserRole := &RetUserRole{UserId: uid}
	db = db.Delete(retUserRole)
	if db.RowsAffected == 0 {
		err = errors.New("删除用户失败")
		return
	}

	_ = addOperationLog(ROLE_RET, operationUid, uid, "删除用户")

	return
}

func GetUserList(page int32, size int32) (userList []*User) {
	if size == 0 {
		size = 10
	}

	if page == 0 {
		page = 1
	}

	offset := (page - 1) * size

	db := dbHelper.GetDB()
	db = db.Limit(size).Offset(offset).Find(&userList)

	return
}

func GetUserTotal() (total int32) {
	db := dbHelper.GetDB()
	db = db.Model(&User{}).Count(&total)

	return
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

type UserDetail struct {
	Uid           int32  `json:"uid"`
	Username      string `json:"username"`
	AddDatetime   string `json:"add_datetime"`
	UptDatetime   string `json:"upt_datetime"`
	RoleId        int32  `json:"role_id"`
	RoleName      string `json:"role_name"`
	AuthorityId   int32  `json:"authority_id"`
	AuthorityName string `json:"authority_name"`
}

/**
获取用户的权限列表
*/
func GetUserDetailInfo(uid int32) (userSession *myHelper.UserSession, err error) {
	db := dbHelper.GetDB()

	db = db.LogMode(false)

	var userDetailList []UserDetail
	db = db.Table("us_user").Select("us_user.id as uid, us_user.username as username, us_user.add_datetime as add_datetime, us_user.upt_datetime as upt_datetime, c.id as role_id, c.role_name as role_name, e.id as authority_id, e.authority_name as authority_name").Joins("LEFT JOIN us_ret_user_role as a ON a.uid = us_user.id").Joins("LEFT JOIN us_role as c ON a.role_id = c.id").Joins("LEFT JOIN us_ret_role_authority as d ON c.id = d.role_id").Joins("LEFT JOIN us_authority as e ON d.authority_id = e.id").Where(
		"us_user.id = ?", uid).Find(&userDetailList)
	if db.Error != nil {
		return
	}

	roleList := make(map[int32]string)
	authorityList := make(map[int32]string)
	userSession = &myHelper.UserSession{}

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

	return
}

//检查是否用户是否有相关权限
func CheckUserAuthority(uid int32, authorityId int32) (exist bool, err error) {
	db := dbHelper.GetDB()

	var res []map[string]interface{}
	db = db.Table("user").Select("us_user.id as uid, us_ret_role_authority.authority_id as authority_id").Joins("LEFT JOIN us_ret_user_role as a ON us_ret_user_role.uid = users.id").Joins("LEFT JOIN us_role as c ON b.role_id = c.id").Joins("LEFT JOIN us_ret_role_authority as d ON c.id = d.role_id").Where("user.uid = ?", uid).Where("authorityId = ?", authorityId).Find(&res)
	if db.Error != nil {
		err = db.Error
		return
	}

	if !db.RecordNotFound() {
		exist = true
		return
	}

	return
}

//检查是否用户是否有相关权限
func GetUserAuthorityList(uid int32) (authorityList map[int32]string, err error) {
	db := dbHelper.GetDB()

	var res []map[string]interface{}
	db = db.Table("user").Select("us_authority.id as authority_id, us_authority.authority_name as authority_name").Joins("LEFT JOIN us_ret_user_role as a ON us_ret_user_role.uid = users.id").Joins("LEFT JOIN us_role as c ON b.role_id = c.id").Joins("LEFT JOIN us_ret_role_authority as d ON c.id = d.role_id").Joins("LEFT JOIN us_authority as e ON d.authority_id = e.id").Where("user.uid = ?", uid).Scan(&res)
	if db.Error != nil {
		err = db.Error
		return
	}

	authorityList = make(map[int32]string)

	for _, info := range res {
		authorityId := myHelper.Int32(info["authority_id"])
		authorityName := myHelper.String(info["authority_name"])
		authorityList[authorityId] = authorityName
	}

	return
}

//检查是否用户是否有相关权限
func GetAuthorityList() (authorityList map[int32]string, err error) {
	db := dbHelper.GetDB()

	var res []map[string]interface{}
	db = db.Select("us_authority.id as authority_id, us_authority.authority_name as authority_name").Scan(&res)
	if db.Error != nil {
		err = db.Error
		return
	}

	authorityList = make(map[int32]string)

	for _, info := range res {
		authorityId := myHelper.Int32(info["authority_id"])
		authorityName := myHelper.String(info["authority_name"])
		authorityList[authorityId] = authorityName
	}

	return
}
