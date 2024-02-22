package web

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Id int64
	UserAgent string
	jwt.RegisteredClaims
}

// JWTKey 因为 JWT Key 不太可能变，所以可以直接写成常量
// 也可以考虑做成依赖注入
var JWTKey = []byte("x12vN8JOO1C74P4DE4seqqgEeGPmWvuJ")
