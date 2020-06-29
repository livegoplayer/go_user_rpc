package helper

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Resp 返回
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type EmptyData struct {
}

// ErrorResp 错误返回值
func ErrorResp(ctx *gin.Context, code int, msg string, data ...interface{}) {
	resp(ctx, code, msg, data...)
}

// SuccessResp 正确返回值
func SuccessResp(ctx *gin.Context, msg string, data ...interface{}) {
	resp(ctx, 0, msg, data...)
}

// resp 返回
func resp(ctx *gin.Context, code int, msg string, data ...interface{}) {
	resp := Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	if len(data) == 1 {
		resp.Data = data[0]
	}

	if len(data) == 0 {
		resp.Data = &EmptyData{}
	}

	ctx.JSON(http.StatusOK, resp)
}

//负责调用panic触发外部的panic处理函数
func CheckError(error error, message ...string) {
	var msg string

	if len(message) == 0 {
		msg = ""
	}

	if error != nil {
		msg = strings.Join(message, " ")
		msg = msg + " error:" + error.Error()
		error = NewError(200, 1, msg)
		panic(error)
	}
}
