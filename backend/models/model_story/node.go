package model_story

import "github.com/kruily/gofastcrud/core/crud"

type Node struct {
	*crud.BaseUUIDEntity
	Content     string `json:"content" gorm:"column:content;not null" validate:"required"`       // 节点内容
	NodeType    string `json:"node_type" gorm:"column:node_type;not null" validate:"required"`   // 节点类型
	ParentID    string `json:"parent_id" gorm:"column:parent_id;size:255"`                       // 父节点ID，可为空
	GraphID     string `json:"graph_id" gorm:"column:graph_id;size:255;index:idx_graph_id(255)"` // 图ID
	Children    []Node `json:"children" gorm:"foreignKey:ParentID"`                              // 子节点列表
	Sort        int32  `json:"sort" gorm:"column:sort"`                                          // 同级排序
	XPos        int32  `json:"x_pos" gorm:"column:x_pos"`                                        // 横坐标
	YPos        int32  `json:"y_pos" gorm:"column:y_pos"`                                        // 纵坐标
	Width       int32  `json:"width" gorm:"column:width"`                                        // 宽度
	Height      int32  `json:"height" gorm:"column:height"`                                      // 高度
	ZIndex      int32  `json:"z_index" gorm:"column:z_index"`                                    // 层级
	IsActive    bool   `json:"is_active" gorm:"column:is_active"`                                // 是否激活
	IsShow      bool   `json:"is_show" gorm:"column:is_show"`                                    // 是否显示
	IsLocked    bool   `json:"is_locked" gorm:"column:is_locked"`                                // 是否锁定
	Color       string `json:"color" gorm:"column:color;size:255"`                               // 颜色
	FontStyle   string `json:"font_style" gorm:"column:font_style;size:255"`                     // 字体样式
	BorderWidth int32  `json:"border_width" gorm:"column:border_width"`                          // 边框宽度
	BorderStyle string `json:"border_style" gorm:"column:border_style;size:255"`                 // 边框样式
	Shape       string `json:"shape" gorm:"column:shape;size:255"`                               // 形状
	ShapeStyle  string `json:"shape_style" gorm:"column:shape_style;size:255"`                   // 形状样式

}

func (Node) Table() string {
	return "nodes"
}
