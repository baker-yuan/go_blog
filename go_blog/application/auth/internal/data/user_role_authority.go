package data

import (
	"context"
	"fmt"

	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
)

// UserRole 账号角色绑定表
type UserRole struct {
	ID     uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:ID"`
	UserID uint32 `gorm:"index:idx_user_id;column:user_id;type:int unsigned;not null;default:0;comment:用户ID"`
	RoleID uint32 `gorm:"index:idx_role_id;column:role_id;type:int unsigned;not null;default:0;comment:角色ID"`
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
}

// TableName 设置 UserRole 结构体对应的数据库表名
func (*UserRole) TableName() string {
	return UserRoleTbName
}

var (
	UserRoleTbName          = "baker_auth_user_role"
	UserRoleFieldId         = "id"
	UserRoleFieldRoleID     = "role_id"
	UserRoleFieldUserID     = "user_id"
	UserRoleFieldCreateUser = "create_user"
	UserRoleFieldCreateTime = "create_time"
)

// RoleAuthority 角色关联的资源和目录表
type RoleAuthority struct {
	ID            uint32           `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:ID"`
	RoleID        uint32           `gorm:"index:idx_role_id;column:role_id;type:int unsigned;not null;comment:角色id"`
	AuthorityType pb.AuthorityType `gorm:"column:authority_type;type:tinyint unsigned;not null;default:0;comment:权限类型 1-菜单 2-资源"`
	AuthorityID   uint32           `gorm:"column:authority_id;type:int unsigned;not null;comment:权限id"` // 资源id/菜单id
	// 公共字段
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
}

// TableName 设置 RoleAuthority 结构体对应的数据库表名
func (*RoleAuthority) TableName() string {
	return RoleAuthorityTbName
}

var (
	RoleAuthorityTbName             = "baker_auth_role_authority"
	RoleAuthorityFieldId            = "id"
	RoleAuthorityFieldAuthorityId   = "authority_id"
	RoleAuthorityFieldAuthorityType = "authority_type"
	RoleAuthorityFieldRoleID        = "role_id"
	RoleAuthorityFieldCreateUser    = "create_user"
	RoleAuthorityFieldCreateTime    = "create_time"
)

// userRoleAuthorityRepo 用户角色权限关联
type userRoleAuthorityRepo struct {
	userRole      *db.GenericDao[UserRole, uint32]
	roleAuthority *db.GenericDao[RoleAuthority, uint32]
}

// NewUserRoleAuthorityRepo 创建
func NewUserRoleAuthorityRepo(data *Data) biz.UserRoleAuthorityRepo {
	return &userRoleAuthorityRepo{
		userRole:      db.NewGenericDao[UserRole, uint32](data.GetDB()),
		roleAuthority: db.NewGenericDao[RoleAuthority, uint32](data.GetDB()),
	}
}

// DeleteRoleAuthority 删除关联关系
func (r *userRoleAuthorityRepo) DeleteRoleAuthority(ctx context.Context, roleId uint32, authorityType pb.AuthorityType) error {
	tx := r.roleAuthority.GetDB(ctx)
	tx = tx.Where(RoleAuthorityFieldRoleID, roleId).Where(RoleAuthorityFieldAuthorityType, authorityType)
	if err := tx.Delete(&RoleAuthority{}).Error; err != nil {
		return err
	}
	return nil
}

// BatchSaveAuthority 批量保存
func (r *userRoleAuthorityRepo) BatchSaveAuthority(ctx context.Context, roleAuthorityList []*biz.RoleAuthority) error {
	saveList := RoleAuthorityEntityListToPO(roleAuthorityList)
	tx := r.roleAuthority.GetDB(ctx)
	if err := tx.Save(saveList).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUserRole 删除关联关系
func (r *userRoleAuthorityRepo) DeleteUserRole(ctx context.Context, userID uint32) error {
	tx := r.userRole.GetDB(ctx)
	tx = tx.Where(UserRoleFieldUserID, userID)
	if err := tx.Delete(&UserRole{}).Error; err != nil {
		return err
	}
	return nil
}

// BatchSaveUserRole 批量保存用户和角色之间的关联
func (r *userRoleAuthorityRepo) BatchSaveUserRole(ctx context.Context, userRoleList []*biz.UserRole) error {
	saveList := UserRoleEntityListToPO(userRoleList)
	tx := r.userRole.GetDB(ctx)
	if err := tx.Save(saveList).Error; err != nil {
		return err
	}
	return nil
}

// GetUserRole 获取用户绑定的角色
func (r *userRoleAuthorityRepo) GetUserRole(ctx context.Context, userIDs []uint32) (biz.UserRoles, error) {
	userRoles := make([]*UserRole, 0)
	tx := r.userRole.GetDB(ctx)
	tx = tx.Where(fmt.Sprintf("%s in ?", UserRoleFieldUserID), userIDs)
	if err := tx.Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return UserRolePOListToEntity(userRoles), nil
}

// GetRoleAuthorityByRoleIDs 获取角色关联的接口权限
func (r *userRoleAuthorityRepo) GetRoleAuthorityByRoleIDs(ctx context.Context, roleIDs []uint32, authorityType pb.AuthorityType) (biz.RoleAuthorityList, error) {
	roleAuthorityList := make([]*RoleAuthority, 0)
	tx := r.roleAuthority.GetDB(ctx)
	tx = tx.Where(fmt.Sprintf("%s in ?", RoleAuthorityFieldRoleID), roleIDs)
	tx = tx.Where(RoleAuthorityFieldAuthorityType, pb.AuthorityType_RESOURCE)
	if err := tx.Find(&roleAuthorityList).Error; err != nil {
		return nil, err
	}
	return RoleAuthorityPOListToEntity(roleAuthorityList), nil
}

// GetUserRoleByUserIDs 查询用户关联的角色
func (r *userRoleAuthorityRepo) GetUserRoleByUserIDs(ctx context.Context, userID uint32) (biz.UserRoles, error) {
	userRoleList := make([]*UserRole, 0)
	tx := r.userRole.GetDB(ctx)
	tx = tx.Where(UserRoleFieldUserID, userID)
	if err := tx.Find(&userRoleList).Error; err != nil {
		return nil, err
	}
	return UserRolePOListToEntity(userRoleList), nil
}
