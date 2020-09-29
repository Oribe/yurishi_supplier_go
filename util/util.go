package util

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// CreateToken 生成uuid(token)
func CreateToken() string {
	uid := uuid.NewV4().String()
	uid = strings.ReplaceAll(uid, "-", "")
	return uid
}
