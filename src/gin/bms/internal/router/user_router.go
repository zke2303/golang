package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/handler"
)

func UserRouter(c *gin.Engine, h handler.UserHandler) {
	userRouter := c.Group("/user")
	{
		userRouter.GET("", h.FindById)
		userRouter.POST("", h.Create)
		userRouter.DELETE("/:id", h.Delete)
		userRouter.PUT("/:id", h.Update)
		userRouter.POST("/page", h.PageQuery)
		userRouter.POST("/login", h.Login)
	}
}
