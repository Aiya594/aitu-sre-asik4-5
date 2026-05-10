output "frontend_url" {
  value = "http://localhost:${var.nginx_port}"
}

output "grafana_url" {
  value = "http://localhost:${var.observability.grafana_port}"
}

output "prometheus_url" {
  value = "http://localhost:${var.observability.prometheus_port}"
}

output "rabbitmq_management" {
  value = "http://localhost:15672"
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

output "payment_service" {
  value = "http://localhost:${var.services.payment_port}"
}

output "profile_service" {
  value = "http://localhost:${var.services.user_profile_port}"
}

output "notification_service" {
  value = "http://localhost:${var.services.notification_port}"
}

