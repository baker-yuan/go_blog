package repository

import (
	"github.com/baker-yuan/go-blog/all_packaged_library/base/db"
	"github.com/baker-yuan/go-blog/all_packaged_library/constant"
	"github.com/baker-yuan/go-blog/all_packaged_library/util"
	"github.com/baker-yuan/go-blog/auth/application/dto"
	"github.com/baker-yuan/go-blog/auth/infrastructure/repository/do"
	"gorm.io/gorm"
)

type MenuRepoImpl struct {
}

func (m *MenuRepoImpl) searchMenu(search dto.MenuSearch) ([]*do.MenuDO, error) {
	// var (
	// 	query   = make(map[string]interface{}, 0)
	// 	menus   = make([]*do.MenuDO, 0)
	// 	current uint32
	// 	total   int64
	// 	tx      *gorm.DB
	// )
	// if search.RoleId != nil {
	//
	// }
	return nil, nil
}

func (m *MenuRepoImpl) listMenusByRoleId(roleId uint32) ([]*do.MenuDO, error) {
	var (
		DB        = db.GetMysqlDb()
		roleMenus = make([]*do.RoleMenu, 0)
		tx        *gorm.DB
		menuIds   = make([]uint32, 0)
		menus     = make([]*do.MenuDO, 0)
	)
	tx = DB.Raw("select * from tb_role_menu where role_id = ? ", roleId).Scan(&roleMenus)
	if tx.Error != nil {
		return nil, tx.Error
	}
	for _, v := range roleMenus {
		menuIds = append(menuIds, v.MenuId)
	}
	tx = DB.Raw("select * from tb_menu where id in ?", menuIds).Scan(&menus)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return menus, nil
}

// ListUserMenus 查看用户菜单
//
// @return 菜单列表
func (m *MenuRepoImpl) ListUserMenus(roleId uint32) ([]*dto.UserMenuDTO, error) {
	var (
		menus       []*do.MenuDO
		catalogs    []*do.MenuDO
		childrenMap map[uint32][]*do.MenuDO
		err         error
	)
	// 查询用户菜单信息
	menus, err = m.listMenusByRoleId(roleId)
	if err != nil {
		return nil, err
	}
	// 获取目录列表
	catalogs = listCatalog(menus)
	// 获取目录下的子菜单
	childrenMap = getMenuMap(menus)
	// 转换前端菜单格式
	return convertUserMenuList(catalogs, childrenMap), nil
}

// 转换用户菜单格式
//
// @catalogs 目录
// @childrenMap 子菜单
func convertUserMenuList(catalogs []*do.MenuDO, childrenMap map[uint32][]*do.MenuDO) []*dto.UserMenuDTO {
	var (
		userMenu []*dto.UserMenuDTO
	)
	for _, item := range catalogs {
		var (
			userMenuDTO   dto.UserMenuDTO // 目录
			list          []dto.UserMenuDTO
			children      []*do.MenuDO // 目录下的子菜单
			existChildren bool
		)
		// 获取目录下的子菜单
		if children, existChildren = childrenMap[item.ID]; existChildren && len(children) != 0 {
			// 多级菜单处理
			util.DeepCopyByJson(item, &userMenuDTO)
			for _, c := range children {
				var (
					childUserMenuDTO dto.UserMenuDTO
				)
				util.DeepCopyByJson(c, &childUserMenuDTO)
				childUserMenuDTO.Hidden = c.IsHidden == constant.LogicDeleteTrue
				list = append(list, childUserMenuDTO)
			}
		} else {
			// 一级菜单处理
			userMenuDTO.Path = item.Path
			userMenuDTO.Component = constant.COMPONENT
			list = append(list, dto.UserMenuDTO{
				Path:      "",
				Name:      item.Name,
				Icon:      item.Icon,
				Component: item.Component,
			})
		}
		userMenuDTO.Hidden = item.IsHidden == constant.LogicDeleteTrue
		userMenuDTO.Children = list
		userMenu = append(userMenu, &userMenuDTO)
	}
	return userMenu
}

// getMenuMap 获取目录下菜单列表
//
// @menuList 菜单列表
// @return 目录下的菜单列表
func getMenuMap(menus []*do.MenuDO) map[uint32][]*do.MenuDO {
	var (
		childrens = make(map[uint32][]*do.MenuDO, 0)
	)
	for _, v := range menus {
		if v.ParentId == 0 {
			continue
		}
		if _, exist := childrens[v.ParentId]; exist {
			childrens[v.ParentId] = append(childrens[v.ParentId], v)
		} else {
			childrens[v.ParentId] = []*do.MenuDO{v}
		}
	}
	return childrens
}

// 获取目录列表
//
// @menuList 菜单列表
// @return 目录列表
func listCatalog(menus []*do.MenuDO) []*do.MenuDO {
	var (
		catalogs []*do.MenuDO
	)
	for _, v := range menus {
		if v.ParentId == 0 {
			catalogs = append(catalogs, v)
		}
	}
	// todo 排序
	return catalogs
}
