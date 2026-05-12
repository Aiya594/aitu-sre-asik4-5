# Quick Start Guide

## Prerequisites

- Docker & Docker Compose (for local development)
- Go 1.20+ (for service development)
- PostgreSQL client (for database access)
- Kubernetes cluster (for K8s deployment)
- Terraform CLI (for infrastructure provisioning)
- Ansible (for configuration management)

## 1. Environment Setup

Clone the repository and navigate to the project directory:

```bash
cd sre_asik4-5
```

Validate environment:

```bash
bash validate_env.sh
```

## 2. Local Development with Docker Compose

Build all services:

```bash
docker-compose build
```

Start the entire stack:

```bash
docker-compose up -d
```

Access the application:

- **Frontend**: http://localhost:80
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (default: admin/admin)

Check service health:

```bash
docker-compose ps
```

View service logs:

```bash
docker-compose logs -f <service-name>
```

Stop all services:

```bash
docker-compose down
```

## 3. Load Testing

Run load tests to simulate traffic:

```bash
bash load_test.sh
```

Monitor metrics in Grafana during load testing.

## 4. Service Access

Access individual microservices:

- **Auth Service**: http://localhost:8001
- **Product Service**: http://localhost:8002
- **Order Service**: http://localhost:8003
- **Payment Service**: http://localhost:8004
- **Notification Service**: http://localhost:8005
- **User Profile Service**: http://localhost:8006
- **Frontend**: http://localhost
- **Grafana**: http://localhost:3000
- **Prometheus**: http://localhost:9090

## Next Steps

- Review [System Architecture](./ARCHITECTURE.md)
- Explore [Deployment Options](./DEPLOYMENT.md)
- Setup [Monitoring](./MONITORING.md)
- Understand [SLI/SLO](./SLI_SLO.md)
