package gocache

import (
	"sync"
	"testing"
)

func TestCache_SetAndGet(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := "testValue"

	cache.Set(key, value)

	retrievedValue, exists := cache.Get(key)

	if !exists {
		t.Errorf("Expected value to exist in cache, but it does not.")
	}

	if retrievedValue != value {
		t.Errorf("Expected value '%s', but got '%s'", value, retrievedValue)
	}
}

func TestCache_GetNonExistentKey(t *testing.T) {
	cache := NewCache()
	key := "nonExistentKey"

	_, exists := cache.Get(key)

	if exists {
		t.Errorf("Expected key to not exist in cache, but it does.")
	}
}

func TestCache_ConcurrentSetAndGet(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := "testValue"
	numGoroutines := 10

	// Concurrently set the value in multiple goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			cache.Set(key, value)
			wg.Done()
		}()
	}

	wg.Wait()

	// Retrieve the value
	retrievedValue, exists := cache.Get(key)

	if !exists {
		t.Errorf("Expected value to exist in cache, but it does not.")
	}

	if retrievedValue != value {
		t.Errorf("Expected value '%s', but got '%s'", value, retrievedValue)
	}
}
