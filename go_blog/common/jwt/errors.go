package jwt

import "errors"

// 配置检查相关异常
var (
	ErrEmptyPrivateKey      = errors.New("私钥不能为空")
	ErrEmptyPublicKey       = errors.New("公钥不能为空")
	ErrInvalidAccessExpiry  = errors.New("访问令牌有效期必须为非负数")
	ErrInvalidRefreshExpiry = errors.New("刷新令牌有效期必须为非负数")
)

// base64加密的公钥/私钥解密异常
var (
	ErrParsePrivateKey = errors.New("解析私钥失败")
	ErrParsePublicKey  = errors.New("解析公钥失败")
)

// 其他异常
var (
	ErrUnexpectedSigningMethod = errors.New("非预期的签名方法")
	ErrInvalidToken            = errors.New("无效的令牌")
	ErrTokenClaims             = errors.New("无法获取令牌声明")
	ErrUserInfoField           = errors.New("userinfo字段缺失或类型不正确")
	ErrUserIDField             = errors.New("userid字段缺失或类型不正确")
	ErrKeyParsing              = errors.New("解析密钥失败")
	ErrKeyTypeNotRSA           = errors.New("密钥类型不是RSA")
)
