package service

import (
	"context"

	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	pb "github.com/baker-yuan/go-blog/protocol/auth"
)

func NewAuthService(article *biz.MenuUsecase) *AuthService {
	return &AuthService{
		article: article,
	}
}

// SearchMenu 菜单搜索
func (a AuthService) SearchMenu(ctx context.Context, req *pb.SearchMenuReq) (*pb.SearchMenuRsp, error) {
	//TODO implement me
	panic("implement me")
}

// AddOrUpdateMenu 添加修改菜单
func (a AuthService) AddOrUpdateMenu(ctx context.Context, req *pb.AddOrUpdateMenuReq) (*pb.AddOrUpdateRsp, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteMenu 删除菜单
func (a AuthService) DeleteMenu(ctx context.Context, req *pb.DeleteMenuReq) (*pb.EmptyRsp, error) {
	//TODO implement me
	panic("implement me")
}

// MenuDetail 菜单详情
func (a AuthService) MenuDetail(ctx context.Context, req *pb.MenuDetailReq) (*pb.Menu, error) {
	//TODO implement me
	panic("implement me")
}

// SearchResource 接口搜索
func (a AuthService) SearchResource(ctx context.Context, req *pb.SearchResourceReq) (*pb.SearchResourceRsp, error) {
	//TODO implement me
	panic("implement me")
}

// AddOrUpdateResource 添加修改接口
func (a AuthService) AddOrUpdateResource(ctx context.Context, req *pb.AddOrUpdateResourceReq) (*pb.AddOrUpdateRsp, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteResource 删除接口
func (a AuthService) DeleteResource(ctx context.Context, req *pb.DeleteResourceReq) (*pb.EmptyRsp, error) {
	//TODO implement me
	panic("implement me")
}

// ResourceDetail 接口详情
func (a AuthService) ResourceDetail(ctx context.Context, req *pb.ResourceDetailReq) (*pb.Resource, error) {
	//TODO implement me
	panic("implement me")
}

// GetEffectiveResource 获取有效状态下的接口
func (a AuthService) GetEffectiveResource(ctx context.Context,
	req *pb.GetEffectiveResourceReq) (*pb.GetEffectiveResourceRsp, error) {
	//TODO implement me
	panic("implement me")
}
