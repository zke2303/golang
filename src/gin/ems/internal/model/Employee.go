package model

import (
	"database/sql"
	"time"
)

type Employee struct {
	Id         uint64       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string       `gorm:"size:40;uniqueKey" json:"username"`
	Password   string       `gorm:"size:20;" json:"password"`
	Nickname   string       `gorm:"size:30;" json:"nickname"`
	Department string       `gorm:"size:20" json:"department"`
	Gender     int8         `gorm:"type:tinyint(2)" json:"gender"`
	Age        uint8        `json:"age"`
	CreateTime time.Time    `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime sql.NullTime `gorm:"column:update_time;autoUpdateTime" json:"update_time,omitempty"`
}
