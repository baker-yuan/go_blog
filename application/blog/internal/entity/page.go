package entity

// Page 页面表
type Page struct {
	ID         uint32    `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:页面id"`
	PageName   string    `gorm:"column:page_name;type:varchar(10);not null;default:'';comment:页面名"`
	PageLabel  string    `gorm:"column:page_label;type:varchar(20);not null;default:'';comment:页面标签"`
	PageCover  string    `gorm:"column:page_cover;type:varchar(255);not null;default:'';comment:页面封面"`
	CreateTime Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 Page 结构体对应的数据库表名
func (Page) TableName() string {
	return "blog_page"
}
