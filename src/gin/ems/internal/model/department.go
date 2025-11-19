package model

import "time"

type Department struct {
	Id        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:30" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
