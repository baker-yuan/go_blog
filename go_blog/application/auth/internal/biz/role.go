package biz

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/common/util"
	pb "github.com/baker-yuan/go-blog/protocol/auth"
)

// Role 角色
type Role struct {
	// 基本信息
	ID       uint32
	Name     string
	Code     string
	Describe string
	//
	IsEnable bool
	// 公共字段
	CreateUser uint32
	UpdateUser uint32
	CreateTime uint32
	UpdateTime uint32
}

type Roles []*Role

// RoleRepo 角色repo
type RoleRepo interface {
	// GetRoleByID 根据角色id集合查询角色
	GetRoleByID(ctx context.Context, id uint32) (*Role, error)
	// GetRoleByIDs 根据角色id集合查询角色
	GetRoleByIDs(ctx context.Context, ids []uint32) (Roles, error)
	// Save 保存角色
	Save(ctx context.Context, role *Role) (uint32, error)
	// UpdateByID 根据ID修改角色
	UpdateByID(ctx context.Context, role *Role) error
	// DeleteByID 根据ID删除角色
	DeleteByID(ctx context.Context, id uint32) error
	// SearchRole 角色搜索
	SearchRole(ctx context.Context, req *pb.SearchRoleReq) (Roles, uint32, error)
}

// RoleUseCase 角色管理
type RoleUseCase struct {
	*CommonUseCase
	repo RoleRepo
}

// NewRoleUseCase 创建角色管理业务逻辑实现
func NewRoleUseCase(
	commonUseCase *CommonUseCase,
	repo RoleRepo,
) *RoleUseCase {
	return &RoleUseCase{
		CommonUseCase: commonUseCase,
		repo:          repo,
	}
}

// RoleDetail 角色详情
func (c *RoleUseCase) RoleDetail(ctx context.Context, req *pb.RoleDetailReq) (*pb.Role, error) {
	role, err := c.repo.GetRoleByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	pbRole := RoleEntityToModel(role)
	return pbRole, nil
}

// SearchRole 角色搜索
func (c *RoleUseCase) SearchRole(ctx context.Context, req *pb.SearchRoleReq) ([]*pb.Role, uint32, error) {
	roles, total, err := c.repo.SearchRole(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	data := make([]*pb.Role, 0)
	for _, role := range roles {
		data = append(data, RoleEntityToModel(role))
	}
	return data, total, nil
}

// AddOrUpdateRole 添加修改角色
func (c *RoleUseCase) AddOrUpdateRole(ctx context.Context, req *pb.AddOrUpdateRoleReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}
	if req.GetId() == 0 {
		return c.addRole(ctx, userID, req)
	} else {
		dbRole, err := c.repo.GetRoleByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}
		return c.updateRole(ctx, dbRole, userID, req)
	}
}

func (c *RoleUseCase) addRole(ctx context.Context, userID uint32, req *pb.AddOrUpdateRoleReq) (uint32, error) {
	role := AddOrUpdateRoleReqToEntity(req)
	role.CreateUser = userID
	role.CreateTime = uint32(time.Now().Unix())
	role.UpdateUser = userID
	role.UpdateTime = uint32(time.Now().Unix())

	lastInsertID, err := c.repo.Save(ctx, role)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, pb.ResourceType_TB_Role,
		"{}", role,
		"新增角色",
	)

	return lastInsertID, nil
}

func (c *RoleUseCase) updateRole(ctx context.Context, dbRole *Role, userID uint32, req *pb.AddOrUpdateRoleReq) (uint32, error) {
	saveRole := AddOrUpdateRoleReqToEntity(req)
	saveRole.CreateUser = dbRole.CreateUser
	saveRole.CreateTime = dbRole.CreateTime
	saveRole.UpdateUser = userID
	saveRole.UpdateTime = uint32(time.Now().Unix())

	if err := c.repo.UpdateByID(ctx, saveRole); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_Role,
		dbRole, saveRole,
		"全字段修改角色",
	)

	return req.GetId(), nil
}

// DeleteRole 删除角色
func (c *RoleUseCase) DeleteRole(ctx context.Context, req *pb.DeleteRoleReq) error {
	role, err := c.repo.GetRoleByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.repo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_Role,
		role, "{}",
		"删除角色",
	)

	return nil
}
