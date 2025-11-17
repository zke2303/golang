package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zhangke/ems/internal/handler"
	"github.com/zhangke/ems/internal/repository"
	"github.com/zhangke/ems/internal/router"
	"github.com/zhangke/ems/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	host := "192.168.138.128"
	port := "3306"
	username := "root"
	password := "123"
	database := "test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败: %v", err))
	}

	// 初始化依赖
	empRepo := repository.NewEmployeeRepository(db)
	empService := service.NewEmployeeService(empRepo)
	empHandler := handler.NewEmployeeHandler(empService)

	// 注册路由
	router.EmployeeRouter(r, empHandler)

	// 启动服务
	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
