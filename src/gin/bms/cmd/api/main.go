package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhang/bms/internal/controller"
	"github.com/zhang/bms/internal/middleware"
	"github.com/zhang/bms/internal/repository"
	"github.com/zhang/bms/internal/router"
	"github.com/zhang/bms/internal/service"
	"github.com/zhang/bms/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 1.初始化日志
	logger.InitLogger("./logs/app.log", "debug")
	// 2.在程序结束前，刷新缓存区
	defer logger.Log.Sync()

	// 3. init router
	r := gin.New()
	r.Use(middleware.GinLogger())
	r.Use(gin.Recovery())

	// 4.链接数据库
	db, err := gorm.Open(sqlite.Open("../../bms.db"), &gorm.Config{})
	if err != nil {
		panic("Database connect fail!")
	}

	// 5
	userRepo := repository.NewIUserRepository(db)
	userService := service.NewIUserService(userRepo)
	userCtl := controller.NewUserController(userService)

	router.UserRouter(r, userCtl)

	r.Run()

}
