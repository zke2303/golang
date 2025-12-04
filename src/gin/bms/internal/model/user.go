package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"column:username;type:varchar(20);not null;uniqueIndex"`
	Password string `json:"password" form:"password" gorm:"column:password;type:varchar(20);not null"`
	Gender   string `json:"gender" form:"gender" gorm:"column:gender;default:male"`
	Age      uint8  `json:"age" form:"age" gorm:"column:age;"`
}
