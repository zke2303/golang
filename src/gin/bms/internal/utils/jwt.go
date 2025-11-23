package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Username string `form:"username"`
	jwt.RegisteredClaims
}

var CustomSecret = []byte("Tentcoo@123")

func GeneratJwt(username string) (string, error) {
	// 1.构建 Claims 对象
	claims := CustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			Issuer:    "Tentcoo",
		},
	}

	// 2.使用 jwt 工具类中的方法生成 jwt.Token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 3. 返回编码后的字符串
	return token.SignedString(CustomSecret)
}

func ParseJwt(tokenString string) (*CustomClaims, error) {
	// 1.使用 jwt 工具类中的方法解析 jwt
	token, err := jwt.ParseWithClaims(
		tokenString, &CustomClaims{}, func(t *jwt.Token) (any, error) {
			return CustomSecret, nil
		},
	)

	// 2.判断是否发生 error
	if err != nil {
		return nil, err
	}
	// 3.判断解析的token中的Claims对象是否是 CustomClaims 类型
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	// 4.token不可用， 或者类型不对
	return nil, errors.New("invalid token")
}
