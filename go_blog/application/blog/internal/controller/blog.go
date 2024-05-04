package controller

import (
	"context"

	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/usecase"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// BlogServiceImpl 接口实现
type BlogServiceImpl struct {
	article       usecase.IArticleUseCase
	category      usecase.ICategoryUseCase
	tag           usecase.ITagUseCase
	friendLink    usecase.IFriendLinkUseCase
	page          usecase.IPageUseCase
	dynamicConfig usecase.IDynamicConfigUseCase
}

// 强制BlogServiceImpl实现pb.BlogApiService
var _ pb.BlogApiService = &BlogServiceImpl{}

// NewBlogServiceImpl 创建接口实现
func NewBlogServiceImpl(
	articleUseCase usecase.IArticleUseCase,
	categoryUseCase usecase.ICategoryUseCase,
	tag usecase.ITagUseCase,
	friendLink usecase.IFriendLinkUseCase,
	page usecase.IPageUseCase,
	dynamicConfig usecase.IDynamicConfigUseCase,
) pb.BlogApiService {
	return &BlogServiceImpl{
		article:       articleUseCase,
		category:      categoryUseCase,
		tag:           tag,
		friendLink:    friendLink,
		page:          page,
		dynamicConfig: dynamicConfig,
	}
}

// SearchArticle 文章搜索
func (m *BlogServiceImpl) SearchArticle(ctx context.Context, req *pb.SearchArticleReq) (*pb.SearchArticleRsp, error) {
	rsp := &pb.SearchArticleRsp{}
	articles, pageTotal, err := m.article.SearchArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = articles
	rsp.PageTotal = pageTotal
	rsp.PageNum = req.GetPageNum()
	rsp.PageSize = req.GetPageSize()
	return rsp, nil
}

// AddOrUpdateArticle 添加修改文章
func (m *BlogServiceImpl) AddOrUpdateArticle(ctx context.Context, req *pb.AddOrUpdateArticleReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := m.article.AddOrUpdateArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteArticle 删除文章
func (m *BlogServiceImpl) DeleteArticle(ctx context.Context, req *pb.DeleteArticleReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := m.article.DeleteArticle(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// ArticleDetail 文章详情
func (m *BlogServiceImpl) ArticleDetail(ctx context.Context, req *pb.ArticleDetailReq) (*pb.Article, error) {
	return m.article.ArticleDetail(ctx, req)
}

// SearchCategory 查询文章分类
func (m *BlogServiceImpl) SearchCategory(ctx context.Context, req *pb.SearchCategoryReq) (*pb.SearchCategoryRsp, error) {
	rsp := &pb.SearchCategoryRsp{}
	categoryList, pageTotal, err := m.category.SearchCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = categoryList
	rsp.PageTotal = pageTotal
	rsp.PageNum = req.GetPageNum()
	rsp.PageSize = req.GetPageSize()
	return rsp, nil
}

// CategoryDetail 文章分类详情
func (m *BlogServiceImpl) CategoryDetail(ctx context.Context, req *pb.CategoryDetailReq) (*pb.Category, error) {
	return m.category.CategoryDetail(ctx, req)
}

// AddOrUpdateCategory 添加修改文章分类
func (m *BlogServiceImpl) AddOrUpdateCategory(ctx context.Context, req *pb.AddOrUpdateCategoryReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := m.category.AddOrUpdateCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteCategory 删除文章分类
func (m *BlogServiceImpl) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := m.category.DeleteCategory(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// SearchTag 查询文章标签
func (m *BlogServiceImpl) SearchTag(ctx context.Context, req *pb.SearchTagReq) (*pb.SearchTagRsp, error) {
	rsp := &pb.SearchTagRsp{}
	tags, pageTotal, err := m.tag.SearchTag(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = tags
	rsp.PageTotal = pageTotal
	rsp.PageNum = req.GetPageNum()
	rsp.PageSize = req.GetPageSize()
	return rsp, nil
}

// TagDetail 文章标签详情
func (m *BlogServiceImpl) TagDetail(ctx context.Context, req *pb.TagDetailReq) (*pb.Tag, error) {
	return m.tag.TagDetail(ctx, req)
}

// AddOrUpdateTag 添加修改文章标签
func (m *BlogServiceImpl) AddOrUpdateTag(ctx context.Context, req *pb.AddOrUpdateTagReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := m.tag.AddOrUpdateTag(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteTag 删除文章标签
func (m *BlogServiceImpl) DeleteTag(ctx context.Context, req *pb.DeleteTagReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := m.tag.DeleteTag(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// SearchFriendLink 查询友链表
func (m *BlogServiceImpl) SearchFriendLink(ctx context.Context, req *pb.SearchFriendLinkReq) (*pb.SearchFriendLinkRsp, error) {
	rsp := &pb.SearchFriendLinkRsp{}
	friendLinks, pageTotal, err := m.friendLink.SearchFriendLink(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = friendLinks
	rsp.PageTotal = pageTotal
	rsp.PageNum = req.GetPageNum()
	rsp.PageSize = req.GetPageSize()
	return rsp, nil
}

// FriendLinkDetail 友链表详情
func (m *BlogServiceImpl) FriendLinkDetail(ctx context.Context, req *pb.FriendLinkDetailReq) (*pb.FriendLink, error) {
	return m.friendLink.FriendLinkDetail(ctx, req)
}

// AddOrUpdateFriendLink 添加修改友链表
func (m *BlogServiceImpl) AddOrUpdateFriendLink(ctx context.Context, req *pb.AddOrUpdateFriendLinkReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := m.friendLink.AddOrUpdateFriendLink(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteFriendLink 删除友链表
func (m *BlogServiceImpl) DeleteFriendLink(ctx context.Context, req *pb.DeleteFriendLinkReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := m.friendLink.DeleteFriendLink(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// SearchPage 查询页面
func (m *BlogServiceImpl) SearchPage(ctx context.Context, req *pb.SearchPageReq) (*pb.SearchPageRsp, error) {
	rsp := &pb.SearchPageRsp{}
	pages, pageTotal, err := m.page.SearchPage(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = pages
	rsp.PageTotal = pageTotal
	rsp.PageNum = req.GetPageNum()
	rsp.PageSize = req.GetPageSize()
	return rsp, nil
}

// PageDetail 页面详情
func (m *BlogServiceImpl) PageDetail(ctx context.Context, req *pb.PageDetailReq) (*pb.Page, error) {
	return m.page.PageDetail(ctx, req)
}

// AddOrUpdatePage 添加修改页面
func (m *BlogServiceImpl) AddOrUpdatePage(ctx context.Context, req *pb.AddOrUpdatePageReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := m.page.AddOrUpdatePage(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeletePage 删除页面
func (m *BlogServiceImpl) DeletePage(ctx context.Context, req *pb.DeletePageReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := m.page.DeletePage(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// SearchDynamicConfig 查询动态配置
func (m *BlogServiceImpl) SearchDynamicConfig(ctx context.Context, req *pb.SearchDynamicConfigReq) (*pb.SearchDynamicConfigRsp, error) {
	rsp := &pb.SearchDynamicConfigRsp{}
	dynamicConfigs, pageTotal, err := m.dynamicConfig.SearchDynamicConfig(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = dynamicConfigs
	rsp.PageTotal = pageTotal
	rsp.PageNum = req.GetPageNum()
	rsp.PageSize = req.GetPageSize()
	return rsp, nil
}

// DynamicConfigDetail 动态配置详情
func (m *BlogServiceImpl) DynamicConfigDetail(ctx context.Context, req *pb.DynamicConfigDetailReq) (*pb.DynamicConfig, error) {
	return m.dynamicConfig.DynamicConfigDetail(ctx, req)
}

// AddOrUpdateDynamicConfig 添加修改动态配置
func (m *BlogServiceImpl) AddOrUpdateDynamicConfig(ctx context.Context, req *pb.AddOrUpdateDynamicConfigReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := m.dynamicConfig.AddOrUpdateDynamicConfig(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteDynamicConfig 删除动态配置
func (m *BlogServiceImpl) DeleteDynamicConfig(ctx context.Context, req *pb.DeleteDynamicConfigReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := m.dynamicConfig.DeleteDynamicConfig(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}
