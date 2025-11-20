package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/model"
	"github.com/zhang/bms/internal/service"
	"github.com/zhang/bms/internal/transport/http/response"
)

type IUserHandler interface {
	FindById(c *gin.Context)
	Create(context *gin.Context)
}

type UserHandlerImpl struct {
	s service.IUserService
}

func NewUserHandler(s service.IUserService) IUserHandler {
	return &UserHandlerImpl{s: s}
}

func (h *UserHandlerImpl) FindById(c *gin.Context) {
	// 1.从请求中获取参数Id
	strId := c.Query("id")
	// 2.转换成 uint64 类型
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		fmt.Println("输入的不是数字")
		response.FailWithMsg(c, http.StatusBadRequest, "请输入数字")
		return
	}

	// 3.调用 service 层
	user, err := h.s.FindById(id)

	if err != nil {
		// 记录日志（实际项目中建议用 logger 库）
		fmt.Printf("查询用户失败: %v\n", err)
		response.FailWithMsg(c, http.StatusInternalServerError, "服务器内部错误")
		return
	}
	// 4. 处理“未找到”的情况
	if user == nil {
		response.FailWithMsg(c, http.StatusNotFound, "用户不存在")
		return
	}

	// 5.返回成功信息
	response.Success(c, user)
}

func (h *UserHandlerImpl) Create(c *gin.Context) {
	// 1.从请求中获取参数, 并将JSON数据转换成 model.user对象
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("数据格式转换错误", err)
		response.FailWithMsg(c, http.StatusBadRequest, "请检查数据格式")
		return
	}

	// 2.调用service层方法
	if err := h.s.Create(&user); err != nil {
		// 实际开发中，这里可能需要判断 err 是“重复存在”还是“系统错误”

		fmt.Println("创建用户失败:", err)
		response.FailWithMsg(c, http.StatusInternalServerError, "创建用户失败")
		return
	}

	// 3. 返回成功 (将生成的 ID 返回给前端)
	// GORM 会自动把生成的 ID 填回 user 结构体中
	response.Success(c, user)

}
