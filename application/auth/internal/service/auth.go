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
