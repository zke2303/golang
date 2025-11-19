package request

type Page struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"pageSize"`
}
