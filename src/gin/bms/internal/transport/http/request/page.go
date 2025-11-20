package request

type Page struct {
	Page     int `form:"page"`
	PageSize int `form:"size"`
}
