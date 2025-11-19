package request

import "time"

type BookQuery struct {
	Title       string    `gorm:"size:30" json:"title"`
	AuthorId    int64     `json:"author_id"`
	Category    uint32    `gorm:"tinyint" json:"category"`
	Status      uint32    `gorm:"tinyint" json:"status"`
	ISBN        string    `json:"ISBN"`
	Price       uint64    `json:"price"`
	PublisherId uint64    `json:"publisher_id"`
	PublishDate time.Time `json:"publish_date"`
	WordCount   uint16    `json:"word_count"`
}
