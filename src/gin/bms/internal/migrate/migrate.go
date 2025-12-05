package main

import (
	"fmt"

	"github.com/zhang/bms/config"
	"github.com/zhang/bms/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DataSource.Mysql.User,
		cfg.DataSource.Mysql.Password,
		cfg.DataSource.Mysql.Host,
		cfg.DataSource.Mysql.Port,
		cfg.DataSource.Mysql.Database,
	)

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(model.User{})
}
