// Package command 命令模式入参
package command

// ObjectPraiseCMD 新增点赞
type ObjectPraiseCMD struct {
	ModuleCode string // 模块标识
	ObjectId   uint32 // 信息ID
	Uid        uint32 // 用户ID
}

// CancelObjectPraiseCMD 取消点赞
type CancelObjectPraiseCMD struct {
	ModuleCode string // 模块标识
	ObjectId   uint32 // 信息ID
	Uid        uint32 // 用户ID
}
