package request

import "time"

type BookQuery struct {
	Title       string    `gorm:"size:30" form:"title"`
	AuthorId    int64     `form:"author_id"`
	Category    uint32    `gorm:"tinyint" form:"category"`
	Status      uint32    `gorm:"tinyint" form:"status"`
	ISBN        string    `form:"ISBN"`
	Price       uint64    `form:"price"`
	PublisherId uint64    `form:"publisher_id"`
	PublishDate time.Time `form:"publish_date"`
	WordCount   uint16    `form:"word_count"`
}
