package jwt

import "fmt"

var (
	userInfoKey   = "userinfo" // 用户信息键，用于JWT的claims
	expirationKey = "exp"      // 过期时间键，用于JWT的claims
	userid        = "userid"   // JWT的claims，存的用户id UserInfo#UserID
)

// 令牌有消息
const (
	// 默认的访问令牌有效期（秒）
	DefaultAccessExpirySec int64 = 3600 // 1小时
	// 默认的刷新令牌有效期（秒）
	DefaultRefreshExpirySec int64 = 2592000 // 30天
)

// UserInfo 存储在JWT中的用户信息
type UserInfo struct {
	UserID string `json:"userid"` // 用户ID
}

// UserInfoBuilder 是用于构建UserInfo的构建器
type UserInfoBuilder struct {
	userInfo UserInfo
}

// NewUserInfoBuilder 创建一个新的UserInfoBuilder实例
func NewUserInfoBuilder() *UserInfoBuilder {
	return &UserInfoBuilder{}
}

// UserID 设置UserInfo的UserID字段
func (b *UserInfoBuilder) UserID(userID uint32) *UserInfoBuilder {
	b.userInfo.UserID = fmt.Sprintf("%d", userID)
	return b
}

// Build 构建并返回最终的UserInfo对象
func (b *UserInfoBuilder) Build() *UserInfo {
	return &b.userInfo
}

// TokenDetails 访问令牌和刷新令牌的详细信息
type TokenDetails struct {
	AccessToken  string // 访问令牌
	RefreshToken string // 刷新令牌
	AtExpires    uint32 // 访问令牌过期时间
	RtExpires    uint32 // 刷新令牌过期时间
}
