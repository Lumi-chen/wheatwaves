package models

import "encoding/json"

type Role int

// 用户表
type UserModel struct {
	MODEL
	NickName      string `gorm:"size: 50" json:"nick_name"` // 昵称
	UserName      string `gorm:"size: 36" json:"user_name"` // 用户名
	Password      string `gorm:"size: 128" json:"password"` // 密码
	Avatar        string `json:"avatar"`                    // 头像
	Role          Role   `json:"role"`                      // 角色权限： 1-管理员、2-博主、3-游客
	IP            string `json:"ip"`                        // IP地址
	Profile       string `json:"profile"`                   // 个人简介
	BlogName      string `json:"blog_name"`                 // 博客空间名称
	GithubAddress string `json:"github_add"`                // github地址
	GiteeAddress  string `json:"gitee_add"`                 // gitee地址
}

const (
	Admin   Role = 1 // 管理员
	Blogger Role = 2 // 博主
	Tourist Role = 3 // 游客
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Role) String() string {
	var str string
	switch r {
	case Admin:
		str = "管理员"
	case Blogger:
		str = "博主"
	case Tourist:
		str = "游客"
	}
	return str
}
