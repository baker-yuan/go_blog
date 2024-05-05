package data

import (
	"github.com/baker-yuan/go-blog/go_blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
)

func ResourcePOsToEntity(dbResources []*Resource) []*biz.Resource {
	res := make([]*biz.Resource, 0)
	for _, dbResource := range dbResources {
		res = append(res, ResourcePOToEntity(dbResource))
	}
	return res
}

func ResourcePOToEntity(dbResource *Resource) *biz.Resource {
	res := &biz.Resource{
		// 基本信息
		ID:         dbResource.ID,
		TreeNodeID: dbResource.TreeNodeID,
		Name:       dbResource.Name,
		Describe:   dbResource.Describe,
		// 个性化
		IsNeedStdRes: dbResource.IsNeedStdRes,
		//
		Status: dbResource.Status,
		// 路径定位
		URL:        dbResource.URL,
		HTTPMethod: dbResource.HTTPMethod,
		// 权限校验
		IsNeedLogin:      dbResource.IsNeedLogin,
		IsNeedPermission: dbResource.IsNeedPermission,
		// 下游服务信息
		Service:  dbResource.Service,
		Method:   dbResource.Method,
		Protocol: dbResource.Protocol,
		// 公共字段
		CreateUserID: dbResource.CreateUserID,
		UpdateUserID: dbResource.UpdateUserID,
		CreateTime:   uint32(dbResource.CreateTime),
		UpdateTime:   uint32(dbResource.UpdateTime),
	}
	return res
}

func ResourceEntityToPO(resource *biz.Resource) *Resource {
	res := &Resource{
		// 基本信息
		ID:         resource.ID,
		TreeNodeID: resource.TreeNodeID,
		Name:       resource.Name,
		Describe:   resource.Describe,
		// 个性化
		IsNeedStdRes: resource.IsNeedStdRes,
		//
		Status: resource.Status,
		// 路径定位
		URL:        resource.URL,
		HTTPMethod: resource.HTTPMethod,
		// 权限校验
		IsNeedLogin:      resource.IsNeedLogin,
		IsNeedPermission: resource.IsNeedPermission,
		// 下游服务信息
		Service:  resource.Service,
		Method:   resource.Method,
		Protocol: resource.Protocol,
		// 公共字段
		CreateUserID: resource.CreateUserID,
		UpdateUserID: resource.UpdateUserID,
		CreateTime:   db.Timestamp(resource.CreateTime),
		UpdateTime:   db.Timestamp(resource.UpdateTime),
	}
	return res
}
