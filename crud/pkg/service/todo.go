package service

import (
	"context"
	"errors"

	"github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"sadlil.com/samples/crud/apis/go/crudapi"
	"sadlil.com/samples/crud/pkg/storage"
)

var _ crudapi.TodoServiceServer = new(TodoServiceImpl)

type TodoServiceImpl struct {
	store storage.Store

	crudapi.UnimplementedTodoServiceServer
}

func NewToDoService() *TodoServiceImpl {
	return &TodoServiceImpl{
		store: storage.Pool(),
	}
}

func (t *TodoServiceImpl) CreateTodo(context.Context, *crudapi.CreateTodoRequest) (*crudapi.CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}

func (t *TodoServiceImpl) ListTodo(context.Context, *crudapi.ListTodoRequest) (*crudapi.ListTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodo not implemented")
}

func (t *TodoServiceImpl) GetTodo(ctx context.Context, req *crudapi.GetTodoRequest) (*crudapi.GetTodoResponse, error) {
	todo, err := t.store.Todo().GetByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "todo not found: %v", err)
		}
		glog.Error("db.GetByID failed, reason %v", err)
		return nil, status.Errorf(codes.Internal, "db.GetByID failed: %v", err)
	}
	return &crudapi.GetTodoResponse{
		Todo: todo,
	}, nil
}

func (t *TodoServiceImpl) UpdateTodo(context.Context, *crudapi.UpdateTodoRequest) (*crudapi.UpdateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}

func (t *TodoServiceImpl) DeleteTodo(context.Context, *crudapi.DeleteTodoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
