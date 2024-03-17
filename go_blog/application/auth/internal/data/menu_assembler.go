package data

import (
	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	"github.com/baker-yuan/go-blog/common/db"
)

func MenuPOsToEntity(dbMenus []*Menu) []*biz.Menu {
	res := make([]*biz.Menu, 0)
	for _, dbMenu := range dbMenus {
		res = append(res, MenuPOToEntity(dbMenu))
	}
	return res
}

func MenuPOToEntity(dbMenu *Menu) *biz.Menu {
	return &biz.Menu{
		// 基本信息
		ID:        dbMenu.ID,
		ParentID:  dbMenu.ParentID,
		Name:      dbMenu.Name,
		Icon:      dbMenu.Icon,
		Describe:  dbMenu.Describe,
		SortValue: dbMenu.SortValue,
		// 前端
		Path:      dbMenu.Path,
		Component: dbMenu.Component,
		// 状态
		IsHidden: bool(dbMenu.IsHidden),
		// 公共字段
		CreateUser: dbMenu.CreateUser,
		UpdateUser: dbMenu.UpdateUser,
		CreateTime: uint32(dbMenu.CreateTime),
		UpdateTime: uint32(dbMenu.UpdateTime),
	}
}

func MenuEntityToPO(menu *biz.Menu) *Menu {
	return &Menu{
		// 基本信息
		ID:        menu.ID,
		ParentID:  menu.ParentID,
		Name:      menu.Name,
		Icon:      menu.Icon,
		Describe:  menu.Describe,
		SortValue: menu.SortValue,
		// 前端
		Path:      menu.Path,
		Component: menu.Component,
		// 状态
		IsHidden: db.BoolBit(menu.IsHidden),
		// 公共字段
		CreateUser: menu.CreateUser,
		UpdateUser: menu.UpdateUser,
		CreateTime: db.Timestamp(menu.CreateTime),
		UpdateTime: db.Timestamp(menu.UpdateTime),
	}
}
