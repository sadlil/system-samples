variable "namespace" {
  default = "monitoring"
}

resource "helm_release" "kube-prometheus" {
  name       = "kube-grafana-loki"
  namespace  = var.namespace
  repository = "https://grafana.github.io/helm-charts"
  chart      = "loki-stack"
}
