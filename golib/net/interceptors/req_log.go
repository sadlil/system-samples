package interceptors

import (
	"context"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

// UnaryRequestLogger ...
func UnaryRequestLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		glog.Infof("Request received for method %v", info.FullMethod)
		resp, err := handler(ctx, req)
		if err != nil {
			glog.Info("Request returning error: %v", err)
		}
		return resp, err
	}
}
