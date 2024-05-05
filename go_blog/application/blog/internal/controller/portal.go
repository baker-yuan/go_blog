package controller

import (
	"context"

	"github.com/baker-yuan/go-blog/go_blog/common/util"
	pb "github.com/baker-yuan/go-blog/go_blog/protocol/blog"
)

// Home 获取站点配置信息、最新推荐博客、分类列表、标签云、随机博客
func (m *BlogServiceImpl) Home(ctx context.Context, req *pb.HomeReq) (*pb.HomeRsp, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return nil, err
	}
	rsp, err := m.portal.Home(ctx, userID, req)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// BlogHomeInfo 查看博客配置信息
func (m *BlogServiceImpl) BlogHomeInfo(ctx context.Context, req *pb.BlogHomeInfoReq) (*pb.BlogHomeInfoRsp, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return nil, err
	}
	rsp, err := m.portal.BlogHomeInfo(ctx, userID, req)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// ArticleSearch 文章搜索
func (m *BlogServiceImpl) ArticleSearch(ctx context.Context, req *pb.ArticleSearchReq) (*pb.ArticleSearchRsp, error) {
	userID, err := util.SessionUtils.GetLoginUserID(ctx)
	if err != nil {
		return nil, err
	}
	rows, total, err := m.portal.ArticleSearch(ctx, userID, req)
	if err != nil {
		return nil, err
	}
	return &pb.ArticleSearchRsp{
		Total: total,
		Rows:  rows,
	}, nil
}
