package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/pkg/errcode"
	"github.com/zhang/bms/internal/service"
	"github.com/zhang/bms/pkg/response"
)

type UserController struct {
	service service.IUserService
}

func NewUserController(service service.IUserService) UserController {
	return UserController{
		service: service,
	}
}

// Insert
// 添加用户
func (ctl UserController) Insert(c *gin.Context) {
	// 1.从请求中获取参数
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		// 参数错误
		response.Error(c,
			http.StatusBadRequest,
			errcode.IllegalParams.Code,
			errcode.IllegalParams.Message,
		)
		return
	}

	// 2.调用 service 层
	if err := ctl.service.Insert(&user); err != nil {
		// 账号已存在
		if errors.Is(err, errcode.Exists) {
			response.Error(c,
				http.StatusBadRequest,
				errcode.Exists.Code,
				errcode.Exists.Message,
			)
			return
		}
		// 数据库内部错误
		response.Error(c,
			http.StatusInternalServerError,
			errcode.InternalServerError.Code,
			errcode.InternalServerError.Message,
		)
		return
	}

	// 3.调用成功
	response.Success(c, nil)
}
