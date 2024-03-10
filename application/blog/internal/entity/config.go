package entity

// Config 网站配置
type Config struct {
	ID         uint32    `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	Code       string    `gorm:"uniqueIndex:uk_code;column:code;type:varchar(20);not null;default:'';comment:唯一编码"`
	Desc       string    `gorm:"column:desc;type:varchar(20);not null;default:'';comment:描述信息"`
	Config     string    `gorm:"column:config;type:mediumtext;not null;comment:配置信息"`
	CreateTime Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 Config 结构体对应的数据库表名
func (Config) TableName() string {
	return "blog_config"
}
