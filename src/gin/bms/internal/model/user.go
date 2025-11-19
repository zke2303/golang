package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:30;uniqueKey" json:"username"`
	Password string `gorm:"size:100" json:"-"`
	Nickname string `gorm:"size:20" json:"nickname"`
	Email    string `gorm:"size:30" json:"email"`
	Gender   uint8  `gorm:"tinyint" json:"gender"`
	Age      uint8  `gorm:"tinyint" json:"age"`
	Icon     string `json:"icon"`
}
