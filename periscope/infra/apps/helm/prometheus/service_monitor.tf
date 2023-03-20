resource "kubernetes_manifest" "prometheus_service_monitor_periscope" {
  manifest = {
    "apiVersion" = "monitoring.coreos.com/v1"
    "kind"       = "ServiceMonitor"
    
    "metadata" = {
      "name"      = "periscope-service-monitor"
      "namespace" = "monitoring"
      "labels" = {
        "release" = "kube-prometheus-stackr"
      }
    }

    "spec" = {
        "endpoints" = [{
            "port" = "metrics"
        }]
        "namespaceSelector" = {
            "any": "true"
        }
        "selector" = {
            "matchLabels": {
                "prometheus.io/scrape" = "true"
            }
        }
    }
    
  }
}
