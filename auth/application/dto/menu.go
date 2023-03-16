package dto

// UserMenuDTO 用户菜单
type UserMenuDTO struct {
	Name      string        `json:"name"`      // 菜单名
	Path      string        `json:"path"`      // 路径
	Component string        `json:"component"` // 组件
	Icon      string        `json:"icon"`      // icon
	Hidden    bool          `json:"hidden"`    // 是否隐藏
	Children  []UserMenuDTO `json:"children"`  // 子菜单列表
}

type MenuSearch struct {
	Name   *string `json:"name"`   // 菜单名
	RoleId *uint32 `json:"roleId"` // 角色id
}
