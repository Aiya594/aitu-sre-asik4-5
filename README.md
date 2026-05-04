# AITU SRE ASIK4-5 - Microservices Platform

A comprehensive microservices platform demonstrating Site Reliability Engineering (SRE) best practices, featuring authentication and order management services with integrated monitoring and infrastructure-as-code deployment.

---



## Overview

AITU SRE ASIK4-5 is a production-ready microservices platform that demonstrates industry best practices for building, deploying, and operating scalable applications. The project includes:

### Service Communication

- **Frontend** → **Nginx**: Static file serving
- **Nginx** → **Services**: Request routing (reverse proxy)
- **Services** → **PostgreSQL**: Data persistence
- **Services** → **Prometheus**: Metrics export
- **Prometheus** ← **Services**: Scrape metrics
- **Grafana** ← **Prometheus**: Visualization

---

## Tech Stack

### Services

| Component | Technology | Version | Purpose |
|-----------|-----------|---------|---------|
| Auth Service | Go + Gin | Latest | User authentication and JWT management |
| Order Service | Go + Gin | Latest | Order management and processing |
| Frontend | HTML/CSS/JS | - | Web UI |
| API Gateway | Nginx | Latest | Reverse proxy and routing |

### Infrastructure

| Component | Technology | Purpose |
|-----------|-----------|---------|
| Database | PostgreSQL 15 | Data persistence |
| Monitoring | Prometheus | Metrics collection |
| Visualization | Grafana | Dashboard and alerting |
| Orchestration | Docker Compose | Local development |
| IaC | Terraform | Infrastructure automation |
| Version Control | Git | Code management |

### Languages & Frameworks

- **Go** - Backend services
- **Gin Framework** - REST API
- **PostgreSQL** - Relational database
- **JavaScript** - Frontend
- **Terraform** - Infrastructure as Code
- **Docker** - Containerization


---

## Quick Start

### Prerequisites

- **Docker** (v20.10+) and **Docker Compose** (v1.29+)
- **Terraform** (v1.0+) - for infrastructure provisioning
- **Go** (v1.18+) - for local service development
- **PostgreSQL Client** - for database management
- Ports available: 80, 5432, 8001, 8002, 9090, 3000

### Option 1: Deploy with Terraform 

```bash
# 1. Navigate to terraform directory
cd terraform/

# 2. Initialize Terraform
terraform init

# 4. Plan deployment
terraform plan

# 5. Apply configuration
terraform apply

# 6. Verify services
docker-compose ps
```

### Option 2: Direct Docker Compose 

```bash
# 1. Start all services
docker-compose up -d

# 2. Verify services
docker-compose ps

# 3. Check service logs
docker-compose logs -f
```

### Access Services

After successful deployment:

| Service | URL | Purpose |
|---------|-----|---------|
| Frontend | http://localhost:80 | Web UI |
| Auth Service | http://localhost:8080 | Authentication API |
| Order Service | http://localhost:8081 | Order API |
| Prometheus | http://localhost:9090 | Metrics collection |
| Grafana | http://localhost:3000 | Dashboards (admin/admin) |
| PostgreSQL | localhost:5432 | Database (admin/admin) |

---

## Services

### Auth Service

**Purpose**: User authentication and JWT token management

**Stack**: Go + Gin + PostgreSQL

**Endpoints**:
- `POST /register` - Register new user
- `POST /login` - User login
- `GET /metrics` - Prometheus metrics

**Environment Variables**:
```
DB_HOST=postgres
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=app
DB_PORT=5432
SERVICE_PORT=8001
```

**Key Features**:
- JWT-based authentication
- Password hashing and validation
- Health check endpoint
- Prometheus metrics integration

### Order Service

**Purpose**: Order management and processing

**Stack**: Go + Gin + PostgreSQL

**Endpoints**:
- `POST /orders` - Create new order
- `GET /orders` - List orders
- `GET /metrics` - Prometheus metrics

**Environment Variables**:
```
DB_HOST=postgres
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=app
DB_PORT=5432
SERVICE_PORT=8002
```

**Key Features**:
- Order CRUD operations
- JWT middleware for authentication
- Health check endpoint
- Prometheus metrics integration

### Frontend

**Purpose**: Web UI for the application

**Stack**: HTML + CSS + JavaScript

---

## Assignment 6: Automation in SRE and Capacity Planning

### Automation Mechanisms

#### 1. Automated Deployment
- **Docker Compose**: Consistent multi-container deployment
- **Terraform Integration**: Infrastructure provisioning
- **Environment Configuration**: Standardized `.env` files with validation

#### 2. Health Checks and Self-Healing
- **HTTP Health Endpoints**: `/health` for all services
- **Docker Health Checks**: Automatic container health verification
- **Restart Policies**: `restart: unless-stopped` for automatic recovery

#### 3. Monitoring-Based Alerting
- **Prometheus Metrics**: CPU, memory, request rates, error rates
- **Alert Rules**: Configured in `terraform/prometheus/alerts.yml`
  - High CPU usage (>80%)
  - Service downtime
  - High error rates (>10%)
  - Container restarts

#### 4. Configuration Validation
- **Pre-deployment Checks**: `validate_env.sh` script validates environment variables
- **Template-based Configs**: `.env.example` files for consistent setup

#### 5. Log-Based Troubleshooting
- **Centralized Logging**: Container logs accessible via `docker-compose logs`
- **Automated Patterns**: Database connection failures, service restarts

### Capacity Planning

