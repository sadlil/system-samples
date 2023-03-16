package main

import (
	"context"

	"github.com/golang/glog"
	assetfs "github.com/philips/go-bindata-assetfs"
	"github.com/spf13/pflag"
	"sadlil.com/samples/crud/apis/go/crudapiv1"
	"sadlil.com/samples/crud/apis/openapi"
	"sadlil.com/samples/crud/pkg/service"
	"sadlil.com/samples/crud/pkg/storage"
	"sadlil.com/samples/golib/application"
	"sadlil.com/samples/golib/server/serverframework"
	"sadlil.com/samples/golib/server/statserver"
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

	redisServerAddress = pflag.StringP("cache_redis_address", "r", "", "The address of redis server to cache data, if not set a lru memory cache will be used")
	redisUsername      = pflag.StringP("cache_redis_user", "", "", "Redis storage username")
	redisPassword      = pflag.StringP("cache_redis_pass", "", "", "Redis storage password")
)

const (
	serviceName = "crudapi.v1.TodoService"
)

func main() {
	application.Init()

	ctx, cancel := context.WithCancel(context.Background())

	err := storage.Init(storage.StorageConfig{
		DatabaseType: storage.StorageType(*storageType),
		Database:     *storageDatabasePath,
		Address:      *storageAddress,
		Username:     *storageUsername,
		Password:     *storagePassword,
	})
	if err != nil {
		glog.Fatalf("Failed to initialize database, reason: %v", err)
	}

	// Intialize servers.
	srv := serverframework.New(
		// Name currently does nothing. But have plan to add support for service
		// registry in the future.
		serverframework.Name(serviceName),
		serverframework.GRPCAddress(*serverGRPCAddr),
		serverframework.HTTPAddress(*serverHTTPAddr),
		serverframework.EnableRequestValidation(),
		// serverframweork.WithUnaryInterceptors(),
		serverframework.EnableRequestCORS(),
		serverframework.WithSwaggerAssetFS(&assetfs.AssetFS{
			Asset:    openapi.Asset,
			AssetDir: openapi.AssetDir,
			Prefix:   "apis/openapi/gen",
		}),
		serverframework.WithStatServer(
			statserver.New(
				statserver.WithMonitoringAddr(*monitoringAddr),
				statserver.WithProfiling(true),
				statserver.WithPromMetric(true),
			),
		),
	)

	srv.RegisterGRPC(
		&crudapiv1.TodoService_ServiceDesc,
		service.NewToDoService(service.TodoServiceOption{
			RedisServerAddress: *redisServerAddress,
			RedisUsername:      *redisUsername,
			RedisPassword:      *redisPassword,
		}),
	).WithHTTP(crudapiv1.RegisterTodoServiceHandler)

	err = srv.Start(ctx)
	if err != nil {
		glog.Fatalf("Failed to start server, reason: %v", err)
	}

	// Wait until Application recives a Terminate signals.
	// Once the SIGKILL/SIGTERM recived we should
	// - stop servers,
	// - stop database,
	// - cancel context just to be sure anyone depending on the
	// context is notified.
	application.ShutdownOnInterrupt(func() {
		_ = srv.Stop(ctx)
		storage.Shutdown()
		cancel()
	})
}
