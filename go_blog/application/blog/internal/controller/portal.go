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
