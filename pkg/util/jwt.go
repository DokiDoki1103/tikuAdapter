package util

import (
	"crypto/rand"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"strconv"
)

var jwtKey []byte

func init() {
	jwtKey = make([]byte, 32) // 生成32字节（256位）的密钥
	if _, err := rand.Read(jwtKey); err != nil {
		panic(err)
	}
}

// GenerateJwt 生成一个jwt
func GenerateJwt(id int32) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(id)),
	})
	jwtStr, err := token.SignedString(jwtKey)
	if err != nil {
		logger.SysLog(err.Error())
		return ""
	}
	return jwtStr
}

// ParseJwtWithClaims 解析 jwt 字符串，返回 Claims 对象
func ParseJwtWithClaims(jwtStr string, options ...jwt.ParserOption) (jwt.Claims, error) {
	mc := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtStr, mc, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	}, options...)
	if err != nil {
		return nil, err
	}
	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）。
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims, nil
}
