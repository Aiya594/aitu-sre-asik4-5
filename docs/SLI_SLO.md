# SLI/SLO Definitions

## Overview

Service Level Indicators (SLIs) and Service Level Objectives (SLOs) form the foundation of SRE practices. They define reliability targets and track system performance against those targets.

## Definitions

### Service Level Indicator (SLI)

**Definition**: A measurable characteristic of a service's behavior.

**Characteristics**:
- Quantifiable and measurable
- Based on user-observed behavior
- Collected continuously from monitoring systems

**Examples**:
- HTTP request success rate
- Request latency (95th percentile)
- Error rate
- Data accuracy

### Service Level Objective (SLO)

**Definition**: A target value or range for an SLI over a specified time window.

**Characteristics**:
- Derived from business requirements
- Achievable with current infrastructure
- Realistic and sustainable
- Communicated to stakeholders

**Examples**:
- 99% availability per month
- Request latency ≤ 200ms (95th percentile)
- Error rate ≤ 1%

### Service Level Agreement (SLA)

**Definition**: A contractual commitment with customers about service reliability.

**Characteristics**:
- More stringent than SLOs
- Include penalties for violations
- Legal/contractual obligations
- Typically SLO + buffer

**Relationship**:
```
SLI (Measured) → SLO (Target) → SLA (Contractual)
                  99%              99.5%
```

## Project SLI/SLO Definitions

### SLI 1: Availability

**Metric**: Percentage of successful requests

**Formula**:
```
Availability = (Successful Requests / Total Requests) * 100
```

**PromQL Query**:
```promql
sum(rate(http_requests_total{status=~"2.."}[5m])) / sum(rate(http_requests_total[5m])) * 100
```

**Measurement Window**: Per minute (aggregated to daily/monthly)

**SLO**: ≥ 99.0%
- Translates to ~7 minutes downtime per month
- Acceptable for non-critical services

**Alerts**:
- Warning: < 99.5% over 5 minutes
- Critical: < 98.0% over 5 minutes

---

### SLI 2: Latency

**Metric**: Response time (95th percentile)

**Formula**:
```
Latency P95 = 95th percentile of request duration
```

**PromQL Query**:
```promql
histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[5m])) by (le))
```

**Measurement Window**: Per 5-minute interval

**SLO**: ≤ 200 milliseconds (P95)

**Breakdown by Service**:
| Service | P95 Target | P99 Target |
|---------|-----------|-----------|
| Auth Service | 50ms | 100ms |
| Product Service | 100ms | 150ms |
| Order Service | 150ms | 250ms |
| Payment Service | 200ms | 300ms |
| Notification Service | 500ms | 1000ms |
| User Profile Service | 100ms | 150ms |

**Alerts**:
- Warning: > 300ms over 5 minutes
- Critical: > 500ms over 5 minutes

---

### SLI 3: Error Rate

**Metric**: Percentage of failed requests

**Formula**:
```
Error Rate = (Failed Requests / Total Requests) * 100
```

**PromQL Query**:
```promql
sum(rate(http_requests_total{status=~"5.."}[5m])) / sum(rate(http_requests_total[5m])) * 100
```

**Measurement Window**: Per minute

**SLO**: ≤ 1.0%

**Types of Errors**:
- **5xx Errors**: Server errors (logs, monitoring)
- **4xx Errors**: Client errors (not counted in SLO)
- **Network Errors**: Connection timeouts
- **Database Errors**: Query failures

**Alerts**:
- Warning: > 0.5% over 5 minutes
- Critical: > 2.0% over 5 minutes

---

### SLI 4: Request Success Rate

**Metric**: Percentage of requests processed successfully

**Formula**:
```
Success Rate = (Successful + Valid Requests) / Total Requests
```

**PromQL Query**:
```promql
sum(rate(http_requests_total{status=~"2.."}[5m])) / sum(rate(http_requests_total[5m])) * 100
```

**Measurement Window**: Per minute

**SLO**: ≥ 99.0%

**Excludes**:
- Intentional rejections (invalid input)
- Rate limiting responses

**Includes**:
- Successful operations
- Valid client-side rejections
- Retryable failures

---

## Error Budget Concept

### What is Error Budget?

Error budget is the amount of time a service can be unavailable while still meeting its SLO.

**Formula**:
```
Error Budget = (1 - SLO Target) * Total Time
```

**Example**:
```
SLO: 99% availability
Time window: 30 days
Error Budget = (1 - 0.99) * 30 days * 24 hours * 60 minutes
            = 0.01 * 43,200 minutes
            = 432 minutes (~7.2 hours)
```

