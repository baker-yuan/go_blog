// Package entity 领域聚合实体
package entity

import (
	"context"
	"errors"
	"time"

	"github.com/baker-yuan/go-blog/all_packaged_library/base/log"
)

// Follow 关注关系实体，followUid 和 uid 为 unique key
type Follow struct {
	id            uint32      // id
	uid           uint32      // 关注者
	followUid     uint32      // 被关注者
	state         FollowState // 关注状态
	createTime    uint32      // 创建时间
	updateTime    uint32      // 更新时间
	countFollower uint32      // 已经关注的人数
}

// setID 设置 UID
func (f *Follow) setID(ctx context.Context, id uint32) error {
	f.id = id
	return nil
}

// setState 设置 State
func (f *Follow) setState(ctx context.Context, state FollowState) error {
	// 获取总数，判断总数是否超过上限
	if state == FollowStateTrue && f.countFollower >= FollowMaxLimitedByFollower {
		return errors.New("当前 Follower 的关注数量超过业务方指定的上限")
	}
	f.state = state
	log.Debug(ctx, "domain - entity - setState state: %+v", state)
	return nil
}

// setModifyTime 设置 modifyTime
func (f *Follow) setModifyTime(ctx context.Context, modifyTime uint32) error {
	f.updateTime = modifyTime
	log.Debug(ctx, "domain - entity - setModifyTime modifyTime: %+v", modifyTime)
	return nil
}

// setUID 设置 uid
func (f *Follow) setUID(ctx context.Context, uID uint32) error {
	if uID == f.followUid {
		return errors.New("setUID: uID 和 followUid 不能相等")
	}
	f.uid = uID
	log.Debug(ctx, "domain - entity - setUID uID: %+v", uID)
	return nil
}

// setFollowUID 设置 followUid
func (f *Follow) setFollowUID(ctx context.Context, followUID uint32) error {
	if followUID == f.uid {
		return errors.New("setFollowUid: followUID 和 uid 不能相等")
	}
	f.followUid = followUID
	log.Debug(ctx, "domain - entity - setFollowUid followUID: %+v", followUID)
	return nil
}

// ChangeFollowStateToTrue 变更 Follow State
func (f *Follow) ChangeFollowStateToTrue(ctx context.Context) error {
	if f.IsFollow() { // 如果要添加关注，则先判断是否已经关注
		return errors.New("已经关注，请勿重复关注")
	}
	err := f.setState(ctx, FollowStateTrue)
	if err != nil {
		return err
	}
	_ = f.setModifyTime(ctx, uint32(time.Now().Second()))
	return nil
}

// IsFollow Check 是否已经关注
func (f *Follow) IsFollow() bool {
	return f.state == FollowStateTrue
}
