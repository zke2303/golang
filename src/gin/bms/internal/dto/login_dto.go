package dto

type LoginDTO struct {
	Username string `form:"username" json:"username" binding:"required,min=6"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}
