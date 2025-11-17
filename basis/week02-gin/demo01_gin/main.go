package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)
import "github.com/thinkerou/favicon"

func main() {
	ginServer := gin.Default()

	ginServer.Use(favicon.New("./go128.png"))

	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "hello gin"})
	})

	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "post user"})
	})

	ginServer.GET("/user", func(context *gin.Context) {
		userId := context.Query("userId")
		username := context.Query("username")
		fmt.Printf("userId: %s username: %s\n", userId, username)
	})

	ginServer.Run(":8080")
}
