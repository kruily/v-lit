package models

import "github.com/kruily/gofastcrud/core/crud"

type User struct {
	*crud.BaseEntity
	Username string `json:"username" gorm:"column:username;type:varchar(255);not null;unique"`
	Password string `json:"password" gorm:"column:password;type:varchar(60);not null"`
	Email    string `json:"email" gorm:"column:email;type:varchar(255);not null;unique"`
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(255);not null;unique"`
}

func (User) Table() string {
	return "users"
}
