package interfaces

import (
	"context"

	"github.com/baker-yuan/go-blog/application/user/application"
	"github.com/baker-yuan/go-blog/application/user/infrastructure/auth"
	pb "github.com/baker-yuan/go-blog/protocol/user"
)

// 强制Users实现UserApiService
var _ pb.UserApiService = &Users{}

// Users struct defines the dependencies that will be used
type Users struct {
	us application.UserAppInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

// NewUsers Users constructor
func NewUsers(us application.UserAppInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Users {
	return &Users{
		us: us,
		rd: rd,
		tk: tk,
	}
}

// SearchUser 用户搜索
func (s *Users) SearchUser(ctx context.Context, req *pb.SearchUserReq) (*pb.SearchUserRsp, error) {
	rsp := &pb.SearchUserRsp{}
	users, pageTotal, err := s.us.SearchUser(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Data = users
	rsp.Total = pageTotal
	return rsp, nil
}

// AddOrUpdateUser 添加修改用户
func (s *Users) AddOrUpdateUser(ctx context.Context, req *pb.AddOrUpdateUserReq) (*pb.AddOrUpdateRsp, error) {
	rsp := &pb.AddOrUpdateRsp{}
	id, err := s.us.AddOrUpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp.Id = id
	return rsp, nil
}

// DeleteUser 删除用户
func (s *Users) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.EmptyRsp, error) {
	rsp := &pb.EmptyRsp{}
	if err := s.us.DeleteUser(ctx, req); err != nil {
		return nil, err
	}
	return rsp, nil
}

// UserDetail 用户详情
func (s *Users) UserDetail(ctx context.Context, req *pb.UserDetailReq) (*pb.User, error) {
	return s.us.UserDetail(ctx, req)
}
