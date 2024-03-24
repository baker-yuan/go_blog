package interfaces

import (
	"context"

	"github.com/baker-yuan/go-blog/application/user/application"
	pb "github.com/baker-yuan/go-blog/protocol/user"
)

type Authenticate struct {
	us application.UserAppInterface
}

// 强制Authenticate实现LoginApiService
var _ pb.LoginApiService = &Authenticate{}

// NewAuthenticate constructor
func NewAuthenticate(uApp application.UserAppInterface) *Authenticate {
	return &Authenticate{
		us: uApp,
	}
}

// Login 登录
func (a Authenticate) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error) {
	return a.us.Login(ctx, req)
}

// Logout 退出
func (a Authenticate) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutRsp, error) {
	return a.us.Logout(ctx, req)
}

// Refresh 刷新token
func (a Authenticate) Refresh(ctx context.Context, req *pb.RefreshReq) (*pb.RefreshRsp, error) {
	return a.us.Refresh(ctx, req)
}
