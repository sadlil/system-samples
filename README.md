# System Design Project Sample

This is a collection of exemple System Design projects that demonstrate the perspectives
and intricacies of building and scaling production systems.

## Objectives

In real-world scenarios, we often encounter various challenges in building and scaling complex systems.
However, publicly sharing these projects and their corresponding implementation design details is seldom
practiced, making it difficult for developers to learn from each other's experiences.

This repository aims to collect a set of "tin-can" projects that may not be production-ready, but effectively
showcase the principles and practices required to build and scale a production system. By sharing these projects,
we can gain insights into the underlying architecture, design patterns, and development strategies used to build
various systems, and reus code and architecture that can be reused to develop production-ready systems.

While it is ideal for each project to have its own dedicated repository, we have chosen to aggregate all projects into a
single repository to make it easier to find and share.

## Completed Projects

- [CRUD](crud/README.md): Sample TODO application that demonstrates Golang best practices, testing, and server monitoring techniques. Built with grpc and grpc-gateway, the project showcases a clear seperation between service handlers and storage backends. The application supports pluggable storage backends, including mysql, sqlite, and in-memory storage. Includes an additional server support to expose endpoints for healthz, statusz, pprof, and prometheus metrics. The API definition and request validation are based on Protocol Buffers. The project supports SQL migration via sql-migrate for mysql and sqlite. Additionally, the project includes a caching layer via Redis/Memory before persistent storages. This projec also includes script to build Docker images and deployment manifests for Kubernetes. [Demo](crud/docs/app_demo.md).<br>

- [Periscope](periscope/README.md) - Monitoring and Logging pipeline for Kubernetes Workloads with Infrastructrue as Code via Terraform. This project aims to establish a monitoring and logging pipeline for Kubernetes workloads, employing Infrastructure as Code principles through Terraform. The project will involve the setup of a Kubernetes cluster, as well as the configuration of a Prometheus and Loki pipeline to gather both system and application-specific metrics and logs. The collection of metrics will be carried out through the use of ServiceMonitor, while the logs will be collected through Promtail. Once collected, the metrics and logs will be visualized through a Grafana dashboard. Moreover, the project will include guidelines for developing custom Grafana dashboards for applications. All the above functionalities will be accomplished through Terraform's Infrastructure as Code capabilities. See the details about the project and the expected visualization in [periscope/README.md].

## Libraries

- [golib](golib/README.md) - Collection of libraries developed to organize and reuse code across multiple projects in this repository. Some of these libraries have been used in production environments.

## Planned Projects

- [GoLink](golink/README.md) - Example URL shortener service.
- [Gossip](gossip/README.md) - Chat application with cli client.
- [LogRaft](lograft/README.md) - Distributed Log delivery service.
- [SSHProxy](sshproxy/README.md) - Auditable ssh proxy for private hosts.
- [Subway](subway/README.md) - Secure tunnel service to private network.
