package role

import "github.com/livegoplayer/go_db_helper/mysql"

const PREFIX = "us_"

//
type Role struct {
	Id          int64  `gorm:"column:id;PRIMARY_KEY" json:"id"`         //
	RoleName    string `gorm:"column:role_name" json:"role_name"`       // 角色名称
	RoleLevel   int64  `gorm:"column:role_level" json:"role_level"`     // 角色等级
	AddDatetime string `gorm:"column:add_datetime" json:"add_datetime"` //
	UptDatetime string `gorm:"column:upt_datetime" json:"upt_datetime"` //
}

func (Role) TableName() string {
	return PREFIX + "role"
}

type RoleQuery struct {
	mysql.Query
}

func (Role) Connect() string {
	return "user"
}
