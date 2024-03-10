package entity

// Category 文章分类表
type Category struct {
	ID          uint32    `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	Name        string    `gorm:"uniqueIndex:uk_name;column:name;type:varchar(50);not null;default:'';comment:文章类型名"`
	Description string    `gorm:"column:description;type:varchar(200);not null;default:'';comment:类型介绍"`
	Sort        int       `gorm:"column:sort;type:int;not null;default:0;comment:排序"`
	Available   uint8     `gorm:"column:available;type:tinyint unsigned;not null;default:0;comment:是否可用 0-不可用 1-可用"`
	CreateTime  Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime  Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 Category 结构体对应的数据库表名
func (Category) TableName() string {
	return "blog_category"
}
