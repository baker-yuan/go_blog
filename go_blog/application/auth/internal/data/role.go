package data

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/common/db"
	"github.com/baker-yuan/go-blog/common/util"
	pb "github.com/baker-yuan/go-blog/protocol/auth"
	"gorm.io/gorm"
)

// Role 角色表
type Role struct {
	// 基本信息
	ID       uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:ID"`
	Name     string `gorm:"uniqueIndex:uk_name;column:name;type:varchar(30);not null;default:'';comment:角色名称"`
	Code     string `gorm:"uniqueIndex:uk_code;column:code;type:varchar(20);not null;default:'';comment:角色编码"`
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
func (r Role) TableName() string {
	return RoleTbName
}

var (
	RoleTbName          = "baker_auth_role"
	RoleFieldID         = "id"
	RoleFieldName       = "name"
	RoleFieldCode       = "code"
	RoleFieldDescribe   = "describe"
	RoleFieldIsEnable   = "is_enable"
	RoleFieldCreateUser = "create_user"
	RoleFieldUpdateUser = "update_user"
	RoleFieldCreateTime = "create_time"
	RoleFieldUpdateTime = "update_time"
)

// RoleRepo 角色
type roleRepo struct {
	*db.GenericDao[Role, uint32]
}

// NewRoleRepo 创建
func NewRoleRepo(data *Data) biz.RoleRepo {
	return &roleRepo{
		GenericDao: &db.GenericDao[Role, uint32]{
			DB: data.GetDB(),
		},
	}
}

func init() {
	registerInitField(initRoleField)
}

var (
	// 全字段修改Role那些字段不修改
	notUpdateRoleField = []string{
		"created_at",
		"create_time",
	}
	updateRoleField []string
)

// InitRoleField 全字段修改
func initRoleField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(Role{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateRoleField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateRoleField...)
	return nil
}

// GetRoleByID 根据角色id查询角色
func (r *roleRepo) GetRoleByID(ctx context.Context, id uint32) (*biz.Role, error) {
	dbRole, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return RolePOToEntity(dbRole), err
}

// GetRoleByIDs 根据角色id集合查询角色
func (r *roleRepo) GetRoleByIDs(ctx context.Context, ids []uint32) (biz.Roles, error) {
	dbRoles, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return RolePOsToEntity(dbRoles), nil
}

// Save 保存角色
func (r *roleRepo) Save(ctx context.Context, role *biz.Role) (uint32, error) {
	if role.ID > 0 {
		return 0, errors.New("illegal argument role id exist")
	}
	dbRole := RoleEntityToPO(role)
	if err := r.GenericDao.Create(ctx, dbRole); err != nil {
		return 0, err
	}
	return dbRole.ID, nil
}

// UpdateByID 根据ID修改角色
func (r *roleRepo) UpdateByID(ctx context.Context, role *biz.Role) error {
	if role.ID == 0 {
		return errors.New("illegal argument role exist")
	}
	dbRole := RoleEntityToPO(role)
	return r.GenericDao.DB.WithContext(ctx).Select(updateRoleField).Updates(dbRole).Error
}

// DeleteByID 根据ID删除角色
func (r *roleRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchRole 角色搜索
func (r *roleRepo) SearchRole(ctx context.Context, req *pb.SearchRoleReq) (biz.Roles, uint32, error) {
	var (
		res       []*Role
		pageTotal int64
	)
	tx, err := db.BuildSearch(
		ctx,
		req.GetSearch(),
		r.GenericDao.DB.WithContext(ctx),
		func(search map[string]*db.SearchValue) {

		},
	)
	if err != nil {
		return nil, 0, err
	}
	tx = tx.Offset(int((req.GetPageNum() - 1) * req.GetPageSize())).
		Limit(int(req.GetPageSize())).Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)
	if err := tx.Error; err != nil {
		return nil, 0, err
	}
	return RolePOsToEntity(res), uint32(pageTotal), nil
}
