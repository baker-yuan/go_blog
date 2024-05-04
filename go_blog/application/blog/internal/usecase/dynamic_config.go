package usecase

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/usecase/assembler"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// DynamicConfigUseCase 动态配置管理
type DynamicConfigUseCase struct {
	*CommonUseCase
	repo IDynamicConfigRepo
}

// NewDynamicConfigUseCase 创建动态配置管理业务逻辑实现
func NewDynamicConfigUseCase(
	commonUseCase *CommonUseCase,
	repo IDynamicConfigRepo,
) *DynamicConfigUseCase {
	return &DynamicConfigUseCase{
		CommonUseCase: commonUseCase,
		repo:          repo,
	}
}

// DynamicConfigDetail 动态配置详情
func (c *DynamicConfigUseCase) DynamicConfigDetail(ctx context.Context, req *pb.DynamicConfigDetailReq) (*pb.DynamicConfig, error) {
	dynamicConfig, err := c.repo.GetDynamicConfigByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	// 补充操作人信息
	uidToNameMap, err := c.UidToName(ctx, dynamicConfig.CreateUserID, dynamicConfig.UpdateUserID)
	if err != nil {
		return nil, err
	}
	pbDynamicConfig := assembler.DynamicConfigEntityToModel(dynamicConfig, uidToNameMap)
	return pbDynamicConfig, nil
}

// SearchDynamicConfig 动态配置搜索
func (c *DynamicConfigUseCase) SearchDynamicConfig(ctx context.Context, req *pb.SearchDynamicConfigReq) ([]*pb.DynamicConfig, uint32, error) {
	dynamicConfigs, total, err := c.repo.SearchDynamicConfig(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	// 补充操作人信息
	uidToNameMap, err := c.UidToName(ctx, dynamicConfigs.GetOperatorUserID()...)
	if err != nil {
		return nil, 0, err
	}
	return assembler.DynamicConfigEntityListToModel(dynamicConfigs, uidToNameMap), total, nil
}

// AddOrUpdateDynamicConfig 添加修改动态配置
func (c *DynamicConfigUseCase) AddOrUpdateDynamicConfig(ctx context.Context, req *pb.AddOrUpdateDynamicConfigReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}
	if req.GetId() == 0 {
		return c.addDynamicConfig(ctx, userID, req)
	} else {
		dbDynamicConfig, err := c.repo.GetDynamicConfigByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}
		return c.updateDynamicConfig(ctx, dbDynamicConfig, userID, req)
	}
}

func (c *DynamicConfigUseCase) addDynamicConfig(ctx context.Context, userID uint32, req *pb.AddOrUpdateDynamicConfigReq) (uint32, error) {
	dynamicConfig := assembler.AddOrUpdateDynamicConfigReqToEntity(req)
	dynamicConfig.CreateUserID = userID
	dynamicConfig.CreateTime = db.Timestamp(time.Now().Unix())
	dynamicConfig.UpdateUserID = userID
	dynamicConfig.UpdateTime = db.Timestamp(time.Now().Unix())

	lastInsertID, err := c.repo.Save(ctx, dynamicConfig)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, pb.ResourceType_TB_DYNAMIC_CONFIG,
		"{}", dynamicConfig,
		"新增动态配置",
	)

	return lastInsertID, nil
}

func (c *DynamicConfigUseCase) updateDynamicConfig(ctx context.Context, dbDynamicConfig *entity.DynamicConfig, userID uint32, req *pb.AddOrUpdateDynamicConfigReq) (uint32, error) {
	saveDynamicConfig := assembler.AddOrUpdateDynamicConfigReqToEntity(req)
	saveDynamicConfig.CreateUserID = dbDynamicConfig.CreateUserID
	saveDynamicConfig.CreateTime = dbDynamicConfig.CreateTime
	saveDynamicConfig.UpdateUserID = userID
	saveDynamicConfig.UpdateTime = db.Timestamp(time.Now().Unix())

	if err := c.repo.UpdateByID(ctx, saveDynamicConfig); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_DYNAMIC_CONFIG,
		dbDynamicConfig, saveDynamicConfig,
		"全字段修改动态配置",
	)

	return req.GetId(), nil
}

// DeleteDynamicConfig 删除动态配置
func (c *DynamicConfigUseCase) DeleteDynamicConfig(ctx context.Context, req *pb.DeleteDynamicConfigReq) error {
	dynamicConfig, err := c.repo.GetDynamicConfigByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.repo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_DYNAMIC_CONFIG,
		dynamicConfig, "{}",
		"删除动态配置",
	)

	return nil
}
