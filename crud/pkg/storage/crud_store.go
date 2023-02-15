package storage

import (
	"gorm.io/gorm"
	"sadlil.com/samples/crud/pkg/storage/memory"
	"sadlil.com/samples/crud/pkg/storage/models"
	"sadlil.com/samples/crud/pkg/storage/persistent"
)

// Store is an interface that abstracts the query functionalities required
// to interact with the underlying tables.
type Store interface {
	// Todo refers to todo table. Additional table can be added as
	// additional method signature.
	Todo() models.TodoQuery
}

var (
	// global is a global state variable that will hold the store.
	global Store
)

func Pool() Store {
	return global
}

type crudStoreImpl struct {
	todoQuery models.TodoQuery

	internalGormDB *gorm.DB
}

// NewCrudStorageForDB returns a new Store for the given Gorm DB.
func NewCrudStorageForDB(db *gorm.DB) Store {
	return &crudStoreImpl{
		internalGormDB: db,
		todoQuery:      persistent.NewTodoQuery(db),
	}
}

// NewCrudStorageForMemory returns a new Store object that uses an in-memory database.
func NewCrudStorageForMemory() Store {
	return &crudStoreImpl{
		todoQuery: memory.NewTodoQuery(),
	}
}

// Todo ...
func (i *crudStoreImpl) Todo() models.TodoQuery {
	return i.todoQuery
}
