# Terraform Deployment Guide

Complete guide for deploying the AITU SRE ASIK4-5 microservices infrastructure using Terraform.

**Version**: 1.0  
**Last Updated**: April 2026  
**Terraform Version**: >= 1.0  
**Status**: Production Ready

---

## Overview

This guide provides step-by-step instructions for using Terraform to provision and manage the AITU SRE ASIK4-5 infrastructure. Terraform generates a complete Docker Compose configuration with environment-specific customization, ensuring reproducible deployments across development, staging, and production environments.


---

## Prerequisites

### Required Software

**Terraform**:
```bash
# Download and install from https://www.terraform.io/downloads.html
terraform version
# Expected: Terraform v1.0 or higher
```

**Docker & Docker Compose**:
```bash
# Check Docker installation
docker --version
# Expected: Docker version 20.10 or higher

# Check Docker Compose
docker-compose --version
# Expected: Docker Compose version 1.29 or higher
```

---

## Quick Start

### 5-Minute Deployment

```bash
# 1. Navigate to terraform directory
cd terraform/

# 2. Initialize Terraform
terraform init

# 3. Review default configuration
terraform plan

# 4. Deploy
terraform apply

# 5. Verify
docker-compose ps
```

**Result**: All services running with default configuration (admin/admin credentials)

### Default Configuration

When running with defaults (no custom terraform.tfvars):

| Parameter | Default Value |
|-----------|---------------|
| Database User | admin |
| Database Password | admin |
| Database Name | app |
| Nginx Port | 80 |
| Auth Service Port | 8080 |
| Order Service Port | 8082 |
| Prometheus Port | 9090 |
| Grafana Port | 3000 |
| Environment | dev |

---
