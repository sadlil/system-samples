package memory

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sadlil.com/samples/crud/apis/go/crudapi"
	"sadlil.com/samples/crud/pkg/storage/models"
)

type store struct {
	namespace string
	c         *cache.Cache
}

func New() (*store, error) {
	return &store{
		namespace: "todo",
		c:         cache.New(0, 0),
	}, nil
}

func (d *store) Todo() models.TodoQuery {
	return d
}

func (t *store) Create(ctx context.Context, todo *crudapi.Todo) (*crudapi.Todo, error) {
	m := &models.Todo{
		ID:          uuid.NewString(),
		Name:        todo.GetName(),
		Description: todo.GetDescription(),
		Priority:    todo.GetPriority(),
		Status:      crudapi.TodoStatus_PENDING.String(),
		Deadline:    todo.GetDeadline().AsDuration(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := t.c.Add(t.key(m.ID), m, 0)
	if err != nil {
		return nil, err
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

func (t *store) List(ctx context.Context, offset, limit int) ([]*crudapi.Todo, error) {
	items := t.c.Items()
	if len(items) == 0 {
		return nil, fmt.Errorf("err: no data")
	}

	var resp []*crudapi.Todo
	for _, it := range items {
		m, ok := it.Object.(*models.Todo)
		if !ok {
			continue
		}
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

	// Support order Desc.
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].GetCreatedAt().AsTime().UnixNano() > resp[j].GetCreatedAt().AsTime().UnixNano()
	})

	if offset > len(resp) {
		return nil, fmt.Errorf("err: no data")
	}

	if offset > 0 && len(resp) >= offset {
		resp = resp[offset:]
	}

	if limit > 0 && len(resp) > limit {
		resp = resp[:limit]
	}
	return resp, nil
}

func (t *store) GetByID(ctx context.Context, id string) (*crudapi.Todo, error) {
	item, found := t.c.Get(t.key(id))
	if !found {
		return nil, fmt.Errorf("err: no data")
	}

	m, ok := item.(*models.Todo)
	if !ok {
		return nil, fmt.Errorf("err: invalid data")
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

func (t *store) Update(ctx context.Context, todo *crudapi.Todo) (*crudapi.Todo, error) {
	item, found := t.c.Get(t.key(todo.Id))
	if !found {
		return nil, fmt.Errorf("err: no data")
	}

	m, ok := item.(*models.Todo)
	if !ok {
		return nil, fmt.Errorf("err: invalid data")
	}

	m.UpdatedAt = time.Now()
	if len(todo.GetName()) != 0 {
		m.Name = todo.GetName()
	}

	if len(todo.GetDescription()) != 0 {
		m.Description = todo.GetDescription()
	}

	if len(todo.GetPriority()) != 0 {
		m.Priority = todo.GetPriority()
	}

	if len(todo.GetStatus().String()) != 0 {
		m.Status = todo.GetStatus().String()
	}

	if todo.Deadline.IsValid() {
		m.Deadline = todo.Deadline.AsDuration()
	}

	t.c.Set(t.key(todo.Id), m, 0)
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

func (t *store) Delete(ctx context.Context, id string) error {
	t.c.Delete(t.key(id))
	return nil
}

func (t *store) key(id string) string {
	return fmt.Sprintf("%s:%s", t.namespace, id)
}
