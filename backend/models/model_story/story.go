package model_story

import "github.com/kruily/gofastcrud/core/crud"

type Story struct {
	*crud.BaseUUIDEntity
	UserID      string `json:"user_id" gorm:"column:user_id;size:255;not null"` // 用户ID
	Name        string `json:"name" gorm:"column:name;size:120;not null"`       // 名称
	Description string `json:"description" gorm:"column:description;size:255"`  // 描述
	Cover       string `json:"cover" gorm:"column:cover;size:255"`              // 封面
}

func (Story) Table() string {
	return "stories"
}
