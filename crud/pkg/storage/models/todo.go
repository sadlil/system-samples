package models

import (
	"context"
	"time"

	"gorm.io/gorm"
	"sadlil.com/samples/crud/apis/go/crudapi"
)

type Todo struct {
	ID string `gorm:"column:id;primarykey"`

	Name        string        `gorm:"column:name"`
	Description string        `gorm:"column:description"`
	Priority    string        `gorm:"column:priority"`
	Deadline    time.Duration `gorm:"column:deadline"`
	Status      string        `gorm:"column:status"`

	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Todo) TableName() string {
	return "todo"
}

//go:generate mockery --name=TodoQuery --filename=todo_mock.go --outpkg=mockstorage --output=../mockstorage --quiet --testonly
type TodoQuery interface {
	Create(ctx context.Context, todo *crudapi.Todo) (*crudapi.Todo, error)
	List(ctx context.Context, offset, limit int) ([]*crudapi.Todo, error)
	GetByID(ctx context.Context, id string) (*crudapi.Todo, error)
	Update(ctx context.Context, todo *crudapi.Todo) (*crudapi.Todo, error)
	Delete(ctx context.Context, id string) error
}
