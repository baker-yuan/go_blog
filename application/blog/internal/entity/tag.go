package entity

// Tag 文章标签表
type Tag struct {
	ID          uint32    `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	Name        string    `gorm:"uniqueIndex:uk_name;column:name;type:varchar(50);not null;default:'';comment:标签名"`
	Description string    `gorm:"column:description;type:varchar(100);not null;default:'';comment:标签描述"`
	Color       string    `gorm:"column:color;type:varchar(64);not null;default:'red';comment:标签颜色"`
	CreateTime  Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime  Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 Tag 结构体对应的数据库表名
func (Tag) TableName() string {
	return "blog_tag"
}
