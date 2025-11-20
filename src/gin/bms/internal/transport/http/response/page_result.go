package response

type PageResult struct {
	Total   int64       `json:"total"`
	Current int         `json:"current"`
	Size    int         `json:"size"`
	Record  interface{} `json:"record"`
}