#### Metrics Collection
- **CPU Usage**: Monitored via Prometheus
- **Memory Utilization**: Container resource tracking
- **Request Rate (RPS)**: API call frequency
- **Error Rate**: Failure percentage monitoring
- **Container Restarts**: Reliability indicators

#### Load Simulation
- **Load Testing Script**: `load_test.sh` for concurrent request simulation
- **Stress Testing**: CPU and API load generation
- **Performance Monitoring**: Response time and throughput analysis

#### Capacity Analysis
- **Order Service**: Identified as most resource-intensive component
- **Database Bottlenecks**: Connection pooling and query optimization needed
- **Scaling Thresholds**: Maximum sustainable request rates under load

#### Scaling Strategies

##### Horizontal Scaling
- **Replicas** were added via Terraform
- **Load Balancing**: Nginx configuration for request distribution
- **Service Discovery**: Container networking for inter-service communication

##### Vertical Scaling
- **Resource Allocation**: Increased CPU/memory limits in Docker configs
- **VM Upgrades**: Terraform variables for infrastructure scaling

##### Database Optimization
- **Connection Pooling**: Efficient database connection management
- **Query Optimization**: Performance tuning for high-load scenarios
- **Resource Tuning**: PostgreSQL configuration optimization

#### Auto-Scaling Considerations
- **Metric-based Scaling**: CPU threshold triggers (proposed)
- **Orchestration Integration**: Kubernetes readiness for production scaling
- **Policy-based Automation**: Declarative scaling rules

### Usage

#### Standard Deployment
```bash
# Validate configuration
./validate_env.sh

# Deploy services
docker-compose up -d

# Check health
curl http://localhost:8080/health
curl http://localhost:8081/health
curl http://localhost:8082/health
```

#### Scaled Deployment
```bash
# Deploy with multiple order service instances
docker-compose -f docker-compose.scaled.yml up -d
```

#### Load Testing
```bash
# Run load simulation
./load_test.sh

# Monitor in Grafana: http://localhost:3000
# Check alerts in Prometheus: http://localhost:9090/alerts
```

#### Monitoring
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (admin/admin)
- **Alert Rules**: View configured alerts and their status

### Deliverables for Assignment 6

#### Updated Source Code
- **Services**: Health checks (`/health`), metrics (`/metrics`), improved configuration
- **Docker Configuration**: Enhanced `docker-compose.yml` with restart policies and health checks
- **Scaled Configuration**: `docker-compose.scaled.yml` for horizontal scaling demonstration

#### Monitoring Configuration
- **Prometheus**: Alert rules in `terraform/prometheus/alerts.yml`
- **Metrics Collection**: CPU, memory, request rates, error rates, restart frequency

#### Automation Scripts
- **Configuration Validation**: `validate_env.sh` for pre-deployment checks
- **Load Testing**: `load_test.sh` for capacity analysis simulation

#### Infrastructure Updates
- **Terraform**: Scaling variables in `variables.tf` for replica configuration
- **Network Configuration**: Updated for multiple service instances

#### Documentation
- **Automation Mechanisms**: Deployment, monitoring, alerting, validation processes
- **Capacity Planning**: Load simulation, scaling strategies, performance analysis
- **Operational Procedures**: Health checks, troubleshooting, scaling procedures

---

## Infrastructure

### Docker Compose

The `docker-compose.yml` file orchestrates all services for local development:

- **Nginx**: Reverse proxy and static file server (port 80)
- **Auth Service**: Go service container (port 8001)
- **Order Service**: Go service container (port 8002)
- **PostgreSQL**: Database container (port 5432)
- **Prometheus**: Metrics collection (port 9090)
- **Grafana**: Dashboard visualization (port 3000)

### Terraform Configuration

Infrastructure as Code for reproducible deployments:

- **main.tf**: Core configuration and resource definitions
- **variables.tf**: Input variables with validation
- **outputs.tf**: Output values for reference
- **docker-compose.tpl**: Template for dynamic generation

**Capabilities**:
- Environment-specific configurations (dev/staging/prod)
- Parameterized resource creation
- Automated file generation
- Remote state management support
- Security-focused (sensitive variable handling)

---


### Database Schema

**Users Table**:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE,
    password TEXT
);
```

**Orders Table**:
```sql
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT,
    product TEXT,
    amount NUMERIC
);
```


---

## Monitoring

### Prometheus

**Purpose**: Time-series metrics collection

**Configuration**: `prometheus/prometheus.yml`

**Default Targets**:
- Auth Service: `http://auth-service:8001/metrics`
- Order Service: `http://order-service:8002/metrics`


### Grafana

**Purpose**: Visualization and dashboarding

**Default Credentials**: admin / admin

**Steps to Add Prometheus as Data Source**:
1. Navigate to http://localhost:3000
2. Login with admin/admin
3. Go to Configuration → Data Sources
4. Add Prometheus: `http://prometheus:9090`
5. Create dashboards using metrics

**Recommended Dashboards**:
- Request rates and latencies
- Error rates and 5xx errors
- Service health and uptime
- Database connection metrics
- Container resource usage

---

## Deployment

### Local Development

```bash
# Start services
docker-compose up -d --build

# View logs
docker-compose logs -f auth-service

# Run this for tables 
docker-compose exec postgres -U admin -d app
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE,
    password TEXT
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT,
    product TEXT,
    amount NUMERIC
);

# Stop services
docker-compose down


```

### Staging/Production with Terraform

See [terraform/DEPLOYMENT.md](terraform/DEPLOYMENT.md) for detailed deployment instructions.

