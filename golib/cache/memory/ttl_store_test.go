package memory

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/sadlil/system-samples/golib/cache"
)

type ttlStoreValue struct {
	ID string
}

func TestTTLStore(t *testing.T) {
	ctx := context.TODO()
	s := NewTTLStore(TTLStoreConfig{
		DefaultExpiration: 0,
		CleanupInterval:   0,
	})

	s.Set(ctx, "foo-key", &ttlStoreValue{ID: "foo"}, &cache.Option{Expiry: time.Second})

	v := &ttlStoreValue{}
	if err := s.Get(ctx, "foo-key", v); err != nil {
		t.Errorf("s.Get(foo-key): got %v, want nil", err)
	}
	if v.ID != "foo" {
		t.Errorf("s.Get(v.ID): got %v, want %v", v.ID, "foo")
	}

	// Fetch
	if err := s.Fetch(ctx, "foo-key", v, &cache.Option{Source: func(ctx context.Context) (interface{}, error) { return nil, fmt.Errorf("error") }}); err != nil {
		t.Errorf("s.Get(foo-key): got %v, want nil", err)
	}
	if v.ID != "foo" {
		t.Errorf("s.Get(v.ID): got %v, want %v", v.ID, "foo")
	}

	time.Sleep(time.Second)
	v = &ttlStoreValue{}
	// Cache should be Clear now.
	if err := s.Get(ctx, "foo-key", v); err != nil {
		if !errors.Is(err, cache.ErrCacheMiss) {
			t.Errorf("s.Get(foo-key): got %v, want cache.ErrCacheMiss", err)
		}
	}

	s.Set(ctx, "foo-key", &ttlStoreValue{ID: "foo"}, &cache.Option{Expiry: time.Hour})
	if err := s.Get(ctx, "foo-key", v); err != nil {
		t.Errorf("s.Get(foo-key): got %v, want nil", err)
	}
	s.Delete(ctx, "foo-key")
	if err := s.Get(ctx, "foo-key", v); err != nil {
		if !errors.Is(err, cache.ErrCacheMiss) {
			t.Errorf("s.Get(foo-key): got %v, want cache.ErrCacheMiss", err)
		}
	}
}
