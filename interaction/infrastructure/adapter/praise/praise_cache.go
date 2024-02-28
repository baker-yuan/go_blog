// Package adapter 端口-适配器模式的 adapter
package adapter

import "context"

// PraiseCache ObjectPraise 的 Cache
type PraiseCache struct {
}

// Save 保存缓存
func (c *PraiseCache) Save(ctx context.Context, moduleCode string, objectId uint32, uid uint32) (err error) {
	return nil
}

// Delete 删除缓存
func (c *PraiseCache) Delete(ctx context.Context, moduleCode string, objectId uint32, uid uint32) (err error) {
	return nil
}
