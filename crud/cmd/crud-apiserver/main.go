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
	serverGRPCAddr = pflag.StringP("server_grpc_addr", "g", ":6001", "gRPC Server Address")
	serverHTTPAddr = pflag.StringP("server_http_addr", "h", ":6002", "HTTP Server Address")

	monitoringAddr = pflag.StringP("monitoring_addr", "m", ":6443", "Address to bind the monitoring server listner")

	storageType         = pflag.StringP("storage_type", "s", "mysql", "Database store for the service")
	storageDatabasePath = pflag.StringP("storage_db_path", "d", "todo_service", "Database name or db path for the service")
	storageAddress      = pflag.StringP("storage_addr", "a", "localhost:3306", "Database address")
	storageUsername     = pflag.StringP("storage_user", "u", "root", "Database storage username")
	storagePassword     = pflag.StringP("storage_pass", "p", "root", "Database storage password")
)

func main() {
	application.Init()

	ctx, cancel := context.WithCancel(context.Background())

	err := storage.Init(storage.StorageConfig{
		Type:     storage.StorageType(*storageType),
		Database: *storageDatabasePath,
		Address:  *storageAddress,
		Username: *storageUsername,
		Password: *storagePassword,
	})
	if err != nil {
		glog.Fatalf("Failed to initialize database, reason: %v", err)
	}

	// Intialize servers.
	srv := serverframework.New(
		serverframework.Name("curdapi.v1.TodoService"),
		serverframework.GRPCAddress(*serverGRPCAddr),
		serverframework.HTTPAddress(*serverHTTPAddr),
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
