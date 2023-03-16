package repo

import "github.com/baker-yuan/go-blog/auth/application/dto"

type MenuRepo interface {
	ListUserMenus(roleId uint32) ([]*dto.UserMenuDTO, error)
}
