module "loki" {
  source = "./loki"
}

module "pometheus" {
  source = "./prometheus"
}

module "grafana" {
  source = "./grafana"
}

module "servicemonitor" {
  source = "./servicemonitors"
}

