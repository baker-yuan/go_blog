package repository

import (
	"context"

	"github.com/baker-yuan/go-blog/application/user/domain/entity"
	pb "github.com/baker-yuan/go-blog/protocol/user"
)

type UserRepository interface {
	// GetUserByID 根据用户id集合查询用户
	GetUserByID(ctx context.Context, id uint32) (*entity.User, error)
	// GetUserByIDs 根据用户id集合查询用户
	GetUserByIDs(ctx context.Context, ids []uint32) (entity.Users, error)
	// Save 保存用户
	Save(ctx context.Context, user *entity.User) (uint32, error)
	// UpdateByID 根据ID修改用户
	UpdateByID(ctx context.Context, user *entity.User) error
	// SearchUser 用户搜索
	SearchUser(ctx context.Context, req *pb.SearchUserReq) (entity.Users, uint32, error)
}
