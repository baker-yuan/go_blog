// Package adapter 端口-适配器模式的 adapter
package adapter

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/db"
	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
	entity "github.com/baker-yuan/go-blog/interaction/domain/entity/praise"
	assembler2 "github.com/baker-yuan/go-blog/interaction/infrastructure/assembler/praise"
	"github.com/baker-yuan/go-blog/interaction/infrastructure/persistence/praise/mysql/dao"
	"github.com/baker-yuan/go-blog/interaction/infrastructure/persistence/praise/mysql/po"
	"gorm.io/gorm"
)

// PraiseAdapter ObjectPraise 的 Adapter
type PraiseAdapter struct {
	praiseCache PraiseCache
	praiseDAO   dao.PraiseDAO
}

func (p PraiseAdapter) TxEnd(ctx context.Context, txFunc func(tx *gorm.DB) error) error {
	return txFunc(db.GetMysqlDb())
	// db := db.GetMysqlDb()
	// // 开始事务
	// tx := db.Begin()
	//
	// // defer func() {
	// // 	if err := recover(); err != nil {
	// // 		// 回滚事务
	// // 		tx.Rollback()
	// // 	}
	// // }()
	//
	// // 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	// if err := txFunc(tx); err != nil {
	// 	// 回滚事务
	// 	tx.Rollback()
	// 	return err
	// }
	//
	// // 提交事务
	// defer tx.Commit()
	// return nil
}

func (p PraiseAdapter) FindByUnique(ctx context.Context, moduleCode string, objectId uint32, uid uint32) (*entity.Praise, error) {
	var (
		praisePO *po.PraisePO
		err      error
	)
	var (
		praiseEntity *entity.Praise
	)
	praisePO, err = p.praiseDAO.FindByUnique(ctx, moduleCode, objectId, uid)
	if err != nil {
		return nil, err
	}
	// PO 转 Entity
	praiseEntity = assembler2.GenPraiseEntity(ctx, praisePO)
	return praiseEntity, nil
}

func (p PraiseAdapter) Save(ctx context.Context, praiseEntity *entity.Praise) func(tx *gorm.DB) error {
	log.Debug(ctx, "infrastructure - adapter - Save PraiseEntity: %v", praiseEntity)
	fun := func(tx *gorm.DB) error {
		var (
			err error
		)
		// 实例化 PraisePO
		praisePO := &po.PraisePO{}
		// 转化器：把 entity 转化成 PO
		assembler2.GenPraisePO(praiseEntity, praisePO)
		log.Debug(ctx, "infrastructure - adapter - Save PraisePO: %+v", praisePO)

		if err = p.praiseDAO.SaveOrUpdate(ctx, tx, praisePO); err != nil {
			return err
		}

		if praiseEntity.Status() == entity.PraiseTrue {
			// 保存缓存
			err = p.praiseCache.Save(ctx, praiseEntity.ModuleCode(), praiseEntity.ObjectId(), praiseEntity.UID())
		} else {
			// 删除缓存
			err = p.praiseCache.Delete(ctx, praiseEntity.ModuleCode(), praiseEntity.ObjectId(), praiseEntity.UID())
		}
		if err != nil {
			log.Error(ctx, "infrastructure - adapter - Save cache.save err: %+v", err)
			return err
		}
		return nil
	}
	return fun
}
