package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 封装响应
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// 成功响应
func Success(data any, m string, c *gin.Context) {
	msg := "成功"
	if len(m) > 0 {
		msg = m
	}
	Result(200, data, msg, c)
}

// 业务报错响应
func Fail(data any, m string, c *gin.Context) {
	Result(500, data, m, c)
}

func FailWithCode(code int, c *gin.Context) {
	msg, ok := ErrorMap[ErrorCode(code)]
	if ok {
		Result(code, map[string]any{}, msg, c)
		return
	}
	Result(code, map[string]any{}, "未知错误", c)
}
