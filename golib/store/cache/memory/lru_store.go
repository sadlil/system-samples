package memory

import (
	"context"
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"
	"golang.org/x/sync/singleflight"
	"sadlil.com/samples/golib/store/cache"
	"sadlil.com/samples/golib/store/cache/internal/copier"
)

// LRUStoreConfig captures configs for NewLRUStore
type LRUStoreConfig struct {
	// Capacity sets the maximum cache items allowed in the cache store, least-recent
	// used items will get removed when capacity is reached.
	Capacity int
}

// NewLRUStore return a namespaced cache.Store that is backed in memory.
func NewLRUStore(cfg LRUStoreConfig) *lruStore {
	if cfg.Capacity <= 0 {
		cfg.Capacity = 1000
	}
	c, _ := lru.New[string, any](cfg.Capacity)
	store := &lruStore{
		cache:        c,
		singleFlight: &singleflight.Group{},
	}
	return store
}

var _ cache.Store = new(lruStore)

// lruStore is an implementation of the cache.Store interface
type lruStore struct {
	cache        *lru.Cache[string, any]
	singleFlight *singleflight.Group
}

func (s *lruStore) Get(ctx context.Context, key string, obj any) error {
	src, ok := s.cache.Get(key)
	if !ok {
		return cache.ErrCacheMiss
	}
	return copier.Copy(src, obj)
}

func (s *lruStore) Set(ctx context.Context, key string, obj any, _ *cache.Option) error {
	s.cache.Add(key, obj)
	return nil
}

func (s *lruStore) Fetch(ctx context.Context, key string, o any, opt *cache.Option) error {
	d, ok := s.cache.Get(key)
	if ok {
		return copier.Copy(d, o)
	}
	return s.refreshEntry(ctx, key, o, opt)
}

func (s *lruStore) Delete(ctx context.Context, key string) error {
	s.cache.Remove(key)
	return nil
}

func (s *lruStore) DeleteAll(ctx context.Context) error {
	s.cache.Purge()
	return nil
}

func (s *lruStore) refreshEntry(ctx context.Context, key string, o any, opt *cache.Option) error {
	d, err, _ := s.singleFlight.Do(key, func() (interface{}, error) {
		if opt.Source != nil {
			obj, err := opt.Source(ctx)
			if err != nil {
				return nil, err
			}
			_ = s.Set(ctx, key, obj, opt)
			return obj, err
		}
		return nil, fmt.Errorf("opt.Source is nil")
	})
	if err != nil {
		return err
	}
	return copier.Copy(d, o)
}
