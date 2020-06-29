package logger

import (
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//文件log单例
var fileLogger *logrus.Logger

//完全日志输出的模式
func GetFileLog() *logrus.Logger {
	if fileLogger != nil {
		return fileLogger
	}

	//设置对某个地方的日志进行分割
	errorWriter, err := rotatelogs.New(
		"../logs/error_%Y%m%d%H%M.log.",
		//每次从这个位置清除分离日志
		rotatelogs.WithLinkName("../logs/error.log"),
		//每30天清除一次日志
		rotatelogs.WithMaxAge(30*24*time.Hour),
		//每一个小时分离一次日志文件
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err != nil {
		//输出到控制台
		_, _ = fmt.Fprint(os.Stdin, "error:"+err.Error())
	}

	//设置对某个地方的日志进行分割
	accessWriter, err := rotatelogs.New(
		"../logs/access_%Y%m%d%H%M.log",
		//每次从这个位置清除分离日志
		rotatelogs.WithLinkName("../logs/access.log"),
		//每30天清除一次日志
		rotatelogs.WithMaxAge(30*24*time.Hour),
		//每一个小时分离一次日志文件
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err != nil {
		//输出到控制台
		_, _ = fmt.Fprint(os.Stdin, "error:"+err.Error())
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.TraceLevel: os.Stdin, // 为不同级别设置不同的输出目的,这些都是ioWriter
		logrus.DebugLevel: os.Stdin, // 为不同级别设置不同的输出目的,这些都是ioWriter
		logrus.InfoLevel:  accessWriter,
		logrus.WarnLevel:  accessWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.FatalLevel: errorWriter,
		logrus.PanicLevel: errorWriter,
	}, &logrus.TextFormatter{})

	fileLogger = logrus.New()
	fileLogger.AddHook(lfHook)

	return fileLogger
}
