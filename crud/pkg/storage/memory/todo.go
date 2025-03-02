package memory

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/sadlil/system-samples/crud/apis/go/crudapiv1"
	"github.com/sadlil/system-samples/crud/pkg/storage/models"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type todoQueryImpl struct {
	namespace string
	c         *cache.Cache
}

func NewTodoQuery() *todoQueryImpl {
	return &todoQueryImpl{
		namespace: "todo",
		c:         cache.New(0, 0),
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
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := t.c.Add(t.key(m.ID), m, 0)
	if err != nil {
		return nil, err
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

func (t *todoQueryImpl) List(ctx context.Context, offset, limit int) ([]*crudapiv1.Todo, error) {
	items := t.c.Items()
	if len(items) == 0 {
		return nil, fmt.Errorf("err: no data")
	}

	var resp []*crudapiv1.Todo
	for _, it := range items {
		m, ok := it.Object.(*models.Todo)
		if !ok {
			continue
		}
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

func (t *todoQueryImpl) GetByID(ctx context.Context, id string) (*crudapiv1.Todo, error) {
	item, found := t.c.Get(t.key(id))
	if !found {
		return nil, fmt.Errorf("err: no data")
	}

	m, ok := item.(*models.Todo)
	if !ok {
		return nil, fmt.Errorf("err: invalid data")
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
	t.c.Delete(t.key(id))
	return nil
}

func (t *todoQueryImpl) key(id string) string {
	return fmt.Sprintf("%s:%s", t.namespace, id)
}
