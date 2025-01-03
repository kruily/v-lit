package model_story

import "github.com/kruily/gofastcrud/core/crud"

type Node struct {
	*crud.BaseUUIDEntity
	Content  string  `json:"content" gorm:"column:content;not null" validate:"required"`     // 节点内容
	NodeType string  `json:"node_type" gorm:"column:node_type;not null" validate:"required"` // 节点类型
	ParentID *string `json:"parent_id" gorm:"column:parent_id"`                              // 父节点ID
	GraphID  string  `json:"graph_id" gorm:"type:text;index:idx_graph_id(255)"`              // 图ID
	Children []Node  `json:"children" gorm:"foreignKey:ParentID"`                            // 子节点列表
	Sort     int32   `json:"sort" gorm:"column:sort"`                                        // 同级排序
}

func (Node) Table() string {
	return "nodes"
}
