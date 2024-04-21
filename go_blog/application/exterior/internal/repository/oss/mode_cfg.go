package oss

import (
	"context"
	"errors"

	oss_model "github.com/baker-yuan/go-blog/go_blog/application/exterior/internal/model/oss"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	oss_pb "github.com/baker-yuan/go-blog/go_blog/protocol/exterior"
	"gorm.io/gorm"
)

// ModeCfg 模块配置
type (
	// IModeCfgRepo 数据存储操作
	IModeCfgRepo interface {
		// GetModeCfgByID 根据模块配置id集合查询模块配置
		GetModeCfgByID(ctx context.Context, id uint32) (*oss_model.ModeCfg, error)
		// GetModeCfgByIDs 根据模块配置id集合查询模块配置
		GetModeCfgByIDs(ctx context.Context, ids []uint32) (oss_model.ModeCfgList, error)
		// Save 保存模块配置
		Save(ctx context.Context, ModeCfg *oss_model.ModeCfg) (uint32, error)
		// UpdateByID 根据ID修改模块配置
		UpdateByID(ctx context.Context, ModeCfg *oss_model.ModeCfg) error
		// DeleteByID 根据ID删除模块配置
		DeleteByID(ctx context.Context, id uint32) error
		// SearchModeCfg 模块配置搜索
		SearchModeCfg(ctx context.Context, req *oss_pb.SearchOssModeCfgReq) (oss_model.ModeCfgList, uint32, error)

		// GetModeCfgByMode 根据文件模块查询模块配置
		GetModeCfgByMode(ctx context.Context, mode string) (*oss_model.ModeCfg, error)
	}
)

// ModeCfgRepo 模块配置
type ModeCfgRepo struct {
	*db.GenericDao[oss_model.ModeCfg, uint32]
}

// NewOssModeCfgRepo 创建
func NewOssModeCfgRepo(gormDB *gorm.DB) *ModeCfgRepo {
	return &ModeCfgRepo{
		GenericDao: db.NewGenericDao[oss_model.ModeCfg, uint32](gormDB),
	}
}

func init() {
	//registerInitField(initModeCfgField)
}

var (
	// 全字段修改ModeCfg那些字段不修改
	notUpdateModeCfgField = []string{
		oss_model.ModeCfgFieldCreateUserID,
		oss_model.ModeCfgFieldCreateTime,
	}
	updateModeCfgField []string
)

// InitOssModeCfgField 全字段修改
func InitOssModeCfgField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(oss_model.ModeCfg{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateModeCfgField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateModeCfgField...)
	return nil
}

// GetModeCfgByID 根据模块配置id查询模块配置
func (r *ModeCfgRepo) GetModeCfgByID(ctx context.Context, id uint32) (*oss_model.ModeCfg, error) {
	dbModeCfg, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbModeCfg, nil
}

// GetModeCfgByIDs 根据模块配置id集合查询模块配置
func (r *ModeCfgRepo) GetModeCfgByIDs(ctx context.Context, ids []uint32) (oss_model.ModeCfgList, error) {
	dbModeCfgs, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return dbModeCfgs, nil
}

// Save 保存模块配置
func (r *ModeCfgRepo) Save(ctx context.Context, ModeCfg *oss_model.ModeCfg) (uint32, error) {
	if ModeCfg.ID > 0 {
		return 0, errors.New("illegal argument ModeCfg id exist")
	}
	if err := r.GenericDao.Create(ctx, ModeCfg); err != nil {
		return 0, err
	}
	return ModeCfg.ID, nil
}

// UpdateByID 根据ID修改模块配置
func (r *ModeCfgRepo) UpdateByID(ctx context.Context, ModeCfg *oss_model.ModeCfg) error {
	if ModeCfg.ID == 0 {
		return errors.New("illegal argument ModeCfg exist")
	}
	return r.GenericDao.GetDB(ctx).Select(updateModeCfgField).Updates(ModeCfg).Error
}

// DeleteByID 根据ID删除模块配置
func (r *ModeCfgRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchModeCfg 模块配置搜索
func (r *ModeCfgRepo) SearchModeCfg(ctx context.Context, req *oss_pb.SearchOssModeCfgReq) (oss_model.ModeCfgList, uint32, error) {
	var (
		res       []*oss_model.ModeCfg
		pageTotal int64
	)
	tx, err := db.BuildSearch(
		ctx,
		req.GetSearch(),
		r.GenericDao.GetDB(ctx),
		func(search map[string]*db.SearchValue) {

		},
	)
	if err != nil {
		return nil, 0, err
	}
	tx = tx.Offset(int((req.GetPageNum() - 1) * req.GetPageSize())).
		Limit(int(req.GetPageSize())).Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)
	if err := tx.Error; err != nil {
		return nil, 0, err
	}
	return res, uint32(pageTotal), nil
}

// GetModeCfgByMode 根据文件模块查询模块配置
func (r *ModeCfgRepo) GetModeCfgByMode(ctx context.Context, mode string) (*oss_model.ModeCfg, error) {
	dbModeCfg := &oss_model.ModeCfg{}
	tx := r.GenericDao.GetDB(ctx)
	tx = tx.Where(oss_model.ModeCfgFieldMode, mode)
	if err := tx.Find(&dbModeCfg).Error; err != nil {
		return nil, err
	}
	return dbModeCfg, nil
}
