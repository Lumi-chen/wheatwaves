package models

import "time"

type TouristModel struct {
	ID             uint      `gorm:"primaryKey"`      // id
	CreatedAt      time.Time `json:"-"`               // 创建时间
	NickName       string    `json:"nick_name"`       // 专栏名称
	OperateAt      time.Time `json:"operate_at"`      // 操作时间
	OperateContent string    `json:"operate_content"` // 操作内容
}
