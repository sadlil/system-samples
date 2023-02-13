package main

import (
	"context"

	"github.com/golang/glog"
	assetfs "github.com/philips/go-bindata-assetfs"
	"github.com/spf13/pflag"
	"sadlil.com/samples/crud/apis/go/crudapi"
	"sadlil.com/samples/crud/apis/openapi"
	"sadlil.com/samples/crud/pkg/service"
	"sadlil.com/samples/crud/pkg/storage"
	"sadlil.com/samples/golib/application"
	"sadlil.com/samples/golib/net/serverframework"
	"sadlil.com/samples/golib/net/statserver"
)

var (
	monitoringAddr = pflag.StringP("monitoring_addr", "m", ":6443", "Address to bind the monitoring server listner")
)

func main() {
	application.Init()

	ctx, cancel := context.WithCancel(context.Background())

	err := storage.Init(storage.StorageConfig{
		Type:     storage.SqLite,
		Database: "/tmp/samples/crud/todo.db",
	})
	if err != nil {
		glog.Fatalf("Failed to initialize database, reason: %v", err)
	}

	// Intialize servers.
	srv := serverframework.New(
		serverframework.Name("curdapi.v1.TodoService"),
		serverframework.GRPCAddress(":6001"),
		serverframework.HTTPAddress(":6002"),
		serverframework.EnableRequestValidation(),
		serverframework.EnableRequestCORS(),
		serverframework.WithSwaggerAssetFS(&assetfs.AssetFS{
			Asset:    openapi.Asset,
			AssetDir: openapi.AssetDir,
			Prefix:   "apis/openapi/apis",
		}),
		serverframework.WithStatServer(
			statserver.New(
				statserver.WithMonitoringAddr(*monitoringAddr),
				statserver.WithProfiling(true),
				statserver.WithPromMetric(true),
			),
		),
	)

	srv.RegisterGRPC(&crudapi.TodoService_ServiceDesc, service.NewToDoService())
	srv.RegisterHTTP(crudapi.RegisterTodoServiceHandler)

	err = srv.Start(ctx)
	if err != nil {
		glog.Fatalf("Failed to start server, reason: %v", err)
	}

	// Wait until Application recives a Terminate signals.
	// Context should be canceld and, srv.Stop will be called to gracefully
	// shutdown the server.
	application.ShutdownOnInterrupt(func() {
		_ = srv.Stop(ctx)
		cancel()
	})
}
