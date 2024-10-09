package repo

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	c         *cache.Cache
	cacheOnce sync.Once
)

// Initialize initializes the cache
func InitCache() *cache.Cache {
	cacheOnce.Do(func() {
		c = cache.New(5*time.Minute, 10*time.Minute)
	})
	return c
}

// Get retrieves a value from the cache
func Get(key string) (interface{}, bool) {
	return c.Get(key)
}

// Set sets a value in the cache
func Set(key string, value interface{}, expiration time.Duration) {
	c.Set(key, value, expiration)
}
