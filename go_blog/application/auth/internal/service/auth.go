package service

import (
	"context"

	"github.com/baker-yuan/go-blog/go_blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
)

// NewAuthService 权限
func NewAuthService(
	menu *biz.MenuUsecase,
	resource *biz.ResourceUseCase,
	treeNode *biz.TreeNodeUseCase,
	role *biz.RoleUseCase,
	userRoleAuthority *biz.UserRoleAuthorityUseCase,
) *AuthService {
	return &AuthService{
		menu:              menu,
		resource:          resource,
		treeNode:          treeNode,
		role:              role,
		userRoleAuthority: userRoleAuthority,
	}
}

// SearchMenu 菜单搜索
func (s *AuthService) SearchMenu(ctx context.Context, req *pb.SearchMenuReq) (*pb.SearchMenuRsp, error) {
	rsp := &pb.SearchMenuRsp{}
	menus, pageTotal, err := s.menu.SearchMenu(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = menus
	rsp.PageTotal = pageTotal
	return rsp, nil
}

// AddOrUpdateMenu 添加修改菜单
func (s *AuthService) AddOrUpdateMenu(ctx context.Context, req *pb.AddOrUpdateMenuReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := s.menu.AddOrUpdateMenu(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteMenu 删除菜单
func (s *AuthService) DeleteMenu(ctx context.Context, req *pb.DeleteMenuReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := s.menu.DeleteMenu(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// MenuDetail 菜单详情
func (s *AuthService) MenuDetail(ctx context.Context, req *pb.MenuDetailReq) (*pb.Menu, error) {
	return s.menu.MenuDetail(ctx, req)
}

// GetMenuTree 菜单树
func (s *AuthService) GetMenuTree(ctx context.Context, req *pb.GetMenuTreeReq) (*pb.GetMenuTreeRsp, error) {
	rsp := &pb.GetMenuTreeRsp{}
	nodes, pageTotal, err := s.menu.GetMenuTree(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Nodes = nodes
	rsp.PageTotal = pageTotal
	return rsp, nil
}

// SearchResource 接口搜索
func (s *AuthService) SearchResource(ctx context.Context, req *pb.SearchResourceReq) (*pb.SearchResourceRsp, error) {
	rsp := &pb.SearchResourceRsp{}
	resources, pageTotal, err := s.resource.SearchResource(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = resources
	rsp.PageTotal = pageTotal
	return rsp, nil
}

// AddOrUpdateResource 添加修改接口
func (s *AuthService) AddOrUpdateResource(ctx context.Context, req *pb.AddOrUpdateResourceReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := s.resource.AddOrUpdateResource(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteResource 删除接口
func (s *AuthService) DeleteResource(ctx context.Context, req *pb.DeleteResourceReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := s.resource.DeleteResource(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// ResourceDetail 接口详情
func (s *AuthService) ResourceDetail(ctx context.Context, req *pb.ResourceDetailReq) (*pb.Resource, error) {
	return s.resource.ResourceDetail(ctx, req)
}

// GetEffectiveResource 获取有效状态下的接口
func (s *AuthService) GetEffectiveResource(ctx context.Context,
	req *pb.GetEffectiveResourceReq) (*pb.GetEffectiveResourceRsp, error) {
	res := &pb.GetEffectiveResourceRsp{}
	data, err := s.resource.GetEffectiveResource(ctx, req)
	if err != nil {
		return nil, err
	}
	res.Data = data
	return res, nil
}

// SearchRole 查询角色
func (s *AuthService) SearchRole(ctx context.Context, req *pb.SearchRoleReq) (*pb.SearchRoleRsp, error) {
	rsp := &pb.SearchRoleRsp{}
	roles, pageTotal, err := s.role.SearchRole(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = roles
	rsp.PageTotal = pageTotal
	return rsp, nil
}

// RoleDetail 角色详情
func (s *AuthService) RoleDetail(ctx context.Context, req *pb.RoleDetailReq) (*pb.Role, error) {
	return s.role.RoleDetail(ctx, req)
}

// AddOrUpdateRole 添加修改角色
func (s *AuthService) AddOrUpdateRole(ctx context.Context, req *pb.AddOrUpdateRoleReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := s.role.AddOrUpdateRole(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteRole 删除角色
func (s *AuthService) DeleteRole(ctx context.Context, req *pb.DeleteRoleReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := s.role.DeleteRole(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// RoleBindMenu 角色绑定菜单权限
func (s *AuthService) RoleBindMenu(ctx context.Context, req *pb.RoleBindMenuReq) (*pb.EmptyRsp, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return nil, err
	}
	rsp := &pb.EmptyRsp{}
	if err := s.userRoleAuthority.RoleBindMenu(ctx, userID, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// RoleBindResource 角色绑定资源权限
func (s *AuthService) RoleBindResource(ctx context.Context, req *pb.RoleBindResourceReq) (*pb.EmptyRsp, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return nil, err
	}
	rsp := &pb.EmptyRsp{}
	if err := s.userRoleAuthority.RoleBindResource(ctx, userID, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// UserBindRole 用户绑定角色
func (s *AuthService) UserBindRole(ctx context.Context, req *pb.UserBindRoleReq) (*pb.EmptyRsp, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return nil, err
	}
	rsp := &pb.EmptyRsp{}
	if err := s.userRoleAuthority.UserBindRole(ctx, userID, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetUserRole 获取用户绑定的角色
func (s *AuthService) GetUserRole(ctx context.Context, req *pb.GetUserRoleReq) (*pb.GetUserRoleRsp, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return nil, err
	}
	rsp := &pb.GetUserRoleRsp{}
	roles, err := s.userRoleAuthority.GetUserRole(ctx, userID, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = roles
	return rsp, nil
}

// GetUserResource 获取用户关联的接口权限
func (s *AuthService) GetUserResource(ctx context.Context, req *pb.GetUserResourceReq) (*pb.GetUserResourceRsp, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return nil, err
	}
	rsp := &pb.GetUserResourceRsp{}
	ids, err := s.userRoleAuthority.GetUserResource(ctx, userID, req)
	if err != nil {
		return nil, err
	}
	rsp.Ids = ids
	return rsp, nil
}

// SearchTreeNode 查询树形结构表
func (s *AuthService) SearchTreeNode(ctx context.Context, req *pb.SearchTreeNodeReq) (*pb.SearchTreeNodeRsp, error) {
	rsp := &pb.SearchTreeNodeRsp{}
	treeNodes, pageTotal, err := s.treeNode.SearchTreeNode(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = treeNodes
	rsp.PageTotal = pageTotal
	return rsp, nil
}

// AddOrUpdateTreeNode 添加修改树形结构表
func (s *AuthService) AddOrUpdateTreeNode(ctx context.Context, req *pb.AddOrUpdateTreeNodeReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := s.treeNode.AddOrUpdateTreeNode(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteTreeNode 删除树形结构表
func (s *AuthService) DeleteTreeNode(ctx context.Context, req *pb.DeleteTreeNodeReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := s.treeNode.DeleteTreeNode(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// TreeNodeDetail 树形结构表详情
func (s *AuthService) TreeNodeDetail(ctx context.Context, req *pb.TreeNodeDetailReq) (*pb.TreeNode, error) {
	return s.treeNode.TreeNodeDetail(ctx, req)
}

// GetTreeByBizType 获取指定业务类型的树形结构
func (s *AuthService) GetTreeByBizType(ctx context.Context, req *pb.GetTreeByBizTypeReq) (*pb.GetTreeByBizTypeRsp, error) {
	rsp := &pb.GetTreeByBizTypeRsp{}
	nodes, err := s.treeNode.GetTreeByBizType(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Nodes = nodes
	return rsp, nil
}
