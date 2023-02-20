package grpctransport

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"sadlil.com/samples/crud/apis/go/crudapi"
	"sadlil.com/samples/crud/pkg/clients"
)

// Client represents a gRPC client
type Client struct {
	Addr     string            // Address of the gRPC server
	dialOpts []grpc.DialOption // Dial options used to create gRPC client connection
	callOpts []grpc.CallOption // Call options used to make gRPC method calls

	conn    *grpc.ClientConn
	service crudapi.TodoServiceClient
}

var _ clients.TodoServiceClient = (*Client)(nil)

// Options is a functional option type for configuring a Client
type Options func(*Client)

// NewClient creates a new gRPC client
func NewClient(opts ...Options) (clients.TodoServiceClient, error) {
	c := &Client{
		Addr: "localhost:6001", // default address
	}

	for _, opt := range opts {
		opt(c)
	}

	// create a gRPC client connection
	gc, err := grpc.Dial(c.Addr, c.dialOpts...)
	if err != nil {
		return nil, err
	}
	c.conn = gc
	c.service = crudapi.NewTodoServiceClient(c.conn)

	return c, nil
}

// CreateTodo calls the CreateTodo method on the gRPC client stub
func (c *Client) CreateTodo(ctx context.Context, in *crudapi.CreateTodoRequest) (*crudapi.CreateTodoResponse, error) {
	return c.service.CreateTodo(ctx, in, c.callOpts...)
}

// ListTodo calls the ListTodo method on the gRPC client stub
func (c *Client) ListTodo(ctx context.Context, in *crudapi.ListTodoRequest) (*crudapi.ListTodoResponse, error) {
	return c.service.ListTodo(ctx, in, c.callOpts...)
}

// GetTodo calls the GetTodo method on the gRPC client stub
func (c *Client) GetTodo(ctx context.Context, in *crudapi.GetTodoRequest) (*crudapi.GetTodoResponse, error) {
	return c.service.GetTodo(ctx, in, c.callOpts...)
}

// UpdateTodo calls the UpdateTodo method on the gRPC client stub
func (c *Client) UpdateTodo(ctx context.Context, in *crudapi.UpdateTodoRequest) (*crudapi.UpdateTodoResponse, error) {
	return c.service.UpdateTodo(ctx, in, c.callOpts...)
}

// DeleteTodo calls the DeleteTodo method on the gRPC client stub
func (c *Client) DeleteTodo(ctx context.Context, in *crudapi.DeleteTodoRequest) (*emptypb.Empty, error) {
	return c.service.DeleteTodo(ctx, in, c.callOpts...)
}

// Close closes the gRPC client connection
func (c *Client) Close() error {
	return c.conn.Close()
}

// WithServerAddress returns an Option to set the server address
func WithServerAddress(s string) Options {
	return func(c *Client) {
		c.Addr = s
	}
}

// WithDialOptions returns an Option to set the gRPC Dial options
func WithDialOptions(do ...grpc.DialOption) Options {
	return func(c *Client) {
		c.dialOpts = append(c.dialOpts, do...)
	}
}

// WithCallOptions returns an Option to set the gRPC Call options
func WithCallOptions(co ...grpc.CallOption) Options {
	return func(c *Client) {
		c.callOpts = append(c.callOpts, co...)
	}
}
