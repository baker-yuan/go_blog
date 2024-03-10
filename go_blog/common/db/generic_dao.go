package db

import (
	"context"

	"gorm.io/gorm"
)

// GenericDao 通用数据存储操作
type GenericDao[M any, P comparable] struct {
	DB *gorm.DB
}

// NewGenericDao 创建一个新的GenericDao实例
func NewGenericDao[M any, P comparable](db *gorm.DB) *GenericDao[M, P] {
	return &GenericDao[M, P]{
		DB: db,
	}
}

// Create 添加记录
func (g *GenericDao[M, P]) Create(ctx context.Context, m *M) error {
	return g.DB.WithContext(ctx).Create(m).Error
}

// DeleteByID 根据主键ID删除记录
func (g *GenericDao[M, P]) DeleteByID(ctx context.Context, id P) error {
	var m M
	return g.DB.WithContext(ctx).Delete(&m, id).Error
}

// UpdateByID 根据主键ID修改记录
func (g *GenericDao[M, P]) UpdateByID(ctx context.Context, m *M) error {
	return g.DB.WithContext(ctx).Updates(m).Error
}

// GetByID 根据主键ID查询记录
func (g *GenericDao[M, P]) GetByID(ctx context.Context, id P) (*M, error) {
	var m M
	if err := g.DB.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// GetByIDs 根据主键ID集合查询记录
func (g *GenericDao[M, P]) GetByIDs(ctx context.Context, ids []P) ([]*M, error) {
	// https://gorm.io/zh_CN/docs/query.html#Struct-amp-Map-条件
	var m = make([]*M, 0)
	if err := g.DB.WithContext(ctx).Where(ids).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// BatchCreate 批量添加记录
func (g *GenericDao[M, P]) BatchCreate(ctx context.Context, ms []*M) error {
	return g.DB.WithContext(ctx).Create(ms).Error
}
