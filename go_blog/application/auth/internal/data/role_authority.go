package data

import (
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
)

// RoleAuthority 角色关联的资源和目录表
type RoleAuthority struct {
	ID            uint32           `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:ID"`
	AuthorityID   uint32           `gorm:"column:authority_id;type:int unsigned;not null;comment:权限id"`
	AuthorityType pb.RoleAuthority `gorm:"column:authority_type;type:tinyint unsigned;not null;default:0;comment:权限类型 1-菜单 2-资源"`
	RoleID        uint32           `gorm:"index:idx_role_id;column:role_id;type:int unsigned;not null;comment:角色id"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
}

// TableName 设置 RoleAuthority 结构体对应的数据库表名
func (RoleAuthority) TableName() string {
	return "baker_auth_role_authority"
}
