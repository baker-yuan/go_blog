package session

import (
	"context"

	"trpc.group/trpc-go/trpc-go"
)

// SessionUtils 会话
type SessionUtils struct{}

// GetLoginUserID 获取登录态中的用户ID
func (c SessionUtils) GetLoginUserID(ctx context.Context) (uint32, error) {
	// todo
	return 0, nil
}

// GetToken 获取token
func (c SessionUtils) GetToken(ctx context.Context) (string, error) {
	metaData := trpc.Message(ctx).ServerMetaData()
	bearToken := string(metaData["X-Access-Token"])
	// todo
	return bearToken, nil
}
