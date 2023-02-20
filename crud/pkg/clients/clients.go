package clients

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"sadlil.com/samples/crud/apis/go/crudapi"
)

// TodoServiceClient is the client API for TodoService service.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *crudapi.CreateTodoRequest) (*crudapi.CreateTodoResponse, error)
	ListTodo(ctx context.Context, in *crudapi.ListTodoRequest) (*crudapi.ListTodoResponse, error)
	GetTodo(ctx context.Context, in *crudapi.GetTodoRequest) (*crudapi.GetTodoResponse, error)
	UpdateTodo(ctx context.Context, in *crudapi.UpdateTodoRequest) (*crudapi.UpdateTodoResponse, error)
	DeleteTodo(ctx context.Context, in *crudapi.DeleteTodoRequest) (*emptypb.Empty, error)
	Close() error
}
