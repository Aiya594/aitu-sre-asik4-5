output "frontend_url" {
  value = "http://localhost:${var.nginx_port}"
}

output "grafana_url" {
  value = "http://localhost:${var.observability.grafana_port}"
}

output "prometheus_url" {
  value = "http://localhost:${var.observability.prometheus_port}"
}

output "auth_service" {
  value = "http://localhost:${var.services.auth_port}"
}

output "order_service" {
  value = "http://localhost:${var.services.order_port}"
}

output "product_service" {
  value = "http://localhost:${var.services.product_port}"
}