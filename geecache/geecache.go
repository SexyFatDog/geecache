package geecache

import (
	"fmt"
	"geecache/geecache/model"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// A Group is a cache namespace and associated data loaded spread over
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

func (g *Group) Get(key string) (model.ByteView, error) {
	if key == "" {
		return model.ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}

	return g.load(key)
}

func (g *Group) load(key string) (value model.ByteView, err error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (model.ByteView, error) {
	bytes, err := g.getter.Get((key))
	if err != nil {
		return model.ByteView{}, err
	}

	value := model.ByteView{B: model.CloneBytes(bytes)}
	g.popularCache(key, value)
	return value, nil
}

func (g *Group) popularCache(key string, value model.ByteView) {
	g.mainCache.add(key, value)
}
