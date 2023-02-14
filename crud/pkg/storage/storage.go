package storage

import (
	"sync"

	"sadlil.com/samples/crud/pkg/storage/models"
	"sadlil.com/samples/crud/pkg/storage/persistent"
)

type StorageType string

const (
	MySQL  StorageType = "mysql"
	SqLite StorageType = "sqlite"
	Memory StorageType = "memory"
)

type StorageConfig struct {
	Type               StorageType
	Address, Database  string
	Username, Password string
}

type Store interface {
	Todo() models.TodoQuery
}

var (
	once       sync.Once
	internalDB Store
)

func Init(cfg StorageConfig) error {
	var err error
	once.Do(func() {
		switch cfg.Type {
		case MySQL, SqLite:
			internalDB, err = persistent.New(
				persistent.WithDBType(string(cfg.Type)),
				persistent.WithAddress(cfg.Address),
				persistent.DatabaseName(cfg.Database),
				persistent.WithUsernamePassword(cfg.Username, cfg.Password),
			)
			if err != nil {
				return
			}
		case Memory:
			// Not implemented.
		}

	})
	return err
}

func Pool() Store {
	return internalDB
}
