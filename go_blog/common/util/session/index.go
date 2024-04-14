package session

import (
	"context"
	"strconv"

	"trpc.group/trpc-go/trpc-go"
	thttp "trpc.group/trpc-go/trpc-go/http"
)

// SessionUtils 会话
type SessionUtils struct{}

// GetLoginUserID 获取登录态中的用户ID
func (c SessionUtils) GetLoginUserID(ctx context.Context) (uint32, error) {
	// https://github.com/trpc-group/trpc-go/blob/main/http/README.zh_CN.md
	//metaData := trpc.Message(ctx).ServerMetaData()
	//codec.Message(ctx).ServerMetaData()
	head := thttp.Head(ctx)
	reqHead := head.Request.Header
	userIDStr := reqHead["X-User-Id"]
	userID, _ := strconv.Atoi(userIDStr[0])
	return uint32(userID), nil
}

// GetToken 获取token
func (c SessionUtils) GetToken(ctx context.Context) (string, error) {
	metaData := trpc.Message(ctx).ServerMetaData()
	bearToken := string(metaData["X-Access-Token"])
	// todo
	return bearToken, nil
}
