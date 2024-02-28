// Package port 端口-适配器模式的 port
package port

import (
	"context"

	entity "github.com/baker-yuan/go-blog/interaction/domain/entity/praise"
	"gorm.io/gorm"
)

// PraisePort 接口
type PraisePort interface {
	// TxEnd 事务
	TxEnd(ctx context.Context, txFunc func(tx *gorm.DB) error) error

	// FindByUnique moduleCode+objectId+uid查询
	FindByUnique(ctx context.Context, moduleCode string, objectId uint32, uid uint32) (*entity.Praise, error)

	// Save 保存
	Save(ctx context.Context, entity *entity.Praise) func(tx *gorm.DB) error
}
