package auth

import (
	"context"

	"github.com/baker-yuan/go-blog/all_packaged_library/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type CurrentUser struct {
	ID       uint32 // 用户ID
	Username string // 用户名
}

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(ctx context.Context) (*CurrentUser, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "failed to get metadata")
	}
	if _, ok := md["username"]; !ok {
		return nil, status.Errorf(codes.DataLoss, "failed to get metadata username")
	}
	if _, ok := md["uid"]; !ok {
		return nil, status.Errorf(codes.DataLoss, "failed to get metadata uid")
	}
	return &CurrentUser{
		Username: md["username"][0],
		ID:       util.StrToUInt32(md["uid"][0]),
	}, nil
}
