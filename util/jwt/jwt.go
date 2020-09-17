package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const signingKey = "{!#(&#)~"

// Claism ...
type Claism struct {
	IP string `json:"ip"`
	jwt.StandardClaims
}

// ValidIP 验证
func (c *Claism) ValidIP(ip string) bool {
	return c.IP == ip
}

// ValidAud 验证一下
func (c *Claism) ValidAud(username interface{}) bool {
	u := username.(string)
	return c.Audience == u
}

// CreateToken 生成Token
func CreateToken(username, ip string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ip":        ip,                // ip地址
		"Audience":  username,          // 受众
		"Id":        id,                // 编号
		"IssuedAt":  time.Now().Unix(), // 签发时间
		"Issuer":    "baiqiao",         // 签发人
		"NotBefore": time.Now().Unix(), // 生效时间
		"Subject":   "login",           // 主题
		// ExpiresAt: time.Now().Add(time.Minute * 30), // 过期时间
	})

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyToken 验证token
func VerifyToken(tokenString, ip string, username interface{}) bool {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claism{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		})

	if claism, ok := token.Claims.(*Claism); ok && token.Valid && claism.ValidIP(ip) && claism.ValidAud(username) {
		// token验证成功
		fmt.Printf(
			"aud:%v ip:%v valid:%v\n",
			claism.Audience,
			claism.IP,
			token.Valid,
		)
		return true
	}
	fmt.Println("token验证失败:", err)
	return false
}
