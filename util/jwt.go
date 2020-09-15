package util

import (
	"github.com/dgrijalva/jwt-go"
)

// const

// CreateToken 生成Token
func CreateToken(username, domain, ip string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"domain":   domain,
		"ip":       ip,
	})

	tokenString, err := token.SignedString("")
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
