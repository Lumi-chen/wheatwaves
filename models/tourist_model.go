package models

import "time"

// 游客模块
type TouristModel struct {
	MODEL
	NickName       string    `gorm:"type:varchar(50)" json:"nick_name"` // 昵称
	OperateTime    time.Time `json:"operate_time"`                      // 操作时间
	OperateContent string    `gorm:"size:255" json:"operate_content"`   // 操作内容
}
