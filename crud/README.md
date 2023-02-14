# CRUD

## Problem Statement

A sample TODO app with plugable storage backend. The system will demonstrate -

- [Protocol Buffer](https://developers.google.com/protocol-buffers) based API contracts.
- Protobuf dependency management with [protodep](https://github.com/stormcat24/protodep).
- Generated [OpenAPI](https://swagger.io/specification/) specification.
- [gRPC](https://grpc.io/) and [HTTP gateway](https://github.com/grpc-ecosystem/grpc-gateway) server for communication.
- Configurable storage backend, support mysql, sqlite or inmemory storage.
- Example database migration management via [sql-migrate](https://github.com/rubenv/sql-migrate).
- Demonstrates a caching layer via redis on top of persistant storage.
- Example for Golang ORM usages with [gorm](https://gorm.io/).
- Sample Golang project structure.
- Examples of multiple testing strategies including unit, e2e testing.
- Support unit testing by auto generated mocks.
- Containerizing and K8s deployment scripts of the system.
- Exposes healthz, statusz, pprof, promethus metrics endpoints.

## Non Goals

This system doesn't consider user management. No authentication, No session management.
Considering all the APIs as public, and owned by everyone.

We will demonstrate Auth in some later examples.

## Highlevel design

## Code Architecture

## Development
