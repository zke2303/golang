package dto

import (
	"time"

	"github.com/zhang/bms/internal/model"
)

// UserRequest: 注册或修改时使用
type UserRequest struct {
	// 使用 binding 标签进行参数校验
	Username string `json:"username" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=6,max=30"`
	Nickname string `json:"nickname" binding:"max=20"`
	Email    string `json:"email" binding:"required,email"` // 强制校验邮箱格式
	Gender   uint8  `json:"gender" binding:"oneof=0 1 2"`   // 限制只能传 0, 1, 2
	Age      uint8  `json:"age" binding:"lte=120"`          // 简单的逻辑校验
	Icon     string `json:"icon"`
}

// UserResponse: 返回给前端的数据
type UserResponse struct {
	ID        uint      `json:"id"`         // 显式返回 ID
	CreatedAt time.Time `json:"created_at"` // 返回创建时间
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Gender    uint8     `json:"gender"`
	Age       uint8     `json:"age"`
	Icon      string    `json:"icon"`
	// 注意：绝对不要在 Response 中包含 Password 字段，即使是 json:"-"
	// 最好是直接把字段删掉，防止误操作
}

// UserUpdateRequest, 用于用户更新信息
type UserUpdateRequest struct {
	// 不需要 binding:"required"，因为是部分更新
	Nickname *string `json:"nickname" binding:"omitempty,max=20"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Gender   *uint8  `json:"gender" binding:"omitempty,oneof=0 1 2"`
	Age      *uint8  `json:"age" binding:"omitempty,lte=120"`
	Icon     *string `json:"icon"`
	// 密码单独处理，如果前端传了值才更新，不传则不改
	Password *string `json:"password" binding:"omitempty,min=6"`
}

// ToModel 属于 UserRequest 的方法，必须和 UserRequest 在同一个包
func (req *UserRequest) ToModel() *model.User {
	return &model.User{
		Username: req.Username,
		// 注意：这里只是简单的数据搬运。
		// 密码加密通常建议在 Service 层调用加密函数后再赋值，或者在这里直接通过 Utils 加密
		Password: req.Password,
		Nickname: req.Nickname,
		Email:    req.Email,
		Gender:   req.Gender,
		Age:      req.Age,
		Icon:     req.Icon,
	}
}

// UserQueryRequest 用户用户的分页查询
type UserQueryRequest struct {
	Nickname string `json:"nickname"`
	Gender   uint8  `json:"gender"`
	Age      uint8  `json:"age"`
}

// ToResponse 是一个普通函数，也可以放在这里，作为“工厂方法”
func ToUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Gender:    user.Gender,
		Age:       user.Age,
		Icon:      user.Icon,
	}
}

// 辅助方法：将 DTO 转为 map[string]interface{} 供 GORM 使用
func (req *UserUpdateRequest) ToMap() map[string]interface{} {
	m := make(map[string]interface{})

	// 只有当字段不为 nil 时，才放入 map
	if req.Nickname != nil {
		m["nickname"] = *req.Nickname
	}
	if req.Email != nil {
		m["email"] = *req.Email
	}
	if req.Gender != nil {
		m["gender"] = *req.Gender
	}
	if req.Age != nil {
		m["age"] = *req.Age
	}
	if req.Icon != nil {
		m["icon"] = *req.Icon
	}
	// 注意：密码通常需要在 Service 层加密后再放入，或者在这里不做处理，Service 层单独判断

	return m
}
