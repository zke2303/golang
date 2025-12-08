package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username;not null;type:varchar(20)"`
	Password string `json:"password" gorm:"column:password; not null;type:varchar(20)"`
	Gender   string `json:"gender" gorm:"column:gender;default:male"`
	Age      uint8  `json:"age" gorm:"column:age;"`
}
