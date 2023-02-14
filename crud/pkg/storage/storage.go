package storage

import (
	"sync"

	"github.com/golang/glog"
	"sadlil.com/samples/crud/pkg/storage/memory"
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
			glog.Infof("Initializing persistent database %s", cfg.Type)
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
			glog.Infof("Initializing in memory storage")
			internalDB, _ = memory.New()
		}

	})
	return err
}

func Pool() Store {
	return internalDB
}
