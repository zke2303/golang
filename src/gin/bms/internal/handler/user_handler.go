package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/dto"
	"github.com/zhang/bms/internal/service"
	"github.com/zhang/bms/internal/transport/http/request"
	"github.com/zhang/bms/internal/transport/http/response"
)

type UserHandler struct {
	s service.IUserService
}

func NewUserHandler(s service.IUserService) UserHandler {
	return UserHandler{s: s}
}

func (h *UserHandler) Login(c *gin.Context) {
	// 1.从请求头中获取username、password
	var login dto.LoginDTO
	err := c.Bind(&login)
	if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, "请输入账号和密码")
		return
	}

	// 2.调用service层，校验登入
	token, err := h.s.Login(&login)
	if err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, "账号或密码错误")
		return
	}
	// 3.登入成功，把token添加到gin.Context中
	c.Set("username", login.Username)
	response.Success(c, token)
}

func (h *UserHandler) FindById(c *gin.Context) {
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

func (h *UserHandler) Create(c *gin.Context) {
	// 1.从请求中获取参数, 并将JSON数据转换成 model.user对象
	var userDTO dto.UserRequest
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		fmt.Println("数据格式转换错误", err)
		response.FailWithMsg(c, http.StatusBadRequest, "请检查数据格式")
		return
	}

	// 2.调用service层方法
	if err := h.s.Create(&userDTO); err != nil {
		// 实际开发中，这里可能需要判断 err 是“重复存在”还是“系统错误”

		fmt.Println("创建用户失败:", err)
		response.FailWithMsg(c, http.StatusInternalServerError, "创建用户失败")
		return
	}

	// 3. 返回成功 (将生成的 ID 返回给前端)
	// GORM 会自动把生成的 ID 填回 user 结构体中
	response.Success(c, nil)
}

func (h *UserHandler) Delete(c *gin.Context) {
	// 1.从请求中获取id
	strId := c.Param("id")
	// 2.将string类型的id转换成uint64
	id, err := strconv.ParseUint(strId, 19, 64)

	if err != nil {
		fmt.Println("输入的id不是数字")
		response.FailWithMsg(c, http.StatusBadRequest, "请输入数字")
		return
	}

	// 3.调用service方法
	if err := h.s.Delete(id); err != nil {
		fmt.Println("删除用户数据失败")
		response.FailWithMsg(c, http.StatusBadRequest, "删除失败")
		return
	}

	// 4.返回删除成功
	response.Success(c, nil)
}

func (h *UserHandler) Update(c *gin.Context) {
	// 1.从请求中获取参数, 并将JSON数据转换成 model.user对象
	var params struct {
		ID uint64 `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&params); err != nil {
		fmt.Println("ID无效")
		response.FailWithMsg(c, http.StatusBadRequest, "ID无效")
		return
	}

	// 2.解析body
	var req dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, http.StatusBadRequest, err.Error())
		return
	}

	// 3.调用service
	err := h.s.Update(params.ID, &req)
	if err != nil {
		response.FailWithMsg(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func (*UserHandler) PageQuery(c *gin.Context) {
	// 1.从请求中获取分页参数
	var page request.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		fmt.Println("分页参数错误")
		response.FailWithMsg(c, http.StatusBadRequest, "请检查页参数")
		return
	}

	// 2.从请求中获取 query 参数
}
