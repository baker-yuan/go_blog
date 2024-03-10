package data

import (
	"context"

	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/common/db"
	pb "github.com/baker-yuan/go-blog/protocol/auth"
)

// Menu 菜单资源表
type Menu struct {
	// 基本信息
	ID        uint32 `gorm:"primary_key;column:id;type:int unsigned auto_increment;comment:菜单id"`
	ParentID  uint32 `gorm:"column:parent_id;type:int unsigned;not null;default:0;comment:父级菜单id"`
	Name      string `gorm:"column:name;type:varchar(20);not null;default:'';comment:菜单名称"`
	Icon      string `gorm:"column:icon;type:varchar(255);not null;default:'';comment:菜单图标"`
	Describe  string `gorm:"column:describe;type:varchar(200);not null;default:'';comment:功能描述"`
	SortValue uint32 `gorm:"column:sort_value;type:int unsigned;not null;default:1;comment:排序"`
	// 前端
	Path      string `gorm:"column:path;type:varchar(255);not null;default:'';comment:对应路由path"`
	Component string `gorm:"column:component;type:varchar(255);not null;default:'';comment:对应路由组件component"`
	// 状态
	IsHidden db.BoolBit `gorm:"column:is_hidden;type:bit(1);not null;default:b'0';comment:是否隐藏"`
	// 公共字段
	CreateUser uint32       `gorm:"column:create_user;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUser uint32       `gorm:"column:update_user;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 AuthMenu 结构体对应的数据库表名
func (Menu) TableName() string {
	return "baker_auth_menu"
}

type menuRepo struct {
}

// NewMenuRepo 菜单
func NewMenuRepo() biz.MenuRepo {
	return &menuRepo{}
}

// GetMenuByID 根据菜单id查询菜单
func (r *menuRepo) GetMenuByID(ctx context.Context, id int) (*biz.Menu, error) {

	return nil, nil
}

// GetMenuByIDs 根据菜单id集合查询菜单
func (r *menuRepo) GetMenuByIDs(ctx context.Context, ids []int) (biz.Menus, error) {

	return nil, nil
}

// SearchMenu 菜单搜索
func (r *menuRepo) SearchMenu(ctx context.Context, req *pb.SearchMenuReq) (biz.Menus, uint32, error) {

	return nil, uint32(0), nil
}

// Save 保存菜单
func (r *menuRepo) Save(ctx context.Context, menu *biz.Menu) (uint32, error) {

	return uint32(0), nil
}

// UpdateByID 根据ID修改菜单
func (r *menuRepo) UpdateByID(ctx context.Context, menu *biz.Menu) error {

	return nil
}

// DeleteByID 根据ID删除菜单
func (r *menuRepo) DeleteByID(ctx context.Context, id int) error {

	return nil
}
