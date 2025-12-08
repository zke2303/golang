package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/controller"
)

func UserRouter(r *gin.Engine, c controller.UserController) {
	userRouter := r.Group("/api/user")
	{
		userRouter.GET("", c.FindByUsername)
		userRouter.POST("", c.Insert)
		userRouter.DELETE("/:id", c.Delete)
	}
}
