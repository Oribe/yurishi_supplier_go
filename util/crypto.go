package util

import (
	"crypto/hmac"
	"crypto/sha256"
)

const (
	salt = "!)3&*yA(."
)

// SignPassword 密码加密
func SignPassword(username string, password string) []byte {
	beforeSign := password + "|" + salt + "|" + username
	mac := hmac.New(sha256.New, []byte(""))
	mac.Write([]byte(beforeSign))
	return mac.Sum(nil)
}
