package biz

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
)

// RoleAuthority 角色关联的资源和目录表
type RoleAuthority struct {
	ID            uint32           // ID
	AuthorityID   uint32           // 菜单id/资源id
	AuthorityType pb.AuthorityType // 权限类型 1-菜单 2-资源
	RoleID        uint32           // 角色id
	// 公共字段
	CreateUserID uint32       // 创建人id
	CreateTime   db.Timestamp // 插入时间
}

type RoleAuthorityList []*RoleAuthority

func (l RoleAuthorityList) GetAuthorityIDs() []uint32 {
	ids := make([]uint32, 0)
	for _, item := range l {
		ids = append(ids, item.AuthorityID)
	}
	return ids
}

// UserRole 账号角色绑定表
type UserRole struct {
	ID     uint32 // ID
	RoleID uint32 // 角色ID
	UserID uint32 // 用户ID
	// 公共字段
	CreateUserID uint32       // 创建人id
	CreateTime   db.Timestamp // 创建时间
}

type UserRoles []*UserRole

// GetRoleIDs 获取角色id
func (r UserRoles) GetRoleIDs() []uint32 {
	ids := make([]uint32, 0)
	for _, item := range r {
		ids = append(ids, item.RoleID)
	}
	return ids
}

type UserRoleAuthorityRepo interface {
	// DeleteRoleAuthority 删除关联关系
	DeleteRoleAuthority(ctx context.Context, roleId uint32, authorityType pb.AuthorityType) error
	// BatchSaveAuthority 批量保存角色和权限之间的关联
	BatchSaveAuthority(ctx context.Context, roleAuthorityList []*RoleAuthority) error
	// DeleteUserRole 删除关联关系
	DeleteUserRole(ctx context.Context, userID uint32) error
	// BatchSaveUserRole 批量保存用户和角色之间的关联
	BatchSaveUserRole(ctx context.Context, userRoleList []*UserRole) error
	// GetUserRole 获取用户绑定的角色
	GetUserRole(ctx context.Context, userIDs []uint32) (UserRoles, error)
	// GetRoleAuthorityByRoleIDs 获取角色关联的接口权限
	GetRoleAuthorityByRoleIDs(ctx context.Context, roleIDs []uint32, authorityType pb.AuthorityType) (RoleAuthorityList, error)
	// GetUserRoleByUserIDs 查询用户关联的角色
	GetUserRoleByUserIDs(ctx context.Context, userID uint32) (UserRoles, error)
}

// UserRoleAuthorityUseCase 用户角色权限关联处理
type UserRoleAuthorityUseCase struct {
	*CommonUseCase
	repo UserRoleAuthorityRepo
}

// NewUserRoleAuthorityUseCase 创建用户角色权限关联处理实现
func NewUserRoleAuthorityUseCase(
	commonUseCase *CommonUseCase,
	repo UserRoleAuthorityRepo,
) *UserRoleAuthorityUseCase {
	return &UserRoleAuthorityUseCase{
		CommonUseCase: commonUseCase,
		repo:          repo,
	}
}

// RoleBindMenu 角色绑定菜单权限
func (c *UserRoleAuthorityUseCase) RoleBindMenu(ctx context.Context, userID uint32, req *pb.RoleBindMenuReq) error {
	if err := db.Transaction(ctx, func(ctx context.Context) error {
		// 删除关联关系
		if err := c.repo.DeleteRoleAuthority(ctx, req.GetRoleId(), pb.AuthorityType_MENU); err != nil {
			return err
		}
		// 重新关联
		authorityList := make([]*RoleAuthority, 0)
		for _, menuID := range req.GetMenuIds() {
			authorityList = append(authorityList, &RoleAuthority{
				AuthorityID:   menuID,
				AuthorityType: pb.AuthorityType_MENU,
				RoleID:        req.GetRoleId(),
				CreateUserID:  userID,
				CreateTime:    db.Timestamp(time.Now().Unix()),
			})
		}
		if err := c.repo.BatchSaveAuthority(ctx, authorityList); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// RoleBindResource 角色绑定资源权限
func (c *UserRoleAuthorityUseCase) RoleBindResource(ctx context.Context, userID uint32, req *pb.RoleBindResourceReq) error {
	if err := db.Transaction(ctx, func(ctx context.Context) error {
		// 删除关联关系
		if err := c.repo.DeleteRoleAuthority(ctx, req.GetRoleId(), pb.AuthorityType_RESOURCE); err != nil {
			return err
		}
		// 重新关联
		authorityList := make([]*RoleAuthority, 0)
		for _, resourceID := range req.GetResourceIds() {
			authorityList = append(authorityList, &RoleAuthority{
				AuthorityID:   resourceID,
				AuthorityType: pb.AuthorityType_RESOURCE,
				RoleID:        req.GetRoleId(),
				CreateUserID:  userID,
				CreateTime:    db.Timestamp(time.Now().Unix()),
			})
		}
		if err := c.repo.BatchSaveAuthority(ctx, authorityList); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// UserBindRole 用户绑定角色
func (c *UserRoleAuthorityUseCase) UserBindRole(ctx context.Context, userID uint32, req *pb.UserBindRoleReq) error {
	if err := db.Transaction(ctx, func(ctx context.Context) error {
		// 删除关联关系
		if err := c.repo.DeleteUserRole(ctx, req.GetUserId()); err != nil {
			return err
		}
		// 重新关联
		userRoleList := make([]*UserRole, 0)
		for _, roleID := range req.GetRoleIds() {
			userRoleList = append(userRoleList, &UserRole{
				UserID:       req.GetUserId(),
				RoleID:       roleID,
				CreateUserID: userID,
				CreateTime:   db.Timestamp(time.Now().Unix()),
			})
		}
		if err := c.repo.BatchSaveUserRole(ctx, userRoleList); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// GetUserRole 获取用户绑定的角色
func (c *UserRoleAuthorityUseCase) GetUserRole(ctx context.Context, loginUserID uint32, req *pb.GetUserRoleReq) ([]*pb.GetUserRoleRsp_Item, error) {
	var (
		userRoleList []*UserRole
		err          error
	)
	userRoleList, err = c.repo.GetUserRole(ctx, req.GetUserIds())
	if err != nil {
		return nil, err
	}
	groupByUserID := util.NewFieldExtractor[*UserRole, uint32]().
		GroupByField(userRoleList, func(userRole *UserRole) uint32 {
			return userRole.UserID
		})
	res := make([]*pb.GetUserRoleRsp_Item, 0)
	for userID, itemList := range groupByUserID {
		item := &pb.GetUserRoleRsp_Item{
			UserId:   userID,
			UserRole: UserRoleEntityListToModel(itemList),
		}
		res = append(res, item)
	}
	return res, nil
}

// GetUserResource 获取用户关联的接口权限
func (c *UserRoleAuthorityUseCase) GetUserResource(ctx context.Context, loginUserID uint32, req *pb.GetUserResourceReq) ([]uint32, error) {
	// 查询用户关联的角色
	userRoleList, err := c.repo.GetUserRoleByUserIDs(ctx, req.GetUserId())
	// 查询角色关联的接口资源
	roleAuthorityList, err := c.repo.GetRoleAuthorityByRoleIDs(ctx, userRoleList.GetRoleIDs(), pb.AuthorityType_RESOURCE)
	if err != nil {
		return nil, err
	}
	return roleAuthorityList.GetAuthorityIDs(), nil
}
