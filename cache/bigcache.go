package cache

import (
	"time"

	"github.com/allegro/bigcache"
)

type BigCacheService struct {
	cache *bigcache.BigCache
}

func NewBigCacheService() *BigCacheService {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Second))
	return &BigCacheService{
		cache: cache,
	}
}

func (b *BigCacheService) Set(key string, value []byte) error {
	return b.cache.Set(key, value)
}

func (b *BigCacheService) Get(key string) ([]byte, error) {
	return b.cache.Get(key)
}
