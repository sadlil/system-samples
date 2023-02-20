package httptransport

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/golang/glog"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"sadlil.com/samples/crud/apis/go/crudapi"
	"sadlil.com/samples/crud/pkg/clients"
)

type Client struct {
	BaseURL string
	client  *http.Client
	headers http.Header
}

var _ clients.TodoServiceClient = (*Client)(nil)

type Options func(*Client)

func NewClient(opts ...Options) (clients.TodoServiceClient, error) {
	c := &Client{
		BaseURL: "http://localhost:6002",
		client: &http.Client{
			Timeout: time.Second * 5,
		},
		headers: make(http.Header),
	}

	for _, opt := range opts {
		opt(c)
	}

	c.BaseURL = strings.TrimRight(c.BaseURL, "/")
	if !strings.HasSuffix(c.BaseURL, "/api/v1/todo") {
		c.BaseURL = c.BaseURL + "/api/v1/todo"
	}

	return c, nil
}

// CreateTodo calls the CreateTodo method on the gRPC client stub
func (c *Client) CreateTodo(ctx context.Context, in *crudapi.CreateTodoRequest) (*crudapi.CreateTodoResponse, error) {
	resp := &crudapi.CreateTodoResponse{}
	if err := c.Post(ctx, "", in, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ListTodo calls the ListTodo method on the gRPC client stub
func (c *Client) ListTodo(ctx context.Context, in *crudapi.ListTodoRequest) (*crudapi.ListTodoResponse, error) {
	resp := &crudapi.ListTodoResponse{}
	if err := c.Get(ctx, "", in, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetTodo calls the GetTodo method on the gRPC client stub
func (c *Client) GetTodo(ctx context.Context, in *crudapi.GetTodoRequest) (*crudapi.GetTodoResponse, error) {
	resp := &crudapi.GetTodoResponse{}
	if err := c.Get(ctx, fmt.Sprintf("%v", in.Id), in, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateTodo calls the UpdateTodo method on the gRPC client stub
func (c *Client) UpdateTodo(ctx context.Context, in *crudapi.UpdateTodoRequest) (*crudapi.UpdateTodoResponse, error) {
	resp := &crudapi.UpdateTodoResponse{}
	if err := c.Put(ctx, fmt.Sprintf("%v", in.Id), in, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteTodo calls the DeleteTodo method on the gRPC client stub
func (c *Client) DeleteTodo(ctx context.Context, in *crudapi.DeleteTodoRequest) (*emptypb.Empty, error) {
	resp := &emptypb.Empty{}
	if err := c.Delete(ctx, fmt.Sprintf("%v", in.Id), in, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Close closes the gRPC client connection
func (c *Client) Close() error {
	c.client.CloseIdleConnections()
	return nil
}

func (c *Client) doRequest(ctx context.Context, method, url string, body, resp proto.Message) error {
	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = protojson.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	glog.Infof("Making HTTP request to url %v, with method %v", url, method)
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header = c.headers
	res, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	return unmarshalResponse(res, resp)
}

func unmarshalResponse(resp *http.Response, into proto.Message) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("http response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp status.Status
		if err := protojson.Unmarshal(body, &errResp); err != nil {
			return fmt.Errorf("http response unmarshal: %w", err)
		}
		return fmt.Errorf("http server error: %v: message: %v", errResp.Code, errResp.Message)
	}
	return protojson.Unmarshal(body, into)
}

func (c *Client) Get(ctx context.Context, path string, body, resp proto.Message) error {
	url := c.BaseURL + path
	return c.doRequest(ctx, http.MethodGet, url, body, resp)
}

func (c *Client) Post(ctx context.Context, path string, body, resp proto.Message) error {
	url := c.BaseURL + path
	return c.doRequest(ctx, http.MethodPost, url, body, resp)
}

func (c *Client) Put(ctx context.Context, path string, body, resp proto.Message) error {
	url := c.BaseURL + path
	return c.doRequest(ctx, http.MethodPut, url, body, resp)
}

func (c *Client) Delete(ctx context.Context, path string, body, resp proto.Message) error {
	url := c.BaseURL + path
	return c.doRequest(ctx, http.MethodDelete, url, body, resp)
}

func WithHeader(key, value string) Options {
	return func(c *Client) {
		c.headers.Add(key, value)
	}
}

func WithHTTPClient(httpClient *http.Client) Options {
	return func(c *Client) {
		c.client = httpClient
	}
}

func WithBaseURL(baseURL string) Options {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}
