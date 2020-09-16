package cache

import (
	"time"

	goCache "github.com/patrickmn/go-cache"
)

// Cache 全局变量
var cache *goCache.Cache

func init() {
	cache = goCache.New(goCache.DefaultExpiration, time.Minute*10)
}

// Set ...
func Set(key string, value interface{}, expiration time.Duration) bool {
	if key == "" {
		return false
	}
	cache.Set(key, value, expiration)
	return true
}
