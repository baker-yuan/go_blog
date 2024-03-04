package controller

import (
	"context"

	pb "github.com/baker-yuan/go-blog/protocol/template"
	"github.com/baker-yuan/go-blog/template/internal/usecase"
	"google.golang.org/protobuf/proto"
)

// TemplateServiceImpl 接口实现
type TemplateServiceImpl struct {
	friendLinkUseCase usecase.IFriendLinkUseCase
}

// NewTemplateServiceImpl 创建接口实现
func NewTemplateServiceImpl(friendLinkUseCase usecase.IFriendLinkUseCase) pb.TemplateApiService {
	return &TemplateServiceImpl{
		friendLinkUseCase: friendLinkUseCase,
	}
}

// SearchFriendLink 查询友链
func (m *TemplateServiceImpl) SearchFriendLink(ctx context.Context, req *pb.SearchFriendLinkReq) (*pb.SearchFriendLinkRsp, error) {
	rsp := &pb.SearchFriendLinkRsp{}
	friendLinks, pageTotal, err := m.friendLinkUseCase.SearchFriendLink(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = friendLinks
	rsp.Total = proto.Uint32(pageTotal)
	return rsp, nil
}

// FriendLinkDetail 友链详情
func (m *TemplateServiceImpl) FriendLinkDetail(ctx context.Context, req *pb.FriendLinkDetailReq) (*pb.FriendLink, error) {
	return m.friendLinkUseCase.FriendLinkDetail(ctx, req)
}

// AddOrUpdateFriendLink 添加修改友链
func (m *TemplateServiceImpl) AddOrUpdateFriendLink(ctx context.Context, req *pb.AddOrUpdateFriendLinkReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := m.friendLinkUseCase.AddOrUpdateFriendLink(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = proto.Uint32(id)
	return rsp, nil
}

// DeleteFriendLink 删除友链
func (m *TemplateServiceImpl) DeleteFriendLink(ctx context.Context, req *pb.DeleteFriendLinkReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := m.friendLinkUseCase.DeleteFriendLink(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}
