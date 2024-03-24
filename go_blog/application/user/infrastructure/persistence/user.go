package persistence

import (
	"github.com/baker-yuan/go-blog/application/user/domain/entity"
)

func UserPosToEntity(users []*User) []*entity.User {
	res := make([]*entity.User, 0)

	return res
}
func UserPoToEntity(user *User) *entity.User {
	res := &entity.User{}
	return res
}

func UserEntityToPo(user *entity.User) *User {
	res := &User{}
	return res
}
