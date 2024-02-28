// Package event 领域事件包
package event

import (
	entity "github.com/baker-yuan/go-blog/interaction/domain/entity/praise"
)

// DomainEvent 领域事件
type DomainEvent interface {
	EventType() DomainEventType
}

// DomainEventType 事件类型
type DomainEventType string

// DomainEventType 的2个类型
const (
	TypePraiseCreated = DomainEventType("praise_created") // 点赞事件
	TypePraiseDeleted = DomainEventType("praise_deleted") // 取消点赞事件
)

// PraiseCreatedEvent Praise 的 Created 事件
type PraiseCreatedEvent struct {
	Praise *entity.Praise
}

// EventType EventType
func (f *PraiseCreatedEvent) EventType() DomainEventType {
	return TypePraiseCreated
}

// PraiseDeletedEvent ObjectPraise 的 Deleted 事件
type PraiseDeletedEvent struct {
	Praise *entity.Praise
}

// EventType EventType
func (f *PraiseDeletedEvent) EventType() DomainEventType {
	return TypePraiseDeleted
}
