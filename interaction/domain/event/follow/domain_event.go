// Package event 领域事件包
package event

import (
	entity "github.com/baker-yuan/go-blog/interaction/domain/entity/follow"
)

// DomainEvent 领域事件
type DomainEvent interface {
	EventType() DomainEventType
}

// DomainEventType 事件类型
type DomainEventType string

// DomainEventType 的2个类型
const (
	TypeFollowCreated = DomainEventType("follow_created") // 关注关系创建事件
	TypeFollowDeleted = DomainEventType("follow_deleted") // 关注关系删除事件
)

// FollowCreatedEvent Follow 的 Created 事件
type FollowCreatedEvent struct {
	Follow *entity.Follow
}

// EventType EventType
func (f *FollowCreatedEvent) EventType() DomainEventType {
	return TypeFollowCreated
}

// FollowDeletedEvent Follow 的 Deleted 事件
type FollowDeletedEvent struct {
	Follow *entity.Follow
}

// EventType EventType
func (f *FollowDeletedEvent) EventType() DomainEventType {
	return TypeFollowDeleted
}
