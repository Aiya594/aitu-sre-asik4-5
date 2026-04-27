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
| Order Service | http://localhost:8082 | Order API |
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

