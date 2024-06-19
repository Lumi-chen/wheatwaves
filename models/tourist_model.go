package models

import "time"

// 游客模块
type TouristModel struct {
	MODEL
	NickName       string    `json:"nick_name"`       // 昵称
	OperateAt      time.Time `json:"operate_at"`      // 操作时间
	OperateContent string    `json:"operate_content"` // 操作内容
}
