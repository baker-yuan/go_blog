package data

import "github.com/baker-yuan/go-blog/common/db"

// Role 角色表
type Role struct {
	// 基本信息
	ID       uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:ID"`
	Code     string `gorm:"uniqueIndex:uk_code;column:code;type:varchar(20);not null;default:'';comment:角色编码"`
	Name     string `gorm:"uniqueIndex:uk_name;column:name;type:varchar(30);not null;default:'';comment:角色名称"`
	Describe string `gorm:"column:describe;type:varchar(100);not null;default:'';comment:功能描述"`
	//
	IsEnable db.BoolBit `gorm:"column:is_enable;type:bit(1);not null;default:b'1';comment:是否启用"`
	// 公共字段
	CreateUser uint32       `gorm:"column:create_user;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUser uint32       `gorm:"column:update_user;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 Role 结构体对应的数据库表名
func (Role) TableName() string {
	return "baker_auth_role"
}
