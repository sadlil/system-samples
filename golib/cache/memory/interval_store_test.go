package memory

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/sadlil/system-samples/golib/cache"
)

type intervalStoreValue struct {
	ID string
}

func TestIntervalStore(t *testing.T) {
	ctx := context.TODO()
	s := NewIntervalStore(IntervalStoreConfig{
		CleanupInterval: time.Second * 2,
	})

	s.Set(ctx, "foo-key", &intervalStoreValue{ID: "foo"}, nil)

	v := &intervalStoreValue{}
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

	<-s.ticker.C
	v = &intervalStoreValue{}
	// Cache should be Clear now.
	if err := s.Get(ctx, "foo-key", v); err != nil {
		if !errors.Is(err, cache.ErrCacheMiss) {
			t.Errorf("s.Get(foo-key): got %v, want cache.ErrCacheMiss", err)
		}
	}

	s.Set(ctx, "foo-key", &intervalStoreValue{ID: "foo"}, nil)
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
