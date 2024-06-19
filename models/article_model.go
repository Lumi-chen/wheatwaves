package models

import "time"

// 文章表
type ArticleModel struct {
	MODEL
	Title        string      `gorm:"size:50" json:"title"`     // 文章标题
	Abstract     string      `gorm:"size:150" json:"abstract"` // 文章简介
	Content      string      `json:"content"`                  // 文章内容
	Category     string      `json:"category"`                 // 文章分类
	CoverID      ImageModel  `json:"cover_id"`                 // 文章封面
	CoverPath    string      `json:"cover_path"`               // 文章封面路径
	LookCount    string      `json:"look_count"`               // 浏览量
	CommentCount string      `json:"comment_count"`            // 评论量
	DiggCount    string      `json:"digg_count"`               // 点赞量
	UserID       uint        `json:"user_id"`                  // 用户ID
	NickName     string      `json:"nick_name" gorm:"size:50"` // 作者昵称
	User         UserModel   `json:"-"`                        // 关联用户
	PublishTime  time.Time   `json:"publish_time"`             // 发布时间
	ColumnID     int         `json:"column_id"`                // 专栏id
	ColumnName   string      `json:"column_name"`              // 专栏名称
	Column       ColumnModel `json:"-"`                        //  关联专栏
	// TagModels    []TagModel `gorm:"many2many:article_tag" json:"tag_model"` // 标签tag
}
