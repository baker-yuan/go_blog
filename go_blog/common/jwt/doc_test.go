package jwt

import (
	"context"
	"fmt"
	"testing"
)

// TestJWT 测试JWT生成和验证逻辑
func TestJWT(t *testing.T) {
	ctx := context.Background()
	base64PrivateKey, base64PublicKey := GenerateBase64EncodedKeys()
	fmt.Println("Private Key Base64:", base64PrivateKey)
	fmt.Println("Public Key Base64:", base64PublicKey)
	// 一、user服务逻辑

	// 创建TokenGenerator配置
	genCfg := &TokenGeneratorConfig{
		PrivateKey:       base64PrivateKey,
		PublicKey:        base64PublicKey,
		AccessExpirySec:  3600,
		RefreshExpirySec: 3600 * 24 * 7,
	}

	// 创建TokenGenerator实例
	tokenGenerator, err := genCfg.Build(ctx)
	if err != nil {
		t.Fatalf("Failed to create token generator: %v", err)
	}

	// 创建用户信息
	userInfo := &UserInfo{
		UserID: "123456",
	}

	// 生成JWT
	tokenDetails, err := tokenGenerator.CreateToken(ctx, userInfo)
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}

	// 二、网关逻辑

	// 创建TokenValidator配置
	valCfg := &TokenValidatorConfig{
		PublicKey: base64PublicKey,
	}

	// 创建TokenValidator实例
	tokenValidator, err := valCfg.Build(ctx)
	if err != nil {
		t.Fatalf("Failed to create token validator: %v", err)
	}

	// 验证JWT
	extractedUserInfo, err := tokenValidator.ValidateAndExtractUserInfo(tokenDetails.AccessToken)
	if err != nil {
		t.Fatalf("Failed to verify token: %v", err)
	}

	// 检查提取的用户信息是否正确
	if extractedUserInfo.UserID != userInfo.UserID {
		t.Errorf("Extracted UserInfo does not match the original. Got %+v, want %+v", extractedUserInfo, userInfo)
	}
}
