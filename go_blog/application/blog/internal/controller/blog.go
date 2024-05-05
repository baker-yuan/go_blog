package controller

import (
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
	portal        usecase.IPortalUseCase
}

// 强制BlogServiceImpl实现pb.BlogApiService
var _ pb.BlogApiService = &BlogServiceImpl{}

// 强制BlogServiceImpl实现pb.BlogPortalApiService
var _ pb.BlogPortalApiService = &BlogServiceImpl{}

// NewBlogServiceImpl 创建接口实现
func NewBlogServiceImpl(
	articleUseCase usecase.IArticleUseCase,
	categoryUseCase usecase.ICategoryUseCase,
	tag usecase.ITagUseCase,
	friendLink usecase.IFriendLinkUseCase,
	page usecase.IPageUseCase,
	dynamicConfig usecase.IDynamicConfigUseCase,
	portal usecase.IPortalUseCase,
) *BlogServiceImpl {
	return &BlogServiceImpl{
		article:       articleUseCase,
		category:      categoryUseCase,
		tag:           tag,
		friendLink:    friendLink,
		page:          page,
		dynamicConfig: dynamicConfig,
		portal:        portal,
	}
}
