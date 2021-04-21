package ret_user_role

import "github.com/livegoplayer/go_db_helper/mysql"

const PREFIX = "us_"

//
type RetUserRole struct {
	Id     int64 `gorm:"column:id;PRIMARY_KEY" json:"id"`                      //
	Uid    int64 `gorm:"column:uid;unique_1_UNIQUE;unique_2_MULTI" json:"uid"` // 用户id
	RoleId int64 `gorm:"column:role_id;unique_1_UNIQUE" json:"role_id"`        // 角色id
}

func (RetUserRole) TableName() string {
	return PREFIX + "ret_user_role"
}

type RetUserRoleQuery struct {
	mysql.Query
}

func (RetUserRole) Connect() string {
	return "user"
}
