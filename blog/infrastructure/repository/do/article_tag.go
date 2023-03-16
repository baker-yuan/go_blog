package do

import "github.com/baker-yuan/go-blog/all_packaged_library/common"

// ArticleTag 标签
type ArticleTag struct {
	common.Model
	ArticleId uint32 `json:"articleId"` // 文章id
	TagId     uint32 `json:"tagId"`     // 标签id
}

func (a ArticleTag) TableName() string {
	return "tb_article_tag"
}
