package oss

import (
	"context"
	"time"

	"github.com/baker-yuan/go-blog/go_blog/application/exterior/config"
	oss_model "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/model/oss"
	oss_repo "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/repository/oss"
	"github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/service/oss/assembler"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	oss_pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
)

// ModeCfg 模块配置
type (
	// IModeCfgUseCase 业务逻辑
	IModeCfgUseCase interface {
		// SearchModeCfg 查询模块配置
		SearchModeCfg(ctx context.Context, req *oss_pb.SearchOssModeCfgReq) ([]*oss_pb.OssModeCfg, uint32, error)
		// ModeCfgDetail 模块配置详情
		ModeCfgDetail(ctx context.Context, req *oss_pb.OssModeCfgDetailReq) (*oss_pb.OssModeCfg, error)
		// AddOrUpdateModeCfg 添加修改模块配置
		AddOrUpdateModeCfg(ctx context.Context, req *oss_pb.AddOrUpdateOssModeCfgReq) (uint32, error)
		// DeleteModeCfg 删除模块配置
		DeleteModeCfg(ctx context.Context, req *oss_pb.DeleteOssModeCfgReq) error
	}
)

// ModeCfgUseCase 模块配置管理
type ModeCfgUseCase struct {
	*CommonUseCase
	cfg  *config.Config
	repo oss_repo.IModeCfgRepo
}

// NewOssModeCfgUseCase 创建模块配置管理业务逻辑实现
func NewOssModeCfgUseCase(
	commonUseCase *CommonUseCase,
	cfg *config.Config,
	repo oss_repo.IModeCfgRepo,
) *ModeCfgUseCase {
	return &ModeCfgUseCase{
		CommonUseCase: commonUseCase,
		cfg:           cfg,
		repo:          repo,
	}
}

// ModeCfgDetail 模块配置详情
func (c *ModeCfgUseCase) ModeCfgDetail(ctx context.Context, req *oss_pb.OssModeCfgDetailReq) (*oss_pb.OssModeCfg, error) {
	ModeCfg, err := c.repo.GetModeCfgByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	// 补充操作人信息
	uidToNameMap, err := c.UidToName(ctx, ModeCfg.CreateUserID, ModeCfg.UpdateUserID)
	if err != nil {
		return nil, err
	}
	pbModeCfg := assembler.ModeCfgEntityToModel(ModeCfg, uidToNameMap)
	return pbModeCfg, nil
}

// SearchModeCfg 模块配置搜索
func (c *ModeCfgUseCase) SearchModeCfg(ctx context.Context, req *oss_pb.SearchOssModeCfgReq) ([]*oss_pb.OssModeCfg, uint32, error) {
	modeCfgList, total, err := c.repo.SearchModeCfg(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	// 补充操作人信息
	uidToNameMap, err := c.UidToName(ctx, modeCfgList.GetOperatorUserID()...)
	if err != nil {
		return nil, 0, err
	}
	return assembler.ModeCfgEntityListToModel(modeCfgList, uidToNameMap), total, nil
}

// AddOrUpdateModeCfg 添加修改模块配置
func (c *ModeCfgUseCase) AddOrUpdateModeCfg(ctx context.Context, req *oss_pb.AddOrUpdateOssModeCfgReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}
	if req.GetId() == 0 {
		return c.addModeCfg(ctx, userID, req)
	} else {
		dbModeCfg, err := c.repo.GetModeCfgByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}
		return c.updateModeCfg(ctx, dbModeCfg, userID, req)
	}
}

func (c *ModeCfgUseCase) addModeCfg(ctx context.Context, userID uint32, req *oss_pb.AddOrUpdateOssModeCfgReq) (uint32, error) {
	modeCfg := assembler.AddOrUpdateModeCfgReqToEntity(req)
	modeCfg.CreateUserID = userID
	modeCfg.CreateTime = db.Timestamp(time.Now().Unix())
	modeCfg.UpdateUserID = userID
	modeCfg.UpdateTime = db.Timestamp(time.Now().Unix())

	lastInsertID, err := c.repo.Save(ctx, modeCfg)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, oss_pb.ResourceType_TB_OSS_FILE_PATH,
		"{}", modeCfg,
		"新增模块配置",
	)

	return lastInsertID, nil
}

func (c *ModeCfgUseCase) updateModeCfg(ctx context.Context, dbModeCfg *oss_model.ModeCfg, userID uint32, req *oss_pb.AddOrUpdateOssModeCfgReq) (uint32, error) {
	saveModeCfg := assembler.AddOrUpdateModeCfgReqToEntity(req)
	saveModeCfg.CreateUserID = dbModeCfg.CreateUserID
	saveModeCfg.CreateTime = dbModeCfg.CreateTime
	saveModeCfg.UpdateUserID = userID
	saveModeCfg.UpdateTime = db.Timestamp(time.Now().Unix())

	if err := c.repo.UpdateByID(ctx, saveModeCfg); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), oss_pb.ResourceType_TB_OSS_FILE_PATH,
		dbModeCfg, saveModeCfg,
		"全字段修改模块配置",
	)

	return req.GetId(), nil
}

// DeleteModeCfg 删除模块配置
func (c *ModeCfgUseCase) DeleteModeCfg(ctx context.Context, req *oss_pb.DeleteOssModeCfgReq) error {
	modeCfg, err := c.repo.GetModeCfgByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.CheckUploadRecordNotExist(ctx, modeCfg.Mode); err != nil {
		return err
	}

	if err := c.repo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), oss_pb.ResourceType_TB_OSS_FILE_PATH,
		modeCfg, "{}",
		"删除模块配置",
	)

	return nil
}
