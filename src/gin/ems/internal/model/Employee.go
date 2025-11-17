package model

import (
	"time"
)

type Employee struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string    `gorm:"size:40;uniqueKey" json:"username"`
	Password   string    `gorm:"size:20;" json:"password"`
	NickName   string    `gorm:"size:30;" json:"nickName"`
	Department string    `gorm:"size:20" json:"department"`
	Gender     int8      `gorm:"type:tinyint(2)" json:"gender"`
	Age        uint8     `json:"age"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}
