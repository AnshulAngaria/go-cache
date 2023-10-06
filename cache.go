package gocache

import "sync"

type Cache struct {
	sync.RWMutex
	store map[string]string
}

func (cache *Cache) Get(key string) string {
	cache.RLock()
	defer cache.RUnlock()

	val, ok := cache.store[key]

	if !ok {
		return ""
	}
	return val
}

func (cache *Cache) Set(key, val string) {
	cache.Lock()
	defer cache.Unlock()

	cache.store[key] = val
}
