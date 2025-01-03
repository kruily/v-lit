package model_story

import "github.com/kruily/gofastcrud/core/crud"

type Story struct {
	*crud.BaseUUIDEntity
	Name        string `json:"name" gorm:"column:name;not null"`
	Author      string `json:"author" gorm:"column:author;not null"`
	Description string `json:"description" gorm:"column:description" `
	Cover       string `json:"cover" gorm:"column:cover"`
}

func (Story) Table() string {
	return "stories"
}
