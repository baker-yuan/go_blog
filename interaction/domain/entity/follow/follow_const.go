// Package entity 领域聚合实体
package entity

// FollowState 关注状态
type FollowState int64

const (
	FollowStateFalse FollowState = 0 // 未关注
	FollowStateTrue  FollowState = 1 // 已关注
	FollowStateBoth  FollowState = 2 // 相互关注
)

// FollowMaxLimitedByFollower 一个 Follower 最大能关联的 Follow 上限
const FollowMaxLimitedByFollower uint32 = 300