### Error Budget Tracking

**Monthly Availability Targets**:
| SLO | Monthly Downtime Budget |
|-----|------------------------|
| 99.0% | 7.2 hours |
| 99.5% | 3.6 hours |
| 99.9% | 43 minutes |
| 99.95% | 21 minutes |
| 99.99% | 4 minutes |

### Using Error Budget

1. **Track Consumption**
   - Monitor actual downtime
   - Compare against budget
   - Plan maintenance accordingly

2. **Decision Making**
   - High budget remaining: Can deploy changes
   - Low budget remaining: Only critical deployments
   - Budget exhausted: Focus on reliability

3. **Incident Response**
   - Each incident consumes error budget
   - Prioritize fixing high-impact incidents
   - Plan preventive measures

---

## SLO Dashboard

### Key Metrics to Display

```
╔════════════════════════════════════════════════════════════╗
║         SLO COMPLIANCE DASHBOARD (30-Day View)             ║
╠════════════════════════════════════════════════════════════╣
║                                                            ║
║  AVAILABILITY                    ERROR RATE                ║
║  ┌─────────────────────┐        ┌─────────────────────┐   ║
║  │ 99.5%               │        │ 0.3%                │   ║
║  │ Target: 99.0% ✓     │        │ Target: ≤ 1.0% ✓   │   ║
║  │ Budget Used: 34%    │        │ Budget Used: 70%    │   ║
║  └─────────────────────┘        └─────────────────────┘   ║
║                                                            ║
║  LATENCY (P95)                   SUCCESS RATE              ║
║  ┌─────────────────────┐        ┌─────────────────────┐   ║
║  │ 145ms               │        │ 99.7%               │   ║
║  │ Target: ≤ 200ms ✓   │        │ Target: ≥ 99.0% ✓   │   ║
║  │ Budget: Excellent   │        │ Budget: Excellent   │   ║
║  └─────────────────────┘        └─────────────────────┘   ║
║                                                            ║
╚════════════════════════════════════════════════════════════╝
```

---

## SLO Implementation

### 1. Instrument Services

Add metrics collection to services:

```go
import "github.com/prometheus/client_golang/prometheus"

var (
    requestCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total HTTP requests",
        },
        []string{"method", "path", "status"},
    )
    
    requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "HTTP request duration",
        },
        []string{"method", "path"},
    )
)
```

### 2. Configure Prometheus

Define scrape targets and alerts:

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'auth-service'
    static_configs:
      - targets: ['localhost:8001']

rule_files:
  - 'alerts.yml'
```

### 3. Create Alert Rules

Define conditions that trigger alerts:

```yaml
groups:
  - name: slo_alerts
    rules:
      - alert: HighErrorRate
        expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.01
        for: 5m
        annotations:
          summary: "High error rate detected"
      
      - alert: HighLatency
        expr: histogram_quantile(0.95, ...) > 0.2
        for: 5m
        annotations:
          summary: "High latency detected"
```

### 4. Monitor SLO Compliance

Create Grafana dashboards to track:
- Current SLI values
- SLO targets
- Error budget consumption
- Historical trends
- Alert status

### 5. Regular Reviews

Conduct weekly/monthly reviews:
- Check if SLOs are being met
- Analyze trends
- Identify patterns
- Adjust thresholds if needed
- Plan improvements

---

## SLO Adjustment Guidelines

### When to Increase SLO Strictness

- Current SLOs consistently exceeded
- Error budget rarely used
- User complaints about reliability
- Competitive pressure

### When to Relax SLOs

- Consistently missing targets
- Infrastructure limitations
- Cost constraints
- Business priorities change

### Review Frequency

- **Monthly**: Check compliance
- **Quarterly**: Analyze trends, adjust if needed
- **Annually**: Major review with stakeholders

---

## Service-Specific Recommendations

### Critical Services
(Auth, Payment, Order)

- Availability: ≥ 99.9%
- Latency: ≤ 100ms (P95)
- Error Rate: ≤ 0.1%

### Standard Services
(Product, Profile)

- Availability: ≥ 99.0%
- Latency: ≤ 200ms (P95)
- Error Rate: ≤ 1.0%

### Non-Critical Services
(Notification)

- Availability: ≥ 95.0%
- Latency: ≤ 500ms (P95)
- Error Rate: ≤ 5.0%

---

## Next Steps

- Setup [Monitoring](./MONITORING.md)
- Implement [Incident Response](./INCIDENT_RESPONSE.md)
- Review [Automation](./AUTOMATION_CAPACITY_PLANNING.md)
