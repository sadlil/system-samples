package persistant

import (
	"context"

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
	return nil, nil
}
