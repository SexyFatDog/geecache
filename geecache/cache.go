package geecache

import (
	"geecache/geecache/lru"
	"geecache/geecache/model"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value model.ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}

	c.lru.Add(key, &value)
}

func (c *cache) get(key string) (value model.ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		if bv, ok := v.(*model.ByteView); ok {
			return *bv, true
		}
	}

	return
}
