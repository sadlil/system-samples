package persistent

import (
	"context"

	"github.com/google/uuid"
	"github.com/sadlil/system-samples/crud/apis/go/crudapiv1"
	"github.com/sadlil/system-samples/crud/pkg/storage/models"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type todoQueryImpl struct {
	gormDB *gorm.DB
}

func NewTodoQuery(db *gorm.DB) *todoQueryImpl {
	return &todoQueryImpl{
		gormDB: db,
	}
}

func (t *todoQueryImpl) Create(ctx context.Context, todo *crudapiv1.Todo) (*crudapiv1.Todo, error) {
	m := &models.Todo{
		ID:          uuid.NewString(),
		Name:        todo.GetName(),
		Description: todo.GetDescription(),
		Priority:    todo.GetPriority(),
		Status:      crudapiv1.TodoStatus_TODO_STATUS_PENDING.String(),
		Deadline:    todo.GetDeadline().AsDuration(),
	}
	tx := t.gormDB.WithContext(ctx).Create(m)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return t.GetByID(ctx, m.ID)
}

func (t *todoQueryImpl) List(ctx context.Context, offset, limit int) ([]*crudapiv1.Todo, error) {
	var todos []*models.Todo
	tx := t.gormDB.WithContext(ctx).Offset(offset).Limit(limit).Order("created_at DESC").Find(&todos)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var resp []*crudapiv1.Todo
	for _, m := range todos {
		resp = append(resp, &crudapiv1.Todo{
			Id:          m.ID,
			Name:        m.Name,
			Description: m.Description,
			Priority:    m.Priority,
			Status:      crudapiv1.TodoStatus(crudapiv1.TodoStatus_value[m.Status]),
			CreatedAt:   timestamppb.New(m.CreatedAt),
			Deadline:    durationpb.New(m.Deadline),
		})
	}
	return resp, nil
}

func (t *todoQueryImpl) GetByID(ctx context.Context, id string) (*crudapiv1.Todo, error) {
	m := &models.Todo{
		ID: id,
	}
	tx := t.gormDB.WithContext(ctx).First(m)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &crudapiv1.Todo{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Priority:    m.Priority,
		Status:      crudapiv1.TodoStatus(crudapiv1.TodoStatus_value[m.Status]),
		CreatedAt:   timestamppb.New(m.CreatedAt),
		Deadline:    durationpb.New(m.Deadline),
	}, nil
}

func (t *todoQueryImpl) Update(ctx context.Context, todo *crudapiv1.Todo) (*crudapiv1.Todo, error) {
	m := &models.Todo{
		ID: todo.GetId(),
	}
	tx := t.gormDB.WithContext(ctx).First(m)
	if tx.Error != nil {
		return nil, tx.Error
	}

	m.Name = todo.GetName()
	m.Description = todo.GetDescription()
	m.Priority = todo.GetPriority()
	m.Status = todo.GetStatus().String()
	m.Deadline = todo.Deadline.AsDuration()

	tx = t.gormDB.WithContext(ctx).Updates(m)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &crudapiv1.Todo{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Priority:    m.Priority,
		Status:      crudapiv1.TodoStatus(crudapiv1.TodoStatus_value[m.Status]),
		CreatedAt:   timestamppb.New(m.CreatedAt),
		Deadline:    durationpb.New(m.Deadline),
	}, nil
}

func (t *todoQueryImpl) Delete(ctx context.Context, id string) error {
	m := &models.Todo{
		ID: id,
	}
	tx := t.gormDB.WithContext(ctx).Delete(m)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
