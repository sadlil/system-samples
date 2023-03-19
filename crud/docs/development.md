# Development

## Local builds

The service is written in golang, and uses buf as its protobuf dependency.
To install all required local tools run `./hack/scripts/install_tools.sh`.

The service uses make with `Makefile` to run lint, build and test rules for the service.

- Run go fmt and lint

```sh
make fmt
```

- Generate Codes

```sh
make gen
```

- Install Code dependencies

```sh
make dep
```

- Build Local Binary

```sh
make install
```

The above command will create a `crud-server` and `todocli` in GOPATH. That can be run using `crud-server --logtostderr`.

- Build Docker Image

```sh
make build.docker VERSION=v0.1
```

The above command will create a docker image with tag `sadlil.com/samples/crud:${VERSION}`.

- Run tests

```sh
make test
```

Above command will run the unit tests for the codebase.

## Server Flags

The server binary includes following flags to provide various configuratoion options.

```go
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
```

## Client Flags

TodoCLI provides following perisstence flags.

```go
 cmd.PersistentFlags().StringP(flagTransport, "t", "http", "Transport to use for the coneection to server, http or grpc")
 cmd.PersistentFlags().StringP(flagServerAddress, "a", "http://localhost:6002", "Server address")
```

Run `todocli --help` to get the list of all available options and child subcommands.
