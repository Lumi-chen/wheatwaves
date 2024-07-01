package models

import (
	"wheatwaves/models/mtype"
)

// 用户表
type UserModel struct {
	MODEL
	NickName      string         `gorm:"type:varchar(50)" json:"nick_name"` // 昵称
	UserName      string         `gorm:"type:varchar(36)" json:"user_name"` // 用户名
	Password      string         `gorm:"type:varchar(128)" json:"password"` // 密码
	AvatarPath    string         `gorm:"size:255" json:"avatar_path"`       // 头像
	Role          mtype.Role     `gorm:"type:int" json:"role"`              // 角色权限： 1-管理员、2-博主、3-游客
	IP            string         `gorm:"size:255" json:"ip"`                // IP地址
	Profile       string         `gorm:"size:255" json:"profile"`           // 个人简介
	BlogName      string         `gorm:"type:varchar(50)" json:"blog_name"` // 博客空间名称
	GithubAddress string         `gorm:"size:255" json:"github_add"`        // github地址
	GiteeAddress  string         `gorm:"size:255" json:"gitee_add"`         // gitee地址
	Articles      []ArticleModel `gorm:"many2many:user_articles"`           // 连接文章表
	// LastLoginTime time.Time      // 最后登陆时间
}
