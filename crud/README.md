# CRUD

## Problem Statement

This project aims to demonstrate the implementation of a sample TODO app with plugable storage backend. The system will showcase the following features:

- Utilization of [Protocol Buffer](https://developers.google.com/protocol-buffers) based API contracts and dependency management via [buf](https://buf.build/).
- Generation of [OpenAPI](https://swagger.io/specification/) specification from the Protobuf API definitions.
- Implementation of [gRPC](https://grpc.io/) and [HTTP gateway](https://github.com/grpc-ecosystem/grpc-gateway) server for communication.
- Configurable storage backend supporting plugable mysql, sqlite or inmemory storage options.
- Example database migration management via [sql-migrate](https://github.com/rubenv/sql-migrate),
- Implementation of a caching layer using redis or in-memory caching on top of persistant storage.
- Examples of Golang [gorm](https://gorm.io/) usage.
- Sample Golang project architecture showcasing best practices and code architecture.
- Multiple testing strategies including unit and end-to-end (e2e) testing, along with support for auto-generated mocks for unit testing.
- Containerizing of the system for deployment on Kubernetes. Scripts for Docker images and K8s deployment manifests
- Implementation of a server that exposes healthz, statusz, pprof, promethus metrics endpoints for server monitoring and debugging.
- CLI application that interacts with server via configurable http or grpc transport.

## Non-Goals

This system does not consider user management, authentication or session management. All APIs are assumed to be public and owned by everyone. These features will be demonstrated in future examples.

## Application Overview

Checkout the application overview and [how the app;ication works here](docs/app_demo.md).

## Highlevel design

Read the detailed design decision of todo service in [docs/design.md](docs/design.md).

![Highlevel architecture](docs/img/design.png "High level todo service architecture")

## Code Overview

- apis - apis directory contains protocol buffer defination for the service and [buf](https://buf.build) generated golang codes and swagger defination. Golang codes resist inside apis/go directory.
- cmd - cmd is the entry point for different binaries. Currenlty have crud-apiserver and todocl
  - crud-apiserver -  runs a server binary that responses to http and grpc.
  - todocli - is a cli application that can connect to the server via http or grpc and perform operations related to todod.

- hack - collection of scripts and configs used in the project

  - hack/docker - contains the docker file to builf and contair the service.
  - hack/k8s - basic Kubernetes yaml files to run deployment, mainly inteneded for testing.
  - hack/scripts - collection of scripts.

- pkg - library codes for the project.
  - pkg/clients - todo service clients.
  - pkg/service - contsins the business logics implementation of the HelloService.
  - pkg/storage - implementation of different persistant storage module, includeing database schema and models.

## Development

Read the development Guideline and local build instructions in [developmen.md](docs/development.md).
