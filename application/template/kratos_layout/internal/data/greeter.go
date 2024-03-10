package data

import (
	"context"
	"kratos-layout/internal/biz"
)

// greeterRepo 数据存储
type greeterRepo struct {
	data *Data
}

// NewGreeterRepo 数据存储
func NewGreeterRepo(data *Data) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
	}
}

// Save 保存
func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

// Update 修改
func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

// FindByID 通过主键id查找
func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

// ListByHello 条件查找
func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

// ListAll 查询所有
func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
