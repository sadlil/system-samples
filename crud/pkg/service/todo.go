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

func (t *TodoServiceImpl) CreateTodo(ctx context.Context, req *crudapi.CreateTodoRequest) (*crudapi.CreateTodoResponse, error) {
	todo, err := t.store.Todo().Create(ctx, req.GetTodo())
	if err != nil {
		glog.Error("db.Create failed, reason %v", err)
		return nil, status.Errorf(codes.Internal, "db.GetByID failed: %v", err)
	}
	return &crudapi.CreateTodoResponse{
		Todo: todo,
	}, nil
}

func (t *TodoServiceImpl) ListTodo(ctx context.Context, req *crudapi.ListTodoRequest) (*crudapi.ListTodoResponse, error) {
	todo, err := t.store.Todo().List(ctx, int(req.GetOffset()), int(req.GetLimit()))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "todo not found: %v", err)
		}
		glog.Error("db.List failed, reason %v", err)
		return nil, status.Errorf(codes.Internal, "db.GetByID failed: %v", err)
	}
	return &crudapi.ListTodoResponse{
		Todos: todo,
	}, nil
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

func (t *TodoServiceImpl) UpdateTodo(ctx context.Context, req *crudapi.UpdateTodoRequest) (*crudapi.UpdateTodoResponse, error) {
	req.Payload.Id = req.Id
	todo, err := t.store.Todo().Update(ctx, req.GetPayload())
	if err != nil {
		glog.Error("db.Update failed, reason %v", err)
		return nil, status.Errorf(codes.Internal, "db.GetByID failed: %v", err)
	}
	return &crudapi.UpdateTodoResponse{
		Todo: todo,
	}, nil
}

func (t *TodoServiceImpl) DeleteTodo(ctx context.Context, req *crudapi.DeleteTodoRequest) (*emptypb.Empty, error) {
	err := t.store.Todo().Delete(ctx, req.GetId())
	if err != nil {
		glog.Error("db.Delete failed, reason %v", err)
		return nil, status.Errorf(codes.Internal, "db.GetByID failed: %v", err)
	}
	return &emptypb.Empty{}, nil
}
