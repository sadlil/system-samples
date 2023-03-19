# System Design Project Sample

This is a curated collection of exemple System Design projects that demonstrate the perspectives
and intricacies of building and scaling production systems from scratch.

## Objectives

In real-world scenarios, we often encounter various challenges in building and scaling complex systems.
However, publicly sharing these projects and their corresponding implementation design details is seldom
practiced, making it difficult for developers to learn from each other's experiences.

This repository aims to collect a set of "tin-can" projects that may not be production-ready, but effectively
showcase the principles and practices required to build and scale a production system. By sharing these projects,
we can gain insights into the underlying architecture, design patterns, and development strategies used to build
various systems.

The primary objectives of this repository are:

- To enable developers to learn from each other's project implementations, including their architectural decisions, performance optimization techniques, and design patterns.
- To provide a reusable code and structure that can be leveraged to develop production-ready systems in the future.

While it is ideal for each project to have its own dedicated repository, we have chosen to aggregate all projects into a single repository to make it easier to find and share.

## Completed Projects

- [CRUD](crud/README.md): Sample TODO application that demonstrates Golang best practices, testing, and server monitoring techniques. Built with grpc and grpc-gateway, the project showcases a clear seperation between service handlers and storage backends. The application supports pluggable storage backends, including mysql, sqlite, and in-memory storage. Includes an additional server support to expose endpoints for healthz, statusz, pprof, and prometheus metrics. The API definition and request validation are based on Protocol Buffers. The project supports SQL migration via sql-migrate for mysql and sqlite. Additionally, the project includes a caching layer via Redis/Memory before persistent storages. This projec also includes script to build Docker images and deployment manifests for Kubernetes. [Demo](crud/docs/app_demo.md).

## Libraries

- [golib](golib/README.md) - Collection of libraries developed to organize and reuse code across multiple projects in this repository. Some of these libraries have been used in production environments.

## Planned Projects

- [Periscope](periscope/README.md) - Monitoring and Logging for Kubernetes Workloads with Infrastructrue as Code.
- [GoLink](golink/README.md) - Example URL shortener service.
- [Gossip](gossip/README.md) - Chat application with cli client.
- [LogRaft](lograft/README.md) - Distributed Log delivery service.
- [SSHProxy](sshproxy/README.md) - Auditable ssh proxy for private hosts.
- [Subway](subway/README.md) - Secure tunnel service to private network.
