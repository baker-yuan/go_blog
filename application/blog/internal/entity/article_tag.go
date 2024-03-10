package entity

// ArticleTag 文章标签关联表
type ArticleTag struct {
	ID         uint32    `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	TagID      uint32    `gorm:"uniqueIndex:uk_tag_id_article_id;index:idx_tag_id;column:tag_id;type:int unsigned;not null;default:0;comment:标签表ID|blog_tag主键"`
	ArticleID  uint32    `gorm:"uniqueIndex:uk_tag_id_article_id;index:idx_article_id;column:article_id;type:int unsigned;not null;default:0;comment:文章ID|biz_article主键"`
	CreateTime Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 ArticleTag 结构体对应的数据库表名
func (ArticleTag) TableName() string {
	return "blog_article_tag"
}
