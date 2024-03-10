package biz

import (
	"context"

	pb "github.com/baker-yuan/go-blog/protocol/auth"
)

// Menu 菜单资源表
type Menu struct {
	// 基本信息
	ID        uint32
	ParentID  uint32
	Name      string
	Icon      string
	Describe  string
	SortValue uint32
	// 前端
	Path      string
	Component string
	// 状态
	IsHidden bool
	// 公共字段
	CreateUser uint32
	UpdateUser uint32
	CreateTime uint32
	UpdateTime uint32
}

type Menus []*Menu

// MenuRepo 菜单
type MenuRepo interface {
	// GetMenuByID 根据菜单id集合查询菜单
	GetMenuByID(ctx context.Context, id int) (*Menu, error)
	// GetMenuByIDs 根据菜单id集合查询菜单
	GetMenuByIDs(ctx context.Context, ids []int) (Menus, error)
	// Save 保存菜单
	Save(ctx context.Context, menu *Menu) (uint32, error)
	// UpdateByID 根据ID修改菜单
	UpdateByID(ctx context.Context, menu *Menu) error
	// DeleteByID 根据ID删除菜单
	DeleteByID(ctx context.Context, id int) error
	// SearchMenu 菜单搜索
	SearchMenu(ctx context.Context, req *pb.SearchMenuReq) (Menus, uint32, error)
}

type MenuUsecase struct {
	repo MenuRepo
}

func NewArticleUsecase(repo MenuRepo) *MenuUsecase {
	return &MenuUsecase{repo: repo}
}

// MenuDetail 菜单详情
func (c *MenuUsecase) MenuDetail(ctx context.Context, req *pb.MenuDetailReq) (*pb.Menu, error) {
	//menu, err := c.menuRepo. GetMenuByID(ctx, int(req.GetId()))
	//if err != nil {
	//	return nil, err
	//}
	//pbMenu := assembler.MenuEntityToModel(menu)
	//return pbMenu, nil

	return nil, nil
}

// SearchMenu 菜单搜索
func (c *MenuUsecase) SearchMenu(ctx context.Context, req *pb.SearchMenuReq) ([]*pb.Menu, uint32, error) {
	//menus, total, err := c.menuRepo. SearchMenu(ctx, req)
	//if err != nil {
	//	return nil, 0, err
	//}
	//data := make([]*pb.Menu, 0)
	//for _, menu := range menus {
	//	data = append(data, assembler.MenuEntityToModel(menu))
	//}
	//return data, total, nil

	return nil, 0, nil
}

// AddOrUpdateMenu 添加修改菜单
func (c *MenuUsecase) AddOrUpdateMenu(ctx context.Context, req *pb.AddOrUpdateMenuReq) (uint32, error) {
	//loginName, err := pkg_util.GetLoginStaffName(ctx)
	//if err != nil {
	//	return 0, err
	//}
	//if err := c.CheckResAuth(ctx, loginName); err != nil {
	//	return 0, err
	//}
	//
	//if req.GetId() == 0 {
	//	return c.addMenu(ctx, loginName, req)
	//} else {
	//	dbMenu, err := c.menuRepo. GetMenuByID(ctx, int(req.GetId()))
	//	if err != nil {
	//		return 0, err
	//	}
	//
	//	return c.updateMenu(ctx, dbMenu, loginName, req)
	//}

	return 0, nil
}

//func (c *MenuUsecase) addMenu(ctx context.Context, loginName string, req *pb.AddOrUpdateMenuReq) (uint32, error) {
//	menu := assembler.AddOrUpdateMenuReqToEntity(req)
//	menu .AddTime = sql.NullTime{Time: time.Now(), Valid: true}
//	menu .AddOperator = loginName
//	menu .LastChgTime = sql.NullTime{Time: time.Now(), Valid: true}
//	menu .LastChgUser = loginName
//
//	lastInsertID, err := c.menuRepo. Save(ctx, menu)
//	if err != nil {
//		return 0, err
//	}
//
//	c.SaveChangeLog(ctx,
//		lastInsertID, pb.ResourceType_RT_,
//		"{}", menu,
//		"新增菜单",
//	)
//
//	return lastInsertID, nil
//}
//
//func (c *MenuUsecase) updateMenu(ctx context.Context, dbMenu *entity.Menu, loginName string, req *pb.AddOrUpdateMenuReq) (uint32, error) {
//	saveMenu := assembler.AddOrUpdateMenuReqToEntity(req)
//	saveMenu.AddTime = dbMenu.AddTime
//	saveMenu.AddOperator = dbMenu.AddOperator
//	saveMenu.LastChgTime = sql.NullTime{Time: time.Now(), Valid: true}
//	saveMenu.LastChgUser = loginName
//
//	if err := c.menuRepo. UpdateByID(ctx, saveMenu); err != nil {
//		return 0, err
//	}
//
//	c.SaveChangeLog(ctx,
//		req.GetId(), pb.ResourceType_RT_,
//		dbMenu, saveMenu,
//		"全字段修改菜单",
//	)
//
//	return req.GetId(), nil
//}

// DeleteMenu 删除菜单
func (c *MenuUsecase) DeleteMenu(ctx context.Context, req *pb.DeleteMenuReq) error {
	//loginName, err := pkg_util.GetLoginStaffName(ctx)
	//if err != nil {
	//	return err
	//}
	//if err := c.CheckResAuth(ctx, loginName); err != nil {
	//	return err
	//}
	//
	//menu, err := c.menuRepo. GetMenuByID(ctx, int(req.GetId()))
	//if err != nil {
	//	return err
	//}
	//
	//if err := c.menuRepo. DeleteByID(ctx, int(req.GetId())); err != nil {
	//	return err
	//}
	//
	//c.SaveChangeLog(ctx,
	//	req.GetId(), pb.ResourceType_RT_,
	//	menu, "{}",
	//	"删除菜单",
	//)

	return nil
}
