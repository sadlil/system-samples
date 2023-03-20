resource "helm_release" "kube-prometheus" {
  name       = "kube-prometheus-stackr"
  namespace  = var.namespace
  repository = "https://prometheus-community.github.io/helm-charts"
  chart      = "kube-prometheus-stack"

  values = [
    "${file("files/values/grafana_ds.yaml")}"
  ]
}
