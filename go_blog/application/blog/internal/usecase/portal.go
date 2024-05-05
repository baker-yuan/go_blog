package usecase

import (
	"context"

	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// PortalUseCase 客户端api
type PortalUseCase struct {
	*CommonUseCase
}

// 强制PortalUseCase实现IPortalUseCase
var _ IPortalUseCase = &PortalUseCase{}

// NewPortalUseCase 创建客户端api逻辑实现
func NewPortalUseCase(
	commonUseCase *CommonUseCase,
) *PortalUseCase {
	return &PortalUseCase{
		CommonUseCase: commonUseCase,
	}
}

// Home 获取站点配置信息、最新推荐博客、分类列表、标签云、随机博客
func (c *PortalUseCase) Home(ctx context.Context, userID uint32, req *pb.HomeReq) (*pb.HomeRsp, error) {
	var (
		newBlogList       = make([]*pb.HomeRsp_HomeBlog, 0)
		randomBlogList    = make([]*pb.HomeRsp_RandomBlog, 0)
		newestCommentList = make([]*pb.HomeRsp_NewestComment, 0)
		tagList           = make([]*pb.HomeRsp_Tag, 0)
		categoryList      = make([]*pb.HomeRsp_Category, 0)
		friendLinkList    = make([]*pb.HomeRsp_FriendLink, 0)
	)

	return &pb.HomeRsp{
		NewBlogList:       newBlogList,
		RandomBlogList:    randomBlogList,
		NewestCommentList: newestCommentList,
		TagList:           tagList,
		CategoryList:      categoryList,
		FriendLinkList:    friendLinkList,
	}, nil
}
