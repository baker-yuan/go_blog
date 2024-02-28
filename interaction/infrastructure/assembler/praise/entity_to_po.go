// Package assembler 数据结构转换器
package assembler

import (
	entity "github.com/baker-yuan/go-blog/interaction/domain/entity/praise"
	"github.com/baker-yuan/go-blog/interaction/infrastructure/persistence/praise/mysql/po"
)

// GenPraisePO 把 Entity 转换成 PO
func GenPraisePO(praiseEntity *entity.Praise, praisePO *po.PraisePO) {
	praisePO.ID = praiseEntity.ID()
	praisePO.ModuleCode = praiseEntity.ModuleCode()
	praisePO.ObjectId = praiseEntity.ObjectId()
	praisePO.UID = praiseEntity.UID()
	praisePO.Status = praiseEntity.Status()
	praisePO.CreateTime = praiseEntity.CreateTime()
	praisePO.UpdateTime = praiseEntity.UpdateTime()
}
