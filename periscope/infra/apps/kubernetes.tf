# module "kube_cluster" {
#   source = "../cluster"
# }

provider "kubernetes" {
  config_path = "~/.kube/config"
  # host                   = module.kube_cluster.cluster_periscope_host
  # client_certificate     = module.kube_cluster.cluster_periscope_client_certificate
  # client_key             = module.kube_cluster.cluster_periscope_client_key
  # cluster_ca_certificate = module.kube_cluster.cluster_periscope_cluster_ca_certificate
}

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
    # host                   = module.kube_cluster.cluster_periscope_host
    # cluster_ca_certificate = module.kube_cluster.cluster_periscope_cluster_ca_certificate
    # client_key             = module.kube_cluster.cluster_periscope_client_key
    # client_certificate     = module.kube_cluster.cluster_periscope_client_certificate
  }
}

module "helm" {
  source = "./helm"
}

module "service" {
  source = "./service"
}