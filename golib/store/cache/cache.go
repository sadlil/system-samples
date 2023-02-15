package cache

import (
	"context"
	"fmt"
	"time"
)

var (
	// ErrCacheMiss is a variable that represents a generic error when a cache read misses a key.
	// This error is commonly used to signal that a value was not found in a cache.
	ErrCacheMiss = fmt.Errorf("cache miss")

	// ErrInvalidObject is a variable that represents a generic error for an invalid object type.
	// This error is commonly used to indicate that an object type is not expected.
	ErrInvalidObject = fmt.Errorf("invalid object")

	//
	ErrUnImplemented = fmt.Errorf("not implemented")
)

// Store defines an interface of CacheStore
type Store interface {
	// Fetch will attempt to read the value associated with a given key from cache. If the key
	// is not found in cache, it will invoke the source function to obtain the value, write it
	// to cache, and unmarshal it into the obj.
	// If the source retruns an error Fetch will retrun the error to the caller.
	//
	//   err = store.Fetch("key", &obj, cache.Option{
	//     Expiry: time.Minute,
	//     Source: func() (any, error) {
	//       return database.Get("key")
	//     },
	//   })
	//
	// Object must be an struct, not any primitive type
	Fetch(ctx context.Context, key string, obj any, opt *Option) error

	// Get will attempt to read the value associated with a given key from cache only,
	// unmarshal it into the provided object. If the key is not found in cache, it will
	// return an ErrCacheMiss error.
	//
	//   store.Get("key", &obj)
	//
	// obj must be an struct, not any primitive type
	Get(ctx context.Context, key string, obj any) error

	// Set will marshal the provided object and write the resulting data to cache with the
	// given expiry time. If the key already exists in cache, it will overwrite the existing
	// value.
	//
	//   obj := struct{Name string}{name:"Name"}
	//   store.Set("key", obj, cache.Option{Expiry: time.Hour})
	//
	// obj must be an struct, not any primitive type
	Set(ctx context.Context, key string, obj any, opt *Option) error

	// Delete will remove/expire the entry associated with the given key from cache.
	Delete(ctx context.Context, key string) error

	// DeleteAll will remove all the entry from the cache. Supported only in some Implementations.
	DeleteAll(ctx context.Context) error
}

// Option represents configurable options for Set/Fetch action
type Option struct {
	// Expiry sets the expiration time (in seconds) on the cache.
	Expiry time.Duration

	// RaceConditionTTL is used in conjunction with the Expiry option. It will
	// prevent race conditions when cache entries expire by preventing multiple
	// processes from simultaneously regenerating the same entry. This option
	// sets the number of seconds that an expired entry can be reused while a new
	// value is being regenerated. It's a good practice to set this value if you
	// use the Expiry option.
	RaceConditionTTL time.Duration

	// Compress indicates whether to gzip the data before sending it to the
	// backend cache store.
	Compress bool

	// Source performs the read to get the data from the data source.
	Source func(ctx context.Context) (interface{}, error)

	// RefreshRate is the percentage chance to pre-refresh the item before it becomes idle.
	// RefreshRate float64

	// OnRefresh is called when Fetch is going to refresh the source object to the cache.
	// It is called with the source object and the Fetch option, and it is expected to return
	// a new Option. This allows you to modify the caching behavior based on the source object.
	OnRefresh func(obj interface{}, opt Option) Option

	// AllowStale allows us to keep the stale cache (during the refresh entry process) if reading from the Source fails.
	// This takes effect only when the Fetch happens during the RaceConditionTTL period (before the cache expires).
	// Therefore, we should have RaceConditionTTL > 0.
	// AllowStale bool

	// StalePeriod defines how long we allow the stale cache to last.
	// StalePeriod should be > RaceConditionTTL to reduce the number of trials to refresh the cache.
	// StalePeriod time.Duration
}

// var (
// 	// defaultRefreshRate defines a percentage of refresh on stale item (0, 1.0)
// 	defaultRefreshRate = 0.3
// )

// // GetRefreshRate ...
// func (o *Option) GetRefreshRate() float64 {
// 	if o.RefreshRate > 0 {
// 		return o.RefreshRate
// 	}

// 	return defaultRefreshRate
// }
