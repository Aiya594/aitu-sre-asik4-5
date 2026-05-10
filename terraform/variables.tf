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
    user_profile_port = number 
    payment_port = number
    notification_port = number
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

# rabbitmq
variable "rabbitmq" {
  type = object({
    user     = string
    password = string
  })

  default = {
    user     = "guest"
    password = "guest"
  }
}

# Scaling configuration
variable "scaling" {
  type = object({
    auth_replicas         = number
    order_replicas        = number
    product_replicas      = number
    payment_replicas      = number
    user_profile_replicas = number
    notification_replicas = number
  })

  default = {
    auth_replicas         = 1
    order_replicas        = 1
    product_replicas      = 1
    payment_replicas      = 1
    user_profile_replicas = 1
    notification_replicas = 1
  }
}