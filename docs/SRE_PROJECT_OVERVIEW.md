# SRE Project Overview

## Title
End-to-End Implementation of Site Reliability Engineering Practices in a Multi-Orchestrated Microservices Infrastructure Using Docker Swarm, Kubernetes, Terraform, and Ansible

## Abstract
This project demonstrates a full SRE lifecycle for a distributed microservices system. It includes:
- Containerized microservices architecture with 6+ services.
- Multi-orchestration using Docker Swarm and Kubernetes.
- Infrastructure provisioning with Terraform for Docker-based resource setup.
- Configuration and deployment automation using Ansible.
- Monitoring with Prometheus and Grafana.
- Simulated incident response and postmortem analysis.
- Capacity planning strategies and automation for reliability.

## System Architecture
- Frontend: Nginx-based web UI.
- API Gateway: Nginx reverse proxy routing requests to backend microservices.
- Microservices: auth-service, product-service, order-service, payment-service, notification-service, user-profile-service.
- Database: PostgreSQL.
- Message broker: RabbitMQ.
- Monitoring: Prometheus and Grafana.
- Orchestration: Docker and Kubernetes.
- Provisioning: Terraform.
- Configuration management: Ansible.

## Deliverables
- Source code for all microservices.
- Docker Compose configuration in `docker-compose.yml`.
- Kubernetes manifests in `k8s/`.
- Terraform files in `terraform/`.
- Ansible playbooks in `ansible/`.
- Monitoring setup in `prometheus/` and `grafana/`.
