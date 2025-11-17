package response

type PageResult[T any] struct {
	Total   int64 `gorm:"total" json:"total"`
	Current int   `gorm:"current" json:"current"`
	Records []T   `gorm:"records" json:"records"`
}
