package gocache

import "hash/fnv"

type shardedCache struct {
	totalShards uint32
	caches      []*Cache
}

func hash(input string) uint32 {
	// Create a new FNV-1a hash object
	h := fnv.New32a()

	// Write the bytes of the input string to the hash
	h.Write([]byte(input))

	// Return the resulting hash value as a uint32
	return h.Sum32()
}

func (sharded *shardedCache) shard(input string) *Cache {
	shard := sharded.caches[hash(input)%sharded.totalShards]
	return shard
}

func (sharded *shardedCache) Get(input string) (string, bool) {
	return sharded.shard(input).Get(input)
}

func (sharded *shardedCache) Set(key, value string) {
	sharded.shard(key).Set(key, value)
}

func newShardedCache(length uint32) *shardedCache {
	sc := &shardedCache{
		totalShards: length,
		caches:      make([]*Cache, length),
	}

	for i := uint32(0); i < length; i++ {
		c := &Cache{
			store: make(map[string]*Value),
		}
		sc.caches[i] = c
	}
	return sc
}
