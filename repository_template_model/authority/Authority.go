package authority

import "github.com/livegoplayer/go_db_helper/mysql"

const PREFIX = "us_"

//
type Authority struct {
	Id            int64  `gorm:"column:id;PRIMARY_KEY" json:"id"`             //
	AuthorityName string `gorm:"column:authority_name" json:"authority_name"` // 权限名称
}

func (Authority) TableName() string {
	return PREFIX + "authority"
}

type AuthorityQuery struct {
	mysql.Query
}

func (Authority) Connect() string {
	return "user"
}
