package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string         `gorm:"type:varchar(20);not null;uniqueIndex:idx_username_del"`
	DeletedAt gorm.DeletedAt `gorm:"uniqueIndex:idx_username_del"`

	Password string `gorm:"type:varchar(20);not null"`
	Gender   string `gorm:"default:male"`
	Age      uint8
}
