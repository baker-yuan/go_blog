// Package entity 领域聚合实体
package entity

import "context"

// PraiseBuilder Praise Entity 的 Builder
type PraiseBuilder struct {
	id         uint32       // id
	moduleCode string       // 模块标识
	objectId   uint32       // 信息ID
	uid        uint32       // 用户ID
	status     PraiseStatus // 点赞状态
	createTime uint32       // 创建时间
	updateTime uint32       // 更新时间
}

// ID build id
func (f *PraiseBuilder) ID(id uint32) *PraiseBuilder {
	f.id = id
	return f
}

// ModuleCode build moduleCode
func (f *PraiseBuilder) ModuleCode(moduleCode string) *PraiseBuilder {
	f.moduleCode = moduleCode
	return f
}

// ObjectId build objectId
func (f *PraiseBuilder) ObjectId(objectId uint32) *PraiseBuilder {
	f.objectId = objectId
	return f
}

// UID build uid
func (f *PraiseBuilder) UID(uid uint32) *PraiseBuilder {
	f.uid = uid
	return f
}

// Status build status
func (f *PraiseBuilder) Status(status PraiseStatus) *PraiseBuilder {
	f.status = status
	return f
}

// CreateTime build createTime
func (f *PraiseBuilder) CreateTime(createTime uint32) *PraiseBuilder {
	f.createTime = createTime
	return f
}

// UpdateTime build updateTime
func (f *PraiseBuilder) UpdateTime(updateTime uint32) *PraiseBuilder {
	f.updateTime = updateTime
	return f
}

// Build 构建一个实体
func (f *PraiseBuilder) Build(ctx context.Context) (*Praise, error) {
	var (
		praise = &Praise{}
		err    error
	)
	_ = praise.setID(ctx, f.id)
	err = praise.setModuleCode(ctx, f.moduleCode)
	if err != nil {
		return nil, err
	}
	_ = praise.setObjectId(ctx, f.objectId)
	_ = praise.setUid(ctx, f.uid)
	_ = praise.setStatus(ctx, f.status)
	_ = praise.setCreateTime(ctx, f.createTime)
	_ = praise.setUpdateTime(ctx, f.updateTime)
	return praise, err
}
