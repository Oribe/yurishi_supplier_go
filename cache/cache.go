package cache

import (
	"reflect"
	"time"

	goCache "github.com/patrickmn/go-cache"
)

const (
	// DefaultExpiration 永不过期
	DefaultExpiration = goCache.DefaultExpiration
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

	v, ok := cache.Get(key)

	if !ok {
		return false
	}

	t := reflect.TypeOf(value).Kind()
	if t == reflect.Struct || t == reflect.Map || t == reflect.Slice {
		return reflect.DeepEqual(value, v)
	}
	return v == value
}

// Get ...
func Get(key string) (interface{}, bool) {
	return cache.Get(key)
}

// Delete ...
func Delete(key string) bool {

	cache.Delete(key)

	_, ok := cache.Get(key)
	if !ok {
		return true
	}
	return false
}
