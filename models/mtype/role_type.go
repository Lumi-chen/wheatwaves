package mtype

import "encoding/json"

type Role int

const (
	Admin        Role = 1 // 管理员
	Blogger      Role = 2 // 博主
	OrdinaryUser Role = 3 // 普通用户
	Tourist      Role = 4 // 游客
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
	case OrdinaryUser:
		str = "普通用户"
	case Tourist:
		str = "游客"
	}
	return str
}
