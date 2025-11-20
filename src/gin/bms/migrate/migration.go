package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zhang/bms/internal/model"
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

	err = db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		fmt.Println("asdasdasd")
		return
	}
}
