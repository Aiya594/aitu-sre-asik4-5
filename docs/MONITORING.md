# Monitoring & Observability

## Overview

Monitoring is a critical component of SRE. This system uses Prometheus for metrics collection and Grafana for visualization and alerting.

## Prometheus

### What is Prometheus?

Prometheus is a time-series database and monitoring system that collects metrics from instrumented applications and stores them with timestamps.

### Configuration

**Location**: `prometheus/prometheus.yml`

The configuration defines:
- **Global settings**: Scrape interval, evaluation interval
- **Scrape configs**: Which targets to collect metrics from
- **Alert rules**: Conditions that trigger alerts

Example configuration:
```yaml
global:
  scrape_interval: 15s      # Default scrape interval
  evaluation_interval: 15s  # How often to evaluate rules
  external_labels:
    monitor: 'sre-system'

scrape_configs:
  - job_name: 'auth-service'
    static_configs:
      - targets: ['localhost:8001']
  
  - job_name: 'order-service'
    static_configs:
      - targets: ['localhost:8003']
  
  # ... other services
```

### Metrics Collected

Each service exposes metrics on the `/metrics` endpoint:

#### Application Metrics
- **http_request_duration_seconds**: HTTP request latency
- **http_requests_total**: Total HTTP requests
- **http_request_size_bytes**: HTTP request size
- **http_response_size_bytes**: HTTP response size

#### Business Metrics
- **orders_created_total**: Total orders created
- **orders_failed_total**: Failed order attempts
- **payment_processed_total**: Successful payments
- **users_authenticated_total**: User authentication count

#### System Metrics
- **process_cpu_seconds_total**: CPU time
- **process_resident_memory_bytes**: Memory usage
- **process_open_fds**: Open file descriptors
- **process_virtual_memory_bytes**: Virtual memory

#### Database Metrics
- **db_connections_open**: Active database connections
- **db_queries_duration_seconds**: Query execution time
- **db_query_errors_total**: Failed database queries

### Accessing Prometheus

**URL**: http://localhost:9090

**Common Operations**:

1. **View Targets**: http://localhost:9090/targets
   - Shows all configured scrape targets
   - Displays health status (UP/DOWN)

2. **Alerts**: http://localhost:9090/alerts
   - View all configured alerts
   - Check alert status (FIRING/PENDING)

3. **Query Metrics**: http://localhost:9090/graph
   - Write custom PromQL queries
   - Visualize metric trends

### PromQL Query Examples

```promql
# Request rate (requests per second)
rate(http_requests_total[5m])

# Error rate percentage
rate(http_requests_total{status="500"}[5m]) / rate(http_requests_total[5m]) * 100

# 95th percentile latency
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))

# Service availability
(rate(http_requests_total{status=~"2.."}[5m]) / rate(http_requests_total[5m])) * 100

# Database connection pool usage
db_connections_open / db_connections_max

# CPU usage
rate(process_cpu_seconds_total[5m]) * 100
```

### Alert Rules

**Location**: `prometheus/alerts.yml`

#### Predefined Alerts

1. **HighErrorRate**
   - Triggers when error rate > 1%
   - Duration: 5 minutes
   - Impact: Service quality degradation

2. **HighLatency**
   - Triggers when 95th percentile latency > 200ms
   - Duration: 5 minutes
   - Impact: User experience degradation

3. **ServiceDown**
   - Triggers when service is unreachable
   - Duration: 1 minute
   - Impact: Complete service unavailability

4. **HighMemoryUsage**
   - Triggers when memory usage > 80%
   - Duration: 5 minutes
   - Impact: Risk of OOM crash

5. **HighCPUUsage**
   - Triggers when CPU usage > 80%
   - Duration: 5 minutes
   - Impact: Performance degradation

## Grafana

### What is Grafana?

Grafana is a visualization platform that creates dashboards from metrics stored in Prometheus.

### Access Grafana

**URL**: http://localhost:3000

**Default Credentials**:
- Username: `admin`
- Password: `admin`

⚠️ **Important**: Change default credentials in production!

### Dashboards

#### 1. SRE Dashboard (`sre-dashboard.json`)

