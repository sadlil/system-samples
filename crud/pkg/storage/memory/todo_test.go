package memory

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/sadlil/system-samples/crud/apis/go/crudapiv1"
	"github.com/sadlil/system-samples/crud/pkg/storage/models"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Through out this whole project I am not using any assertion libraries
// following https://google.github.io/styleguide/go/decisions#assertion-libraries.

func TestMemoryStoreCreate(t *testing.T) {
	// The uuid output will be 73757065-722d-4365-a375-72652d757569
	uuid.SetRand(strings.NewReader("super-secure-uuid"))
	// After the tests are ran unset the Random reader
	defer uuid.SetRand(nil)

	store := NewTodoQuery()

	in := &crudapiv1.Todo{
		Name:        "TODO",
		Description: "Hello Store!",
		Priority:    "P1",
		Status:      crudapiv1.TodoStatus_TODO_STATUS_PENDING,
		Deadline:    durationpb.New(time.Hour),
	}
	resp, err := store.Create(context.TODO(), in)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	if resp.CreatedAt.AsTime().IsZero() {
		t.Errorf("resp.CreatedAt: got %v, want time.Time", resp.CreatedAt)
	}

	if diff := cmp.Diff(in, resp, protocmp.Transform(), protocmp.IgnoreFields(resp, "id", "created_at")); diff != "" {
		t.Errorf("cmp.Diff: %v", diff)
	}

	// Check data is present in memory store
	items := store.c.Items()
	if len(items) != 1 {
		t.Errorf("len(items): got %v, expected 1", len(items))
	}

	item, found := items["todo:73757065-722d-4365-a375-72652d757569"]
	if !found {
		t.Errorf("items[key]: got %v, expected true", found)
	}

	model, ok := item.Object.(*models.Todo)
	if !ok {
		t.Errorf("Object.(*models.Todo): got %v, expected true", ok)
	}

	if model.Name != "TODO" {
		t.Errorf("model.Name: got %q, expected %q", model.Name, "TODO")
	}
}

func TestMemoryStoreList(t *testing.T) {
	store := NewTodoQuery()

	data := []*models.Todo{
		{
			ID:        uuid.NewString(),
			Name:      "TODO-1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.NewString(),
			Name:      "TODO-2",
			CreatedAt: time.Now().Add(-time.Hour),
			UpdatedAt: time.Now().Add(-time.Hour),
		},
		{
			ID:        uuid.NewString(),
			Name:      "TODO-3",
			CreatedAt: time.Now().Add(time.Hour),
			UpdatedAt: time.Now().Add(time.Hour),
		},
	}
	for _, d := range data {
		store.c.Add(store.key(d.ID), d, 0)
	}

	tests := []struct {
		name            string
		offset, limit   int
		expectErr       bool
		expectedRespLen int
		expectedName    string
	}{
		{
			name:            "HappyPath",
			expectedRespLen: 3,
			expectedName:    "TODO-3",
		},
		{
			name:            "WithOffsetAndLimit",
			offset:          1,
			limit:           10,
			expectedRespLen: 2,
			expectedName:    "TODO-1",
		},
		{
			name:            "WithOffsetAndSmallLimit",
			offset:          1,
			limit:           1,
			expectedRespLen: 1,
			expectedName:    "TODO-1",
		},
		{
			name:            "WithLimitOnly",
			offset:          0,
			limit:           1,
			expectedRespLen: 1,
			expectedName:    "TODO-3",
		},
		{
			name:      "WithLimitOnly",
			offset:    100,
			limit:     1,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := store.List(context.Background(), test.offset, test.limit)
			if err != nil {
				if !test.expectErr {
					t.Fatalf("store.List: got %v, expected no error", err)
				}
				return
			}

			if len(resp) != test.expectedRespLen {
				t.Errorf("len(resp): got %v, expected %v", len(resp), test.expectedRespLen)
			}

			// We sorted the result internally, so the todo that is in the future is the first item
			if resp[0].Name != test.expectedName {
				t.Errorf("resp[0].Name: got %q, expected %q", resp[0].Name, test.expectedName)
			}
		})
	}
}

func TestMemoryStoreGet(t *testing.T) {
	store := NewTodoQuery()

	data := []*models.Todo{
		{
			ID:        uuid.NewString(),
			Name:      "TODO-1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.NewString(),
			Name:      "TODO-2",
			CreatedAt: time.Now().Add(-time.Hour),
			UpdatedAt: time.Now().Add(-time.Hour),
		},
		{
			ID:        uuid.NewString(),
			Name:      "TODO-3",
			CreatedAt: time.Now().Add(time.Hour),
			UpdatedAt: time.Now().Add(time.Hour),
		},
	}

	for _, d := range data {
		store.c.Add(store.key(d.ID), d, 0)
	}

	tests := []struct {
		name         string
		Id           string
		expectErr    bool
		expectedName string
	}{
		{
			name:         "HappyPath",
			Id:           data[0].ID,
			expectedName: "TODO-1",
		},
		{
			name:      "IDNotPresentInDB",
			Id:        "a-not-so-secure-uuid",
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := store.GetByID(context.Background(), test.Id)
			if err != nil {
				if !test.expectErr {
					t.Fatalf("store.List: got %v, expected no error", err)
				}
				return
			}

			if resp.Name != test.expectedName {
				t.Errorf("resp.Name: got %q, expected %q", resp.Name, test.expectedName)
			}
		})
	}
}

func TestMemoryStoreUpdate(t *testing.T) {
	store := NewTodoQuery()

	in := &crudapiv1.Todo{
		Name:        "TODO",
		Description: "Hello Store!",
		Priority:    "P1",
		Status:      crudapiv1.TodoStatus(crudapiv1.TodoStatus_value["PENDING"]),
		Deadline:    durationpb.New(time.Hour),
	}
	resp, err := store.Create(context.TODO(), in)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	resp.Status = crudapiv1.TodoStatus_TODO_STATUS_DONE
	resp, err = store.Update(context.TODO(), resp)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	resp, err = store.GetByID(context.Background(), resp.Id)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	if resp.Status != crudapiv1.TodoStatus_TODO_STATUS_DONE {
		t.Errorf("resp.Status: got %v, want %v", resp.Status.String(), crudapiv1.TodoStatus_TODO_STATUS_DONE.String())
	}
}

func TestMemoryStoreDelete(t *testing.T) {
	store := NewTodoQuery()

	in := &crudapiv1.Todo{
		Name:        "TODO",
		Description: "Hello Store!",
		Priority:    "P1",
		Status:      crudapiv1.TodoStatus(crudapiv1.TodoStatus_value["PENDING"]),
		Deadline:    durationpb.New(time.Hour),
	}
	resp, err := store.Create(context.TODO(), in)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	err = store.Delete(context.TODO(), resp.Id)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	_, found := store.c.Get(store.key(resp.Id))
	if found {
		t.Errorf("c.Get: got %v, want %v", found, false)
	}
}
