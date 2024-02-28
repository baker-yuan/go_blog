package po

import "github.com/baker-yuan/go-blog/all_packaged_library/do"

// ArticlePO 文章
type ArticlePO struct {
	do.Model
	UserId         uint32 `json:"userId"`         // 作者
	CategoryId     uint32 `json:"categoryId"`     // 文章分类
	ArticleCover   string `json:"articleCover"`   // 文章缩略图
	ArticleTitle   string `json:"articleTitle"`   // 标题
	ArticleContent string `json:"articleContent"` // 内容
	Type           uint32 `json:"type"`           // 文章类型
	OriginalUrl    string `json:"originalUrl"`    // 原文链接
	IsTop          uint32 `json:"isTop"`          // 是否置顶
	IsDelete       uint32 `json:"isDelete"`       // 是否删除
	Status         uint32 `json:"status"`         // 文章状态 1-公开 2-私密 3-评论可见
}

func (a ArticlePO) TableName() string {
	return "tb_article"
}
