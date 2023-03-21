variable "namespace" {
  default = "monitoring"
}

resource "kubernetes_config_map" "example" {
  metadata {
    name      = "grafana-custom-dashboards"
    namespace = var.namespace
    labels = {
      "grafana_dashboard" = "1"
    }
  }

  data = {
    "loki_dashboard.json"      = "${file("files/grafana/loki_dashboard.json")}"
    "periscope_dashboard.json" = "${file("files/grafana/periscope_dashboard.json")}"
  }
}