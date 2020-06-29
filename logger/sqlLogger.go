package logger

import (
	"github.com/sirupsen/logrus"
)

var mysqlLogger *logrus.Logger

func getMysqlLogger() *logrus.Logger {
	if mysqlLogger != nil {
		return mysqlLogger
	}

	//初始化配置文件
	SwitchLogModeFmtByType(MYSQL)
	//根据配置文件初始化logger
	mysqlLogger = setDefaultLogger()

	return mysqlLogger
}
