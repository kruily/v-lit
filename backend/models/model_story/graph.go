package model_story

import "github.com/kruily/gofastcrud/core/crud"

type Graph struct {
	*crud.BaseUUIDEntity
	UserID      string `json:"user_id" gorm:"column:user_id;size:255;not null"`       // 用户ID
	StoryID     string `json:"story_id" gorm:"column:story_id;size:255;not null"`     // 故事id
	GraphName   string `json:"graph_name" gorm:"column:graph_name;size:120;not null"` // 图名称
	GraphRemark string `json:"graph_remark" gorm:"column:graph_remark;size:255"`      // 图备注
	Nodes       []Node `json:"nodes" gorm:"foreignKey:GraphID"`                       // 节点列表
	GraphType   string `json:"graph_type" gorm:"column:graph_type;size:255"`          // 图类型
}

func (Graph) Table() string {
	return "graphs"
}
