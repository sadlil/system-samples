# Periscope

## Problem Statement

Monitoring and Logging pipeline for Kubernetes Workloads with Infrastructrue as Code via Terraform. This project aims to establish a monitoring and logging pipeline for Kubernetes workloads, employing Infrastructure as Code principles through Terraform. The project will involve the setup of a Kubernetes cluster, as well as the configuration of a Prometheus and Loki pipeline to gather both system and application-specific metrics. The collection of metrics will be carried out through the use of ServiceMonitor, while the logs will be collected through Promtail. Once collected, the metrics and logs will be visualized through a Grafana dashboard. Moreover, the project will include guidelines for developing custom Grafana dashboards for applications. All the above functionalities will be accomplished through Terraform's Infrastructure as Code capabilities.

### Golas

Setup Infrastructure as Code that

- turns up a Kubeneretes cluster via minikube, can be extended to reuse other clusters as well.
- deploys a demo application that publishes application Logs and custom Promethues metrics.
- Promethues Pipeline to Collect Application and System Metrics via ServiceMonitor.
- Loki Pipeline to collect Logs via Promtail.
- Grafana Dashboard with custom application dashboard to display system and application Metrics from Promethues.
- Grafana Dashboard to display application logs from Loki.

Planned Improvements:

- Setup Distributed tracing for the application.
- Setup custom alerts.
