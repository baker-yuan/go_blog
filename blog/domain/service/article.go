package service

import (
	"github.com/baker-yuan/go-blog/blog/application/cqe"
	"github.com/baker-yuan/go-blog/blog/application/dto"
)

// ArticleService 文章服务
type ArticleService struct {
	repo port.ArticleRepo
}

func NewArticleService(repo port.ArticleRepo) ArticleService {
	return ArticleService{
		repo: repo,
	}
}

// ListArticleBacks 分页查询后台文章
//
// param condition 条件
// return 文章列表
func (a *ArticleService) ListArticleBacks(condition cqe.ConditionVO, currentPage uint32, pageSize uint32) ([]*dto.ArticleBackDTO, uint32, error) {
	var (
		err      error
		articles []*dto.ArticleBackDTO
		total    uint32
	)
	// 分页查询文章
	articles, total, err = a.repo.ListArticleBacks(condition, currentPage, pageSize)
	if err != nil {
		return nil, 0, err
	}
	// 查询文章点赞量和浏览量

	// 封装点赞量和浏览量
	return articles, total, nil
}
