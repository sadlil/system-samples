package serverframework

import (
	assetfs "github.com/philips/go-bindata-assetfs"
	"github.com/sadlil/system-samples/golib/server/interceptors"
	"github.com/sadlil/system-samples/golib/server/statserver"
	"google.golang.org/grpc"
)

type Option func(s *Server)

func Name(n string) Option {
	return func(s *Server) {
		s.Name = n
	}
}

func GRPCAddress(addr string) Option {
	return func(s *Server) {
		s.GRPCAddress = addr
	}
}

func HTTPAddress(addr string) Option {
	return func(s *Server) {
		s.HTTPAddress = addr
	}
}

func WithStatServer(srv *statserver.Server) Option {
	return func(s *Server) {
		s.StatServer = srv
	}
}

// BeforeStart run funcs before service starts.
func BeforeStart(fn func() error) Option {
	return func(s *Server) {
		s.BeforeStart = append(s.BeforeStart, fn)
	}
}

// BeforeStop run funcs before service stops.
func BeforeStop(fn func() error) Option {
	return func(s *Server) {
		s.BeforeStop = append(s.BeforeStop, fn)
	}
}

// AfterStart run funcs after service starts.
func AfterStart(fn func() error) Option {
	return func(s *Server) {
		s.AfterStart = append(s.AfterStart, fn)
	}
}

// AfterStop run funcs after service stops.
func AfterStop(fn func() error) Option {
	return func(s *Server) {
		s.AfterStop = append(s.AfterStop, fn)
	}
}

func WithUnaryInterceptors(i ...grpc.UnaryServerInterceptor) Option {
	return func(s *Server) {
		s.UnaryInterceptors = append(s.UnaryInterceptors, i...)
	}
}

func WithStreamInterceptors(i ...grpc.StreamServerInterceptor) Option {
	return func(s *Server) {
		s.StreamInterceptors = append(s.StreamInterceptors, i...)
	}
}

func EnableRequestValidation() Option {
	return func(s *Server) {
		s.UnaryInterceptors = append(s.UnaryInterceptors, interceptors.ValidateUnaryRequest())
	}
}

func EnableRequestCORS() Option {
	return func(s *Server) {
		s.enableCORS = true
	}
}

func WithSwaggerAssetFS(a *assetfs.AssetFS) Option {
	return func(s *Server) {
		s.swaggerAsset = a
	}
}
