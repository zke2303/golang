package request

type EmployeeDTO struct {
	Username   *string `json:"username" form:"username"`
	Password   *string `json:"password" form:"password"`
	Nickname   *string `json:"nickname" form:"nickname"`
	Department *string `json:"department" form:"department"`
	Gender     *string `json:"gender" form:"gender"`
	Age        *int    `json:"age" form:"age"`
}
