package persistent

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/durationpb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sadlil.com/samples/crud/apis/go/crudapi"
	"sadlil.com/samples/crud/pkg/storage/models"
)

// // Through out this whole project I am not using any assertion libraries
// // following https://google.github.io/styleguide/go/decisions#assertion-libraries.

func TestPersistentStoreCreate(t *testing.T) {
	store, _ := getTestDB(t)

	in := &crudapi.Todo{
		Name:        "TODO",
		Description: "Hello Store!",
		Priority:    "P1",
		Status:      crudapi.TodoStatus(crudapi.TodoStatus_value["PENDING"]),
		Deadline:    durationpb.New(time.Hour),
	}
	resp, err := store.Todo().Create(context.TODO(), in)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	resp, err = store.Todo().GetByID(context.TODO(), resp.Id)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	if resp.CreatedAt.AsTime().IsZero() {
		t.Errorf("resp.CreatedAt: got %v, want time.Time", resp.CreatedAt)
	}

	if diff := cmp.Diff(in, resp, protocmp.Transform(), protocmp.IgnoreFields(resp, "id", "created_at")); diff != "" {
		t.Errorf("cmp.Diff: %v", diff)
	}
}

func TestPersistentStoreList(t *testing.T) {
	store, testDB := getTestDB(t)

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
		testDB.Create(d)
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
			offset:          -1,
			limit:           -1,
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
			name:      "WithHigherLimitOnly",
			offset:    100,
			limit:     1,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, err := store.Todo().List(context.Background(), test.offset, test.limit)
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
			if test.expectedRespLen > 0 && resp[0].Name != test.expectedName {
				t.Errorf("resp[0].Name: got %q, expected %q", resp[0].Name, test.expectedName)
			}
		})
	}
}

func TestPersistentStoreGet(t *testing.T) {
	store, testDB := getTestDB(t)

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
		testDB.Create(d)
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
			resp, err := store.Todo().GetByID(context.Background(), test.Id)
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

func TestPersistentStoreUpdate(t *testing.T) {
	store, _ := getTestDB(t)

	in := &crudapi.Todo{
		Name:        "TODO",
		Description: "Hello Store!",
		Priority:    "P1",
		Status:      crudapi.TodoStatus(crudapi.TodoStatus_value["PENDING"]),
		Deadline:    durationpb.New(time.Hour),
	}
	resp, err := store.Todo().Create(context.TODO(), in)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	resp.Status = crudapi.TodoStatus_DONE
	resp, err = store.Todo().Update(context.TODO(), resp)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	resp, err = store.Todo().GetByID(context.Background(), resp.Id)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	if resp.Status != crudapi.TodoStatus_DONE {
		t.Errorf("resp.Status: got %v, want %v", resp.Status.String(), crudapi.TodoStatus_DONE.String())
	}
}

func TestPersistentStoreDelete(t *testing.T) {
	store, testDB := getTestDB(t)

	in := &crudapi.Todo{
		Name:        "TODO",
		Description: "Hello Store!",
		Priority:    "P1",
		Status:      crudapi.TodoStatus(crudapi.TodoStatus_value["PENDING"]),
		Deadline:    durationpb.New(time.Hour),
	}
	resp, err := store.Todo().Create(context.TODO(), in)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	err = store.Todo().Delete(context.TODO(), resp.Id)
	if err != nil {
		t.Fatalf("store.Create: got %v, expected nil error", err)
	}

	var md models.Todo
	tx := testDB.First(&md).Where("id = ?", resp.Id)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			t.Fatalf("testDB.First: got %v, want %v", tx.Error, false)
		}
	}
}

func getTestDB(t *testing.T) (*db, *gorm.DB) {
	t.Helper()

	f, err := os.CreateTemp(os.TempDir(), "todo_service.db")
	if err != nil {
		t.Fatalf("Failed to create temp sqlite database file, reason %v", err)
	}

	d, err := gorm.Open(
		sqlite.Open(f.Name()),
		&gorm.Config{
			Logger: newAppLogger(),
		},
	)
	if err != nil {
		t.Fatalf("Failed to open temp sqlite database connection, reason %v", err)
	}

	// Create the tables for the test database
	err = d.AutoMigrate(&models.Todo{})
	if err != nil {
		t.Fatalf("Failed to open temp sqlite database connection, reason %v", err)
	}

	return &db{gormDB: d}, d
}
