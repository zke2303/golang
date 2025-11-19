package model

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name     string    `json:"name"`
	Age      uint8     `gorm:"tinyint" json:"age"`
	Gender   uint8     `gorm:"tinyint" json:"gender"`
	Birthday time.Time `gorm:"datetime" json:"birthday"`
}
