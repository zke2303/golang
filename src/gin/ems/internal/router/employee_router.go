package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/transport/http/handler"
)

func EmployeeRouter(r *gin.Engine, h *handler.EmployeeHandler) {
	employeeRouter := r.Group("/employee")
	{
		employeeRouter.GET("/", h.FindById)
		employeeRouter.GET("/list", h.List)
	}
}
