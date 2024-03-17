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
		// SearchArticle 文章搜索
		SearchArticle(ctx context.Context, req *pb.SearchArticleReq) (entity.Articles, uint32, error)
	}
)

// Category 文章分类
type (
	// ICategoryUseCase 业务逻辑
	ICategoryUseCase interface {
		// SearchCategory 查询文章分类
		SearchCategory(ctx context.Context, req *pb.SearchCategoryReq) ([]*pb.Category, uint32, error)
		// CategoryDetail 文章分类详情
		CategoryDetail(ctx context.Context, req *pb.CategoryDetailReq) (*pb.Category, error)
		// AddOrUpdateCategory 添加修改文章分类
		AddOrUpdateCategory(ctx context.Context, req *pb.AddOrUpdateCategoryReq) (uint32, error)
		// DeleteCategory 删除文章分类
		DeleteCategory(ctx context.Context, req *pb.DeleteCategoryReq) error
	}

	// ICategoryRepo 数据存储操作
	ICategoryRepo interface {
		// GetCategoryByID 根据文章分类id集合查询文章分类
		GetCategoryByID(ctx context.Context, id uint32) (*entity.Category, error)
		// GetCategoryByIDs 根据文章分类id集合查询文章分类
		GetCategoryByIDs(ctx context.Context, ids []uint32) (entity.CategoryList, error)
		// Save 保存文章分类
		Save(ctx context.Context, category *entity.Category) (uint32, error)
		// UpdateByID 根据ID修改文章分类
		UpdateByID(ctx context.Context, category *entity.Category) error
		// DeleteByID 根据ID删除文章分类
		DeleteByID(ctx context.Context, id uint32) error
		// SearchCategory 文章分类搜索
		SearchCategory(ctx context.Context, req *pb.SearchCategoryReq) (entity.CategoryList, uint32, error)
	}
)
