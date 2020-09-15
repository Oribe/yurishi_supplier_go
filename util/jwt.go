package util

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const signingKey = "{!#(&#)~"

// Claism ...
type Claism struct {
	Username string `json:"username"`
	Domain   string `json:"domain"`
	IP       string `json:"ip"`

	un string
	dm string
	ip string
}

// Valid 验证
func (c *Claism) Valid() error {
	if c.Username != c.un || c.Domain != c.dm || c.IP != c.ip {
		return errors.New("无效的Token")
	}
	return nil
}

// CreateToken 生成Token
func CreateToken(username, domain, ip string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"domain":   domain,
		"ip":       ip,
	})

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyToken 验证token
func VerifyToken(tokenString, username, domain, ip string) bool {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claism{
			un: username,
			dm: domain,
			ip: ip,
		},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		})

	if claism, ok := token.Claims.(*Claism); ok && token.Valid {
		fmt.Printf("%v\n %v\n %v\n %v\n", claism.Username, claism.Domain, claism.IP, token.Valid)
		return true
	}
	fmt.Println("token验证失败:", err)
	return false
}
