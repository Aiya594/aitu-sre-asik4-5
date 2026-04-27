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
  name = "microservices_network"
}

# =========================
# POSTGRESQL (DB VM)
# =========================
resource "docker_container" "postgres" {
  name  = "postgres"
  image = "postgres:15"

  env = [
    "POSTGRES_USER=admin",
    "POSTGRES_PASSWORD=admin",
    "POSTGRES_DB=app"
  ]

  volumes {
    host_path      = abspath("${path.module}/db_init/init.sql")
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
    "DB_HOST=postgres",
    "DB_USER=admin",
    "DB_PASSWORD=admin",
    "DB_NAME=app",
    "DB_MODE=disable",
    "DB_PORT=5432"
  ]

  ports {
    internal = 8080
    external = 8080
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
    "DB_HOST=postgres",
    "DB_USER=admin",
    "DB_PASSWORD=admin",
    "DB_NAME=app",
    "DB_MODE=disable",
    "DB_PORT=5432"
  ]

  ports {
    internal = 8082
    external = 8082
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
    internal = 9090
    external = 9090
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
    internal = 3000
    external = 3000
  }

  networks_advanced {
    name = docker_network.app_network.name
  }
}

# =========================
# FRONTEND (NGINX)
# =========================
resource "docker_container" "nginx" {
  name  = "frontend"
  image = "nginx:latest"

  ports {
    internal = 80
    external = 80
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