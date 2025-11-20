package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/handler"
)

func UserRouter(c *gin.Engine, h handler.IUserHandler) {
	userRouter := c.Group("/user")
	{
		userRouter.GET("", h.FindById)
		userRouter.POST("", h.Create)
	}
}
