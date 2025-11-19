package response

type PageResult struct {
	Total   uint32      `json:"total"`
	Current uint32      `json:"current"`
	Size    uint32      `json:"size"`
	Record  interface{} `json:"record"`
}
