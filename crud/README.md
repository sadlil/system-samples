# CRUD

## Problem Statement

This project aims to demonstrate the implementation of a sample TODO app with plugable storage backend. The system will showcase the following features:

- Utilization of [Protocol Buffer](https://developers.google.com/protocol-buffers) based API contracts and dependency management via [protodep](https://github.com/stormcat24/protodep).
- Generation of [OpenAPI](https://swagger.io/specification/) specification from the Protobuf API definitions.
- Implementation of [gRPC](https://grpc.io/) and [HTTP gateway](https://github.com/grpc-ecosystem/grpc-gateway) server for communication.
- Configurable storage backend supporting plugable mysql, sqlite or inmemory storage options.
- Example database migration management via [sql-migrate](https://github.com/rubenv/sql-migrate),
- Implementation of a caching layer using redis or in-memory caching on top of persistant storage, as well as examples of Golang [gorm](https://gorm.io/) usage.
- Sample Golang project architecture showcasing best practices and code architecture.
- Multiple testing strategies including unit and end-to-end (e2e) testing, along with support for auto-generated mocks for unit testing.
- Containerizing of the system for deployment on Kubernetes. Scripts for Docker images and K8s deployment manifests
- Implementation of a server that exposes healthz, statusz, pprof, promethus metrics endpoints for server monitoring and debugging.

## Non-Goals

This system does not consider user management, authentication or session management. All APIs are assumed to be public and owned by everyone. These features will be demonstrated in future examples.

## Highlevel design

## Code Architecture

## Development
