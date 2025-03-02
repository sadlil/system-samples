package memory

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sadlil/system-samples/golib/cache"
	"golang.org/x/sync/singleflight"
)

// TTLStoreConfig captures configs for ttlStore
type TTLStoreConfig struct {
	DefaultExpiration, CleanupInterval time.Duration
}

// NewTTLStore return a namespaced cache.Store that is backed in memory
func NewTTLStore(cfg TTLStoreConfig) *ttlStore {
	return &ttlStore{
		cfg:       cfg,
		refresher: &singleflight.Group{},
		cache:     gocache.New(cfg.DefaultExpiration, cfg.CleanupInterval),
	}
}

var _ cache.Store = new(ttlStore)

// ttlStore is an implementation of the cache.Store interface
type ttlStore struct {
	cfg       TTLStoreConfig
	refresher *singleflight.Group
	cache     *gocache.Cache
}

func (s *ttlStore) Get(ctx context.Context, key string, obj any) error {
	src, ok := s.cache.Get(key)
	if !ok {
		return cache.ErrCacheMiss
	}
	return copier.Copy(obj, src)
}

func (s *ttlStore) Set(ctx context.Context, key string, obj any, opt *cache.Option) error {
	if opt == nil {
		opt = &cache.Option{Expiry: s.cfg.DefaultExpiration}
	}
	s.cache.Add(key, obj, opt.Expiry)
	return nil
}

func (s *ttlStore) Fetch(ctx context.Context, key string, o any, opt *cache.Option) error {
	d, ok := s.cache.Get(key)
	if ok {
		return copier.Copy(o, d)
	}
	return s.refreshEntry(ctx, key, o, opt)
}

func (s *ttlStore) Delete(ctx context.Context, key string) error {
	s.cache.Delete(key)
	return nil
}

func (s *ttlStore) DeleteAll(ctx context.Context) error {
	s.cache.Flush()
	return nil
}

func (s *ttlStore) refreshEntry(ctx context.Context, key string, o any, opt *cache.Option) error {
	d, err, _ := s.refresher.Do(key, func() (interface{}, error) {
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
	return copier.Copy(o, d)
}
