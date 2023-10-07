package gocache

import "sync"

type Cache struct {
	sync.RWMutex
	store map[string]*Value
}

type Value struct {
	sync.Mutex
	value string
}

func (cache *Cache) Get(key string) (string, bool) {
	cache.RLock()

	val, ok := cache.store[key]
	cache.RUnlock()

	if !ok {
		return "", false
	}
	return val.value, true
}

func (cache *Cache) Set(key, val string) {
	cache.Lock()
	defer cache.Unlock()

	cache.store[key] = &Value{value: val}
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]*Value),
	}
}
