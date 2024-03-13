package session

import "context"

// SessionUtils 会话
type SessionUtils struct{}

// GetLoginUserID 获取登录态中的用户ID
func (c SessionUtils) GetLoginUserID(ctx context.Context) (uint32, error) {
	// todo
	return 0, nil
}
