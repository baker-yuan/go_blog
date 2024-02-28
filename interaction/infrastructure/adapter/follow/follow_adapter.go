// Package adapter 端口-适配器模式的 adapter
package adapter

import (
	"context"

	entity "github.com/baker-yuan/go-blog/interaction/domain/entity/follow"
	"github.com/baker-yuan/go-blog/interaction/infrastructure/persistence/follow/mysql/dao"
)

// FollowAdapter Follow 的 Adapter
type FollowAdapter struct {
	followDAO   dao.FollowDAO
	followCache FollowCache
}

func (f FollowAdapter) TxEnd(ctx context.Context, txFunc func() error) error {
	// TODO implement me
	panic("implement me")
}

func (f FollowAdapter) FindByUnique(ctx context.Context, uID, followUID uint32) (*entity.Follow, error) {
	// var (
	// 	DB = db.GetMysqlDb()
	// )
	// // 实例化 FollowPO
	// followPO := po.FollowPO{}

	// DB.Table("")
	panic(1)
}

func (f FollowAdapter) GetFolloweeCount(ctx context.Context, uID uint32, followState entity.FollowState) (uint32, error) {
	// TODO implement me
	panic("implement me")
}

func (f FollowAdapter) Save(ctx context.Context, entity *entity.Follow) func() error {
	// TODO implement me
	panic("implement me")
}
