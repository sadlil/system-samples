# CRUD

## Problem Statement
A sample CRUD app with MySQL as storage engine. The system will demonstrate -
 - [Protocol Buffer](https://developers.google.com/protocol-buffers) based API contract.
 - Protobuf dependency management with [protodep](https://github.com/stormcat24/protodep).
 - Generated [OpenAPI](https://swagger.io/specification/) specification.
 - [gRPC](https://grpc.io/) and an [HTTP gateway](https://github.com/grpc-ecosystem/grpc-gateway) server for communication.
 - Example MySQL migration management.
 - Example for Golang ORM usages with [gorm](https://gorm.io/).
 - Sample Golang project structure.
 - Examples of multiple testing strategies including unit, e2e testing.
 - Support unit testing by auto generated mocks.
 - Containerizing and K8s deployment scripts of the system.
 - Sample monitoring.

## Non Goals
This system doesn't consider user management. No authentication, No session management.
Considering all the APIs as public, and owned by everyone.

We will demonstrate Auth in some later examples.

