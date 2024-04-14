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

// FriendLinkUseCase 友链表管理
type FriendLinkUseCase struct {
	*CommonUseCase
	repo IFriendLinkRepo
}

// NewFriendLinkUseCase 创建友链表管理业务逻辑实现
func NewFriendLinkUseCase(
	commonUseCase *CommonUseCase,
	repo IFriendLinkRepo,
) *FriendLinkUseCase {
	return &FriendLinkUseCase{
		CommonUseCase: commonUseCase,
		repo:          repo,
	}
}

// FriendLinkDetail 友链表详情
func (c *FriendLinkUseCase) FriendLinkDetail(ctx context.Context, req *pb.FriendLinkDetailReq) (*pb.FriendLink, error) {
	friendLink, err := c.repo.GetFriendLinkByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	// 补充操作人信息
	uidToNameMap, err := c.UidToName(ctx, friendLink.CreateUserID, friendLink.UpdateUserID)
	if err != nil {
		return nil, err
	}
	pbFriendLink := assembler.FriendLinkEntityToModel(friendLink, uidToNameMap)
	return pbFriendLink, nil
}

// SearchFriendLink 友链表搜索
func (c *FriendLinkUseCase) SearchFriendLink(ctx context.Context, req *pb.SearchFriendLinkReq) ([]*pb.FriendLink, uint32, error) {
	friendLinks, total, err := c.repo.SearchFriendLink(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	// 补充操作人信息
	uidToNameMap, err := c.UidToName(ctx, friendLinks.GetOperatorUserID()...)
	if err != nil {
		return nil, 0, err
	}
	return assembler.FriendLinkEntityListToModel(friendLinks, uidToNameMap), total, nil
}

// AddOrUpdateFriendLink 添加修改友链表
func (c *FriendLinkUseCase) AddOrUpdateFriendLink(ctx context.Context, req *pb.AddOrUpdateFriendLinkReq) (uint32, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return 0, err
	}
	if req.GetId() == 0 {
		return c.addFriendLink(ctx, userID, req)
	} else {
		dbFriendLink, err := c.repo.GetFriendLinkByID(ctx, req.GetId())
		if err != nil {
			return 0, err
		}
		return c.updateFriendLink(ctx, dbFriendLink, userID, req)
	}
}

func (c *FriendLinkUseCase) addFriendLink(ctx context.Context, userID uint32, req *pb.AddOrUpdateFriendLinkReq) (uint32, error) {
	friendLink := assembler.AddOrUpdateFriendLinkReqToEntity(req)
	friendLink.CreateUserID = userID
	friendLink.CreateTime = db.Timestamp(time.Now().Unix())
	friendLink.UpdateUserID = userID
	friendLink.UpdateTime = db.Timestamp(time.Now().Unix())

	lastInsertID, err := c.repo.Save(ctx, friendLink)
	if err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		lastInsertID, pb.ResourceType_TB_FRIEND_LINK,
		"{}", friendLink,
		"新增友链表",
	)

	return lastInsertID, nil
}

func (c *FriendLinkUseCase) updateFriendLink(ctx context.Context, dbFriendLink *entity.FriendLink, userID uint32, req *pb.AddOrUpdateFriendLinkReq) (uint32, error) {
	saveFriendLink := assembler.AddOrUpdateFriendLinkReqToEntity(req)
	saveFriendLink.CreateUserID = dbFriendLink.CreateUserID
	saveFriendLink.CreateTime = dbFriendLink.CreateTime
	saveFriendLink.UpdateUserID = userID
	saveFriendLink.UpdateTime = db.Timestamp(time.Now().Unix())

	if err := c.repo.UpdateByID(ctx, saveFriendLink); err != nil {
		return 0, err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_FRIEND_LINK,
		dbFriendLink, saveFriendLink,
		"全字段修改友链表",
	)

	return req.GetId(), nil
}

// DeleteFriendLink 删除友链表
func (c *FriendLinkUseCase) DeleteFriendLink(ctx context.Context, req *pb.DeleteFriendLinkReq) error {
	friendLink, err := c.repo.GetFriendLinkByID(ctx, req.GetId())
	if err != nil {
		return err
	}

	if err := c.repo.DeleteByID(ctx, req.GetId()); err != nil {
		return err
	}

	c.SaveChangeLog(ctx,
		req.GetId(), pb.ResourceType_TB_FRIEND_LINK,
		friendLink, "{}",
		"删除友链表",
	)

	return nil
}
