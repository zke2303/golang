package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/handler"
	"github.com/zhang/bms/internal/repository"
	"github.com/zhang/bms/internal/router"
	"github.com/zhang/bms/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dbHost := "172.16.140.128"
	dbPort := "3306"
	dbUsername := "root"
	dbPassword := "123"
	dbDatabase := "test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbDatabase)

	// 新建日志器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出位置
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	router.BootRouter(r, bookHandler)
	router.UserRouter(r, userHandler)

	if err := r.Run("localhost:8080"); err != nil {
		fmt.Println("服务器启动失败")
		return
	}
}
