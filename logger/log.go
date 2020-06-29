package logger

import (
	"github.com/sirupsen/logrus"
)

var ltype logType
var Logger *logrus.Logger

//主动log的一些方法
func Panic(message string) {
	Logger = GetLogger()
	text, _ := logrus.PanicLevel.MarshalText()
	Logger.Panicf("{{level,%s},{levelNo,%s},{message:%s}",
		text, logrus.PanicLevel, message)
}

func Fatal(message string) {
	Logger = GetLogger()
	text, _ := logrus.FatalLevel.MarshalText()
	Logger.Fatalf("{{level,%s},{levelNo,%s},{message:%s},{clientIP,%15s}",
		text, logrus.FatalLevel, message)
}

func Error(message string) {
	Logger = GetLogger()
	text, _ := logrus.ErrorLevel.MarshalText()
	Logger.Errorf("{{level,%s},{levelNo,%s},{message:%s},{clientIP,%15s}",
		text, logrus.ErrorLevel, message)
}

func Warning(message string) {
	Logger = GetLogger()
	text, _ := logrus.WarnLevel.MarshalText()
	Logger.Warnf("{{level,%s},{levelNo,%s},{message:%s},{clientIP,%15s}",
		text, logrus.WarnLevel, message)
}

func Info(message string) {
	Logger = GetLogger()
	text, _ := logrus.InfoLevel.MarshalText()
	Logger.Infof("{{level,%s},{levelNo,%s},{message:%s},{clientIP,%15s}",
		text, logrus.InfoLevel, message)
}

func Debug(message string) {
	Logger = GetLogger()
	text, _ := logrus.DebugLevel.MarshalText()
	Logger.Debugf("{{level,%s},{levelNo,%s},{message:%s},{clientIP,%15s}",
		text, logrus.DebugLevel, message)
}

func Trace(message string) {
	Logger = GetLogger()
	text, _ := logrus.TraceLevel.MarshalText()
	Logger.Tracef("{{level,%s},{levelNo,%s},{message:%s},{clientIP,%15s}",
		text, logrus.TraceLevel, message)
}

func GetLogger() *logrus.Logger {
	if ltype == Config.LType && Logger != nil {
		return Logger
	}
	switch Config.LType {
	case CONSOLE:
		Logger = getConsoleLogger()
	case MYSQL:
		Logger = getMysqlLogger()
	case FILE:
		Logger = GetFileLog()
	}

	return Logger
}
