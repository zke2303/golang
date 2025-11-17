package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhangke/ems/internal/service"
)

type EmployeeHandler struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeHandler(s service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		EmployeeService: s,
	}
}

func (h *EmployeeHandler) FindById(c *gin.Context) {
	employee, err := h.EmployeeService.FindById(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {

	err := h.EmployeeService.Create(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "success")
}
