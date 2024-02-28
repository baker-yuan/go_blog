// Package dao DAO 层
package dao

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/db"
	"github.com/baker-yuan/go-blog/interaction/infrastructure/persistence/praise/mysql/po"
	"gorm.io/gorm"
)

// PraiseDAO ObjectPraise 的 DAO
type PraiseDAO struct {
}

func (p *PraiseDAO) FindByUnique(ctx context.Context, moduleCode string, objectId uint32, uid uint32) (*po.PraisePO, error) {
	var (
		db       = db.GetMysqlDb()
		praisePO = &po.PraisePO{}
		tx       *gorm.DB
	)
	tx = db.Model(&po.PraisePO{}).
		Where("module_code = ? and object_id = ? and uid = ? ", moduleCode, objectId, uid).
		Find(&praisePO)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return praisePO, nil
}

func (p *PraiseDAO) SaveOrUpdate(ctx context.Context, tx *gorm.DB, praisePO *po.PraisePO) error {
	if res := tx.Model(&po.PraisePO{}).Save(praisePO); res != nil {
		return res.Error
	}
	return nil
}

func (p *PraiseDAO) RemoveByUnique(ctx context.Context, tx *gorm.DB, id uint32) error {
	condition := &po.PraisePO{
		ID: id,
	}
	res := tx.Model(&po.PraisePO{}).Delete(condition)
	if res != nil {
		return res.Error
	}
	return nil
}
