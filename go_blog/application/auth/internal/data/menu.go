package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
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
	CreateTime   db.Timestamp `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:插入时间"`
	UpdateTime   db.Timestamp `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:修改时间"`
}

// TableName 设置 AuthMenu 结构体对应的数据库表名
func (m Menu) TableName() string {
	return MenuTbName
}

var (
	MenuTbName          = "baker_auth_menu"
	MenuFieldID         = "id"
	MenuFieldName       = "name"
	MenuFieldDescribe   = "describe"
	MenuFieldPath       = "path"
	MenuFieldComponent  = "component"
	MenuFieldIsHidden   = "is_hidden"
	MenuFieldSortValue  = "sort_value"
	MenuFieldIcon       = "icon"
	MenuFieldParentID   = "parent_id"
	MenuFieldCreateUser = "create_user"
	MenuFieldUpdateUser = "update_user"
	MenuFieldCreateTime = "create_time"
	MenuFieldUpdateTime = "update_time"
)

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

// GetMenuTree 查询菜单树
func (r *menuRepo) GetMenuTree(ctx context.Context, req *pb.GetMenuTreeReq) (biz.Menus, uint32, error) {
	var (
		menus     []*Menu
		pageTotal int64
	)
	tx := r.GenericDao.DB.WithContext(ctx)
	if len(req.GetName()) != 0 {
		tx = tx.Where(fmt.Sprintf("%s like ?", MenuFieldName), "%"+req.GetName()+"%")
	}
	// 对一级目录翻页
	tx = tx.Where(MenuFieldParentID, 0)
	if req.GetPageNum() > 0 && req.GetPageSize() > 0 {
		tx = tx.Offset(int((req.GetPageNum() - 1) * req.GetPageSize())).
			Limit(int(req.GetPageSize())).
			Order(fmt.Sprintf("%s asc", MenuFieldSortValue)).
			Find(&menus).
			Offset(-1).
			Limit(-1).
			Count(&pageTotal)
	} else {
		tx = tx.
			Order(fmt.Sprintf("%s asc", MenuFieldSortValue)).
			Find(&menus).
			Count(&pageTotal)
	}

	if err := tx.Error; err != nil {
		return nil, 0, err
	}
	return MenuPOsToEntity(menus), 0, nil
}

// GetSubMenu 补充子节点
func (r *menuRepo) GetSubMenu(ctx context.Context, parentIDs []uint32) (biz.Menus, error) {
	var subMenus []*Menu
	// 递归查询子菜单
	err := r.findSubMenus(ctx, parentIDs, &subMenus)
	if err != nil {
		return nil, err
	}
	return MenuPOsToEntity(subMenus), err
}

// findSubMenus 是一个递归辅助函数，用于查找和累积子菜单
func (r *menuRepo) findSubMenus(ctx context.Context, parentIDs []uint32, subMenus *[]*Menu) error {
	// 根据父ID查询子菜单
	children, err := r.getMenuByParentIDs(ctx, parentIDs)
	if err != nil {
		return err
	}
	// 如果没有子菜单，递归结束
	if len(children) == 0 {
		return nil
	}
	// 将子菜单添加到结果列表
	*subMenus = append(*subMenus, children...)

	// 为下一层递归准备新的父 ID 列表
	var nextParentIDs []uint32
	for _, child := range children {
		nextParentIDs = append(nextParentIDs, child.ID)
	}
	// 递归查询下一层子菜单
	return r.findSubMenus(ctx, nextParentIDs, subMenus)
}

// getMenuByParentIDs 根据父节点ID查找
func (r *menuRepo) getMenuByParentIDs(ctx context.Context, parentIDs []uint32) ([]*Menu, error) {
	var res []*Menu
	tx := r.GenericDao.DB.WithContext(ctx)
	tx = tx.Where(fmt.Sprintf("%s in ?", MenuFieldParentID), parentIDs)
	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
