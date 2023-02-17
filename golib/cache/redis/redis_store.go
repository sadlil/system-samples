package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
	"sadlil.com/samples/golib/cache"
)

type StoreConfig struct {
	Namespace string
}

// NewCacheStore return a namespaced cache.Store that is backed in redis.
func NewCacheStore(client redis.Cmdable, cfg StoreConfig) *redisStore {
	store := &redisStore{
		client:       client,
		cfg:          cfg,
		singleFlight: &singleflight.Group{},
	}
	return store
}

var _ cache.Store = new(redisStore)

// redisStore is an implementation of the cache.Store interface
type redisStore struct {
	cfg          StoreConfig
	singleFlight *singleflight.Group
	client       redis.Cmdable
}

func (r *redisStore) Fetch(ctx context.Context, key string, obj any, opt *cache.Option) error {
	entryObj, err := r.readEntry(ctx, key)
	// When data not found or redis err, reach to data source
	if err != nil {
		return r.loadEntry(ctx, key, obj, opt)
	}
	return entryObj.Unmarshal(obj)
}

func (r *redisStore) Get(ctx context.Context, key string, obj any) error {
	entryObj, err := r.readEntry(ctx, key)
	if err != nil {
		return err
	}
	return entryObj.Unmarshal(obj)

}

func (r *redisStore) Set(ctx context.Context, key string, obj any, opt *cache.Option) error {
	entryObj, err := newRedisEntry(obj, opt)
	if err != nil {
		return err
	}

	data, err := json.Marshal(entryObj)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, r.cfg.NamespacedKey(key), data, opt.Expiry).Err()

}

func (r *redisStore) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, r.cfg.NamespacedKey(key)).Err()
}

// DeleteAll is not supported for redis
func (r *redisStore) DeleteAll(ctx context.Context) error {
	return fmt.Errorf("redis.DeleteAll is not suppported: %w", cache.ErrUnImplemented)
}

func (s redisStore) readEntry(ctx context.Context, key string) (*redisEntry, error) {
	data, err := s.client.Get(ctx, s.cfg.NamespacedKey(key)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, cache.ErrCacheMiss
		}
		return nil, err
	}

	entryObj := &redisEntry{}
	return entryObj, json.Unmarshal(data, entryObj)
}

func (r *redisStore) loadEntry(ctx context.Context, key string, target any, opt *cache.Option) error {
	obj, err, _ := r.singleFlight.Do(key, func() (interface{}, error) {
		obj, err := opt.Source(ctx)
		if err != nil {
			return nil, err
		}

		refreshOpt := *opt
		if opt.OnRefresh != nil {
			// OnRefresh should not be happening async. Any changes to obj will be reflected in the cached object.
			refreshOpt = opt.OnRefresh(obj, refreshOpt)
		}
		go func() {
			_ = r.Set(context.Background(), key, obj, &refreshOpt)
		}()
		return obj, err
	})
	if err != nil {
		return err
	}
	return copier.Copy(target, obj)
}

func (c *StoreConfig) NamespacedKey(key string) string {
	if len(c.Namespace) > 0 {
		return fmt.Sprintf("%s:%s", c.Namespace, key)
	}
	return key
}

// DeNamespacedKey returns the key without namespace prefixed
func (c *StoreConfig) DeNamespacedKey(key string) string {
	if len(c.Namespace) > 0 && strings.HasPrefix(key, fmt.Sprintf("%v:", c.Namespace)) {
		return strings.TrimPrefix(key, fmt.Sprintf("%v:", c.Namespace))
	}
	return key
}
