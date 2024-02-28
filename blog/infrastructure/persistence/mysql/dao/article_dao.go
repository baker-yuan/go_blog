// Package dao DAO 层
package dao

import (
	"github.com/baker-yuan/go-blog/blog/infrastructure/persistence/mysql/po"
)

// ArticleDAO Article 的 DAO
type ArticleDAO struct {
	FollowPO po.ArticlePO
}
