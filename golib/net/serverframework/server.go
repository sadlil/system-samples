package serverframework

import (
	"context"
	"fmt"
	"mime"
	"net"
	"net/http"
	"strings"

	"github.com/golang/glog"
	grpcprom "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	assetfs "github.com/philips/go-bindata-assetfs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"sadlil.com/samples/golib/net/serverframework/rpcregistry"
	"sadlil.com/samples/golib/net/statserver"
)

type Server struct {
	Name        string
	GRPCAddress string
	HTTPAddress string
	EnableCORS  bool

	StatServer *statserver.Server
	*rpcregistry.RPCServiceRegistry

	swaggerAsset *assetfs.AssetFS
	enableCORS   bool

	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor

	// Before and After funcs
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	grpc *grpc.Server
	http *http.Server
}

func New(opts ...Option) *Server {
	s := &Server{
		StatServer:         statserver.New(),
		RPCServiceRegistry: rpcregistry.New(),
		UnaryInterceptors:  []grpc.UnaryServerInterceptor{grpcprom.UnaryServerInterceptor},
		StreamInterceptors: []grpc.StreamServerInterceptor{grpcprom.StreamServerInterceptor},
	}

	for _, opt := range opts {
		opt(s)
	}

	s.grpc = grpc.NewServer(
		grpc.ChainStreamInterceptor(
			s.StreamInterceptors...,
		),
		grpc.ChainUnaryInterceptor(
			s.UnaryInterceptors...,
		),
	)
	grpcprom.EnableHandlingTimeHistogram()
	grpcprom.Register(s.grpc)
	return s
}

func (s *Server) Start(ctx context.Context) error {
	for _, fn := range s.BeforeStart {
		if err := fn(); err != nil {
			glog.Errorf("Error running BeforeStart: reason %v", err)
			return err
		}
	}

	if err := s.start(ctx); err != nil {
		return err
	}

	for _, fn := range s.AfterStart {
		if err := fn(); err != nil {
			glog.Errorf("Error running AfterStart: reason %v", err)
			return err
		}
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	var err error

	for _, fn := range s.BeforeStop {
		if err = fn(); err != nil {
			glog.Errorf("Error running BeforeStop: reason %v", err)
		}
	}

	if err = s.stop(ctx); err != nil {
		return err
	}

	for _, fn := range s.AfterStop {
		if err = fn(); err != nil {
			glog.Errorf("Error running AfterStop: reason %v", err)
		}
	}
	return err
}

func (s *Server) Run(ctx context.Context) (err error) {
	if err = s.Start(ctx); err != nil {
		return err
	}

	<-ctx.Done()
	return s.Stop(ctx)
}

func (s *Server) start(ctx context.Context) error {
	if s.StatServer != nil {
		s.StatServer.Start(ctx)
	}

	if err := s.startGPRCServer(ctx); err != nil {
		return err
	}

	if err := s.startHTTPServer(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Server) stop(ctx context.Context) error {
	// StatServer should be stopped after the grpc and http servers
	// are stopped.
	defer s.StatServer.Stop(ctx)

	s.http.Shutdown(ctx)
	s.grpc.GracefulStop()
	return nil
}

func (s *Server) startGPRCServer(ctx context.Context) error {
	if len(s.RPCServiceRegistry.Services()) == 0 {
		glog.Infof("No gRPC Service Registered.")
		return fmt.Errorf("no gRPC Service registered")
	}

	for _, svc := range s.RPCServiceRegistry.Services() {
		s.grpc.RegisterService(svc.Desc, svc.Impl)
	}

	l, err := net.Listen("tcp", s.GRPCAddress)
	if err != nil {
		return err
	}

	glog.Infof("Starting grpc server at %s", s.GRPCAddress)
	for svc, info := range s.grpc.GetServiceInfo() {
		for _, minfo := range info.Methods {
			glog.Infof("Route registered to server: %v %v\n", svc, minfo.Name)
		}
	}
	go func() {
		err := s.grpc.Serve(l)
		if err != nil {
			glog.Fatalf("Failed to start grpc server, reason %s", err)
		}
	}()
	return nil
}

func (s *Server) startHTTPServer(ctx context.Context) error {
	if len(s.RPCServiceRegistry.HTTPServices()) == 0 {
		glog.Infof("No HTTP Service handler Registered. Skipping HTTP Server.")
		return nil
	}

	if len(s.HTTPAddress) == 0 {
		glog.Infof("No HTTP Server address Provided")
		return fmt.Errorf("no HTTPAddr provided, but HTTP service registered")
	}

	clientConn, err := grpc.Dial(
		"localhost"+s.GRPCAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
		}),
	)
	for _, f := range s.RPCServiceRegistry.HTTPServices() {
		if err = f(ctx, mux, clientConn); err != nil {
			glog.Errorf("failed to register http handler: %v", err)
			return err
		}
	}

	m := http.NewServeMux()
	m.Handle("/", mux)

	if s.swaggerAsset != nil {
		s.serveSwagger(m)
	}

	handler := m
	if s.enableCORS {
		allowCORS(m)
	}

	s.http = &http.Server{
		Handler: handler,
		Addr:    s.HTTPAddress,
	}

	glog.Infof("Starting http server at %s", s.HTTPAddress)
	go func() {
		_ = s.http.ListenAndServe()
	}()
	return nil
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

func (s *Server) serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	jsonServer := http.FileServer(s.swaggerAsset)
	prefix := "/swagger/"
	mux.Handle(prefix, http.StripPrefix(prefix, jsonServer))
}
