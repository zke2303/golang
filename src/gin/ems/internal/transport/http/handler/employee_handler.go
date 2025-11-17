package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/service"
	"github.com/zhang/bms/internal/transport/http/request"
	"github.com/zhang/bms/internal/transport/http/response"
	"gorm.io/gorm"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) EmployeeHandler {
	return EmployeeHandler{service: service}
}

func (h EmployeeHandler) FindById(c *gin.Context) {
	strId := c.Query("id")
	if strId == "" {
		fmt.Println("id不能为空")
		response.Fail(c, http.StatusBadRequest, "id不能为空")
		return
	}

	// 将 strId 转换成 uint64
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		fmt.Println("请输入数字")
		response.Fail(c, http.StatusBadRequest, "请输入数字")
		return
	}

	// 调用 service 层
	employee, err := h.service.FindById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, http.StatusNotFound, "员工未找到")
			return
		}
		fmt.Println("查询数据库错误:", err)
		response.Fail(c, http.StatusInternalServerError, "查询数据库失败")
		return
	}

	response.Success(c, employee)
}

func (h EmployeeHandler) List(c *gin.Context) {
	// 1.将查询参数从请求中提取出来
	var query request.EmployeeQuery
	err := c.ShouldBindJSON(&query)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "JSON格式转换异常, 请检查请求参数")
		return
	}

	// 2.查询数据库
	records, err := h.service.List(query)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail(c, http.StatusNotFound, "员工未找到")
			return
		}
		fmt.Println("查询数据库错误:", err)
		response.Fail(c, http.StatusInternalServerError, "查询数据库失败")
		return
	}

	response.Success(c, records)
}
