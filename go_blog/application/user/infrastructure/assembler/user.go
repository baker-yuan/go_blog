package assembler

import (
	"github.com/baker-yuan/go-blog/application/user/domain/entity"
	"github.com/baker-yuan/go-blog/application/user/infrastructure/persistence"
)

func UserPosToEntity(users []*persistence.User) []*entity.User {
	res := make([]*entity.User, 0)

	return res
}
func UserPoToEntity(user *persistence.User) *entity.User {
	res := &entity.User{}
	return res
}

func UserEntityToPo(user *entity.User) *persistence.User {
	res := &persistence.User{}
	return res
}
