package common

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	BaseClaim
	jwt.RegisteredClaims
}

// BaseClaim 用于保存token中的信息
type BaseClaim struct {
	ID       int32
	UserName string
}
