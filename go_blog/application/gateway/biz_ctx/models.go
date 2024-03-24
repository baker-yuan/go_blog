package biz_ctx

// NodeStatus 节点状态类型
type NodeStatus uint8

const (
	// Running 节点运行中状态
	Running NodeStatus = 1
	// Down 节点不可用状态
	Down NodeStatus = 2
)

// IService 服务
type IService interface {
	Instances() []IInstance // 节点
}

// IInstance 节点实例
type IInstance interface {
	IAttributes         // 属性 key=string v=string
	ID() string         // ID
	IP() string         // IP
	Port() int          // 端口
	Addr() string       // 地址 ip/url:端口
	Status() NodeStatus // 节点状态
	Up()                // 设置节点可用
	Down()              // 设置节点不可用
}

// Attrs 属性集合
type Attrs map[string]string

// IAttributes 属性接口
type IAttributes interface {
	GetAttrs() Attrs
	GetAttrByName(name string) (string, bool)
}
