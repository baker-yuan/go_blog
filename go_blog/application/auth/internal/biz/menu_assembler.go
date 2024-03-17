package biz

import pb "github.com/baker-yuan/go-blog/protocol/auth"

// MenuEntityToModel entity转pb
func MenuEntityToModel(menu *Menu) *pb.Menu {
	modelRes := &pb.Menu{
		// 基本信息
		Id:        menu.ID,
		ParentId:  menu.ParentID,
		Name:      menu.Name,
		Icon:      menu.Icon,
		Describe:  menu.Describe,
		SortValue: menu.SortValue,
		// 前端
		Path:      menu.Path,
		Component: menu.Component,
		// 状态
		IsHidden: menu.IsHidden,
		// 公共字段
		CreateUser: menu.CreateUser,
		UpdateUser: menu.UpdateUser,
		CreateTime: menu.CreateTime,
		UpdateTime: menu.UpdateTime,
	}
	return modelRes
}

// AddOrUpdateMenuReqToEntity pb转entity
func AddOrUpdateMenuReqToEntity(pbMenu *pb.AddOrUpdateMenuReq) *Menu {
	entityRes := &Menu{
		// 基本信息
		ID:        pbMenu.Id,
		ParentID:  pbMenu.ParentId,
		Name:      pbMenu.Name,
		Icon:      pbMenu.Icon,
		Describe:  pbMenu.Describe,
		SortValue: pbMenu.SortValue,
		// 前端
		Path:      pbMenu.Path,
		Component: pbMenu.Component,
		// 状态
		IsHidden: pbMenu.IsHidden,
	}
	return entityRes
}
