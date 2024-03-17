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
	CreateUserID uint32       `gorm:"column:create_user_id;type:int unsigned;not null;default:0;comment:创建人id"`
	UpdateUserID uint32       `gorm:"column:update_user_id;type:int unsigned;not null;default:0;comment:更新人id"`
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 AuthMenu 结构体对应的数据库表名
func (Menu) TableName() string {
	return "baker_auth_menu"
}

// menuRepo 菜单
type menuRepo struct {
	*db.GenericDao[Menu, uint32]
}

// NewMenuRepo 菜单
func NewMenuRepo(data *Data) biz.MenuRepo {
	return &menuRepo{
		GenericDao: &db.GenericDao[Menu, uint32]{
			DB: data.GetDB(),
		},
	}
}

func init() {
	registerInitField(initMenuField)
}

var (
	// 全字段修改Menu那些字段不修改
	notUpdateMenuField = []string{
		"create_time",
		"create_user_id",
	}
	updateMenuField []string
)

// InitMenuField 全字段修改
func initMenuField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&Menu{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateMenuField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateMenuField...)
	return nil
}

// GetMenuByID 根据菜单id查询菜单
func (r *menuRepo) GetMenuByID(ctx context.Context, id uint32) (*biz.Menu, error) {
	dbMenu, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return MenuPOToEntity(dbMenu), err
}

// GetMenuByIDs 根据菜单id集合查询菜单
func (r *menuRepo) GetMenuByIDs(ctx context.Context, ids []uint32) (biz.Menus, error) {
	dbMenus, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return MenuPOsToEntity(dbMenus), nil
}

// Save 保存菜单
func (r *menuRepo) Save(ctx context.Context, menu *biz.Menu) (uint32, error) {
	if menu.ID > 0 {
		return 0, errors.New("illegal argument menu id exist")
	}
	dbMenu := MenuEntityToPO(menu)
	if err := r.GenericDao.Create(ctx, dbMenu); err != nil {
		return 0, err
	}
	return menu.ID, nil
}

// UpdateByID 根据ID修改菜单
func (r *menuRepo) UpdateByID(ctx context.Context, menu *biz.Menu) error {
	if menu.ID == 0 {
		return errors.New("illegal argument menu exist")
	}
	dbMenu := MenuEntityToPO(menu)
	return r.GenericDao.DB.WithContext(ctx).Select(updateMenuField).Updates(dbMenu).Error
}

// DeleteByID 根据ID删除菜单
func (r *menuRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchMenu 菜单搜索
func (r *menuRepo) SearchMenu(ctx context.Context, req *pb.SearchMenuReq) (biz.Menus, uint32, error) {
	var (
		res       []*Menu
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
	return MenuPOsToEntity(res), uint32(pageTotal), nil
}
