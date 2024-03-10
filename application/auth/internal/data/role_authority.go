package data

type AuthorityType string

const (
	MENU     = "MENU"
	RESOURCE = "RESOURCE"
)

// RoleAuthority 角色关联的资源和目录表
type RoleAuthority struct {
	ID            uint32        `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:ID"`
	AuthorityID   uint32        `gorm:"column:authority_id;type:int unsigned;not null;comment:权限id"`
	AuthorityType AuthorityType `gorm:"column:authority_type;type:varchar(10);not null;default:'MENU';comment:权限类型 MENU-菜单 RESOURCE-资源"`
	RoleID        uint32        `gorm:"index:idx_role_id;column:role_id;type:int unsigned;not null;comment:角色id"`
	// 公共字段
	CreateUser uint32 `gorm:"column:create_user;type:int unsigned;not null;default:0;comment:创建人ID"`
	CreateTime uint32 `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
}

// TableName 设置 RoleAuthority 结构体对应的数据库表名
func (RoleAuthority) TableName() string {
	return "baker_auth_role_authority"
}
