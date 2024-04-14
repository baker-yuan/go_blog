package entity

import "github.com/baker-yuan/go-blog/go_blog/common/db"

// Moment 动态表
type Moment struct {
	ID          uint32  `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	Content     string  `gorm:"column:content;type:longtext;not null;comment:动态内容"`
	Likes       uint32  `gorm:"column:likes;type:int unsigned;not null;default:0;comment:点赞数量"`
	IsPublished BoolBit `gorm:"column:is_published;type:bit(1);not null;default:b'0';comment:是否公开"`
	IsDeleted   BoolBit `gorm:"column:is_deleted;type:bit(1);not null;default:b'0';comment:是否删除"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 Moment 结构体对应的数据库表名
func (Moment) TableName() string {
	return "blog_moment"
}
