package repository

import (
	"context"

	"github.com/baker-yuan/go-blog/go_blog/application/user/domain/entity"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/user"
)

type UserRepository interface {
	// SaveAdmin 保存管理员用户
	SaveAdmin(ctx context.Context, user *entity.User) (uint32, error)
	// UpdateAdminByID 根据ID修改管理员用户
	UpdateAdminByID(ctx context.Context, user *entity.User) error

	// GetUserByID 根据用户id集合查询用户
	GetUserByID(ctx context.Context, id uint32) (*entity.User, error)
	// GetUserByIDs 根据用户id集合查询用户
	GetUserByIDs(ctx context.Context, ids ...uint32) (entity.Users, error)
	// SearchUser 用户搜索
	SearchUser(ctx context.Context, req *pb.SearchUserReq) (entity.Users, uint32, error)

	// GetUserByEmailAndPassword 通过用户名查找用户，并且验证密码
	GetUserByEmailAndPassword(ctx context.Context, username string, password string) (*entity.User, error)
}
