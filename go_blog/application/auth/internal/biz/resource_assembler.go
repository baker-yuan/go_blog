package biz

import (
	"strings"

	pb "github.com/baker-yuan/go-blog/go_blog/protocol/auth"
)

// ResourceEntityListToModel entity转pb
func ResourceEntityListToModel(resourceList []*Resource, uidToNameMap map[uint32]string, treeNodeMap map[uint32][]*pb.TreeNode) []*pb.Resource {
	res := make([]*pb.Resource, 0)
	for _, item := range resourceList {
		res = append(res, ResourceEntityToModel(item, uidToNameMap, treeNodeMap))
	}
	return res
}

// ResourceEntityToModel entity转pb
func ResourceEntityToModel(resource *Resource, uidToNameMap map[uint32]string, treeNodeMap map[uint32][]*pb.TreeNode) *pb.Resource {
	modelRes := &pb.Resource{
		// 基本信息
		Id:         resource.ID,
		TreeNodeId: resource.TreeNodeID,
		Name:       resource.Name,
		Describe:   resource.Describe,
		//
		Status: resource.Status,
		// 路径定位
		Url:        resource.URL,
		HttpMethod: resource.HTTPMethod,
		// 权限校验
		IsNeedLogin:      resource.IsNeedLogin,
		IsNeedPermission: resource.IsNeedPermission,
		// 下游服务信息
		Service: resource.Service,
		Method:  resource.Method,
		// 公共字段
		CreateUserId: resource.CreateUserID,
		UpdateUserId: resource.UpdateUserID,
		CreateTime:   resource.CreateTime,
		UpdateTime:   resource.UpdateTime,
	}
	if node, exist := treeNodeMap[resource.TreeNodeID]; exist {
		treeNodePath := make([]string, 0)
		for i := len(node) - 1; i >= 0; i-- {
			treeNodePath = append(treeNodePath, node[i].Name)
		}
		modelRes.TreeNodePath = strings.Join(treeNodePath, "/")
	}
	modelRes.CreateUser = uidToNameMap[resource.CreateUserID]
	modelRes.UpdateUser = uidToNameMap[resource.UpdateUserID]
	return modelRes
}

// AddOrUpdateResourceReqToEntity pb转entity
func AddOrUpdateResourceReqToEntity(pbResource *pb.AddOrUpdateResourceReq) *Resource {
	entityRes := &Resource{
		// 基本信息
		ID:         pbResource.Id,
		TreeNodeID: pbResource.TreeNodeId,
		Name:       pbResource.Name,
		Describe:   pbResource.Describe,
		//
		Status: pbResource.Status,
		// 路径定位
		URL:        pbResource.Url,
		HTTPMethod: pbResource.HttpMethod,
		// 权限校验
		IsNeedLogin:      pbResource.IsNeedLogin,
		IsNeedPermission: pbResource.IsNeedPermission,
		// 下游服务信息
		Service: pbResource.Service,
		Method:  pbResource.Method,
	}
	return entityRes
}
