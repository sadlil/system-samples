package clients

import (
	"context"

	"github.com/sadlil/system-samples/crud/apis/go/crudapiv1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TodoServiceClient is the client API for TodoService service.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *crudapiv1.CreateTodoRequest) (*crudapiv1.CreateTodoResponse, error)
	ListTodo(ctx context.Context, in *crudapiv1.ListTodoRequest) (*crudapiv1.ListTodoResponse, error)
	GetTodo(ctx context.Context, in *crudapiv1.GetTodoRequest) (*crudapiv1.GetTodoResponse, error)
	UpdateTodo(ctx context.Context, in *crudapiv1.UpdateTodoRequest) (*crudapiv1.UpdateTodoResponse, error)
	DeleteTodo(ctx context.Context, in *crudapiv1.DeleteTodoRequest) (*emptypb.Empty, error)
	Close() error
}
