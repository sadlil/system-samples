package redis

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"sadlil.com/samples/golib/cache"
)

type redisStoreValue struct {
	ID string
}

func TestTTLStore(t *testing.T) {
	ctx := context.TODO()

	r := miniredis.RunT(t)
	cmd := redis.NewClient(&redis.Options{
		Addr:       r.Addr(),
		PoolSize:   5,
		MaxRetries: 5,
		DB:         0,
	})

	s := NewCacheStore(cmd, StoreConfig{
		Namespace: "foo",
	})

	s.Set(ctx, "foo-key", &redisStoreValue{ID: "foo"}, &cache.Option{Expiry: time.Second})

	v := &redisStoreValue{}
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
	v = &redisStoreValue{}
	// Cache should be Clear now.
	if err := s.Get(ctx, "foo-key", v); err != nil {
		if !errors.Is(err, cache.ErrCacheMiss) {
			t.Errorf("s.Get(foo-key): got %v, want cache.ErrCacheMiss", err)
		}
	}

	s.Set(ctx, "foo-key", &redisStoreValue{ID: "foo"}, &cache.Option{Expiry: time.Hour})
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

func TestNamespacedKey(t *testing.T) {
	cfg := StoreConfig{Namespace: "ns"}
	if cfg.NamespacedKey("key") != "ns:key" {
		t.Errorf("NamespacedKey: got %v, expected ns:key", cfg.NamespacedKey("key"))
	}
}

func TestDeNamespacedKey(t *testing.T) {
	cfg := StoreConfig{Namespace: "ns"}
	if cfg.DeNamespacedKey("ns:key") != "key" {
		t.Errorf("NamespacedKey: got %v, expected ns:key", cfg.NamespacedKey("key"))
	}
}
