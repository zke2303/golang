package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/config"
	"github.com/zhang/bms/internal/controller"
	"github.com/zhang/bms/internal/middleware"
	"github.com/zhang/bms/internal/repository"
	"github.com/zhang/bms/internal/router"
	"github.com/zhang/bms/internal/service"
	"github.com/zhang/bms/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 读取配置文件
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	// 1.初始化日志
	logger.InitLogger("./logs/app.log", "debug")
	// 2.在程序结束前，刷新缓存区
	defer logger.Log.Sync()

	// 3. init router
	r := gin.New()
	r.Use(middleware.GinLogger())
	r.Use(gin.Recovery())

	// 4.连接dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DataSource.Mysql.User,
		cfg.DataSource.Mysql.Password,
		cfg.DataSource.Mysql.Host,
		cfg.DataSource.Mysql.Port,
		cfg.DataSource.Mysql.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connect fail!")
	}

	// 5
	userRepo := repository.NewIUserRepository(db)
	userService := service.NewIUserService(userRepo)
	userCtl := controller.NewUserController(userService)

	router.UserRouter(r, userCtl)

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	r.Run(addr)

}
