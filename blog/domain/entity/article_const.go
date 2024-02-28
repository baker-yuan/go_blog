package entity

// ArticleStatus 文章状态枚举
type ArticleStatus uint32

const (
	ArticleStatusPublic = 1 // 公开
	ArticleStatusSecret = 2 // 私密
	ArticleStatusDraft  = 3 // 草稿
)
