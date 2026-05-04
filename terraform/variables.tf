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
    product_port = number
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

# Scaling configuration
variable "scaling" {
  type = object({
    auth_replicas    = number
    order_replicas   = number
    product_replicas = number
  })
  default = {
    auth_replicas    = 1
    order_replicas   = 2  # Order service is resource-intensive
    product_replicas = 1
  }
}