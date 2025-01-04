package model_story

import "github.com/kruily/gofastcrud/core/crud"

type Card struct {
	*crud.BaseUUIDEntity
	Name     string `json:"name" gorm:"column:name;size:120;not null"`           // 名称
	Content  string `json:"content" gorm:"column:content;not null;type:text"`    // 内容
	StoryID  string `json:"story_id" gorm:"column:story_id;size:255;not null"`   // 故事ID
	CardType string `json:"card_type" gorm:"column:card_type;size:255;not null"` // 卡片类型
}

func (Card) Table() string {
	return "cards"
}
