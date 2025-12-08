package controller

import (
	"errors"
	"net/http"
	"strconv"

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

// FindByUsername
// 根据用户名查询用户信息
func (ctl UserController) FindByUsername(c *gin.Context) {
	// 1.从请求中获取用户名
	username := c.Query("username")
	if username == "" {
		response.ErrorWithMsg(c, errcode.IllegalParams.Code, "Username cannot be empty.")
		return
	}

	// 2.调用service层方法
	user, err := ctl.service.FindByUsername(username)
	if err != nil {
		if errors.Is(err, errcode.NotFound) {
			response.ErrorWithMsg(c, errcode.NotFound.Code, "用户名不存在")
			return
		}
		response.ErrorWithMsg(c, errcode.InternalServerError.Code, err.Error())
		return
	}
	// 3.返回查询结果
	response.Success(c, user)
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
		defer c.Set("msg", err.Error())
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

// Delete 根据用户id删除用户信息
func (ctl UserController) Delete(c *gin.Context) {
	// 1.从 请求中获取用户id
	idStr := c.Param("id")

	// 2.转换成 uint64 格式
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ErrorWithMsg(c, errcode.IllegalParams.Code, "请输入数字")
		return
	}
	// 3.调用 service 层
	if err := ctl.service.Delete(id); err != nil {
		response.ErrorWithMsg(c, errcode.Exists.Code, "请输入正确的id")
		return
	}

	// 4.返回成功信息
	response.Success(c, nil)
}
