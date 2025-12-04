package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, JsonResponse{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, JsonResponse{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

// Fail 业务失败（HTTP 状态码永远 200，业务码自定义）
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, JsonResponse{
		Code: code,
		Msg:  msg,
	})
}

// FailWithHTTPCode 当真的需要返回 4xx/5xx 时使用（比如 401、404、500，少用！）
func FailWithHTTPCode(c *gin.Context, httpStatus, code int, msg string) {
	c.JSON(httpStatus, JsonResponse{
		Code: code,
		Msg:  msg,
	})
}
