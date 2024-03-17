package biz

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/common/util"
	pb "github.com/baker-yuan/go-blog/protocol/auth"
)

// Resource 接口资源
type Resource struct {
	// 基本信息
	ID       uint32
	ParentID uint32
	Name     string
	Describe string
	//
	ResourceType pb.ResourceResourceType
	Status       pb.ResourceStatus
	// 路径定位
	URL        string
	HTTPMethod string
	// 权限校验
	IsNeedLogin      pb.NeedLogin
	IsNeedPermission pb.NeedPermission
	// 下游服务信息
	Service string
	Method  string
	// 公共字段
	CreateUser uint32
	UpdateUser uint32
	CreateTime uint32
	UpdateTime uint32
}

// Resources 接口集合
type Resources []*Resource

type ResourceRepo interface {
	// GetResourceByID 根据接口id集合查询接口
	GetResourceByID(ctx context.Context, id uint32) (*Resource, error)
	// GetResourceByIDs 根据接口id集合查询接口
	GetResourceByIDs(ctx context.Context, ids []uint32) (Resources, error)
	// Save 保存接口
	Save(ctx context.Context, resource *Resource) (uint32, error)
	// UpdateByID 根据ID修改接口
	UpdateByID(ctx context.Context, resource *Resource) error
	// DeleteByID 根据ID删除接口
	DeleteByID(ctx context.Context, id uint32) error
	// SearchResource 接口搜索
	SearchResource(ctx context.Context, req *pb.SearchResourceReq) (Resources, uint32, error)
	// GetEffectiveResource 获取有效状态下的接口
	GetEffectiveResource(ctx context.Context, req *pb.GetEffectiveResourceReq) (Resources, error)
}

// ResourceUseCase 接口业务实现
type ResourceUseCase struct {
	*CommonUseCase
	repo ResourceRepo
}

// NewResourceUseCase 创建接口管理service
func NewResourceUseCase(
	commonUseCase *CommonUseCase,
	repo ResourceRepo,
) *ResourceUseCase {
	return &ResourceUseCase{
		CommonUseCase: commonUseCase,
		repo:          repo,
	}
}

// ResourceDetail 接口详情
func (c *ResourceUseCase) ResourceDetail(ctx context.Context, req *pb.ResourceDetailReq) (*pb.Resource, error) {
	resource, err := c.repo.GetResourceByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	pbResource := ResourceEntityToModel(resource)
	return pbResource, nil
}

// SearchResource 接口搜索
func (c *ResourceUseCase) SearchResource(ctx context.Context, req *pb.SearchResourceReq) ([]*pb.Resource, uint32, error) {
	resources, total, err := c.repo.SearchResource(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	data := make([]*pb.Resource, 0)
	for _, resource := range resources {
		data = append(data, ResourceEntityToModel(resource))
	}
	return data, total, nil
}

// AddOrUpdateResource 添加修改接口
func (c *ResourceUseCase) AddOrUpdateResource(ctx context.Context, req *pb.AddOrUpdateResourceReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}

	if req.GetId() == 0 {
		return c.addResource(ctx, userID, req)
	} else {
		dbResource, err := c.repo.GetResourceByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}

		return c.updateResource(ctx, dbResource, userID, req)
	}
}

func (c *ResourceUseCase) addResource(ctx context.Context, userID uint32, req *pb.AddOrUpdateResourceReq) (uint32, error) {
	resource := AddOrUpdateResourceReqToEntity(req)
	resource.CreateUser = userID
	resource.CreateTime = uint32(time.Now().Unix())
	resource.UpdateUser = userID
	resource.UpdateTime = uint32(time.Now().Unix())

	lastInsertID, err := c.repo.Save(ctx, resource)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, pb.ResourceType_TB_RESOURCE,
		"{}", resource,
		"新增接口",
	)

	return lastInsertID, nil
}

func (c *ResourceUseCase) updateResource(ctx context.Context, dbResource *Resource, userID uint32, req *pb.AddOrUpdateResourceReq) (uint32, error) {
	saveResource := AddOrUpdateResourceReqToEntity(req)
	saveResource.CreateUser = dbResource.CreateUser
	saveResource.CreateTime = dbResource.CreateTime
	saveResource.UpdateUser = userID
	saveResource.UpdateTime = uint32(time.Now().Unix())

	if err := c.repo.UpdateByID(ctx, saveResource); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_RESOURCE,
		dbResource, saveResource,
		"全字段修改接口",
	)

	return req.GetId(), nil
}

// DeleteResource 删除接口
func (c *ResourceUseCase) DeleteResource(ctx context.Context, req *pb.DeleteResourceReq) error {
	resource, err := c.repo.GetResourceByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.repo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_RESOURCE,
		resource, "{}",
		"删除接口",
	)

	return nil
}

// GetEffectiveResource 获取有效状态下的接口
func (c *ResourceUseCase) GetEffectiveResource(ctx context.Context,
	req *pb.GetEffectiveResourceReq) ([]*pb.Resource, error) {
	resources, err := c.repo.GetEffectiveResource(ctx, req)
	if err != nil {
		return nil, err
	}
	data := make([]*pb.Resource, 0)
	for _, resource := range resources {
		data = append(data, ResourceEntityToModel(resource))
	}
	return data, nil
}
