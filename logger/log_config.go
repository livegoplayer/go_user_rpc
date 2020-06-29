package logger

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type logConfig struct {
	Format *logrus.TextFormatter
	Out    io.Writer
	LType  logType
	Level  logrus.Level
}

type logType uint32

const (
	CONSOLE logType = iota
	FILE
	MYSQL
)

//默认初始化
func init() {
	SwitchLogModeFmtByType(MYSQL)
}

var (
	Config *logConfig = &logConfig{}
)

//根据传进来的logType初始化log模式
func SwitchLogModeFmtByType(t logType) {
	//默认的初始化
	Config.Level = logrus.InfoLevel
	switch t {
	case CONSOLE:
		ChangeToConsoleMode()
	case FILE:
		ChangeToFileMode()
	case MYSQL:
		ChangeToMysqlMode()
	}
}

//初始化输出模式为文件
func ChangeToFileMode() {
	Config.LType = FILE
	logFilePath := "/../logs"
	logFileName := "/../file_store_log"

	//日志文件目录
	fileName := path.Join(logFilePath, logFileName)

	//定位文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("配置日志文件出错", err)
	}
	Config.Format = &logrus.TextFormatter{}
	Config.Out = src
}

//初始化输出模式为mysql
func ChangeToMysqlMode() {
	Config.LType = MYSQL
	out := &MysqlWriter{}
	Config.Format = &logrus.TextFormatter{}
	Config.Out = out
}

//初始化输出模式为console
func ChangeToConsoleMode() {
	Config.LType = CONSOLE
	Config.Format = &logrus.TextFormatter{}
	Config.Out = os.Stdin
}
