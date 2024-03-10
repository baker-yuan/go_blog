// Package service 实现了 api 定义的服务层，类似 DDD 的 application 层，
// 处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
package service

import (
	pb "github.com/baker-yuan/go-blog/application/blog/api/blog/v1"
	"github.com/baker-yuan/go-blog/application/blog/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlogService)

// 强制编译时检查 *BlogService 类型是否实现了 BlogApiService 接口
var _ pb.BlogApiService = (*BlogService)(nil)

type BlogService struct {
	pb.UnimplementedBlogApi

	article *biz.ArticleUsecase
}
