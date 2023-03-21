# Uncomment after https://github.com/scott-the-programmer/terraform-provider-minikube/issues/57
# module "minikube_cluster" {
#   source = "./cluster"
# }

module "applications" {
  source = "./apps"
}