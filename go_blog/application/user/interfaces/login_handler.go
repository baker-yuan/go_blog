package interfaces

import (
	"context"

	"github.com/baker-yuan/go-blog/application/user/application"
	"github.com/baker-yuan/go-blog/application/user/infrastructure/auth"
	pb "github.com/baker-yuan/go-blog/protocol/user"
)

type Authenticate struct {
	us application.UserAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

// 强制Authenticate实现LoginApiService
var _ pb.LoginApiService = &Authenticate{}

// NewAuthenticate constructor
func NewAuthenticate(uApp application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Authenticate {
	return &Authenticate{
		us: uApp,
		rd: rd,
		tk: tk,
	}
}

// Login 登录
func (a Authenticate) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRsp, error) {
	//TODO implement me
	panic("implement me")
}

// Logout 退出
func (a Authenticate) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutRsp, error) {
	//TODO implement me
	panic("implement me")
}

// Refresh 刷新token
func (a Authenticate) Refresh(ctx context.Context, req *pb.RefreshReq) (*pb.RefreshRsp, error) {
	//TODO implement me
	panic("implement me")
}
