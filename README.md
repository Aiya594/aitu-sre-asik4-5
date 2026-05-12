# End-to-End Site Reliability Engineering (SRE) Implementation

A comprehensive implementation of Site Reliability Engineering principles applied to a distributed microservices-based system with multi-platform orchestration, infrastructure-as-code provisioning, real-time monitoring, and incident management.

---

##  Documentation Index

Start with the documentation that matches your role or current task:

###  Getting Started
- **[Quick Start Guide](./docs/QUICK_START.md)** - Deploy system in 5 minutes
- **[System Architecture](./docs/ARCHITECTURE.md)** - Understand the design
- **[Technology Stack](#technology-stack)** - Tools and frameworks used

###  Deployment & Operations
- **[Deployment Options](./docs/DEPLOYMENT.md)** - Docker Compose,  Kubernetes


###  Monitoring & SRE
- **[Monitoring & Observability](./docs/MONITORING.md)** - Prometheus & Grafana setup
- **[SLI/SLO Definitions](./docs/SLI_SLO.md)** - Service Level Indicators & Objectives


---

## Quick Start

### Deploy in 3 Steps:

```bash
# 1. Build services
docker-compose build

# 2. Start system
docker-compose up -d

# 3. Access services
# Frontend: http://localhost
# Grafana: http://localhost:3000
# Prometheus: http://localhost:9090
```

**Full setup instructions:** See [Quick Start Guide](./docs/QUICK_START.md)

---

## Overview

This project demonstrates a **complete SRE lifecycle** including:

✅ **6+ Microservices** - Production-ready Go services  
✅ **Multi-Orchestration** - Docker Compose, Swarm, Kubernetes  
✅ **Infrastructure as Code** - Terraform + Ansible  
✅ **Real-time Monitoring** - Prometheus + Grafana  
✅ **SLI/SLO Tracking** - 99% availability, ≤200ms latency, ≤1% error rate  
✅ **Incident Management** - Detection, response, and postmortem analysis  
✅ **Automation & Scaling** - Health checks, auto-recovery, capacity planning  
✅ **Complete Documentation** - 9 detailed guides covering all aspects

---

## Architecture

```
┌──────────────┐
│ End User     │
└──────┬───────┘
       │
┌──────▼────────────┐
│  Frontend (Nginx) │
└──────┬────────────┘
       │
┌──────▼──────────────────────────────────────────┐
│         Microservices (6+ Services)             │
│  Auth | Product | Order | Payment | ...         │
└──────┬──────────────────────────────────────────┘
       │
┌──────▼──────────────────────────────────────────┐
│    Data Layer & Message Broker                  │
│  PostgreSQL | RabbitMQ | Redis                  │
└─────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────┐
│    Monitoring Stack                             │
│  Prometheus → Grafana                           │
└─────────────────────────────────────────────────┘
```

**Detailed architecture:** [System Architecture](./docs/ARCHITECTURE.md)

---

## Microservices

| Service | Purpose | Port |
|---------|---------|------|
| **Auth Service** | User authentication & JWT | 8001 |
| **Product Service** | Product catalog | 8002 |
| **Order Service** | Order processing | 8003 |
| **Payment Service** | Payment handling | 8004 |
| **Notification Service** | Email/alerts | 8005 |
| **User Profile Service** | User profiles | 8006 |

---

## Technology Stack

| Category | Technology |
|----------|-----------|
| **Backend** | Go 1.20+, Gin Framework |
| **Database** | PostgreSQL 15 |
| **Message Queue** | RabbitMQ |
| **Containers** | Docker & Docker Compose |
| **Orchestration** | Docker Swarm & Kubernetes |
| **Infrastructure** | Terraform & Ansible |
| **Monitoring** | Prometheus & Grafana |
| **Frontend** | Nginx & HTML/CSS/JS |

---

## Service Access

After starting services with `docker-compose up -d`:

| Service | URL |
|---------|-----|
| Frontend | http://localhost |
| Auth Service | http://localhost:8001 |
| Product Service | http://localhost:8002 |
| Order Service | http://localhost:8003 |
| Payment Service | http://localhost:8004 |
| Notification Service | http://localhost:8005 |
| User Profile Service | http://localhost:8006 |
| **Prometheus** | http://localhost:9090 |
| **Grafana** | http://localhost:3000 |

**Grafana credentials:** `admin` / `admin` (⚠️ change in production)

---

## Project Structure

```
sre_asik4-5/
├── docs/                      #  Detailed documentation
│   ├── QUICK_START.md
│   ├── ARCHITECTURE.md
│   ├── DEPLOYMENT.md
│   ├── INFRASTRUCTURE_AS_CODE.md
│   ├── MONITORING.md
│   ├── SLI_SLO.md
│   ├── INCIDENT_RESPONSE.md
│   ├── AUTOMATION_CAPACITY_PLANNING.md
│   └── TROUBLESHOOTING.md
├── auth-service/              #  Microservices
├── product-service/
├── order-service/
├── payment-service/
├── notification-service/
├── user-profile-service/
├── frontend/                  #  Web interface
├── k8s/                       #  Kubernetes manifests
├── terraform/                 #  Infrastructure code
├── ansible/                   #  Configuration management
├── prometheus/                #  Monitoring config
├── grafana/                   #  Dashboards
├── docker-compose.yml         #  Local development
├── validate_env.sh            # Environment validation
├── load_test.sh               # Load testing
└── README.md                  # This file
```

---

## Quick Commands

### Local Development

```bash
# Build and start
docker-compose build
docker-compose up -d

# View status
docker-compose ps

# View logs
docker-compose logs -f service-name

# Stop all
docker-compose down
```

### Load Testing

```bash
bash load_test.sh
```

### Environment Validation

```bash
bash validate_env.sh
```

---


---

## Default Credentials

**Grafana:**
- Username: `admin`
- Password: `admin`
- Change in production!



---


---

## Deliverables Checklist

✅ 6+ microservices (Go)  
✅ Docker Compose configuration  
✅ Kubernetes manifests (k8s/)  
✅ Terraform IaC (terraform/)  
✅ Ansible playbooks (ansible/)  
✅ Prometheus alerts (prometheus/)  
✅ Grafana dashboards (grafana/)  
✅ Load testing scripts (load_test.sh)  
✅ Comprehensive documentation (docs/)  

---
