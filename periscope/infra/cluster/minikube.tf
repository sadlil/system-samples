provider "minikube" {
  kubernetes_version = "v1.26.1"
}

resource "minikube_cluster" "periscope" {
  cluster_name      = "periscope-k8s"
  driver            = "docker"
  container_runtime = "docker"
  delete_on_failure = true
  force_systemd     = true
  memory            = 8000
  subnet            = "192.168.50.0"
  apiserver_ips     = ["127.0.0.1", "localhost", "192.168.50.1"]
}
