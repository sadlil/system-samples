package models

import (
	"context"
	"time"

	"github.com/sadlil/system-samples/crud/apis/go/crudapiv1"
	"gorm.io/gorm"
)

// Todo represents a Todo item in the database.
type Todo struct {
	// ID is a unique identifier for the Todo item.
	ID string `gorm:"column:id;primarykey"`
	// Name is the name of the Todo item.
	Name string `gorm:"column:name"`
	// Description is a longer description of the Todo item.
	Description string `gorm:"column:description"`
	// Priority is the priority of the Todo item, such as "P0", "p1".
	Priority string `gorm:"column:priority"`
	// Deadline is the deadline for the Todo item, specified as a duration from the time it was created.
	Deadline time.Duration `gorm:"column:deadline"`
	// Status is the current status of the Todo item, such as "Incomplete" or "Complete".
	Status string `gorm:"column:status"`

	// CreatedAt is the time that the Todo item was created.
	CreatedAt time.Time `gorm:"column:created_at"`
	// UpdatedAt is the time that the Todo item was last updated.
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// DeletedAt is the time that the Todo item was deleted, or null if it hasn't been deleted.
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

// TableName returns the name of the database table that corresponds to the Todo struct.
func (Todo) TableName() string {
	return "todo"
}

// TodoQuery is an interface that defines the methods required to interact with a Todo model in a CRUD API.
//
//go:generate mockery --name=TodoQuery --filename=todo_mock.go --outpkg=mockstorage --output=../mockstorage --quiet --testonly
type TodoQuery interface {
	// Create creates a new Todo record in the database and returns the created Todo record.
	// It takes a context and a pointer to a Todo object to be created.
	Create(ctx context.Context, todo *crudapiv1.Todo) (*crudapiv1.Todo, error)

	// List returns a list of Todo records from the database.
	// It takes a context, an offset, and a limit to specify the range of records to return.
	List(ctx context.Context, offset, limit int) ([]*crudapiv1.Todo, error)

	// GetByID returns a single Todo record from the database by its ID.
	// It takes a context and the ID of the Todo record to retrieve.
	GetByID(ctx context.Context, id string) (*crudapiv1.Todo, error)

	// Update updates an existing Todo record in the database and returns the updated record.
	// It takes a context and a pointer to a Todo object to be updated.
	Update(ctx context.Context, todo *crudapiv1.Todo) (*crudapiv1.Todo, error)

	// Delete deletes a single Todo record from the database by its ID.
	// It takes a context and the ID of the Todo record to be deleted.
	Delete(ctx context.Context, id string) error
}
