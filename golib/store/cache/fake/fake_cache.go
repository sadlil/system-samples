package fake

import (
	"context"

	"sadlil.com/samples/golib/store/cache"
	"sadlil.com/samples/golib/store/cache/internal/copier"
)

// FakeStore is an implementation of the cache.Store interface, used in testing.
type FakeStore struct {
	ReadKeys   []string
	WriteKeys  []string
	FetchKeys  []string
	DeleteKeys []string
	Cache      map[string]any
}

// Compile time check of cache.Store compatibility
var _ cache.Store = new(FakeStore)

func NewCacheStore() *FakeStore {
	return &FakeStore{Cache: map[string]any{}}
}

// Get implements cache.Store
func (s *FakeStore) Get(ctx context.Context, key string, obj any) error {
	s.ReadKeys = append(s.ReadKeys, key)
	if s.Cache != nil {
		if v, ok := s.Cache[key]; ok {
			return copier.Copy(v, obj)
		}
	}
	return cache.ErrCacheMiss
}

// Set implements cache.Store
func (s *FakeStore) Set(ctx context.Context, key string, obj any, opt *cache.Option) error {
	s.WriteKeys = append(s.WriteKeys, key)
	if s.Cache != nil {
		if obj == nil {
			delete(s.Cache, key)
		} else {
			s.Cache[key] = obj
		}
	}
	return nil
}

// Fetch implements cache.Store
func (s *FakeStore) Fetch(ctx context.Context, key string, obj any, opt *cache.Option) error {
	s.FetchKeys = append(s.FetchKeys, key)

	if s.Cache != nil {
		dbobj, ok := s.Cache[key]
		if ok {
			return copier.Copy(dbobj, obj)
		}
	}

	dbobj, err := opt.Source(ctx)
	if err != nil {
		return err
	}
	_ = s.Set(ctx, key, obj, opt)
	return copier.Copy(dbobj, obj)
}

// Delete expires the mentioned key.
func (s *FakeStore) Delete(ctx context.Context, key string) error {
	s.DeleteKeys = append(s.DeleteKeys, key)

	delete(s.Cache, key)
	return nil
}

// Reset clears all keys
func (s *FakeStore) DeleteAll(ctx context.Context) error {
	s.ReadKeys = []string{}
	s.WriteKeys = []string{}
	s.FetchKeys = []string{}
	s.DeleteKeys = []string{}
	s.Cache = make(map[string]interface{})
	return nil
}
