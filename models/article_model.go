package models

import (
	"time"
)

// 文章表
type ArticleModel struct {
	MODEL
	Title        string      `gorm:"type:varchar(50)" json:"title"`       // 文章标题
	Abstract     string      `gorm:"size:255" json:"abstract"`            // 文章简介
	Content      string      `gorm:"type:longtext" json:"content"`        // 文章内容
	Category     string      `gorm:"type:varchar(50)" json:"category"`    // 文章分类
	CoverID      string      `gorm:"type:varchar(36)" json:"cover_id"`    // 文章封面
	CoverPath    string      `gorm:"size:255" json:"cover_path"`          // 文章封面路径
	LookCount    int         `json:"look_count"`                          // 浏览量
	CommentCount int         `json:"comment_count"`                       // 评论量
	DiggCount    int         `json:"digg_count"`                          // 点赞量
	UserID       string      `gorm:"type:varchar(36)" json:"user_id"`     // 用户ID
	NickName     string      `gorm:"type:varchar(50)" json:"nick_name"`   // 作者昵称
	User         UserModel   `json:"-"`                                   // 关联用户
	PublishTime  time.Time   `json:"publish_time"`                        // 发布时间
	ColumnID     string      `gorm:"type:varchar(36)" json:"column_id"`   // 专栏id
	ColumnName   string      `gorm:"type:varchar(50)" json:"column_name"` // 专栏名称
	Column       ColumnModel `json:"-"`                                   //  关联专栏
	Tags         []TagModel  `gorm:"many2many:article_tag"`               // 关联标签表
	// TagModels    []TagModel `gorm:"many2many:article_tag" json:"tag_model"` // 标签tag
}