Main dashboard showing system health overview:

**Service Health Section**:
- Service UP/DOWN status
- Pod restart counts
- Deployment replica status

**Request Metrics Section**:
- Request rate (RPS)
- Request latency (p50, p95, p99)
- Error rate

**Resource Utilization Section**:
- CPU usage per service
- Memory usage per service
- Disk I/O metrics

**Database Metrics Section**:
- Connection pool status
- Query latency
- Slow query count

**Alerts Section**:
- Active alerts
- Alert history
- Alert statistics

### Creating Custom Dashboards

1. Click **"+"** icon → **Dashboard**
2. Add panels with visualization
3. Select Prometheus as data source
4. Write PromQL query
5. Configure visualization options
6. Save dashboard

Example panel queries:

**Request Rate**:
```promql
rate(http_requests_total[5m])
```

**Service Latency**:
```promql
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))
```

**Error Rate**:
```promql
rate(http_requests_total{status=~"5.."}[5m])
```

### Alert Notifications

Grafana can send alerts via:
- Email
- Slack
- PagerDuty
- Webhooks
- Opsgenie

**To Configure**:
1. Go to Alerting → Notification channels
2. Add new channel (e.g., Slack)
3. Configure channel settings
4. Attach to alert rules

### Grafana Best Practices

1. **Organization**: Group related metrics in sections
2. **Labeling**: Use clear panel titles
3. **Colors**: Red for errors, yellow for warnings, green for good
4. **Time Ranges**: Set appropriate time windows (5m for alerts, 1h for trends)
5. **Refresh Rate**: Auto-refresh every 30-60 seconds
6. **Documentation**: Add descriptions to complex panels

## Metrics Best Practices

### Naming Conventions

```
<namespace>_<subsystem>_<name>_<unit>

Examples:
- http_request_duration_seconds
- db_connections_open
- orders_created_total
- cache_hit_ratio
```

### Metric Types

1. **Counter**: Monotonically increasing values
   - Total requests, total errors
   - Never decreases

2. **Gauge**: Can increase or decrease
   - Current memory usage, active connections
   - Can have any value

3. **Histogram**: Observations in defined buckets
   - Request latency, request size
   - Calculates quantiles

4. **Summary**: Similar to histogram
   - Percentile calculations on client side

### SLO Tracking

Monitor SLOs continuously:

```promql
# Availability SLO (target: 99%)
(count(up{job=~".*-service"} == 1) / count(up{job=~".*-service"})) * 100

# Latency SLO (target: <= 200ms)
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m])) < 0.2

# Error Rate SLO (target: <= 1%)
(rate(http_requests_total{status=~"5.."}[5m]) / rate(http_requests_total[5m])) * 100 < 1
```

## Troubleshooting Monitoring

### Prometheus Issues

**Problem**: Targets showing DOWN
- Check network connectivity
- Verify service is running
- Check metrics endpoint (/metrics)
- Review Prometheus logs

**Problem**: No data in Prometheus
- Verify scrape_interval setting
- Check data retention setting
- Confirm targets are UP
- Review service metrics implementation

### Grafana Issues

**Problem**: No data in dashboard
- Verify Prometheus datasource is configured
- Check query syntax
- Confirm metrics exist in Prometheus
- Review time range selection

**Problem**: Alerts not firing
- Verify alert rules are configured
- Check alert evaluation interval
- Review alert condition logic
- Check notification channel settings

## Data Retention

### Prometheus Storage

Default settings:
- **Retention time**: 15 days
- **Retention size**: Unlimited

Modify retention:
```bash
# In docker-compose.yml
prometheus:
  command:
    - '--storage.tsdb.retention.time=30d'
    - '--storage.tsdb.retention.size=50GB'
```

## Monitoring Checklist

- [ ] All services configured in Prometheus
- [ ] All targets showing UP status
- [ ] Grafana dashboards created
- [ ] Alert rules defined
- [ ] Alert notification channels configured
- [ ] Alert thresholds validated
- [ ] SLI/SLO metrics tracked
- [ ] Retention policies set
- [ ] Backup strategy implemented
- [ ] On-call runbooks created

