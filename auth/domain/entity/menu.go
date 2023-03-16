package do

import "github.com/baker-yuan/go-blog/all_packaged_library/common"

// Menu 菜单
type Menu struct {
	common.Model
	Name      string `json:"name"`      // 菜单名
	Path      string `json:"path"`      // 路径
	Component string `json:"component"` // 组件
	Icon      string `json:"icon"`      // icon
	OrderNum  uint32 `json:"orderNum"`  // 排序
	ParentId  uint32 `json:"parentId"`  // 父id
	IsHidden  uint32 `json:"isHidden"`  // 是否隐藏
}
