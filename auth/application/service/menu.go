package service

import (
	"github.com/baker-yuan/go-blog/auth/application/dto"
	"github.com/baker-yuan/go-blog/auth/domain/repo"
)

// MenuService 菜单服务
type MenuService struct {
	menuRepo repo.MenuRepo
}

func NewMenuService(repo repo.MenuRepo) *MenuService {
	return &MenuService{
		menuRepo: repo,
	}
}

// ListUserMenus 查看用户菜单
//
// @return 菜单列表
func (m *MenuService) ListUserMenus(roleId uint32) ([]*dto.UserMenuDTO, error) {
	return m.menuRepo.ListUserMenus(roleId)
}
