package usecase

import (
	"context"

	pb "github.com/baker-yuan/go-blog/protocol/template"
	"github.com/baker-yuan/go-blog/template/internal/usecase/assembler"
)

// FriendLinkUseCase 友链管理
type FriendLinkUseCase struct {
	ICommonUseCase
	friendLinkRepo IFriendLinkRepo
}

// NewFriendLinkUseCase 创建友链管理service
func NewFriendLinkUseCase(
	commonUseCase ICommonUseCase,
	friendLinkRepo IFriendLinkRepo,
) *FriendLinkUseCase {
	return &FriendLinkUseCase{
		ICommonUseCase: commonUseCase,
		friendLinkRepo: friendLinkRepo,
	}
}

// FriendLinkDetail 友链详情
func (c *FriendLinkUseCase) FriendLinkDetail(ctx context.Context, req *pb.FriendLinkDetailReq) (*pb.FriendLink, error) {
	friendLink, err := c.friendLinkRepo.GetFriendLinkByID(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	pbFriendLink := assembler.FriendLinkEntityToModel(friendLink)
	return pbFriendLink, nil
}

// SearchFriendLink 友链搜索
func (c *FriendLinkUseCase) SearchFriendLink(ctx context.Context, req *pb.SearchFriendLinkReq) ([]*pb.FriendLink, uint32, error) {
	friendLinks, total, err := c.friendLinkRepo.SearchFriendLink(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	data := make([]*pb.FriendLink, 0)
	for _, friendLink := range friendLinks {
		data = append(data, assembler.FriendLinkEntityToModel(friendLink))
	}
	return data, total, nil
}

// AddOrUpdateFriendLink 添加修改友链
func (c *FriendLinkUseCase) AddOrUpdateFriendLink(ctx context.Context, req *pb.AddOrUpdateFriendLinkReq) (uint32, error) {
	//loginName, err := pkg_util.GetLoginStaffName(ctx)
	//if err != nil {
	//	return 0, err
	//}
	//if err := c.CheckResAuth(ctx, loginName); err != nil {
	//	return 0, err
	//}
	//
	//if req.GetId() == 0 {
	//	return c.addFriendLink(ctx, loginName, req)
	//} else {
	//	dbFriendLink, err := c.friendLinkRepo.GetFriendLinkByID(ctx, int(req.GetId()))
	//	if err != nil {
	//		return 0, err
	//	}
	//
	//	return c.updateFriendLink(ctx, dbFriendLink, loginName, req)
	//}

	return 0, nil
}

//
//func (c *FriendLinkUseCase) addFriendLink(ctx context.Context, loginName string, req *pb.AddOrUpdateFriendLinkReq) (uint32, error) {
//	friendLink := assembler.AddOrUpdateFriendLinkReqToEntity(req)
//	friendLink.AddTime = sql.NullTime{Time: time.Now(), Valid: true}
//	friendLink.AddOperator = loginName
//	friendLink.LastChgTime = sql.NullTime{Time: time.Now(), Valid: true}
//	friendLink.LastChgUser = loginName
//
//	lastInsertID, err := c.friendLinkRepo.Save(ctx, friendLink)
//	if err != nil {
//		return 0, err
//	}
//
//	c.SaveChangeLog(ctx,
//		lastInsertID, pb.ResourceType_RT_,
//		"{}", friendLink,
//		"新增友链",
//	)
//
//	return lastInsertID, nil
//}
//
//func (c *FriendLinkUseCase) updateFriendLink(ctx context.Context, dbFriendLink *entity.FriendLink, loginName string, req *pb.AddOrUpdateFriendLinkReq) (uint32, error) {
//	saveFriendLink := assembler.AddOrUpdateFriendLinkReqToEntity(req)
//	saveFriendLink.AddTime = dbFriendLink.AddTime
//	saveFriendLink.AddOperator = dbFriendLink.AddOperator
//	saveFriendLink.LastChgTime = sql.NullTime{Time: time.Now(), Valid: true}
//	saveFriendLink.LastChgUser = loginName
//
//	if err := c.friendLinkRepo.UpdateByID(ctx, saveFriendLink); err != nil {
//		return 0, err
//	}
//
//	c.SaveChangeLog(ctx,
//		req.GetId(), pb.ResourceType_RT_,
//		dbFriendLink, saveFriendLink,
//		"全字段修改友链",
//	)
//
//	return req.GetId(), nil
//}

// DeleteFriendLink 删除友链
func (c *FriendLinkUseCase) DeleteFriendLink(ctx context.Context, req *pb.DeleteFriendLinkReq) error {
	//loginName, err := pkg_util.GetLoginStaffName(ctx)
	//if err != nil {
	//	return err
	//}
	//if err := c.CheckResAuth(ctx, loginName); err != nil {
	//	return err
	//}
	//
	//friendLink, err := c.friendLinkRepo.GetFriendLinkByID(ctx, int(req.GetId()))
	//if err != nil {
	//	return err
	//}
	//
	//if err := c.friendLinkRepo.DeleteByID(ctx, int(req.GetId())); err != nil {
	//	return err
	//}
	//
	//c.SaveChangeLog(ctx,
	//	req.GetId(), pb.ResourceType_RT_,
	//	friendLink, "{}",
	//	"删除友链",
	//)

	return nil
}
