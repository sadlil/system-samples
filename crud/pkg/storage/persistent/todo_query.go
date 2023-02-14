package persistent

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"sadlil.com/samples/crud/apis/go/crudapi"
	"sadlil.com/samples/crud/pkg/storage/models"
)

type todoQueryImpl struct {
	gormDB *gorm.DB
}

func (d *db) Todo() models.TodoQuery {
	return &todoQueryImpl{gormDB: d.gormDB.Session(&gorm.Session{})}
}

func (t *todoQueryImpl) Create(ctx context.Context, todo *crudapi.Todo) (*crudapi.Todo, error) {
	m := &models.Todo{
		ID:          uuid.NewString(),
		Name:        todo.GetName(),
		Description: todo.GetDescription(),
		Priority:    todo.GetPriority(),
		Status:      crudapi.TodoStatus_PENDING.String(),
		Deadline:    todo.GetDeadline().AsDuration(),
	}
	tx := t.gormDB.WithContext(ctx).Create(m)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return t.GetByID(ctx, m.ID)
}

func (t *todoQueryImpl) List(ctx context.Context, offset, limit int) ([]*crudapi.Todo, error) {
	var todos []*models.Todo
	tx := t.gormDB.WithContext(ctx).Find(&todos).Offset(offset).Limit(limit).Order("created_at DESC")
	if tx.Error != nil {
		return nil, tx.Error
	}

	var resp []*crudapi.Todo
	for _, m := range todos {
		resp = append(resp, &crudapi.Todo{
			Id:          m.ID,
			Name:        m.Name,
			Description: m.Description,
			Priority:    m.Priority,
			Status:      crudapi.TodoStatus(crudapi.TodoStatus_value[m.Status]),
			CreatedAt:   timestamppb.New(m.CreatedAt),
			Deadline:    durationpb.New(m.Deadline),
		})
	}
	return resp, nil
}

func (t *todoQueryImpl) GetByID(ctx context.Context, id string) (*crudapi.Todo, error) {
	m := &models.Todo{
		ID: id,
	}
	tx := t.gormDB.WithContext(ctx).First(m)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &crudapi.Todo{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Priority:    m.Priority,
		Status:      crudapi.TodoStatus(crudapi.TodoStatus_value[m.Status]),
		CreatedAt:   timestamppb.New(m.CreatedAt),
		Deadline:    durationpb.New(m.Deadline),
	}, nil
}

func (t *todoQueryImpl) Update(ctx context.Context, todo *crudapi.Todo) (*crudapi.Todo, error) {
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

	return &crudapi.Todo{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Priority:    m.Priority,
		Status:      crudapi.TodoStatus(crudapi.TodoStatus_value[m.Status]),
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
