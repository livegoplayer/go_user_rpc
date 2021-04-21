package user

import "github.com/livegoplayer/go_db_helper/mysql"

const PREFIX = "us_"

//
type User struct {
	Id          int64  `gorm:"column:id;PRIMARY_KEY" json:"id"`                                //
	Username    string `gorm:"column:username;unique_1_UNIQUE;unique_2_MULTI" json:"username"` // 用户名
	Password    string `gorm:"column:password;unique_1_UNIQUE" json:"password"`                // 密码
	AddDatetime string `gorm:"column:add_datetime" json:"add_datetime"`                        //
	UptDatetime string `gorm:"column:upt_datetime" json:"upt_datetime"`                        //
}

func (User) TableName() string {
	return PREFIX + "user"
}

type UserQuery struct {
	mysql.Query
}

func (User) Connect() string {
	return "user"
}
