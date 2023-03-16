package repo

import (
	"github.com/baker-yuan/go-blog/blog/application/cqe"
	"github.com/baker-yuan/go-blog/blog/application/dto"
	"github.com/baker-yuan/go-blog/blog/domain/entity"
)

type ArticleRepo interface {
	// FindById 通过文章ID查找
	FindById(articleID uint32) (*entity.Article, error)
	Save(article *entity.Article) error

	// ListArchives 文章归档
	ListArchives(currentPage uint32, pageSize uint32) (articles []*dto.ArchiveDTO, total uint32, err error)

	// ListArticleBacks 分页查询后台文章
	ListArticleBacks(condition cqe.ConditionVO, currentPage uint32, pageSize uint32) ([]*dto.ArticleBackDTO, uint32, error)
}
