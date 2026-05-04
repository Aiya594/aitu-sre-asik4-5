variable "project_name" {
  type    = string
  default = "microservices-sre"
}

# Database
variable "database" {
  type = object({
    host     = string
    port     = number
    name     = string
    user     = string
    password = string
    mode     = string 
  })
}

# Service ports
variable "services" {
  type = object({
    auth_port  = number
    order_port = number
  })
}

variable "nginx_port" {
  type    = number
  default = 80
}

# Observability
variable "observability" {
  type = object({
    prometheus_port = number
    grafana_port    = number
  })
}

# Network configuration
variable "network" {
  type = object({
    name = string
  })
  default = {
    name = "microservices-network"
  }
}