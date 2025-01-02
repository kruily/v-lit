package models

import "github.com/kruily/gofastcrud/core/crud"

// User 用户模型
type User struct {
	*crud.BaseEntity
	Username string `json:"username" gorm:"column:username;type:varchar(255);not null;unique"`
	Password string `json:"-" gorm:"column:password;type:varchar(60);not null"`
	Email    string `json:"email" gorm:"column:email;type:varchar(255);not null;unique"`
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(255);not null;unique"`
}

func (User) Table() string {
	return "users"
}

// UserRegisterRequest 用户注册请求结构体
type UserRegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"` // 必填，长度限制
	Password string `json:"password" binding:"required,min=6,max=60"` // 必填，长度限制
	Email    string `json:"email" binding:"omitempty,email"`          // 可选，格式校验
	Phone    string `json:"phone" binding:"omitempty,len=11"`         // 可选，格式校验
}

// UserRegisterResponse 用户注册响应结构体
type UserLoginedResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// UserLoginRequest 用户登录请求结构体
type UserLoginRequest struct {
	Phone    string `json:"phone" binding:"omitempty,len=11"`
	Email    string `json:"email" binding:"omitempty,email"`
	Username string `json:"username" binding:"omitempty"`
	Password string `json:"password" binding:"required"`
}
