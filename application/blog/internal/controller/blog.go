package controller

import (
	"context"

	"github.com/baker-yuan/go-blog/application/blog/internal/usecase"
	pb "github.com/baker-yuan/go-blog/protocol/blog"
)

// BlogServiceImpl 接口实现
type BlogServiceImpl struct {
	articleUseCase usecase.IArticleUseCase
}

// NewBlogServiceImpl 创建接口实现
func NewBlogServiceImpl(articleUseCase usecase.IArticleUseCase) pb.BlogApiService {
	return &BlogServiceImpl{
		articleUseCase: articleUseCase,
	}
}

// SearchArticle 文章搜索
func (m *BlogServiceImpl) SearchArticle(ctx context.Context, req *pb.SearchArticleReq) (*pb.SearchArticleRsp, error) {
	rsp := &pb.SearchArticleRsp{}
	articles, pageTotal, err := m.articleUseCase.SearchArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = articles
	rsp.Total = pageTotal
	return rsp, nil
}

// AddOrUpdateArticle 添加修改文章
func (m *BlogServiceImpl) AddOrUpdateArticle(ctx context.Context, req *pb.AddOrUpdateArticleReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := m.articleUseCase.AddOrUpdateArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteArticle 删除文章
func (m *BlogServiceImpl) DeleteArticle(ctx context.Context, req *pb.DeleteArticleReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := m.articleUseCase.DeleteArticle(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// ArticleDetail 文章详情
func (m *BlogServiceImpl) ArticleDetail(ctx context.Context, req *pb.ArticleDetailReq) (*pb.Article, error) {
	return m.articleUseCase.ArticleDetail(ctx, req)
}
