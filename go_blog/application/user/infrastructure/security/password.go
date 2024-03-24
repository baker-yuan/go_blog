package security

import "golang.org/x/crypto/bcrypt"

// Hash 密码hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword 校验密码
//
// hashedPassword hash后的密码
// password 用户输入的密码
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
