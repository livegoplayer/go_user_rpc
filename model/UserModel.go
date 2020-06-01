package model

import (
	"go_user_rpc/dbHelper"
)

type operation_type int

const (
	UNDEFINED_RET operation_type = iota
	USER_RET
	ROLE_RET
	AUTHORITY_RET
)

type User struct {
	Model
	ID             int    `gorm:"column:uid"` //会被自动认为是主键
	Name           string `gorm:"column:username"`
	Password       string
	AddDatetime    int `gorm:"column:add_datetime;-"`
	UpdateDatetime int `gorm:"column:update_datetime;-"`
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
	ID          int            //会被自动认为是主键
	Category    operation_type `gorm:"column:cat"`
	Uid         int            `gorm:"column:uid"`
	Message     string         `gorm:"column:message"`
	OperatorUid int            `gorm:"column:operator_uid"`
	AddDatetime int            `gorm:"column:add_datetime"`
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

func addOperationLog(category operation_type, operationUid int, targetUid int, message string) error {
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

	_ = addOperationLog(USER_RET, operationUid, uid, "增加用户")

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
