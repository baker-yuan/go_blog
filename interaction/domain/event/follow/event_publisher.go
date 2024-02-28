// Package event 领域事件包
package event

import "context"

// publisher publisher的结构体
type publisher struct {
	// key=事件类型 value=处理事件函数
	subscribers map[DomainEventType][]Handler
}

// Handler event 回调的 Handler
type Handler func(context.Context, DomainEvent)

// Subscribe 订阅事件
func (d *publisher) Subscribe(eventType DomainEventType, handler Handler) {
	_, ok := d.subscribers[eventType]
	if !ok {
		d.subscribers[eventType] = []Handler{}
	}
	d.subscribers[eventType] = append(d.subscribers[eventType], handler)
}

// Publish 发布事件
func (d *publisher) Publish(ctx context.Context, event DomainEvent) {
	handlers := d.subscribers[event.EventType()]
	for _, handler := range handlers {
		handler(ctx, event)
	}
}

var (
	// Publisher 单例
	Publisher = &publisher{
		subscribers: make(map[DomainEventType][]Handler),
	}
)
