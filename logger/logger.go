package logger

import (
	"github.com/sirupsen/logrus"
)

//这是一个公用方法
func setDefaultLogger() *logrus.Logger {
	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = Config.Out

	//设置日志级别
	logger.SetLevel(Config.Level)

	//设置日志格式
	logger.SetFormatter(Config.Format)

	return logger
}
