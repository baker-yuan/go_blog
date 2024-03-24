package jwt

import (
	"context"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt"
)

// TokenGeneratorConfig 用于生成JWT的配置，业务直接使用当前结构体映射配置文件，然后调用Build方法构成jwt工具类
type TokenGeneratorConfig struct {
	PrivateKey       string `json:"private_key" mapstructure:"private_key" yaml:"private_key"`                      // RSA私钥，base64编码
	PublicKey        string `json:"public_key" mapstructure:"public_key" yaml:"public_key"`                         // RSA公钥，base64编码
	AccessExpirySec  int64  `json:"access_expiry_sec" mapstructure:"access_expiry_sec" yaml:"access_expiry_sec"`    // 访问令牌的有效期（秒）
	RefreshExpirySec int64  `json:"refresh_expiry_sec" mapstructure:"refresh_expiry_sec" yaml:"refresh_expiry_sec"` // 刷新令牌的有效期（秒）
}

// check 配置检查
func (cfg *TokenGeneratorConfig) check(ctx context.Context) error {
	if len(cfg.PrivateKey) == 0 {
		return ErrEmptyPrivateKey
	}
	if len(cfg.PublicKey) == 0 {
		return ErrEmptyPublicKey
	}
	if cfg.AccessExpirySec == 0 {
		cfg.AccessExpirySec = DefaultAccessExpirySec
	}
	if cfg.RefreshExpirySec == 0 {
		cfg.AccessExpirySec = DefaultRefreshExpirySec
	}
	return nil
}

// Build 创建一个新的TokenUtil实例
func (cfg *TokenGeneratorConfig) Build(ctx context.Context) (*TokenGenerator, error) {
	if err := cfg.check(ctx); err != nil {
		return nil, err
	}
	return NewTokenGenerator(ctx, cfg)
}

// TokenGenerator 用于生成JWT
type TokenGenerator struct {
	privateKey              *rsa.PrivateKey // RSA私钥
	publicKey               *rsa.PublicKey  // RSA工钥
	accessTokenDurationSec  int64           // 访问令牌的有效期（秒）
	refreshTokenDurationSec int64           // 刷新令牌的有效期（秒）
}

// NewTokenGenerator 创建一个新的TokenGenerator实例
func NewTokenGenerator(ctx context.Context, cfg *TokenGeneratorConfig) (*TokenGenerator, error) {
	privateKey, err := parseRSAPrivateKeyFromBase64(cfg.PrivateKey)
	if err != nil {
		return nil, ErrParsePrivateKey
	}
	publicKey, err := parseRSAPublicKeyFromBase64(cfg.PublicKey)
	if err != nil {
		return nil, ErrParsePublicKey
	}

	return &TokenGenerator{
		privateKey:              privateKey,
		publicKey:               publicKey,
		accessTokenDurationSec:  cfg.AccessExpirySec,
		refreshTokenDurationSec: cfg.RefreshExpirySec,
	}, nil
}

// CreateToken 生成新的访问令牌和刷新令牌
func (t *TokenGenerator) CreateToken(ctx context.Context, userInfo *UserInfo) (*TokenDetails, error) {
	now := time.Now()
	td := &TokenDetails{
		AtExpires: uint32(now.Add(time.Second * time.Duration(t.accessTokenDurationSec)).Unix()),
		RtExpires: uint32(now.Add(time.Second * time.Duration(t.refreshTokenDurationSec)).Unix()),
	}

	atClaims := jwt.MapClaims{
		userInfoKey:   userInfo,
		expirationKey: td.AtExpires,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	var err error
	td.AccessToken, err = at.SignedString(t.privateKey)
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{
		userInfoKey:   userInfo,
		expirationKey: td.RtExpires,
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodRS256, rtClaims)
	td.RefreshToken, err = rt.SignedString(t.privateKey)
	if err != nil {
		return nil, err
	}

	return td, nil
}

// RefreshToken 使用刷新令牌获取新的访问令牌
func (t *TokenGenerator) RefreshToken(ctx context.Context, refreshTokenString string) (*TokenDetails, error) {
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
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

	userInfo, err := extractUserInfo(claims)
	if err != nil {
		return nil, err
	}

	newAccessToken, err := t.CreateToken(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	return &TokenDetails{
		AccessToken: newAccessToken.AccessToken,
		AtExpires:   newAccessToken.AtExpires,
	}, nil
}
