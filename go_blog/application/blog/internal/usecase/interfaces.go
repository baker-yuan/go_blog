// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/entity"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

type (
	// ICommonUseCase 通用service
	ICommonUseCase interface {
		// SaveChangeLog 保存变更日志
		SaveChangeLog(ctx context.Context,
			resourceID uint32, resourceType pb.ResourceType,
			changeBefore interface{}, changeAfter interface{},
			notes string)
		// UidToName 根据用户ID获取用户名称
		UidToName(ctx context.Context, userIDs ...uint32) (map[uint32]string, error)
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

// Tag 文章标签
type (
	// ITagUseCase 业务逻辑
	ITagUseCase interface {
		// SearchTag 查询文章标签
		SearchTag(ctx context.Context, req *pb.SearchTagReq) ([]*pb.Tag, uint32, error)
		// TagDetail 文章标签详情
		TagDetail(ctx context.Context, req *pb.TagDetailReq) (*pb.Tag, error)
		// AddOrUpdateTag 添加修改文章标签
		AddOrUpdateTag(ctx context.Context, req *pb.AddOrUpdateTagReq) (uint32, error)
		// DeleteTag 删除文章标签
		DeleteTag(ctx context.Context, req *pb.DeleteTagReq) error
	}

	// ITagRepo 数据存储操作
	ITagRepo interface {
		// GetTagByID 根据文章标签id集合查询文章标签
		GetTagByID(ctx context.Context, id uint32) (*entity.Tag, error)
		// GetTagByIDs 根据文章标签id集合查询文章标签
		GetTagByIDs(ctx context.Context, ids []uint32) (entity.Tags, error)
		// Save 保存文章标签
		Save(ctx context.Context, tag *entity.Tag) (uint32, error)
		// UpdateByID 根据ID修改文章标签
		UpdateByID(ctx context.Context, tag *entity.Tag) error
		// DeleteByID 根据ID删除文章标签
		DeleteByID(ctx context.Context, id uint32) error
		// SearchTag 文章标签搜索
		SearchTag(ctx context.Context, req *pb.SearchTagReq) (entity.Tags, uint32, error)
	}
)

// FriendLink 友链表
type (
	// IFriendLinkUseCase 业务逻辑
	IFriendLinkUseCase interface {
		// SearchFriendLink 查询友链表
		SearchFriendLink(ctx context.Context, req *pb.SearchFriendLinkReq) ([]*pb.FriendLink, uint32, error)
		// FriendLinkDetail 友链表详情
		FriendLinkDetail(ctx context.Context, req *pb.FriendLinkDetailReq) (*pb.FriendLink, error)
		// AddOrUpdateFriendLink 添加修改友链表
		AddOrUpdateFriendLink(ctx context.Context, req *pb.AddOrUpdateFriendLinkReq) (uint32, error)
		// DeleteFriendLink 删除友链表
		DeleteFriendLink(ctx context.Context, req *pb.DeleteFriendLinkReq) error
	}

	// IFriendLinkRepo 数据存储操作
	IFriendLinkRepo interface {
		// GetFriendLinkByID 根据友链表id集合查询友链表
		GetFriendLinkByID(ctx context.Context, id uint32) (*entity.FriendLink, error)
		// GetFriendLinkByIDs 根据友链表id集合查询友链表
		GetFriendLinkByIDs(ctx context.Context, ids []uint32) (entity.FriendLinks, error)
		// Save 保存友链表
		Save(ctx context.Context, friendLink *entity.FriendLink) (uint32, error)
		// UpdateByID 根据ID修改友链表
		UpdateByID(ctx context.Context, friendLink *entity.FriendLink) error
		// DeleteByID 根据ID删除友链表
		DeleteByID(ctx context.Context, id uint32) error
		// SearchFriendLink 友链表搜索
		SearchFriendLink(ctx context.Context, req *pb.SearchFriendLinkReq) (entity.FriendLinks, uint32, error)
	}
)

// Page 页面
type (
	// IPageUseCase 业务逻辑
	IPageUseCase interface {
		// SearchPage 查询页面
		SearchPage(ctx context.Context, req *pb.SearchPageReq) ([]*pb.Page, uint32, error)
		// PageDetail 页面详情
		PageDetail(ctx context.Context, req *pb.PageDetailReq) (*pb.Page, error)
		// AddOrUpdatePage 添加修改页面
		AddOrUpdatePage(ctx context.Context, req *pb.AddOrUpdatePageReq) (uint32, error)
		// DeletePage 删除页面
		DeletePage(ctx context.Context, req *pb.DeletePageReq) error
	}

	// IPageRepo 数据存储操作
	IPageRepo interface {
		// GetPageByID 根据页面id集合查询页面
		GetPageByID(ctx context.Context, id uint32) (*entity.Page, error)
		// GetPageByIDs 根据页面id集合查询页面
		GetPageByIDs(ctx context.Context, ids []uint32) (entity.Pages, error)
		// Save 保存页面
		Save(ctx context.Context, page *entity.Page) (uint32, error)
		// UpdateByID 根据ID修改页面
		UpdateByID(ctx context.Context, page *entity.Page) error
		// DeleteByID 根据ID删除页面
		DeleteByID(ctx context.Context, id uint32) error
		// SearchPage 页面搜索
		SearchPage(ctx context.Context, req *pb.SearchPageReq) (entity.Pages, uint32, error)
	}
)

// DynamicConfig 动态配置
type (
	// IDynamicConfigUseCase 业务逻辑
	IDynamicConfigUseCase interface {
		// SearchDynamicConfig 查询动态配置
		SearchDynamicConfig(ctx context.Context, req *pb.SearchDynamicConfigReq) ([]*pb.DynamicConfig, uint32, error)
		// DynamicConfigDetail 动态配置详情
		DynamicConfigDetail(ctx context.Context, req *pb.DynamicConfigDetailReq) (*pb.DynamicConfig, error)
		// AddOrUpdateDynamicConfig 添加修改动态配置
		AddOrUpdateDynamicConfig(ctx context.Context, req *pb.AddOrUpdateDynamicConfigReq) (uint32, error)
		// DeleteDynamicConfig 删除动态配置
		DeleteDynamicConfig(ctx context.Context, req *pb.DeleteDynamicConfigReq) error
	}

	// IDynamicConfigRepo 数据存储操作
	IDynamicConfigRepo interface {
		// GetDynamicConfigByID 根据动态配置id集合查询动态配置
		GetDynamicConfigByID(ctx context.Context, id uint32) (*entity.DynamicConfig, error)
		// GetDynamicConfigByIDs 根据动态配置id集合查询动态配置
		GetDynamicConfigByIDs(ctx context.Context, ids []uint32) (entity.DynamicConfigs, error)
		// Save 保存动态配置
		Save(ctx context.Context, dynamicConfig *entity.DynamicConfig) (uint32, error)
		// UpdateByID 根据ID修改动态配置
		UpdateByID(ctx context.Context, dynamicConfig *entity.DynamicConfig) error
		// DeleteByID 根据ID删除动态配置
		DeleteByID(ctx context.Context, id uint32) error
		// SearchDynamicConfig 动态配置搜索
		SearchDynamicConfig(ctx context.Context, req *pb.SearchDynamicConfigReq) (entity.DynamicConfigs, uint32, error)
	}
)

// Portal 客户端接口
type (
	// IPortalUseCase 业务逻辑
	IPortalUseCase interface {
		// Home 获取站点配置信息、最新推荐博客、分类列表、标签云、随机博客
		Home(ctx context.Context, userID uint32, req *pb.HomeReq) (*pb.HomeRsp, error)
	}
)
