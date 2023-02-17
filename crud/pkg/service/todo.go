package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"sadlil.com/samples/crud/apis/go/crudapi"
	"sadlil.com/samples/crud/pkg/storage"
	"sadlil.com/samples/golib/cache"
	"sadlil.com/samples/golib/cache/memory"
	redisstore "sadlil.com/samples/golib/cache/redis"
)

const (
	defaultLimit = 1000
)

var _ crudapi.TodoServiceServer = new(TodoServiceImpl)

type TodoServiceImpl struct {
	store storage.Store
	cache cache.Store
	opt   TodoServiceOption

	crudapi.UnimplementedTodoServiceServer
}

type TodoServiceOption struct {
	RedisServerAddress           string
	RedisUsername, RedisPassword string
}

func NewToDoService(opt TodoServiceOption) *TodoServiceImpl {
	var cache cache.Store
	cache = memory.NewLRUStore(memory.LRUStoreConfig{Capacity: 200})
	if opt.RedisServerAddress != "" {
		glog.Infof("Redis server address found: %v", opt.RedisServerAddress)
		r := redis.NewClient(&redis.Options{
			Addr:       opt.RedisServerAddress,
			ClientName: "crudapi.v1.TodoService",
			Username:   opt.RedisUsername,
			Password:   opt.RedisPassword,
		})
		cache = redisstore.NewCacheStore(r, redisstore.StoreConfig{Namespace: "crudapi.v1.TodoService"})
	}

	return &TodoServiceImpl{
		cache: cache,
		store: storage.Pool(),
		opt:   opt,
	}
}

func (t *TodoServiceImpl) CreateTodo(ctx context.Context, req *crudapi.CreateTodoRequest) (*crudapi.CreateTodoResponse, error) {
	todo, err := t.store.Todo().Create(ctx, req.GetTodo())
	if err != nil {
		glog.Errorf("db.Create failed, reason %v", err)
		return nil, status.Errorf(codes.Internal, "db.Create failed: %v", err)
	}
	return &crudapi.CreateTodoResponse{
		Todo: todo,
	}, nil
}

func (t *TodoServiceImpl) ListTodo(ctx context.Context, req *crudapi.ListTodoRequest) (*crudapi.ListTodoResponse, error) {
	if req.GetLimit() == 0 {
		req.Limit = defaultLimit
	}

	var todos []*crudapi.Todo
	// Try featching from the cache first, if not found in cache read from the database.
	// I understand we have support for memory as storage backend, and putting a cache infront of
	// memory store doesn't make sense. But This is a sample of doing things, in prodduction we are defenetly
	// using mysql/psql, and putting a redis infront of it.
	err := t.cache.Fetch(ctx,
		fmt.Sprintf("todo:list:offset:%v:limit:%v", req.GetOffset(), req.GetLimit()),
		&todos,
		&cache.Option{
			Expiry: time.Minute,
			Source: func(ctx context.Context) (interface{}, error) {
				todo, err := t.store.Todo().List(ctx, int(req.GetOffset()), int(req.GetLimit()))
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.NotFound, "todo not found: %v", err)
					}
					glog.Errorf("db.List failed, reason %v", err)
					return nil, status.Errorf(codes.Internal, "db.List failed: %v", err)
				}
				return todo, nil
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &crudapi.ListTodoResponse{
		Todos: todos,
	}, nil
}

func (t *TodoServiceImpl) GetTodo(ctx context.Context, req *crudapi.GetTodoRequest) (*crudapi.GetTodoResponse, error) {
	var todo *crudapi.Todo
	err := t.cache.Fetch(ctx,
		fmt.Sprintf("todo:get:id:%v", req.Id),
		todo,
		&cache.Option{
			Expiry: time.Minute,
			Source: func(ctx context.Context) (interface{}, error) {
				todo, err := t.store.Todo().GetByID(ctx, req.GetId())
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, status.Errorf(codes.NotFound, "todo not found: %v", err)
					}
					glog.Errorf("db.GetByID failed, reason %v", err)
					return nil, status.Errorf(codes.Internal, "db.GetByID failed: %v", err)
				}
				return todo, err
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &crudapi.GetTodoResponse{
		Todo: todo,
	}, nil
}

func (t *TodoServiceImpl) UpdateTodo(ctx context.Context, req *crudapi.UpdateTodoRequest) (*crudapi.UpdateTodoResponse, error) {
	req.Payload.Id = req.Id
	todo, err := t.store.Todo().Update(ctx, req.GetPayload())
	if err != nil {
		glog.Errorf("db.Update failed, reason %v", err)
		return nil, status.Errorf(codes.Internal, "db.Update failed: %v", err)
	}

	_ = t.cache.Delete(ctx, fmt.Sprintf("todo:get:id:%v", req.Id))

	return &crudapi.UpdateTodoResponse{
		Todo: todo,
	}, nil
}

func (t *TodoServiceImpl) DeleteTodo(ctx context.Context, req *crudapi.DeleteTodoRequest) (*emptypb.Empty, error) {
	err := t.store.Todo().Delete(ctx, req.GetId())
	if err != nil {
		glog.Errorf("db.Delete failed, reason %v", err)
		return nil, status.Errorf(codes.Internal, "db.Delete failed: %v", err)
	}
	_ = t.cache.Delete(ctx, fmt.Sprintf("todo:get:id:%v", req.Id))
	return &emptypb.Empty{}, nil
}
