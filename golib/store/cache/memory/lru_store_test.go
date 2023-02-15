package memory

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"sadlil.com/samples/golib/store/cache"
)

type lruStoreValue struct {
	ID string
}

func TestLRUStore(t *testing.T) {
	ctx := context.TODO()
	s := NewLRUStore(LRUStoreConfig{
		Capacity: 2,
	})

	s.Set(ctx, "foo-key-1", &lruStoreValue{ID: "foo-1"}, nil)
	v := &lruStoreValue{}
	if err := s.Get(ctx, "foo-key-1", v); err != nil {
		t.Errorf("s.Get(foo-key-1): got %v, want nil", err)
	}
	if v.ID != "foo-1" {
		t.Errorf("s.Get(v.ID): got %v, want %v", v.ID, "foo")
	}

	// Fetch
	if err := s.Fetch(ctx, "foo-key-1", v, &cache.Option{Source: func(ctx context.Context) (interface{}, error) { return nil, fmt.Errorf("error") }}); err != nil {
		t.Errorf("s.Get(foo-key): got %v, want nil", err)
	}

	// Push aditional data, foo-key-1 should be evicted.
	s.Set(ctx, "foo-key-2", &lruStoreValue{ID: "foo-2"}, nil)
	s.Set(ctx, "foo-key-3", &lruStoreValue{ID: "foo-2"}, nil)

	v = &lruStoreValue{}
	// Cache should be Clear now.
	if err := s.Get(ctx, "foo-key-1", v); err != nil {
		if !errors.Is(err, cache.ErrCacheMiss) {
			t.Errorf("s.Get(foo-key): got %v, want cache.ErrCacheMiss", err)
		}
	}

	s.Delete(ctx, "foo-key-2")
	if err := s.Get(ctx, "foo-key-2", v); err != nil {
		if !errors.Is(err, cache.ErrCacheMiss) {
			t.Errorf("s.Get(foo-key): got %v, want cache.ErrCacheMiss", err)
		}
	}
}
