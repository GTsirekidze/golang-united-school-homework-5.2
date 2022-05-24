package cache

import "time"

type Cache struct {
	mp map[string]string
}

func NewCache() Cache {
	return Cache{}
}

func (caches Cache) Get(key string) (string, bool) {
	val, ok := caches.mp[key]
	return val, ok
}

func (caches *Cache) Put(key, value string) {
	caches.mp[key] = value
}

func (caches Cache) Keys() []string {
	var ans []string
	for k, _ := range caches.mp {
		ans = append(ans, k)
	}

	return ans
}

func (caches *Cache) PutTill(key, value string, deadline time.Time) {
	caches.mp[key] = value

	start := time.Now().UnixNano() / int64(time.Millisecond)
	end := deadline.UnixNano() / int64(time.Millisecond)
	time.Sleep(time.Duration(end - start))

	delete(caches.mp, key)
}
