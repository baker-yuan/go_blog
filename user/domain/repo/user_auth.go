package repo

import "github.com/baker-yuan/go-blog/user/domain/entity"

type UserRepo interface {
	// FindByUsername 通过用户名查找
	FindByUsername(username string) (*entity.User, error)
}
