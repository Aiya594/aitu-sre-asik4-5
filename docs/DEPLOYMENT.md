# Deployment Options

## Overview

This project supports three deployment strategies, each with different characteristics and use cases.

## 1. Docker Compose (Local Development)

Docker Compose is ideal for local development and testing. All services run on a single machine with shared networking.

### Prerequisites
- Docker installed
- Docker Compose installed (v1.29+)
- 4GB+ available RAM

### Setup

Navigate to project root:
```bash
cd sre_asik4-5
```

Build all service images:
```bash
docker-compose build
```

Start the entire stack:
```bash
docker-compose up -d
```

Verify all services are running:
```bash
docker-compose ps
```

Expected output:
```
NAME                      COMMAND                  SERVICE             STATUS              PORTS
sre_asik4-5-auth-1       "/app"                   auth-service        Up 2 minutes        0.0.0.0:8001->8001/tcp
sre_asik4-5-product-1    "/app"                   product-service     Up 2 minutes        0.0.0.0:8002->8002/tcp
sre_asik4-5-order-1      "/app"                   order-service       Up 2 minutes        0.0.0.0:8003->8003/tcp
sre_asik4-5-payment-1    "/app"                   payment-service     Up 2 minutes        0.0.0.0:8004->8004/tcp
sre_asik4-5-notification-1 "/app"               notification-service Up 2 minutes        0.0.0.0:8005->8005/tcp
sre_asik4-5-profile-1    "/app"                   user-profile-service Up 2 minutes      0.0.0.0:8006->8006/tcp
sre_asik4-5-frontend-1   "nginx -g daemon off"   frontend            Up 2 minutes        0.0.0.0:80->80/tcp
sre_asik4-5-postgres-1   "docker-entrypoint.s…"  postgres            Up 2 minutes        0.0.0.0:5432->5432/tcp
sre_asik4-5-rabbitmq-1   "docker-entrypoint.s…"  rabbitmq            Up 2 minutes        0.0.0.0:5672->5672/tcp
sre_asik4-5-prometheus-1 "/bin/prometheus --c…"  prometheus          Up 2 minutes        0.0.0.0:9090->9090/tcp
sre_asik4-5-grafana-1    "/run.sh"                grafana             Up 2 minutes        0.0.0.0:3000->3000/tcp
```

### Useful Commands

View logs for a specific service:
```bash
docker-compose logs -f auth-service
```

View logs for all services:
```bash
docker-compose logs -f
```

Stop all services:
```bash
docker-compose stop
```

Restart a specific service:
```bash
docker-compose restart order-service
```

Remove all containers (but keep volumes):
```bash
docker-compose down
```

Remove all containers and volumes:
```bash
docker-compose down -v
```

Scale a service to multiple replicas:
```bash
docker-compose up -d --scale order-service=3
```


## 2. Kubernetes

Kubernetes is recommended for production deployments. It provides advanced orchestration, auto-scaling, and self-healing capabilities.

### Prerequisites
- Kubernetes cluster running (v1.20+)
- kubectl configured
- Docker images pushed to registry

### Namespace Setup

Create a dedicated namespace:
```bash
kubectl create namespace sre-app
kubectl config set-context --current --namespace=sre-app
```

### Deploy Services

Apply all Kubernetes manifests:
```bash
kubectl apply -f k8s/
```

Verify all deployments:
```bash
kubectl get deployments
kubectl get pods
kubectl get services
```

Expected pods (should be running):
```
NAME                                  READY   STATUS    RESTARTS
auth-deployment-5d8c9f7d8f-9kl3m      1/1     Running   0
product-deployment-7b6c8d4e2f-4x9p2   1/1     Running   0
order-deployment-6a9d7e5f1c-8l2q3     1/1     Running   0
payment-deployment-8c2e9f4d3a-5m7r8   1/1     Running   0
notification-deployment-9d3f0g5e4b-6n8s9 1/1 Running 0
profile-deployment-7e1g2h6f5c-7o9t0   1/1     Running   0
postgres-deployment-4c5d6e7f8g-2p1u1  1/1     Running   0
rabbitmq-deployment-5d6e7f8g9h-3q2v2  1/1     Running   0
prometheus-deployment-6e7f8g9h0i-4r3w3 1/1    Running   0
grafana-deployment-7f8g9h0i1j-5s4x4   1/1     Running   0
```

### Access Services

Get service endpoints:
```bash
kubectl get services
```

Port-forward to access locally:
```bash
# Access Grafana
kubectl port-forward svc/grafana 3000:3000

# Access Prometheus
kubectl port-forward svc/prometheus 9090:9090

# Access Frontend
kubectl port-forward svc/nginx 80:80
```

### Scaling Services

Scale a deployment:
```bash
kubectl scale deployment order-service --replicas=3
```

View replicas:
```bash
kubectl get pods -l app=order-service
```

### Rolling Updates

Update image:
```bash
kubectl set image deployment/order-service \
  order-service=myregistry/order-service:v2 \
  --record
```

Monitor rollout:
```bash
kubectl rollout status deployment/order-service
```

Rollback if needed:
```bash
kubectl rollout undo deployment/order-service
```

### Logs and Debugging

View pod logs:
```bash
kubectl logs pod-name
kubectl logs -f pod-name  # Follow logs
```

Describe pod details:
```bash
kubectl describe pod pod-name
```

Execute command in pod:
```bash
kubectl exec -it pod-name -- bash
```

### Resource Management

View resource usage:
```bash
kubectl top nodes
kubectl top pods
```

### Cleanup

Delete all deployments:
```bash
kubectl delete -f k8s/
```

Delete namespace:
```bash
kubectl delete namespace sre-app
```

## Comparison: Deployment Options

| Feature | Docker Compose | Docker Swarm | Kubernetes |
|---------|---|---|---|
| **Use Case** | Development | Small clusters | Production |
| **Learning Curve** | Easy | Medium | Steep |
| **Scalability** | Single node | Limited | Excellent |
| **Auto-scaling** | ❌ | ❌ | ✅ |
| **Self-healing** | ❌ | ✅ | ✅ |
| **Storage Orchestration** | Basic | Basic | Advanced |
| **Networking** | Simple | Good | Excellent |
| **Multi-cloud** | ❌ | ❌ | ✅ |
| **Production Ready** | ❌ | ✅ | ✅ |

## Deployment Decision Tree

```
                        Ready to Deploy?
                              |
                    ┌─────────┴──────────┐
                    |                    |
              Local Development    Production/Multi-node
                    |                    |
            Use Docker Compose    Ready for cloud?
                                        |
                            ┌───────────┴────────────┐
                            |                        |
                      Small cluster          Large/multi-cloud
                            |                        |
                    Use Docker Swarm      Use Kubernetes
```

