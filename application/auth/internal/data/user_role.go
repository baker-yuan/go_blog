package data

// UserRole 账号角色绑定表
type UserRole struct {
	ID     uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:ID"`
	RoleID uint32 `gorm:"index:idx_role_id;column:role_id;type:int unsigned;not null;default:0;comment:角色ID"`
	UserID uint32 `gorm:"index:idx_user_id;column:user_id;type:int unsigned;not null;default:0;comment:用户ID"`
	// 公共字段
	CreateUser uint32 `gorm:"column:create_user;type:int unsigned;not null;default:0;comment:创建人ID"`
	CreateTime uint32 `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
}

// TableName 设置 UserRole 结构体对应的数据库表名
func (UserRole) TableName() string {
	return "baker_auth_user_role"
}
