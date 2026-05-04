terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

provider "docker" {}

resource "docker_network" "app_network" {
  name = var.network.name
}

# =========================
# POSTGRESQL (DB VM)
# =========================
resource "docker_container" "postgres" {
  name  = "postgres"
  image = "postgres:15"

  env = [
    "POSTGRES_USER=${var.database.user}",
    "POSTGRES_PASSWORD=${var.database.password}",
    "POSTGRES_DB=${var.database.name}"
  ]

  volumes {
    host_path      = abspath("${path.module}/db_init")
    container_path = "/docker-entrypoint-initdb.d"
  }

  networks_advanced {
    name = docker_network.app_network.name
  }
}

# =========================
# AUTH SERVICE
# =========================
resource "docker_container" "auth" {
  name  = "auth-service"
  image = "auth-service:latest"

  env = [
    "DB_HOST=${var.database.host}",
    "DB_USER=${var.database.user}",
    "DB_PASSWORD=${var.database.password}",
    "DB_NAME=${var.database.name}",
    "DB_MODE=${var.database.mode}",
    "DB_PORT=${var.database.port}"
  ]

  ports {
    internal = var.services.auth_port
    external = var.services.auth_port
  }

  networks_advanced {
    name = docker_network.app_network.name
  }

  depends_on = [docker_container.postgres]
}

# =========================
# ORDER SERVICE
# =========================
resource "docker_container" "order" {
  name  = "order-service"
  image = "order-service:latest"

  env = [
    "DB_HOST=${var.database.host}",
    "DB_USER=${var.database.user}",
    "DB_PASSWORD=${var.database.password}",
    "DB_NAME=${var.database.name}",
    "DB_MODE=${var.database.mode}",
    "DB_PORT=${var.database.port}"
  ]

  ports {
    internal = var.services.order_port
    external = var.services.order_port
  }

  networks_advanced {
    name = docker_network.app_network.name
  }

  depends_on = [docker_container.postgres]
}

# =========================
# PROMETHEUS
# =========================
resource "docker_container" "prometheus" {
  name  = "prometheus"
  image = "prom/prometheus"

  ports {
    internal = var.observability.prometheus_port
    external = var.observability.prometheus_port
  }

  volumes {
    host_path      = abspath("${path.module}/prometheus/prometheus.yml")
    container_path = "/etc/prometheus/prometheus.yml"
  }

  networks_advanced {
    name = docker_network.app_network.name
  }
}

# =========================
# GRAFANA
# =========================
resource "docker_container" "grafana" {
  name  = "grafana"
  image = "grafana/grafana"

  ports {
    internal = var.observability.grafana_port
    external = var.observability.grafana_port
  }

  networks_advanced {
    name = docker_network.app_network.name
  }
}

# =========================
# FRONTEND (NGINX)
# =========================
resource "docker_container" "nginx" {
  name  = "${var.project_name}-frontend"
  image = "nginx:latest"

  ports {
    internal = var.nginx_port
    external = var.nginx_port
  }

  volumes {
    host_path      = abspath("${path.module}/frontend/html")
    container_path = "/usr/share/nginx/html"
  }

  volumes {
    host_path      = abspath("${path.module}/frontend/nginx.conf")
    container_path = "/etc/nginx/nginx.conf"
  }

  networks_advanced {
    name = docker_network.app_network.name
  }
}