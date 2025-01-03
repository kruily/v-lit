package model_story

import "github.com/kruily/gofastcrud/core/crud"

type Chapter struct {
	*crud.BaseUUIDEntity
	VolumeID string `json:"volume_id" gorm:"column:volume_id;not null"`       // 卷ID
	Sort     int32  `json:"sort" gorm:"column:sort"`                          // 同级排序
	Name     string `json:"name" gorm:"column:name;not null"`                 // 名称
	Desc     string `json:"desc" gorm:"column:desc"`                          // 描述
	Content  string `json:"content" gorm:"column:content;not null;type:text"` // 内容，支持长文本
}

func (Chapter) Table() string {
	return "chapters"
}
