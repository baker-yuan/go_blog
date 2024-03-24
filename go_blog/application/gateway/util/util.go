package util

import "github.com/baker-yuan/go-blog/common/jwt"

var (
	JwtUtil *jwt.TokenValidator
)

func Init(tokenValidator *jwt.TokenValidator) {
	JwtUtil = tokenValidator
}
