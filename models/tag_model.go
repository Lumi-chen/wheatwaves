package models

type TagModel struct {
	MODEL
	Name     string `json:"name"`      // 专栏名称
	TagColor string `json:"tag_color"` // 标签颜色
}
