// Package entity 领域聚合实体
package entity

// PraiseStatus 点赞状态 1-已点赞 0-未点赞
type PraiseStatus int64

const (
	PraiseFalse PraiseStatus = 0 // 未点赞
	PraiseTrue  PraiseStatus = 1 // 已点赞
)
