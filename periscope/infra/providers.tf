terraform {
  required_providers {
    kubernetes = {
      source = "hashicorp/kubernetes"
    }
    minikube = {
      source  = "scott-the-programmer/minikube"
      version = "0.2.3"
    }
  }
}