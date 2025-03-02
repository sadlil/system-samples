package storage

import (
	"github.com/golang/glog"
	"github.com/sadlil/system-samples/crud/pkg/storage/memory"
	"github.com/sadlil/system-samples/crud/pkg/storage/models"
	"github.com/sadlil/system-samples/crud/pkg/storage/persistent"
	"gorm.io/gorm"
)

// Store is an interface that abstracts the query functionalities required
// to interact with the underlying tables.
type Store interface {
	// Todo refers to todo table. Additional table can be added as
	// additional method signature.
	Todo() models.TodoQuery
}

// global is a global state variable that will hold the store.
var global *crudStoreImpl

type crudStoreImpl struct {
	todoQuery models.TodoQuery

	internalGormDB *gorm.DB
}

// NewCrudStorageForDB returns a new Store for the given Gorm DB.
func NewCrudStorageForDB(db *gorm.DB, idle, open int) *crudStoreImpl {
	sql, err := db.DB()
	if err != nil {
		glog.Errorf("Failed to get sql.DB from gorm.DB, reason: %v", err)
	}

	if sql != nil {
		err = sql.Ping()
		if err != nil {
			glog.Errorf("Failed to ping sql.DB, reason: %v", err)
		}
		sql.SetMaxIdleConns(idle)
		sql.SetMaxOpenConns(open)
	}

	return &crudStoreImpl{
		internalGormDB: db,
		todoQuery:      persistent.NewTodoQuery(db),
	}
}

// NewCrudStorageForMemory returns a new Store object that uses an in-memory database.
func NewCrudStorageForMemory() *crudStoreImpl {
	return &crudStoreImpl{
		todoQuery: memory.NewTodoQuery(),
	}
}

func (c *crudStoreImpl) Shutdown() {
	glog.Infof("Shutting down database connections")
	if c.internalGormDB != nil {
		glog.Infof("Shutting down sql database connection")
		sql, err := c.internalGormDB.DB()
		if err != nil {
			glog.Errorf("Failed to get sql.DB from gorm.DB, reason: %v", err)
			return
		}
		err = sql.Close()
		if err != nil {
			glog.Errorf("Failed to Close sql.DB, reason: %v", err)
			return
		}
		glog.Infof("Database connection closed.")
	}
}

// Todo ...
func (i *crudStoreImpl) Todo() models.TodoQuery {
	return i.todoQuery
}
