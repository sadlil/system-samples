module "pometheus" {
  source = "./prometheus"
}

module "servicemonitor" {
  source = "./servicemonitors"
}

module "loki" {
  source = "./loki"
}