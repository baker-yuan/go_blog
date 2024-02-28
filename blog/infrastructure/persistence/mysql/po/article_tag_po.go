package po

import "github.com/baker-yuan/go-blog/all_packaged_library/do"

// ArticleTagPO 标签
type ArticleTagPO struct {
	do.Model
	ArticleId uint32 `json:"articleId"` // 文章id
	TagId     uint32 `json:"tagId"`     // 标签id
}

func (a ArticleTagPO) TableName() string {
	return "tb_article_tag"
}
