// Package po 持久化对象
package po

import entity "github.com/baker-yuan/go-blog/interaction/domain/entity/praise"

// PraisePO Praise的PO
type PraisePO struct {
	ID         uint32              `gorm:"primaryKey"` // id
	ModuleCode string              // 模块标识
	ObjectId   uint32              // 信息ID
	UID        uint32              // 用户ID
	Status     entity.PraiseStatus // 点赞状态
	CreateTime uint32              // 创建时间
	UpdateTime uint32              // 更新时间
}

func (p PraisePO) TableName() string {
	return "tb_praise"
}
