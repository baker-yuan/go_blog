package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"github.com/golang-jwt/jwt"
)

// GenerateBase64EncodedKeys 生成RSA密钥对并返回Base64编码的私钥和公钥。 该函数用于生成用于签名和验证JWT的密钥对。
func GenerateBase64EncodedKeys() (string, string) {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // 建议使用2048位密钥长度
	if err != nil {
		panic(err)
	}

	// 将私钥转换为ASN.1 PKCS#1 DER编码
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// 对私钥进行Base64编码
	privBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privDER,
	}
	privPEM := pem.EncodeToMemory(&privBlock)
	privBase64 := base64.StdEncoding.EncodeToString(privPEM)

	// 将公钥转换为ASN.1 PKIX DER编码
	pubDER, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	// 对公钥进行Base64编码
	pubBlock := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubDER,
	}
	pubPEM := pem.EncodeToMemory(&pubBlock)
	pubBase64 := base64.StdEncoding.EncodeToString(pubPEM)
	return privBase64, pubBase64
}

// parseRSAPrivateKeyFromBase64 解析Base64编码的RSA私钥
func parseRSAPrivateKeyFromBase64(base64Key string) (*rsa.PrivateKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return nil, ErrParsePrivateKey
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, ErrKeyParsing
	}

	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, ErrParsePrivateKey
	}

	return privKey, nil
}

// parseRSAPublicKeyFromBase64 解析Base64编码的RSA公钥
func parseRSAPublicKeyFromBase64(base64Key string) (*rsa.PublicKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return nil, ErrParsePublicKey
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, ErrKeyParsing
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, ErrParsePublicKey
	}

	switch pubKey := pubKey.(type) {
	case *rsa.PublicKey:
		return pubKey, nil
	default:
		return nil, ErrKeyTypeNotRSA
	}
}

// extractUserInfo 从JWT声明中提取用户信息
func extractUserInfo(claims jwt.MapClaims) (*UserInfo, error) {
	userInfoMap, ok := claims[userInfoKey].(map[string]interface{})
	if !ok {
		return nil, ErrUserInfoField
	}

	userID, ok := userInfoMap[userid].(string)
	if !ok {
		return nil, ErrUserIDField
	}
	return &UserInfo{UserID: userID}, nil
}
