package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangke/ems/internal/handler"
)

func EmployeeRouter(r *gin.Engine, h *handler.EmployeeHandler) {
	employeeRouter := r.Group("/employee")
	{
		employeeRouter.GET("", h.FindById)
		employeeRouter.POST("", h.CreateEmployee)
	}
}
