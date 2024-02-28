package constant

// 分页查询相关
const (
	// 当前页码
	CURRENT = "current"
	//  默认页码
	DEFAULT_CURRENT = 1

	// 页码条数
	SIZE = "size"
	// 默认条数
	DEFAULT_SIZE = 10
)

// LogicDelete 逻辑删除枚举
type LogicDelete uint32

// 逻辑删除相关
const (
	LogicDeleteFalse = 0 // 否
	LogicDeleteTrue  = 1 // 是
)

// 其他
const (
	// 前端组件名
	COMPONENT = "Layout"
)
