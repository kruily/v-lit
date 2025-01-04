package model_story

import "github.com/kruily/gofastcrud/core/crud"

type Volume struct {
	*crud.BaseUUIDEntity
	GraphID  string    `json:"graph_id" gorm:"column:graph_id;size:255;not null"` // 图ID
	Sort     int32     `json:"sort" gorm:"column:sort"`                           // 同级排序
	Name     string    `json:"name" gorm:"column:name;size:120;not null"`         // 名称
	Desc     string    `json:"desc" gorm:"column:desc;size:255"`                  // 描述
	Chapters []Chapter `json:"chapters" gorm:"foreignKey:VolumeID"`               // 章节列表
}

func (Volume) Table() string {
	return "volumes"
}
