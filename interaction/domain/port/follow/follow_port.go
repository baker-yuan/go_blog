// Package port 端口-适配器模式的 port
package port

import (
	"context"

	entity "github.com/baker-yuan/go-blog/interaction/domain/entity/follow"
)

// FollowPort 接口
type FollowPort interface {
	// TxEnd 事务
	TxEnd(ctx context.Context, txFunc func() error) error

	// FindByUnique 通过 Unique Key 查
	FindByUnique(ctx context.Context, uID, followUID uint32) (*entity.Follow, error)

	// GetFolloweeCount 通过 uID 查 Follow 数量
	GetFolloweeCount(ctx context.Context, uID uint32, followState entity.FollowState) (uint32, error)

	// Save 增和改
	Save(ctx context.Context, entity *entity.Follow) func() error
}
