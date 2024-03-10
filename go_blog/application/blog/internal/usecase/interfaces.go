// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/baker-yuan/go-blog/application/blog/internal/entity"
	pb "github.com/baker-yuan/go-blog/protocol/blog"
)

type (
	// ICommonUseCase 通用service
	ICommonUseCase interface {
		// SaveChangeLog 保存变更日志
		SaveChangeLog(ctx context.Context,
			resourceID uint32, resourceType pb.ResourceType,
			changeBefore interface{}, changeAfter interface{},
			notes string)
	}
)

type (
	// IArticleUseCase 文章service
	IArticleUseCase interface {
		// SearchArticle 查询文章
		SearchArticle(ctx context.Context, req *pb.SearchArticleReq) ([]*pb.Article, uint32, error)
		// ArticleDetail 文章详情
		ArticleDetail(ctx context.Context, req *pb.ArticleDetailReq) (*pb.Article, error)
		// AddOrUpdateArticle 添加修改文章
		AddOrUpdateArticle(ctx context.Context, req *pb.AddOrUpdateArticleReq) (uint32, error)
		// DeleteArticle 删除文章
		DeleteArticle(ctx context.Context, req *pb.DeleteArticleReq) error
	}
	// IArticleRepo 文章repo
	IArticleRepo interface {
		// SearchArticle 文章搜索
		SearchArticle(ctx context.Context, req *pb.SearchArticleReq) (entity.Articles, uint32, error)
		// GetArticleByID 根据文章id集合查询文章
		GetArticleByID(ctx context.Context, id uint32) (*entity.Article, error)
		// GetArticleByIDs 根据文章id集合查询文章
		GetArticleByIDs(ctx context.Context, ids []uint32) (entity.Articles, error)
		// Save 保存文章
		Save(ctx context.Context, article *entity.Article) (uint32, error)
		// UpdateByID 根据ID修改文章
		UpdateByID(ctx context.Context, article *entity.Article) error
		// DeleteByID 根据ID删除文章
		DeleteByID(ctx context.Context, id uint32) error
	}
)
