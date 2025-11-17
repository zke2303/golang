package model

import "time"

type Employee struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"size:40;unique" json:"username"`
	Password   string    `gorm:"size:20" json:"password,omitempty"`
	Nickname   string    `gorm:"size:20" json:"nickname"`
	Department string    `gorm:"size:20" json:"department"`
	Gender     int8      `json:"gender"`
	Age        int       `json:"age"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
}
