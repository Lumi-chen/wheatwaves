package models

// 专栏表
type ColumnModel struct {
	MODEL
	Name      string `json:"name"`       // 专栏名称
	Abstract  string `json:"abstract"`   // 专栏简介
	CoverID   string `json:"cover_id"`   // 封面id
	CoverPath string `json:"cover_path"` // 封面路径
}
