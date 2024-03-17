package biz

import pb "github.com/baker-yuan/go-blog/protocol/auth"

// ResourceEntityToModel entity转pb
func ResourceEntityToModel(resource *Resource) *pb.Resource {
	modelRes := &pb.Resource{
		// 基本信息
		Id:       resource.ID,
		ParentId: resource.ParentID,
		Name:     resource.Name,
		Describe: resource.Describe,
		//
		ResourceType: resource.ResourceType,
		Status:       resource.Status,
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
		CreateUser: resource.CreateUser,
		UpdateUser: resource.UpdateUser,
		CreateTime: resource.CreateTime,
		UpdateTime: resource.UpdateTime,
	}
	return modelRes
}

// AddOrUpdateResourceReqToEntity pb转entity
func AddOrUpdateResourceReqToEntity(pbResource *pb.AddOrUpdateResourceReq) *Resource {
	entityRes := &Resource{
		// 基本信息
		ID:       pbResource.Id,
		ParentID: pbResource.ParentId,
		Name:     pbResource.Name,
		Describe: pbResource.Describe,
		//
		ResourceType: pbResource.ResourceType,
		Status:       pbResource.Status,
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
