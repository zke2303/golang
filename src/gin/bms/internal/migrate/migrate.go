package main

import (
	"github.com/zhang/bms/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("../../bms.db"), &gorm.Config{})
	db.AutoMigrate(model.User{})
}
