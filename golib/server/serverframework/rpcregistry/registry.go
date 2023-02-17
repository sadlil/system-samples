package rpcregistry

import (
	"context"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type (
	// ServiceDef represents a grpc service endpoint
	ServiceDef struct {
		Desc *grpc.ServiceDesc
		Impl any
	}

	// HTTPProxy represens a http proxy endpoint
	HTTPProxy func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
)

type RPCServiceRegistry struct {
	mu sync.Mutex

	services  []ServiceDef
	httpProxy []HTTPProxy
}

func New() *RPCServiceRegistry {
	return &RPCServiceRegistry{
		services:  make([]ServiceDef, 0),
		httpProxy: make([]HTTPProxy, 0),
	}
}

func (r *RPCServiceRegistry) RegisterGRPC(sd *grpc.ServiceDesc, impl any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.services = append(r.services, ServiceDef{Desc: sd, Impl: impl})
}

func (r *RPCServiceRegistry) RegisterHTTP(def HTTPProxy) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.httpProxy = append(r.httpProxy, def)
}

func (r *RPCServiceRegistry) Services() []ServiceDef {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.services
}

func (r *RPCServiceRegistry) HTTPServices() []HTTPProxy {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.httpProxy
}
