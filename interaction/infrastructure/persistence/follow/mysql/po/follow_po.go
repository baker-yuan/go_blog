// Package po 持久化对象
package po

import entity "github.com/baker-yuan/go-blog/interaction/domain/entity/follow"

// FollowPO Follow 的 PO
type FollowPO struct {
	ID            uint32             // id
	UID           uint32             // 关注者
	FollowUID     uint32             // 被关注者
	State         entity.FollowState // 关注状态
	CreateTime    uint32             // 创建时间
	UpdateTime    uint32             // 更新时间
	CountFollower uint32             // 已经关注的人数
}
