package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{}
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}

func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	})
}

func Fail(c *gin.Context, code int) {
	c.JSON(code, Result{
		Code: code,
		Msg:  "fail",
		Data: nil,
	})
}

func FailWithMsg(c *gin.Context, code int, msg string) {
	c.JSON(code, Result{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
