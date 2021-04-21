package user_operation_log

import "github.com/livegoplayer/go_db_helper/mysql"

const PREFIX = "us_"

//
type UserOperationLog struct {
	Id          int64  `gorm:"column:id;PRIMARY_KEY" json:"id"`         //
	Cat         int64  `gorm:"column:cat" json:"cat"`                   // 0: 未定义类型的操作 1：用户相关 2：角色相关 3：权限相关
	Uid         int64  `gorm:"column:uid" json:"uid"`                   // 被操作者uid
	Message     string `gorm:"column:message" json:"message"`           // 日志记录
	OperatorUid int64  `gorm:"column:operator_uid" json:"operator_uid"` // 操作者uid
	AddDatetime string `gorm:"column:add_datetime" json:"add_datetime"` // 操作时间
}

func (UserOperationLog) TableName() string {
	return PREFIX + "user_operation_log"
}

type UserOperationLogQuery struct {
	mysql.Query
}

func (UserOperationLog) Connect() string {
	return "user"
}
