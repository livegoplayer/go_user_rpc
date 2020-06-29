package helper

//处理全局panic的返回值，重写gin.Recover中间件的内容
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 错误处理的结构体
type Error struct {
	//只是说他们两个差不多，没用到
	Resp       `json:"-"`
	StatusCode int         `json:"-"`
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

var (
	ServerError = NewError(http.StatusInternalServerError, 1, "系统异常，请稍后重试!")
	NotFound    = NewError(http.StatusNotFound, 1, http.StatusText(http.StatusNotFound))
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
		Data:       &EmptyData{},
	}
}

// 404处理
func HandleNotFound(c *gin.Context) {
	err := NotFound
	c.JSON(err.StatusCode, err)
	return
}

// 服务异常处理
func HandleServerError(c *gin.Context) {
	err := ServerError
	c.JSON(err.StatusCode, err)
	return
}

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *Error
				if e, ok := err.(*Error); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = OtherError(e.Error())
				} else {
					Err = ServerError
				}
				// 记录一个错误的日志
				c.JSON(Err.StatusCode, Err)
				fmt.Print()
				return
			}
		}()
		c.Next()
	}
}
