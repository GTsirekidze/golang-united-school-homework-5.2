package cache

import (
	"time"
)

type Cache struct {
	mp map[string]localCache
}

type localCache struct {
	value       string
	deadline    time.Time
	isPermanent bool
}

func NewCache() Cache {
	return Cache{make(map[string]localCache)}
}

func (caches *Cache) cleenUp(key string) {
	if !caches.mp[key].isPermanent || caches.mp[key].deadline.Before(time.Now()) {
		delete(caches.mp, key)
	}
}

func (caches *Cache) Get(key string) (string, bool) {
	caches.cleenUp(key)
	val, ok := caches.mp[key]
	return val.value, ok
}

func (caches *Cache) Put(key, value string) {
	caches.mp[key] = localCache{value, time.Now(), true}
}

func (caches Cache) Keys() []string {
	var ans []string
	for k, _ := range caches.mp {
		caches.cleenUp(k)
		ans = append(ans, k)
	}

	return ans
}

func (caches *Cache) PutTill(key, value string, deadline time.Time) {
	caches.mp[key] = localCache{value, deadline, false}
}
