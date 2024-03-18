package dao

import "github.com/baker-yuan/go-blog/common/db"

// CanalConfig canal监听配置表
type CanalConfig struct {
	// 基本信息
	ID            uint32     `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:主键"`
	DBName        string     `gorm:"unique_index:uk_db_name_tb_name;column:db_name;type:varchar(255);not null;default:'';comment:库名"`
	TBName        string     `gorm:"unique_index:uk_db_name_tb_name;column:tb_name;type:varchar(255);not null;default:'';comment:表名"`
	MonitorInsert db.BoolBit `gorm:"column:monitor_insert;type:bit(1);not null;default:b'0';comment:是否监听插入"`
	MonitorUpdate db.BoolBit `gorm:"column:monitor_update;type:bit(1);not null;default:b'0';comment:是否监听修改"`
	MonitorDelete db.BoolBit `gorm:"column:monitor_delete;type:bit(1);not null;default:b'0';comment:是否监听删除"`
	// 变更通知

	// 公共字段
	CreateUser uint32       `gorm:"column:create_user;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUser uint32       `gorm:"column:update_user;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 CanalConfig 结构体对应的数据库表名
func (CanalConfig) TableName() string {
	return "baker_canal_config"
}
