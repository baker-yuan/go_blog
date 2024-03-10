// Package service 实现了 api 定义的服务层，类似 DDD 的 application 层，
// 处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
package service

import (
	"github.com/baker-yuan/go-blog/application/auth/internal/biz"
	pb "github.com/baker-yuan/go-blog/protocol/auth"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAuthService)

// 强制编译时检查 *AuthService 类型是否实现了 AuthApiService 接口
var _ pb.AuthApiService = (*AuthService)(nil)

type AuthService struct {
	pb.UnimplementedAuthApi

	article *biz.MenuUsecase
}
