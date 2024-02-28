// Package entity 领域聚合实体
package entity

import (
	"context"
	"errors"
	"time"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
)

// Praise 点赞实体
type Praise struct {
	id         uint32       // id
	moduleCode string       // 模块标识
	objectId   uint32       // 信息ID
	uid        uint32       // 用户ID
	status     PraiseStatus // 点赞状态
	createTime uint32       // 创建时间
	updateTime uint32       // 更新时间
}

// ID 获取 ID
func (p *Praise) ID() uint32 {
	return p.id
}

// setID 设置 ID
func (p *Praise) setID(ctx context.Context, id uint32) error {
	p.id = id
	log.Debug(ctx, "domain - entity - setID id: %+v", id)
	return nil
}

// ModuleCode 获取 moduleCode
func (p *Praise) ModuleCode() string {
	return p.moduleCode
}

// setModuleCode 设置 moduleCode
func (p *Praise) setModuleCode(ctx context.Context, moduleCode string) error {
	if len(moduleCode) == 0 {
		return errors.New("moduleCode参数异常")
	}

	p.moduleCode = moduleCode
	log.Debug(ctx, "domain - entity - setModuleCode moduleCode: %s", moduleCode)
	return nil
}

// ObjectId 获取 objectId
func (p *Praise) ObjectId() uint32 {
	return p.objectId
}

// setObjectId 设置 objectId
func (p *Praise) setObjectId(ctx context.Context, objectId uint32) error {
	p.objectId = objectId
	log.Debug(ctx, "domain - entity - setObjectId objectId: %d", objectId)
	return nil
}

// UID 获取 uid
func (p *Praise) UID() uint32 {
	return p.uid
}

// setUid 设置 uid
func (p *Praise) setUid(ctx context.Context, uid uint32) error {
	p.uid = uid
	log.Debug(ctx, "domain - entity - setUid uid: %d", uid)
	return nil
}

// Status 获取 status
func (p *Praise) Status() PraiseStatus {
	return p.status
}

// setStatus 设置 status
func (p *Praise) setStatus(ctx context.Context, status PraiseStatus) error {
	p.status = status
	log.Debug(ctx, "domain - entity - setStatus status: %d", status)
	return nil
}

// CreateTime 获取 createTime
func (p *Praise) CreateTime() uint32 {
	return p.createTime
}

// setCreateTime 设置 createTime
func (p *Praise) setCreateTime(ctx context.Context, createTime uint32) error {
	p.createTime = createTime
	log.Debug(ctx, "domain - entity - setCreateTime createTime: %+v", createTime)
	return nil
}

// UpdateTime 获取 updateTime
func (p *Praise) UpdateTime() uint32 {
	return p.updateTime
}

// setCreateTime 设置 updateTime
func (p *Praise) setUpdateTime(ctx context.Context, createTime uint32) error {
	p.createTime = createTime
	log.Debug(ctx, "domain - entity - setUpdateTime updateTime: %+v", createTime)
	return nil
}

// IsPraise Check 是否已经点赞
func (p *Praise) IsPraise() bool {
	return p.status == PraiseTrue
}

// Praise 点赞
func (p *Praise) Praise(ctx context.Context) error {
	if p.IsPraise() { // 如果要点赞，则先判断是否已经点赞
		return errors.New("已经点赞，请勿重复点赞")
	}
	err := p.setStatus(ctx, PraiseTrue)
	if err != nil {
		return err
	}
	_ = p.setUpdateTime(ctx, uint32(time.Now().Second()))
	return nil
}

// CancelPraise 取消点赞
func (p *Praise) CancelPraise(ctx context.Context) error {
	if !p.IsPraise() { // 如果要点赞，则先判断是否已经点赞
		return errors.New("未点赞，请先点赞")
	}
	err := p.setStatus(ctx, PraiseFalse)
	if err != nil {
		return err
	}
	_ = p.setUpdateTime(ctx, uint32(time.Now().Second()))
	return nil
}
