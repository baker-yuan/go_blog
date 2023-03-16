package do

import "github.com/baker-yuan/go-blog/all_packaged_library/common"

// MenuDO 菜单
type MenuDO struct {
	common.Model
	Name      string `json:"name"`      // 菜单名
	Path      string `json:"path"`      // 路径
	Component string `json:"component"` // 组件
	Icon      string `json:"icon"`      // icon
	OrderNum  uint32 `json:"orderNum"`  // 排序
	ParentId  uint32 `json:"parentId"`  // 父id
	IsHidden  uint32 `json:"isHidden"`  // 是否隐藏
}

func (a MenuDO) TableName() string {
	return "tb_menu"
}

// RoleMenu 角色菜单
type RoleMenu struct {
	Id     uint32 `json:"id"`     // 主键id
	RoleId uint32 `json:"roleId"` // 角色id
	MenuId uint32 `json:"menuId"` // 菜单id
}

func (a RoleMenu) TableName() string {
	return "tb_role_menu"
}
