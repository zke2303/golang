package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/handler"
)

func BootRouter(c *gin.Engine, h handler.IBookHandler) {
	bookRouter := c.Group("/book")
	{
		bookRouter.POST("/", h.Create)
		bookRouter.DELETE("/:id", h.Delete)
		bookRouter.GET("/", h.FindById)
		bookRouter.PUT("", h.Update)
		bookRouter.GET("/page", h.PageQuery)
	}
}
