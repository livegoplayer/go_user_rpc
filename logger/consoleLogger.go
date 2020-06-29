package logger

import (
	"github.com/sirupsen/logrus"
)

var consoleLogger *logrus.Logger

func getConsoleLogger() *logrus.Logger {
	if consoleLogger != nil {
		return consoleLogger
	}

	//初始化配置文件
	SwitchLogModeFmtByType(CONSOLE)
	//根据配置文件初始化logger
	consoleLogger = setDefaultLogger()

	return consoleLogger
}
