package main

import (
	"fmt"

	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/repository"
	"github.com/zhang/bms/internal/router"
	"github.com/zhang/bms/internal/service"
	"github.com/zhang/bms/internal/transport/http/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// 初始化 Gin 引擎
	r := gin.Default()

	// 数据库配置
	dbHost := "192.168.138.128"
	dbPort := "3306"
	dbUser := "root"
	dbPassword := "123"
	dbName := "test"

	// 拼接 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	fmt.Println(dsn)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // Log level: Info 会打印所有 SQL
			IgnoreRecordNotFoundError: true,        // 忽略记录未找到的错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("数据库连接异常:", err)
		return
	}

	// 初始化 Repository、Service、Handler
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	router.EmployeeRouter(r, &employeeHandler)

	// 启动服务
	if err := r.Run("localhost:8080"); err != nil {
		fmt.Println("服务器启动异常:", err)
		return
	}
}
