package memory

import (
	"context"
	"fmt"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/jinzhu/copier"
	"github.com/sadlil/system-samples/golib/cache"
	"golang.org/x/sync/singleflight"
)

// IntervalStoreConfig captures configs for NewLRUStore
type IntervalStoreConfig struct {
	// Capacity sets the maximum cache items allowed in the cache store, least-recent
	// used items will get removed when capacity is reached.
	Capacity int

	// Cleaner sets a regular ticker that the cache store will be cleared at each
	// tick of the cleaner.

	CleanupInterval time.Duration
}

// NewIntervalStore return a namespaced cache.Store that is backed in memory
// and refreshes its states at a regular interval.
func NewIntervalStore(cfg IntervalStoreConfig) *intervalStore {
	if cfg.Capacity <= 0 {
		cfg.Capacity = 1000
	}

	c, _ := lru.New[string, any](cfg.Capacity)
	lru.New[string, IntervalStoreConfig](1)
	store := &intervalStore{
		cache:        c,
		singleFlight: &singleflight.Group{},
		ticker:       time.NewTicker(cfg.CleanupInterval),
	}
	go store.cleanUp()
	return store
}

var _ cache.Store = new(intervalStore)

// lruStore is an implementation of the cache.Store interface
type intervalStore struct {
	cache        *lru.Cache[string, any]
	singleFlight *singleflight.Group
	ticker       *time.Ticker
}

func (s *intervalStore) Get(ctx context.Context, key string, obj any) error {
	src, ok := s.cache.Get(key)
	if !ok {
		return cache.ErrCacheMiss
	}
	return copier.Copy(obj, src)
}

func (s *intervalStore) Set(ctx context.Context, key string, obj any, _ *cache.Option) error {
	s.cache.Add(key, obj)
	return nil
}

func (s *intervalStore) Fetch(ctx context.Context, key string, o any, opt *cache.Option) error {
	d, ok := s.cache.Get(key)
	if ok {
		return copier.Copy(o, d)
	}
	return s.refreshEntry(ctx, key, o, opt)
}

func (s *intervalStore) Delete(ctx context.Context, key string) error {
	s.cache.Remove(key)
	return nil
}

func (s *intervalStore) DeleteAll(ctx context.Context) error {
	s.cache.Purge()
	return nil
}

func (s *intervalStore) refreshEntry(ctx context.Context, key string, o any, opt *cache.Option) error {
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
	return copier.Copy(o, d)
}

func (s *intervalStore) cleanUp() {
	for range s.ticker.C {
		s.cache.Purge()
	}
}
