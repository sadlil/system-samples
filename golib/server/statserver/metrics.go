// Package statserver provides support for metrics and debugging in server binaries.
// This package simplifies the process of setting up essential promhttp monitoring
// and debugging components, by exposing a http server.
package statserver

import (
	"bytes"
	"context"
	"net/http"
	"net/http/pprof"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Server wraps a http.Server for metric and debug endpoints.
type Server struct {
	EnableProfiling bool
	EnableMetric    bool
	MonitoringAddr  string

	buf  bytes.Buffer
	http *http.Server
}

type ServerOptions func(*Server)

func New(opts ...ServerOptions) *Server {
	s := &Server{
		EnableProfiling: false,
		EnableMetric:    false,
		MonitoringAddr:  ":6443",
		buf:             bytes.Buffer{},
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithMonitoringAddr(addr string) ServerOptions {
	return func(s *Server) {
		s.MonitoringAddr = addr
	}
}

func WithProfiling(e bool) ServerOptions {
	return func(s *Server) {
		s.EnableProfiling = e
	}
}

func WithPromMetric(e bool) ServerOptions {
	return func(s *Server) {
		s.EnableMetric = e
	}
}

func WithAdditionalData(data string) ServerOptions {
	return func(s *Server) {
		s.buf.WriteString(data)
	}
}

// Start and Run a monitoring server in the provided address.
func (s *Server) Start(ctx context.Context) {
	if s.EnableProfiling || s.EnableMetric {
		// Monitoring server is separated to a different port so that we don't accidentally
		// open them to public.
		monitoringServer := http.NewServeMux()
		s.http = &http.Server{
			Addr:    s.MonitoringAddr,
			Handler: monitoringServer,
		}
		glog.Infoln("Starting metric server at", s.http.Addr)
		monitoringServer.Handle("/healthz", http.HandlerFunc(HealthzHandler))
		monitoringServer.Handle("/statusz", StatuszHandler(s.buf))
		if s.EnableProfiling {
			glog.Infoln("Profiling is enabled at", s.http.Addr+"/debug/pprof")
			monitoringServer.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
			monitoringServer.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
			monitoringServer.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
			monitoringServer.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
			monitoringServer.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
		}

		if s.EnableMetric {
			glog.Infoln("Prometheus metric is enabled at", s.http.Addr+"/metrics")
			monitoringServer.Handle("/metrics", promhttp.Handler())
		}
		go func() {
			_ = s.http.ListenAndServe()
		}()
	}
}

// Stop gracefully shuts down the server without interrupting any
// active connections. Shutdown works by first closing all open
// listeners, then closing all idle connections, and then waiting
// indefinitely for connections to return to idle and then shut down.
// If the provided context expires before the shutdown is complete,
// Shutdown returns the context's error, otherwise it returns any
// error returned from closing the Server's underlying Listener(s).
//
// When Shutdown is called, Serve, ListenAndServe, and
// ListenAndServeTLS immediately return ErrServerClosed. Make sure the
// program doesn't exit and waits instead for Shutdown to return.
//
// Shutdown does not attempt to close nor wait for hijacked
// connections such as WebSockets. The caller of Shutdown should
// separately notify such long-lived connections of shutdown and wait
// for them to close, if desired. See RegisterOnShutdown for a way to
// register shutdown notification functions.
//
// Once Shutdown has been called on a server, it may not be reused;
// future calls to methods such as Serve will return ErrServerClosed.
func (s *Server) Stop(ctx context.Context) {
	if s.http != nil {
		_ = s.http.Shutdown(ctx)
	}
}
