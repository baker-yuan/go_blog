package biz

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/common/util"
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
	CreateUserID uint32
	UpdateUserID uint32
	CreateTime   uint32
	UpdateTime   uint32
}

type Menus []*Menu

// MenuRepo 菜单
type MenuRepo interface {
	// GetMenuByID 根据菜单id集合查询菜单
	GetMenuByID(ctx context.Context, id uint32) (*Menu, error)
	// GetMenuByIDs 根据菜单id集合查询菜单
	GetMenuByIDs(ctx context.Context, ids []uint32) (Menus, error)
	// Save 保存菜单
	Save(ctx context.Context, menu *Menu) (uint32, error)
	// UpdateByID 根据ID修改菜单
	UpdateByID(ctx context.Context, menu *Menu) error
	// DeleteByID 根据ID删除菜单
	DeleteByID(ctx context.Context, id uint32) error
	// SearchMenu 菜单搜索
	SearchMenu(ctx context.Context, req *pb.SearchMenuReq) (Menus, uint32, error)
}

// MenuUsecase 菜单业务实现
type MenuUsecase struct {
	*CommonUseCase
	repo MenuRepo
}

// NewMenuUsecase 菜单业务实现
func NewMenuUsecase(
	commonUseCase *CommonUseCase,
	repo MenuRepo,
) *MenuUsecase {
	return &MenuUsecase{
		CommonUseCase: commonUseCase,
		repo:          repo,
	}
}

// MenuDetail 菜单详情
func (c *MenuUsecase) MenuDetail(ctx context.Context, req *pb.MenuDetailReq) (*pb.Menu, error) {
	menu, err := c.repo.GetMenuByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	pbMenu := MenuEntityToModel(menu)
	return pbMenu, nil
}

// SearchMenu 菜单搜索
func (c *MenuUsecase) SearchMenu(ctx context.Context, req *pb.SearchMenuReq) ([]*pb.Menu, uint32, error) {
	menus, total, err := c.repo.SearchMenu(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	data := make([]*pb.Menu, 0)
	for _, menu := range menus {
		data = append(data, MenuEntityToModel(menu))
	}
	return data, total, nil
}

// AddOrUpdateMenu 添加修改菜单
func (c *MenuUsecase) AddOrUpdateMenu(ctx context.Context, req *pb.AddOrUpdateMenuReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}
	if req.GetId() == 0 {
		return c.addMenu(ctx, userID, req)
	} else {
		dbMenu, err := c.repo.GetMenuByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}
		return c.updateMenu(ctx, dbMenu, userID, req)
	}
}

func (c *MenuUsecase) addMenu(ctx context.Context, userID uint32, req *pb.AddOrUpdateMenuReq) (uint32, error) {
	menu := AddOrUpdateMenuReqToEntity(req)
	menu.CreateUserID = userID
	menu.CreateTime = uint32(time.Now().Unix())
	menu.UpdateUserID = userID
	menu.UpdateTime = uint32(time.Now().Unix())

	lastInsertID, err := c.repo.Save(ctx, menu)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, pb.ResourceType_TB_MENU,
		"{}", menu,
		"新增菜单",
	)

	return lastInsertID, nil
}

func (c *MenuUsecase) updateMenu(ctx context.Context, dbMenu *Menu, userID uint32, req *pb.AddOrUpdateMenuReq) (uint32, error) {
	saveMenu := AddOrUpdateMenuReqToEntity(req)
	saveMenu.CreateUserID = dbMenu.CreateUserID
	saveMenu.CreateTime = dbMenu.CreateTime
	saveMenu.UpdateUserID = userID
	saveMenu.UpdateTime = uint32(time.Now().Unix())

	if err := c.repo.UpdateByID(ctx, saveMenu); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_MENU,
		dbMenu, saveMenu,
		"全字段修改菜单",
	)

	return req.GetId(), nil
}

// DeleteMenu 删除菜单
func (c *MenuUsecase) DeleteMenu(ctx context.Context, req *pb.DeleteMenuReq) error {
	menu, err := c.repo.GetMenuByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.repo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_MENU,
		menu, "{}",
		"删除菜单",
	)
	return nil
}
