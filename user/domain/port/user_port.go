// Package port 端口-适配器模式的 port
package port

import (
	"context"

	"github.com/baker-yuan/go-blog/user/domain/entity"
)

// UserPort 接口
type UserPort interface {
	// FindByUsername 通过用户名查找
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	// ListUsers 用户搜索
	ListUsers(ctx context.Context, current uint32, size uint32, nickname string, loginType uint32) ([]*entity.User, error)
}
