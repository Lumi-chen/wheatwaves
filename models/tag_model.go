package models

type TagModel struct {
	MODEL
	Name     string `gorm:"type:varchar(50)" json:"name"`      // 标签名称
	TagColor string `gorm:"type:varchar(10)" json:"tag_color"` // 标签颜色
}
