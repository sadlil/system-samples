# module "minikube_cluster" {
#   source = "./cluster"
# }

module "applications" {
  source = "./apps"
}