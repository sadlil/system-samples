variable "periscope_service_image_tag" {
  default     = "v0.1"
  description = "Docker Image tag for the service"
}


resource "kubernetes_deployment" "periscope_service_deployment" {
  metadata {
    name = "periscope-service"
    labels = {
      "k8s.io/app" = "periscope-service"
    }
  }

  spec {
    replicas = 5
    selector {
      match_labels = {
        "k8s.io/app" = "periscope-service"
      }
    }
    template {
      metadata {
        labels = {
          "k8s.io/app" = "periscope-service"
        }
      }
      spec {
        container {
          image = "github.com/sadlil/system-samples/periscope:${var.periscope_service_image_tag}"
          name  = "periscope-service"
          port {
            name           = "metrics"
            container_port = 6443
          }

          command = ["/go/bin/periscope"]
          args    = ["--logtostderr"]
          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "periscope_service_service" {
  metadata {
    name = "periscope-service"
    labels = {
      "k8s.io/app" = "periscope-service"
      "prometheus.io/scrape" : "true"
    }
  }
  spec {
    selector = {
      "k8s.io/app" = "periscope-service"
    }
    port {
      name        = "metrics"
      port        = 6443
      target_port = 6443
    }
  }
}
