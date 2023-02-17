package interceptors

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestValidateUnaryRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *testRequest
		code codes.Code
	}{
		{
			name: "HappyPath",
			req:  &testRequest{isvalid: true},
			code: codes.OK,
		},
		{
			name: "ValidationFail",
			req:  &testRequest{isvalid: false},
			code: codes.InvalidArgument,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			validator := ValidateUnaryRequest()

			_, err := validator(
				context.Background(),
				test.req,
				&grpc.UnaryServerInfo{},
				grpc.UnaryHandler(func(ctx context.Context, req interface{}) (interface{}, error) {
					return struct{}{}, nil
				}),
			)

			if status.Code(err) != test.code {
				t.Errorf("validator: got %v, expected %v", err, test.code)
			}
		})
	}
}

type testRequest struct {
	isvalid bool
}

func (t *testRequest) Validate() error {
	if t.isvalid {
		return nil
	}
	return fmt.Errorf("err: validation failed")
}
