package models

import "encoding/json"

type ImageModel struct {
	MODEL
	Path        string      `gorm:"size:255" json:"path"`          // 图片路径
	Hash        string      `json:"hash"`                          // 图片hash值，用于去重
	Name        string      `gorm:"type:varchar(100)" json:"name"` // 文件名称
	PurposeType PurposeType `gorm:"type:int" json:"purpose_type"`  // 图片用途
}

type PurposeType int

const (
	Avatar  PurposeType = 1 // 头像
	Cover   PurposeType = 2 // 封面
	Content PurposeType = 3 // 文章内容
)

func (p PurposeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p PurposeType) String() string {
	var str string
	switch p {
	case Avatar:
		str = "头像"
	case Cover:
		str = "封面"
	case Content:
		str = "文章内容"
	}
	return str
}
