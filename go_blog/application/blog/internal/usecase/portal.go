package usecase

import (
	"context"

	"github.com/baker-yuan/go-blog/go_blog/application/blog/config"
	"github.com/baker-yuan/go-blog/go_blog/application/blog/internal/usecase/assembler"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
	"trpc.group/trpc-go/trpc-go"
)

// PortalUseCase 客户端api
type PortalUseCase struct {
	cfg *config.Config
	*CommonUseCase
	articleRepo IArticleRepo
}

// 强制PortalUseCase实现IPortalUseCase
var _ IPortalUseCase = &PortalUseCase{}

// NewPortalUseCase 创建客户端api逻辑实现
func NewPortalUseCase(
	cfg *config.Config,
	commonUseCase *CommonUseCase,
	articleRepo IArticleRepo,
) *PortalUseCase {
	return &PortalUseCase{
		cfg:           cfg,
		CommonUseCase: commonUseCase,
		articleRepo:   articleRepo,
	}
}

// Home 获取站点配置信息、最新推荐博客、分类列表、标签云、随机博客
func (c *PortalUseCase) Home(ctx context.Context, userID uint32, req *pb.HomeReq) (*pb.HomeRsp, error) {
	var (
		newBlogList       = make([]*pb.HomeRsp_HomeBlog, 0)      // 最新博客
		randomBlogList    = make([]*pb.HomeRsp_RandomBlog, 0)    // 推荐博客
		newestCommentList = make([]*pb.HomeRsp_NewestComment, 0) // 最新评论
		tagList           = make([]*pb.HomeRsp_Tag, 0)           // 标签
		categoryList      = make([]*pb.HomeRsp_Category, 0)      // 分类
		friendLinkList    = make([]*pb.HomeRsp_FriendLink, 0)    // 友情链接
	)

	// 最新博客
	searchNewBlog := func() error {
		newestArticleList, err := c.articleRepo.GetNewestArticle(ctx, c.cfg.NewBlogLimitNum)
		if err != nil {
			return err
		}
		newBlogList = assembler.ArticleEntityToHomeHomeBlogModel(newestArticleList)
		return nil
	}
	// 推荐博客
	searchRandomBlog := func() error {
		recommendArticleList, err := c.articleRepo.GetRecommendArticle(ctx, c.cfg.RandomBlogLimitNum)
		if err != nil {
			return err
		}
		randomBlogList = assembler.ArticleEntityToHomeRandomBlogModel(recommendArticleList)
		return nil
	}

	if err := trpc.GoAndWait(searchNewBlog, searchRandomBlog); err != nil {
		return nil, err
	}

	return &pb.HomeRsp{
		NewBlogList:       newBlogList,
		RandomBlogList:    randomBlogList,
		NewestCommentList: newestCommentList,
		TagList:           tagList,
		CategoryList:      categoryList,
		FriendLinkList:    friendLinkList,
	}, nil
}

// BlogHomeInfo 查看博客配置信息
func (c *PortalUseCase) BlogHomeInfo(ctx context.Context, id uint32, req *pb.BlogHomeInfoReq) (*pb.BlogHomeInfoRsp, error) {
	res := &pb.BlogHomeInfoRsp{
		WebsiteConfig: &pb.WebsiteConfig{},
		PageList:      make([]*pb.BlogHomeInfoRsp_Page, 0),
	}

	return res, nil
}

// ArticleSearch 文章搜索
func (c *PortalUseCase) ArticleSearch(ctx context.Context, userID uint32, req *pb.ArticleSearchReq) ([]*pb.ArticleSearchItem, uint32, error) {
	rsp := make([]*pb.ArticleSearchItem, 0)
	rsp = append(rsp, &pb.ArticleSearchItem{})
	return rsp, 1, nil
}
