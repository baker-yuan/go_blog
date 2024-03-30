package security

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// TestGeneratePassword 测试密码生成函数
func TestGeneratePassword(t *testing.T) {
	password := "123456"
	hashedPassword, err := GeneratePassword(password)
	if err != nil {
		t.Errorf("Failed to generate password: %v", err)
	}

	err = VerifyPassword(hashedPassword, password)
	if err != nil {
		t.Errorf("Failed to verify password: %v", err)
	}
}

// TestVerifyPassword 测试密码验证函数
func TestVerifyPassword(t *testing.T) {
	password := "123456"
	wrongPassword := "wrongpassword"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("bcrypt.GenerateFromPassword failed: %v", err)
	}

	fmt.Printf("hashedPassword: %s", hashedPassword)

	// 测试正确的密码
	err = VerifyPassword(string(hashedPassword), password)
	if err != nil {
		t.Errorf("VerifyPassword failed with correct password: %v", err)
	}

	// 测试错误的密码
	err = VerifyPassword(string(hashedPassword), wrongPassword)
	if err == nil {
		t.Errorf("VerifyPassword should have failed with wrong password")
	}
}
