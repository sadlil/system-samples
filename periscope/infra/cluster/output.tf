output "cluster_periscope_host" {
  value       = minikube_cluster.periscope.host
  description = "Minikube Kubernetes API Server URL"
}

output "cluster_periscope_client_certificate" {
  value       = minikube_cluster.periscope.client_certificate
  description = "Minikube Kubernetes Client Certificate"
}

output "cluster_periscope_client_key" {
  value       = minikube_cluster.periscope.client_key
  description = "Minikube Kubernetes Client Key"
}

output "cluster_periscope_cluster_ca_certificate" {
  value       = minikube_cluster.periscope.cluster_ca_certificate
  description = "Minikube Kubernetes CA certificate"
}
