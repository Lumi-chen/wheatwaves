package models

// 专栏表
type ColumnModel struct {
	MODEL
	Name      string         `gorm:"type:varchar(50)" json:"name"`     // 专栏名称
	Abstract  string         `gorm:"size:255" json:"abstract"`         // 专栏简介
	CoverID   string         `gorm:"type:varchar(36)" json:"cover_id"` // 封面id
	CoverPath string         `gorm:"size:255" json:"cover_path"`       // 封面路径
	Articles  []ArticleModel `gorm:"many2many:column_articles"`        // 连接文章表
}
