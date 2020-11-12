package util

import (
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

// CreateToken 生成uuid(token)
func CreateToken() string {
	uid := uuid.NewV4().String()
	uid = strings.ReplaceAll(uid, "-", "")
	return uid
}

// NowTime 当前时间
func NowTime() string {
	return time.Now().Format(timeFormat)
}

// ParseTimestamp 格式化时间戳
func ParseTimestamp(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(timeFormat)
}
