package cache

import "time"

type Cache struct {
	value string
}

type Caches struct {
	mp map[string]Cache
}

func NewCache(value string) Cache {
	return Cache{value}
}

func (caches Caches) Get(key string) (string, bool) {
	return caches.mp[key].value, true
}

func (caches *Caches) Put(key, value string) {
	caches.mp[key] = Cache{value}
}

func (caches Caches) Keys() []string {
	var ans []string
	for k, _ := range caches.mp {
		ans = append(ans, k)
	}

	return ans
}

func (caches *Caches) PutTill(key, value string, deadline time.Time) {
	caches.mp[key] = Cache{value}

	start := time.Now().UnixNano() / int64(time.Millisecond)
	end := deadline.UnixNano() / int64(time.Millisecond)
	time.Sleep(time.Duration(end - start))

	delete(caches.mp, key)
}
