package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// reponse 基础响应体结构
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Success 成功响应
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Result{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// Error 失败响应
// httpCode: HTTP 状态码
// errCode: 业务错误码
// message: 错误描述
func Error(c *gin.Context, httpCode int, errCode int, message string) {
	c.JSON(httpCode, Result{
		Code:    errCode,
		Message: message,
		Data:    nil,
	})
}

// ErrorWithMsg 只传业务错误码,使用预定义的消息
func ErrorWithMsg(c *gin.Context, errCode int, message string) {
	c.JSON(http.StatusOK, Result{
		Code:    errCode,
		Message: message,
		Data:    nil,
	})
}
