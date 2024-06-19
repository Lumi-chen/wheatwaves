package models

type TagModel struct {
	MODEL
	Name     string `json:"name"`      // 标签名称
	TagColor string `json:"tag_color"` // 标签颜色
}
