# System Architecture

## Overview Architecture

```
┌─────────────────────────────────────────────────────────┐
│                     End Users                           │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│           Frontend (Nginx Reverse Proxy)                │
└────────────────────┬────────────────────────────────────┘
                     │
        ┌────────────┴────────────┐
        │                         │
┌───────▼────────┐      ┌────────▼─────────┐
│   API Routes   │      │  Static Assets   │
└───────┬────────┘      └──────────────────┘
        │
┌───────▼──────────────────────────────────────────────────────┐
│              Microservices Layer                             │
├────────────────────────────────────────────────────────────┤
│  • Authentication Service  (User login & JWT)              │
│  • Product Service        (Catalog management)             │
│  • Order Service          (Order processing)               │
│  • Payment Service        (Payment handling)               │
│  • Notification Service   (Email/alert simulation)         │
│  • User Profile Service   (User data management)           │
└───────┬──────────────────────────────────────────────────────┘
        │
┌───────┴──────────────────────────────────────────────────────┐
│              Supporting Infrastructure                       │
├────────────────────────────────────────────────────────────┤
│  • PostgreSQL Database    (Data persistence)               │
│  • RabbitMQ               (Message broker)                 │
│  • Prometheus             (Metrics collection)             │
│  • Grafana                (Dashboards & alerting)          │
└─────────────────────────────────────────────────────────────┘
```

## Microservices

### Core Services

| Service | Purpose | Technology | Port | Responsibility |
|---------|---------|-----------|------|-----------------|
| **Auth Service** | User authentication and JWT token management | Go + Gin | 8001 | Authenticate users, issue JWTs, manage sessions |
| **Product Service** | Product catalog management | Go + Gin | 8002 | CRUD operations for products, inventory management |
| **Order Service** | Order processing and management | Go + Gin | 8003 | Create orders, track status, manage order lifecycle |
| **Payment Service** | Payment handling and transaction simulation | Go + Gin | 8004 | Process payments, validate transactions, logs |
| **Notification Service** | Email and alert notifications | Go + Gin | 8005 | Send notifications, manage alert rules |
| **User Profile Service** | User profile and data management | Go + Gin | 8006 | Manage user profiles, preferences, personal data |

### Supporting Components

- **Frontend**: Nginx-based web interface with HTML/CSS/JavaScript
  - Serves static assets
  - Reverse proxy for API requests
  - Load balancing

- **Database**: PostgreSQL 15 for persistent data storage
  - Relational data structure
  - ACID compliance
  - Connection pooling

- **Message Broker**: RabbitMQ for asynchronous communication
  - Event-driven architecture
  - Service decoupling
  - Reliable message delivery

- **API Gateway**: Nginx reverse proxy for request routing
  - Single entry point
  - Request routing to services
  - SSL/TLS termination

## Service Communication Patterns

```
User Request Flow:
─────────────────
1. Client → Frontend (Nginx)
2. Frontend → Service API (e.g., Auth Service)
3. Service → Database (PostgreSQL) or Message Broker (RabbitMQ)
4. Response → Frontend → Client

Inter-Service Communication:
──────────────────────────
1. Synchronous: Direct HTTP REST API calls
2. Asynchronous: Message queue via RabbitMQ
   Example: Order Service → RabbitMQ → Notification Service

Metrics Collection:
──────────────────
1. Services expose /metrics endpoint
2. Prometheus scrapes metrics from all services
3. Grafana visualizes metrics from Prometheus
4. Alerts trigger based on metric thresholds
```

## Deployment Architecture

The system supports three orchestration approaches:

### 1. Docker Compose (Local Development)
- Single-node deployment
- All services in containers
- Shared network bridge
- Volume management for data persistence

### 2. Docker Swarm
- Multi-node clustering
- Service replication and load balancing
- Built-in service discovery
- Rolling updates

### 3. Kubernetes
- Enterprise-grade orchestration
- Auto-scaling and self-healing
- Declarative management
- Advanced networking and storage

## Infrastructure Components

### Monitoring Stack

```
Services
   ↓
Prometheus (Scrape metrics)
   ↓
Time-Series Database
   ↓
Grafana (Visualize)
   ↓
Alerting Rules (Alert on thresholds)
```

### Automation Stack

```
Terraform (Provision)
   ↓
Cloud Infrastructure (VMs, Networks)
   ↓
Ansible (Configure)
   ↓
Services Deployment
   ↓
Monitoring Setup
```

## Data Flow

### Request Flow
```
1. Client Request → Frontend (Nginx)
2. Frontend routes to appropriate service
3. Service processes request
4. Database query/update
5. Response returned to client
6. Metrics exported to Prometheus
```

### Event-Driven Flow
```
1. Service generates event
2. Event published to RabbitMQ
3. Message Broker queues event
4. Consumer service receives event
5. Consumer processes event
6. State updated in database
```

## Technology Stack

- **Backend**: Go 1.20+ with Gin Framework
- **Database**: PostgreSQL 15
- **Message Broker**: RabbitMQ
- **Containers**: Docker
- **Orchestration**: Docker Compose, Docker Swarm, Kubernetes
- **Infrastructure**: Terraform
- **Configuration**: Ansible
- **Monitoring**: Prometheus + Grafana
- **Frontend**: Nginx + HTML/CSS/JavaScript

## Scalability Considerations

### Horizontal Scaling
- Multiple service replicas
- Load balancing via Nginx or Kubernetes
- Stateless service design

### Vertical Scaling
- Increased CPU/Memory allocation
- Database optimization
- Connection pool tuning

### Database Scaling
- Read replicas for read-heavy workloads
- Partitioning for large tables
- Query optimization

