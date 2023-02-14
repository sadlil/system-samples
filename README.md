# System Design Project Sample

Collection of example System Design projects.

## What is it?

In real life we work on different projects, but rarely we have the ability to share them
publicly to learn from each others. There are lots of system design articles/books/videos
but most of them doesn't come with implementations to share the underlying details.

This repo is here to collect some tin-can projects that are far from production ready, but displays
and shares some perspective of a scalable production system with implementation details. The
two purpose of this repo is:

  1. Share some implementation level perspective of a system to learn from and improve ourselves,
  2. Reuse some code and structure to build other production ready projects.

Ideally each project should be in its own repo, but we have decided to collect all the projects
in one repo to make it easier to find and share.

## Completed Projects

- [CRUD](crud/README.md): Sample TODO application backed by grpc, http gateway proxy and storage.
  - Displays Golang common practices, testing and server monitoring techniques.
  - Example code architecture. Clear distingtion between service handler and storage backends.
  - Supports plugable storage backends for any of - mysql, sqlite or inmemory.
  - Exposes healthz, statusz, pprof, promethus metrics endpoints.
  - Protocol buffer based API defination and request validation.
  - Supports SQL migration via sql-migrate for mysql, sqlite.
  - Caching mechanisms via radis for persistant storages namely mysql.

## Libraries

- [golib](golib/README.md) - Collection of libraries developed for organizing and
reusing througout multiple projects in this repository. Some of these libraries are
uses in production in some capacity.

## Planned Projects

- [GoLink](golink/README.md) - Example Scalable URL shortener service.
- [Gossip](gossip/README.md) - Chat application with cli client.
- [LogRaft](lograft/README.md) - Distributed Log delivery service.
- [SSHProxy](sshproxy/README.md) - Auditable ssh proxy for private hosts.
- [Subway](subway/README.md) - Secure tunnel service to private network.

## Contributions

If you would like to contribute to the repo, or to own a project, please open a pull request.
