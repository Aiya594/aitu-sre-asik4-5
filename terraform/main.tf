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
    host_path      = abspath("${path.module}/../db_init")
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
  # count=var.scaling.auth_replicas

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
  # count=var.scaling.order_replicas


  name  = "order-service"
  image = "order-service:latest"

  env = [
    "DB_HOST=${var.database.host}",
    "DB_USER=${var.database.user}",
    "DB_PASSWORD=${var.database.password}",
    "DB_NAME=${var.database.name}",
    "DB_MODE=${var.database.mode}",
    "DB_PORT=${var.database.port}",
    "PRODUCT_SERVICE_URL=http://product-service:8082",
    "PAYMENT_SERVICE_URL=http://payment-service:8083",
    "RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/"
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
# PRODUCT SERVICE
# =========================
resource "docker_container" "product" {
  # count=var.scaling.product_replicas


  name  = "product-service"
  image = "product-service:latest"

  env = [
    "DB_HOST=${var.database.host}",
    "DB_USER=${var.database.user}",
    "DB_PASSWORD=${var.database.password}",
    "DB_NAME=${var.database.name}",
    "DB_MODE=${var.database.mode}",
    "DB_PORT=${var.database.port}"
  ]

  ports {
    internal = var.services.product_port
    external = var.services.product_port
  }

  networks_advanced {
    name = docker_network.app_network.name
  }

  depends_on = [docker_container.postgres]
}

# =========================
# PROFILE
# =========================

resource "docker_container" "profile" {
  # count = var.scaling.user_profile_replicas

  name  = "profile-service"
  image = "profile-service:latest"

  env = [
    "DB_HOST=${var.database.host}",
    "DB_USER=${var.database.user}",
    "DB_PASSWORD=${var.database.password}",
    "DB_NAME=${var.database.name}",
    "DB_MODE=${var.database.mode}",
    "DB_PORT=${var.database.port}"
  ]

  ports {
    internal = var.services.user_profile_port
    external = var.services.user_profile_port 
  }

  networks_advanced {
    name = docker_network.app_network.name
  }

  depends_on = [
    docker_container.postgres
  ]
}

# =========================
# PAYMENT
# =========================
resource "docker_container" "payment" {
  # count = var.scaling.payment_replicas

  name  = "payment-service"
  image = "payment-service:latest"

  env = [
    "DB_HOST=${var.database.host}",
    "DB_USER=${var.database.user}",
    "DB_PASSWORD=${var.database.password}",
    "DB_NAME=${var.database.name}",
    "DB_MODE=${var.database.mode}",
    "DB_PORT=${var.database.port}"
  ]

  ports {
    internal = var.services.payment_port
    external = var.services.payment_port 
  }

  networks_advanced {
    name = docker_network.app_network.name
  }

  depends_on = [
    docker_container.postgres
  ]
}

# =========================
# NOTIFICATION
# =========================

resource "docker_container" "notification" {
  # count = var.scaling.notification_replicas

  name  = "notification-service"
  image = "notification-service:latest"

  env = [
    "DB_HOST=${var.database.host}",
    "DB_USER=${var.database.user}",
    "DB_PASSWORD=${var.database.password}",
    "DB_NAME=notificationdb",
    "DB_PORT=${var.database.port}",
    "RABBITMQ_URL=amqp://${var.rabbitmq.user}:${var.rabbitmq.password}@rabbitmq:5672/"
  ]

  ports {
    internal = var.services.notification_port
    external = var.services.notification_port
  }

  networks_advanced {
    name = docker_network.app_network.name
  }

  depends_on = [
    docker_container.postgres,
    docker_container.rabbitmq
  ]
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
    host_path      = abspath("${path.module}/../prometheus/prometheus.yml")
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
  name  = "frontend"
  image = "nginx:latest"

  ports {
    internal = var.nginx_port
    external = var.nginx_port
  }

  volumes {
    host_path      = abspath("${path.module}/../frontend/html")
    container_path = "/usr/share/nginx/html"
  }

  volumes {
    host_path      = abspath("${path.module}/../frontend/nginx.conf")
    container_path = "/etc/nginx/nginx.conf"
  }

  depends_on = [
    docker_container.auth,
    docker_container.order,
    docker_container.product,
    docker_container.payment,
    docker_container.profile
  ]

  networks_advanced {
    name = docker_network.app_network.name
  }
}

# =========================
# RABBITMQ
# =========================
resource "docker_container" "rabbitmq" {
  name  = "rabbitmq"
  image = "rabbitmq:3-management"

  env = [
    "RABBITMQ_DEFAULT_USER=${var.rabbitmq.user}",
    "RABBITMQ_DEFAULT_PASS=${var.rabbitmq.password}"
  ]

  ports {
    internal = 5672
    external = 5672
  }

  ports {
    internal = 15672
    external = 15672
  }

  networks_advanced {
    name = docker_network.app_network.name
  }
}