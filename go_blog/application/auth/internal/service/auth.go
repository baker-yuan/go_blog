package service

import (
	"context"

	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	pb "github.com/baker-yuan/go-blog/protocol/auth"
)

func NewAuthService(menu *biz.MenuUsecase, resource *biz.ResourceUseCase) *AuthService {
	return &AuthService{
		menu:     menu,
		resource: resource,
	}
}

// SearchMenu 菜单搜索
func (a AuthService) SearchMenu(ctx context.Context, req *pb.SearchMenuReq) (*pb.SearchMenuRsp, error) {
	rsp := &pb.SearchMenuRsp{}
	menus, pageTotal, err := a.menu.SearchMenu(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = menus
	rsp.Total = pageTotal
	return rsp, nil
}

// AddOrUpdateMenu 添加修改菜单
func (a AuthService) AddOrUpdateMenu(ctx context.Context, req *pb.AddOrUpdateMenuReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := a.menu.AddOrUpdateMenu(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteMenu 删除菜单
func (a AuthService) DeleteMenu(ctx context.Context, req *pb.DeleteMenuReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := a.menu.DeleteMenu(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// MenuDetail 菜单详情
func (a AuthService) MenuDetail(ctx context.Context, req *pb.MenuDetailReq) (*pb.Menu, error) {
	return a.menu.MenuDetail(ctx, req)
}

// SearchResource 接口搜索
func (a AuthService) SearchResource(ctx context.Context, req *pb.SearchResourceReq) (*pb.SearchResourceRsp, error) {
	rsp := &pb.SearchResourceRsp{}
	resources, pageTotal, err := a.resource.SearchResource(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = resources
	rsp.Total = pageTotal
	return rsp, nil
}

// AddOrUpdateResource 添加修改接口
func (a AuthService) AddOrUpdateResource(ctx context.Context, req *pb.AddOrUpdateResourceReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := a.resource.AddOrUpdateResource(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteResource 删除接口
func (a AuthService) DeleteResource(ctx context.Context, req *pb.DeleteResourceReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := a.resource.DeleteResource(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// ResourceDetail 接口详情
func (a AuthService) ResourceDetail(ctx context.Context, req *pb.ResourceDetailReq) (*pb.Resource, error) {
	return a.resource.ResourceDetail(ctx, req)
}

// GetEffectiveResource 获取有效状态下的接口
func (a AuthService) GetEffectiveResource(ctx context.Context,
	req *pb.GetEffectiveResourceReq) (*pb.GetEffectiveResourceRsp, error) {
	res := &pb.GetEffectiveResourceRsp{}
	data, err := a.resource.GetEffectiveResource(ctx, req)
	if err != nil {
		return nil, err
	}
	res.Data = data
	return res, nil
}
