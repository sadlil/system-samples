package persistant

import (
	"context"

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

func (t *todoQueryImpl) GetByID(ctx context.Context, id string) (*crudapi.Todo, error) {
	m := &models.Todo{
		ID: id,
	}
	tx := t.gormDB.First(m)
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
