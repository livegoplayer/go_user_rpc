package user_detail

import (
	"github.com/livegoplayer/go_db_helper/mysql"
)

const PREFIX = "us_"

type UserDetail struct {
	Uid           int64  `json:"uid"`
	Username      string `json:"username"`
	AddDatetime   string `json:"add_datetime"`
	UptDatetime   string `json:"upt_datetime"`
	RoleId        int64  `json:"role_id"`
	RoleName      string `json:"role_name"`
	AuthorityId   int64  `json:"authority_id"`
	AuthorityName string `json:"authority_name"`
}

func (UserDetail) TableName() string {
	return PREFIX + "user"
}

type UserQuery struct {
	mysql.Query
}

func (UserDetail) Connect() string {
	return "user"
}
