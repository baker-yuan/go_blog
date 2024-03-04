// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	pb "github.com/baker-yuan/go-blog/protocol/template"
	"github.com/baker-yuan/go-blog/template/internal/entity"
)

type (
	// IFriendLinkUseCase 友链service
	IFriendLinkUseCase interface {
		// SearchFriendLink 查询友链
		SearchFriendLink(ctx context.Context, req *pb.SearchFriendLinkReq) ([]*pb.FriendLink, uint32, error)
		// FriendLinkDetail 友链详情
		FriendLinkDetail(ctx context.Context, req *pb.FriendLinkDetailReq) (*pb.FriendLink, error)
		// AddOrUpdateFriendLink 添加修改友链
		AddOrUpdateFriendLink(ctx context.Context, req *pb.AddOrUpdateFriendLinkReq) (uint32, error)
		// DeleteFriendLink 删除友链
		DeleteFriendLink(ctx context.Context, req *pb.DeleteFriendLinkReq) error
	}
	// IFriendLinkRepo 友链repo
	IFriendLinkRepo interface {
		// SearchFriendLink 友链搜索
		SearchFriendLink(ctx context.Context, req *pb.SearchFriendLinkReq) (entity.FriendLinks, uint32, error)
		// GetFriendLinkByID 根据友链id集合查询友链
		GetFriendLinkByID(ctx context.Context, id int) (*entity.FriendLink, error)
		// GetFriendLinkByIDs 根据友链id集合查询友链
		GetFriendLinkByIDs(ctx context.Context, ids []int) (entity.FriendLinks, error)
		// Save 保存友链
		Save(ctx context.Context, friendLink *entity.FriendLink) (uint32, error)
		// UpdateByID 根据ID修改友链
		UpdateByID(ctx context.Context, friendLink *entity.FriendLink) error
		// DeleteByID 根据ID删除友链
		DeleteByID(ctx context.Context, id int) error
	}
)

type (
	// ICommonUseCase 通用service
	ICommonUseCase interface {
		// SaveChangeLog 保存变更日志
		SaveChangeLog(ctx context.Context,
			resourceID uint32, resourceType pb.ResourceType,
			changeBefore interface{}, changeAfter interface{},
			notes string)
	}
)
