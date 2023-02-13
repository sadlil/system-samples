package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"sadlil.com/samples/crud/apis/go/crudapi"
)

var _ crudapi.TodoServiceServer = new(TodoServiceImpl)

type TodoServiceImpl struct {
	// Include db, cache here

	crudapi.UnimplementedTodoServiceServer
}

func NewToDoService() *TodoServiceImpl {
	return &TodoServiceImpl{}
}

func (TodoServiceImpl) CreateTodo(context.Context, *crudapi.CreateTodoRequest) (*crudapi.CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (TodoServiceImpl) ListTodo(context.Context, *crudapi.ListTodoRequest) (*crudapi.ListTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodo not implemented")
}
func (TodoServiceImpl) GetTodo(context.Context, *crudapi.GetTodoRequest) (*crudapi.GetTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (TodoServiceImpl) UpdateTodo(context.Context, *crudapi.UpdateTodoRequest) (*crudapi.UpdateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (TodoServiceImpl) DeleteTodo(context.Context, *crudapi.DeleteTodoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
