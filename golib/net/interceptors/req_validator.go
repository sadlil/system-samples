package interceptors

import (
	"context"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type validationInterface interface {
	Validate() error
}

// ValidateUnaryRequest ...
func ValidateUnaryRequest() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		reqV, ok := req.(validationInterface)
		if ok {
			err := reqV.Validate()
			if err != nil {
				glog.Errorf("Failed to validate request, err: %v", err)
				return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %v", err)
			}
		}

		return handler(ctx, req)
	}
}
