package service

import (
	"context"

	pb "github.com/baker-yuan/go-blog/application/blog/api/blog/v1"
	"github.com/baker-yuan/go-blog/application/blog/internal/biz"
)

func NewBlogService(article *biz.ArticleUsecase) *BlogService {
	return &BlogService{
		article: article,
	}
}

// SearchArticle 文章搜索
func (b *BlogService) SearchArticle(ctx context.Context, req *pb.SearchArticleReq) (*pb.SearchArticleRsp, error) {
	rsp := &pb.SearchArticleRsp{}
	articles, pageTotal, err := b.article.SearchArticle(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = articles
	rsp.Total = pageTotal
	return rsp, nil
}

// AddOrUpdateArticle 添加修改文章
func (b *BlogService) AddOrUpdateArticle(ctx context.Context, req *pb.AddOrUpdateArticleReq) (*pb.AddOrUpdateRsp, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteArticle 删除文章
func (b *BlogService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleReq) (*pb.EmptyRsp, error) {
	//TODO implement me
	panic("implement me")
}

// ArticleDetail 文章详情
func (b *BlogService) ArticleDetail(ctx context.Context, req *pb.ArticleDetailReq) (*pb.Article, error) {
	//TODO implement me
	panic("implement me")
}
