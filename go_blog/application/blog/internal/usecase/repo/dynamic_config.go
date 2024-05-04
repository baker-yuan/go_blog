package repo

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	"github.com/baker-yuan/go-blog/go_blog/common/db"
	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
	"gorm.io/gorm"
)

// DynamicConfigRepo 动态配置
type DynamicConfigRepo struct {
	*db.GenericDao[entity.DynamicConfig, uint32]
}

// NewDynamicConfigRepo 创建
func NewDynamicConfigRepo(gormDB *gorm.DB) *DynamicConfigRepo {
	return &DynamicConfigRepo{
		GenericDao: db.NewGenericDao[entity.DynamicConfig, uint32](gormDB),
	}
}

func init() {
	registerInitField(initDynamicConfigField)
}

var (
	// 全字段修改DynamicConfig那些字段不修改
	notUpdateDynamicConfigField = []string{
		entity.DynamicConfigFieldCreateUserId,
		entity.DynamicConfigFieldCreateTime,
	}
	updateDynamicConfigField []string
)

// InitDynamicConfigField 全字段修改
func initDynamicConfigField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(entity.DynamicConfig{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateDynamicConfigField = util.NewSliceUtils[string]().SliceRemove(columns, notUpdateDynamicConfigField...)
	return nil
}

// GetDynamicConfigByID 根据动态配置id查询动态配置
func (r *DynamicConfigRepo) GetDynamicConfigByID(ctx context.Context, id uint32) (*entity.DynamicConfig, error) {
	dbDynamicConfig, err := r.GenericDao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dbDynamicConfig, nil
}

// GetDynamicConfigByIDs 根据动态配置id集合查询动态配置
func (r *DynamicConfigRepo) GetDynamicConfigByIDs(ctx context.Context, ids []uint32) (entity.DynamicConfigs, error) {
	dbDynamicConfigs, err := r.GenericDao.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return dbDynamicConfigs, nil
}

// Save 保存动态配置
func (r *DynamicConfigRepo) Save(ctx context.Context, dynamicConfig *entity.DynamicConfig) (uint32, error) {
	if dynamicConfig.ID > 0 {
		return 0, errors.New("illegal argument dynamicConfig id exist")
	}
	if err := r.GenericDao.Create(ctx, dynamicConfig); err != nil {
		return 0, err
	}
	return dynamicConfig.ID, nil
}

// UpdateByID 根据ID修改动态配置
func (r *DynamicConfigRepo) UpdateByID(ctx context.Context, dynamicConfig *entity.DynamicConfig) error {
	if dynamicConfig.ID == 0 {
		return errors.New("illegal argument dynamicConfig exist")
	}
	return r.GenericDao.GetDB(ctx).Select(updateDynamicConfigField).Updates(dynamicConfig).Error
}

// DeleteByID 根据ID删除动态配置
func (r *DynamicConfigRepo) DeleteByID(ctx context.Context, id uint32) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// SearchDynamicConfig 动态配置搜索
func (r *DynamicConfigRepo) SearchDynamicConfig(ctx context.Context, req *pb.SearchDynamicConfigReq) (entity.DynamicConfigs, uint32, error) {
	var (
		res       []*entity.DynamicConfig
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
