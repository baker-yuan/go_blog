package jwt

import (
	"context"
	"crypto/rsa"

	"github.com/golang-jwt/jwt"
)

// TokenValidatorConfig 用于验证JWT的配置，业务直接使用当前结构体映射配置文件，然后调用Build方法构成jwt工具类
type TokenValidatorConfig struct {
	PublicKey string `json:"public_key" mapstructure:"public_key" yaml:"public_key"` // RSA公钥，base64编码
}

// check 配置检查
func (cfg *TokenValidatorConfig) check(ctx context.Context) error {
	if len(cfg.PublicKey) == 0 {
		return ErrEmptyPublicKey
	}
	return nil
}

// Build 创建一个新的TokenUtil实例
func (cfg *TokenValidatorConfig) Build(ctx context.Context) (*TokenValidator, error) {
	if err := cfg.check(ctx); err != nil {
		return nil, err
	}
	return NewTokenValidator(ctx, cfg)
}

// TokenValidator 用于验证JWT
type TokenValidator struct {
	publicKey *rsa.PublicKey // RSA公钥
}

// NewTokenValidator 创建一个新的TokenValidator实例
func NewTokenValidator(ctx context.Context, cfg *TokenValidatorConfig) (*TokenValidator, error) {
	publicKey, err := parseRSAPublicKeyFromBase64(cfg.PublicKey)
	if err != nil {
		return nil, ErrParsePublicKey
	}

	return &TokenValidator{
		publicKey: publicKey,
	}, nil
}

// ValidateAndExtractUserInfo 验证JWT并提取用户信息
func (t *TokenValidator) ValidateAndExtractUserInfo(tokenString string) (*UserInfo, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return t.publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrTokenClaims
	}

	return extractUserInfo(claims)
}
