package storage

import "context"

type Interface interface {
	Get(ctx context.Context, key string) (string, error)
	List(ctx context.Context, key string) ([]Keypair, error)
	Create(ctx context.Context, key, val string) error
	Update(ctx context.Context, key, val string) error
	BatchDelete(ctx context.Context, keys []string) error
	Watch(ctx context.Context, key string) <-chan WatchResponse
}

// WatchResponse Watch方法通道传输的数据
type WatchResponse struct {
	Events   []Event
	Error    error
	Canceled bool
}

// Keypair 数据
type Keypair struct {
	Key   string // 目录
	Value string // 值
}

// Event 事件
type Event struct {
	Keypair           // 数据
	Type    EventType // 事件类型
}

// EventType 事件类型
type EventType string

var (
	EventTypePut    EventType = "put"    // 新增、修改
	EventTypeDelete EventType = "delete" // 删除
)
