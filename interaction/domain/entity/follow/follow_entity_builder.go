package entity

import (
	"context"
)

// FollowBuilder Follow Entity 的 Builder
type FollowBuilder struct {
	id              uint32      // id
	uid             uint32      // 关注者
	followUid       uint32      // 被关注者
	state           FollowState // 关注状态
	createTime      uint32      // 创建时间
	updateTime      uint32      // 更新时间
	countByFollower uint32
}

// ID build id
func (f *FollowBuilder) ID(id uint32) *FollowBuilder {
	f.id = id
	return f
}

// UID build uid
func (f *FollowBuilder) UID(uid uint32) *FollowBuilder {
	f.uid = uid
	return f
}

// FollowUID build uid
func (f *FollowBuilder) FollowUID(followUID uint32) *FollowBuilder {
	f.followUid = followUID
	return f
}

// State build state
func (f *FollowBuilder) State(state FollowState) *FollowBuilder {
	f.state = state
	return f
}

// CreateTime build createTime
func (f *FollowBuilder) CreateTime(createTime uint32) *FollowBuilder {
	f.createTime = createTime
	return f
}

// UpdateTime build updateTime
func (f *FollowBuilder) UpdateTime(modifyTime uint32) *FollowBuilder {
	f.updateTime = modifyTime
	return f
}

// CountByFollower build CountByFollower
func (f *FollowBuilder) CountByFollower(countByFollower uint32) *FollowBuilder {
	f.countByFollower = countByFollower
	return f
}

// Build 构建一个实体
func (f *FollowBuilder) Build(ctx context.Context) (follow *Follow, err error) {
	follow = &Follow{}
	_ = follow.setID(ctx, f.id)

	err = follow.setFollowUID(ctx, f.followUid)
	if err != nil {
		return nil, err
	}

	return follow, nil
}
