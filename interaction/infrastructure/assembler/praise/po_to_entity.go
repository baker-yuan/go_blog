// Package assembler 数据结构转换器
package assembler

import (
	"context"

	entity "github.com/baker-yuan/go-blog/interaction/domain/entity/praise"
	"github.com/baker-yuan/go-blog/interaction/infrastructure/persistence/praise/mysql/po"
)

// ZeroUINT32 uint32 0 值
const ZeroUINT32 uint32 = 0

// GenPraiseEntity 构造 Follow Entity
func GenPraiseEntity(ctx context.Context, praisePO *po.PraisePO) *entity.Praise {
	if praisePO.ID == ZeroUINT32 {
		return nil
	}
	praiseBuilder := &entity.PraiseBuilder{}
	praiseBuilder.ID(praisePO.ID)
	praiseBuilder.ModuleCode(praisePO.ModuleCode)
	praiseBuilder.ObjectId(praisePO.ObjectId)
	praiseBuilder.UID(praisePO.UID)
	praiseBuilder.Status(praisePO.Status)
	praiseBuilder.CreateTime(praisePO.CreateTime)
	praiseBuilder.UpdateTime(praisePO.UpdateTime)
	praiseEntity, _ := praiseBuilder.Build(ctx)
	return praiseEntity
}
