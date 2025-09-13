# Observability & Monitoring - MCP Ultra

## 📋 Visão Geral

O MCP Ultra implementa um sistema de observabilidade completo baseado em OpenTelemetry, fornecendo **tracing distribuído**, **métricas de negócio** e **logging estruturado** para monitoramento e debugging em produção.

## 🔍 Recursos de Observabilidade

### ✅ Implementações
- **📊 Métricas Prometheus** - Métricas de aplicação e infraestrutura
- **🔍 Tracing Distribuído** - Rastreamento de requests end-to-end
- **📝 Logging Estruturado** - Logs centralizados com correlação
- **🏥 Health Monitoring** - Monitoramento de saúde da aplicação
- **⚡ Performance Metrics** - Latência, throughput e disponibilidade
- **🛡️ Security Monitoring** - Auditoria e segurança

## 📊 Métricas Disponíveis

### Métricas de Aplicação

#### HTTP Metrics
```
# Request counter
http_requests_total{method="GET", endpoint="/api/tasks", status="200"} 1500

# Request duration histogram  
http_request_duration_seconds{method="GET", endpoint="/api/tasks"} 0.125

# Response size histogram
http_response_size_bytes{method="GET", endpoint="/api/tasks"} 2048
```

#### Business Metrics
```
# Task operations
tasks_created_total{priority="high", assignee="team-a"} 45
tasks_completed_total{priority="high", assignee="team-a"} 38
task_processing_duration_seconds{priority="high"} 3600

# Error tracking
errors_total{component="database", type="connection_timeout"} 3
errors_total{component="external_api", type="rate_limit"} 12
```

#### System Metrics
```
# Runtime metrics
go_goroutines 42
go_memstats_alloc_bytes 15728640
go_memstats_gc_duration_seconds{quantile="0.5"} 0.000123

# Custom metrics
cache_hit_ratio{cache="redis"} 0.85
connection_pool_active{service="database"} 8
connection_pool_idle{service="database"} 2
```

### Configuração de Métricas

```yaml
# config/telemetry.yaml
telemetry:
  service_name: "mcp-ultra"
  service_version: "1.0.0"
  environment: "production"
  enabled: true
  
  metrics:
    enabled: true
    port: 8080
    path: "/metrics"
    push_gateway: "http://pushgateway:9091"
    push_interval: "30s"
    
  tracing:
    enabled: true
    sampling_rate: 1.0  # Sample all traces in dev
    jaeger_endpoint: "http://jaeger:14268/api/traces"
    otlp_endpoint: "http://otel-collector:4318"
    batch_timeout: "5s"
    batch_size: 100
    
  logging:
    level: "info"
    format: "json"
    output: "stdout"
    correlation_enabled: true
```

## 🔍 Tracing Distribuído

### Instrumentação Automática

```go
// HTTP middleware automático
func main() {
    telemetryService, _ := observability.NewTelemetryService(config, logger)
    telemetryService.Start(ctx)
    
    // Router com instrumentação automática
    router := httphandlers.NewRouter(taskService, flagManager, healthService, logger)
    instrumentedRouter := telemetryService.HTTPMiddleware()(router)
    
    http.ListenAndServe(":8080", instrumentedRouter)
}
```

### Tracing Manual

```go
func processTask(ctx context.Context, taskID uuid.UUID) error {
    tracer := otel.Tracer("task-processor")
    
    // Create span
    ctx, span := tracer.Start(ctx, "process_task",
        trace.WithAttributes(
            attribute.String("task.id", taskID.String()),
            attribute.String("task.type", "data_processing"),
        ),
    )
    defer span.End()
    
    // Database operation (automatically instrumented)
    task, err := taskRepo.GetByID(ctx, taskID)
    if err != nil {
        span.RecordError(err)
        span.SetStatus(codes.Error, "Failed to fetch task")
        return err
    }
    
    // Add task details to span
    span.SetAttributes(
        attribute.String("task.title", task.Title),
        attribute.String("task.status", string(task.Status)),
        attribute.String("task.priority", string(task.Priority)),
    )
    
    // Child span for business logic
    err = processTaskLogic(ctx, task)
    if err != nil {
        span.RecordError(err)
        return err
    }
    
    span.SetStatus(codes.Ok, "Task processed successfully")
    return nil
}

func processTaskLogic(ctx context.Context, task *domain.Task) error {
    tracer := otel.Tracer("task-processor")
    
    ctx, span := tracer.Start(ctx, "task_business_logic",
        trace.WithAttributes(
            attribute.String("operation", "validate_and_execute"),
        ),
    )
    defer span.End()
    
    // Business logic here...
    time.Sleep(100 * time.Millisecond) // Simulate processing
    
    return nil
}
```

### Trace Context Propagation

```go
// Propagação via HTTP headers
func makeExternalRequest(ctx context.Context, url string) error {
    req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
    
    // Inject trace context
    otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    return nil
}
```

### Exemplo de Trace

```json
{
  "trace_id": "4bf92f3577b34da6a3ce929d0e0e4736",
  "span_id": "00f067aa0ba902b7",
  "operation_name": "POST /api/tasks",
  "start_time": "2025-09-12T18:50:58.123Z",
  "duration": "125ms",
  "tags": {
    "http.method": "POST",
    "http.url": "/api/tasks",
    "http.status_code": 201,
    "user.id": "user-789",
    "task.priority": "high"
  },
  "spans": [
    {
      "span_id": "0192837465abcdef",
      "parent_id": "00f067aa0ba902b7",
      "operation_name": "validate_task_data",
      "duration": "15ms",
      "tags": {
        "validation.result": "success"
      }
    },
    {
      "span_id": "abcdef0192837465", 
      "parent_id": "00f067aa0ba902b7",
      "operation_name": "database.insert",
      "duration": "85ms",
      "tags": {
        "db.statement": "INSERT INTO tasks...",
        "db.connection_id": "conn-123"
      }
    },
    {
      "span_id": "1234567890abcdef",
      "parent_id": "00f067aa0ba902b7", 
      "operation_name": "publish_event",
      "duration": "25ms",
      "tags": {
        "event.type": "task.created",
        "event.bus": "nats"
      }
    }
  ]
}
```

## 📝 Logging Estruturado

### Configuração de Logs

```go
func setupLogging(config LoggingConfig) (*zap.Logger, error) {
    var zapConfig zap.Config
    
    if config.Environment == "production" {
        zapConfig = zap.NewProductionConfig()
    } else {
        zapConfig = zap.NewDevelopmentConfig()
    }
    
    zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
    zapConfig.OutputPaths = []string{"stdout"}
    zapConfig.ErrorOutputPaths = []string{"stderr"}
    
    // Add correlation fields
    zapConfig.InitialFields = map[string]interface{}{
        "service":     config.ServiceName,
        "version":     config.Version,
        "environment": config.Environment,
    }
    
    return zapConfig.Build()
}
```

### Correlação de Logs

```go
func correlatedLogger(ctx context.Context) *zap.Logger {
    logger := zap.L()
    
    // Extract trace context
    span := trace.SpanFromContext(ctx)
    if span.SpanContext().IsValid() {
        logger = logger.With(
            zap.String("trace_id", span.SpanContext().TraceID().String()),
            zap.String("span_id", span.SpanContext().SpanID().String()),
        )
    }
    
    // Extract user context
    if userID := getUserID(ctx); userID != "" {
        logger = logger.With(zap.String("user_id", userID))
    }
    
    // Extract request context
    if requestID := getRequestID(ctx); requestID != "" {
        logger = logger.With(zap.String("request_id", requestID))
    }
    
    return logger
}
```

### Exemplo de Log Estruturado

```json
{
  "timestamp": "2025-09-12T18:50:58.123Z",
  "level": "info",
  "message": "Task created successfully",
  "service": "mcp-ultra",
  "version": "1.0.0",
  "environment": "production",
  "trace_id": "4bf92f3577b34da6a3ce929d0e0e4736",
  "span_id": "00f067aa0ba902b7",
  "user_id": "user-789",
  "request_id": "req-456",
  "task": {
    "id": "task-123",
    "title": "Process user data",
    "priority": "high",
    "assignee": "team-a"
  },
  "performance": {
    "duration_ms": 125,
    "database_calls": 3,
    "cache_hits": 2
  }
}
```

## 📊 Dashboards e Alertas

### Grafana Dashboard

```json
{
  "dashboard": {
    "title": "MCP Ultra - Application Metrics",
    "panels": [
      {
        "title": "Request Rate",
        "type": "graph",
        "targets": [
          {
            "expr": "rate(http_requests_total[5m])",
            "legendFormat": "{{method}} {{endpoint}}"
          }
        ]
      },
      {
        "title": "Response Time P95",
        "type": "graph", 
        "targets": [
          {
            "expr": "histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))",
            "legendFormat": "P95 Latency"
          }
        ]
      },
      {
        "title": "Error Rate",
        "type": "singlestat",
        "targets": [
          {
            "expr": "rate(http_requests_total{status=~\"5..\"}[5m]) / rate(http_requests_total[5m]) * 100",
            "legendFormat": "Error Rate %"
          }
        ]
      },
      {
        "title": "Business Metrics",
        "type": "graph",
        "targets": [
          {
            "expr": "rate(tasks_created_total[5m])",
            "legendFormat": "Tasks Created/sec"
          },
          {
            "expr": "rate(tasks_completed_total[5m])", 
            "legendFormat": "Tasks Completed/sec"
          }
        ]
      }
    ]
  }
}
```

### Prometheus Alerting Rules

```yaml
# alerts.yml
groups:
- name: mcp-ultra-application
  rules:
  - alert: HighErrorRate
    expr: (rate(http_requests_total{status=~"5.."}[5m]) / rate(http_requests_total[5m])) > 0.05
    for: 2m
    labels:
      severity: critical
    annotations:
      summary: "High error rate detected"
      description: "Error rate is {{ $value | humanizePercentage }} for the last 5 minutes"

  - alert: HighLatency
    expr: histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m])) > 1.0
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "High request latency"
      description: "95th percentile latency is {{ $value }}s"

  - alert: DatabaseConnectionPoolExhausted
    expr: connection_pool_active{service="database"} / connection_pool_max{service="database"} > 0.9
    for: 1m
    labels:
      severity: warning
    annotations:
      summary: "Database connection pool nearly exhausted"
      description: "Connection pool usage is at {{ $value | humanizePercentage }}"

- name: mcp-ultra-business
  rules:
  - alert: TaskProcessingBacklog
    expr: tasks_created_total - tasks_completed_total > 1000
    for: 10m
    labels:
      severity: warning
    annotations:
      summary: "Task processing backlog detected"
      description: "{{ $value }} tasks in backlog"

  - alert: ComplianceViolation
    expr: increase(compliance_violations_total[1h]) > 0
    for: 0s
    labels:
      severity: critical
    annotations:
      summary: "Compliance violation detected"
      description: "{{ $value }} compliance violations in the last hour"
```

## 🔧 Configuração de Monitoramento

### Docker Compose - Stack de Observabilidade

```yaml
# docker-compose.observability.yml
version: '3.8'

services:
  # Prometheus
  prometheus:
    image: prom/prometheus:latest
    ports:
      - "9090:9090"
    volumes:
      - ./config/prometheus.yml:/etc/prometheus/prometheus.yml
      - ./config/alerts.yml:/etc/prometheus/alerts.yml
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--web.console.libraries=/etc/prometheus/console_libraries'
      - '--web.console.templates=/etc/prometheus/consoles'
      - '--web.enable-lifecycle'

  # Grafana
  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin123
    volumes:
      - grafana-storage:/var/lib/grafana
      - ./config/grafana/dashboards:/etc/grafana/provisioning/dashboards
      - ./config/grafana/datasources:/etc/grafana/provisioning/datasources

  # Jaeger
  jaeger:
    image: jaegertracing/all-in-one:latest
    ports:
      - "16686:16686"
      - "14268:14268"
    environment:
      - COLLECTOR_OTLP_ENABLED=true

  # OpenTelemetry Collector
  otel-collector:
    image: otel/opentelemetry-collector:latest
    ports:
      - "4317:4317"   # OTLP gRPC
      - "4318:4318"   # OTLP HTTP
      - "8888:8888"   # Prometheus metrics
    volumes:
      - ./config/otel-collector.yml:/etc/otel-collector-config.yml
    command: ["--config=/etc/otel-collector-config.yml"]

  # AlertManager
  alertmanager:
    image: prom/alertmanager:latest
    ports:
      - "9093:9093"
    volumes:
      - ./config/alertmanager.yml:/etc/alertmanager/alertmanager.yml

volumes:
  grafana-storage:
```

### Prometheus Configuration

```yaml
# config/prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

rule_files:
  - "alerts.yml"

alerting:
  alertmanagers:
    - static_configs:
        - targets:
          - alertmanager:9093

scrape_configs:
  - job_name: 'mcp-ultra'
    static_configs:
      - targets: ['mcp-ultra:8080']
    metrics_path: /metrics
    scrape_interval: 15s

  - job_name: 'mcp-ultra-health'
    static_configs:
      - targets: ['mcp-ultra:8080']
    metrics_path: /health
    scrape_interval: 30s

  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']
```

### OpenTelemetry Collector Config

```yaml
# config/otel-collector.yml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317
      http:
        endpoint: 0.0.0.0:4318

processors:
  batch:
    timeout: 5s
    send_batch_size: 100

  resource:
    attributes:
      - key: service.instance.id
        from_attribute: host.name
        action: insert

exporters:
  jaeger:
    endpoint: jaeger:14250
    tls:
      insecure: true

  prometheus:
    endpoint: "0.0.0.0:8888"
    const_labels:
      env: "production"

  logging:
    loglevel: debug

service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [resource, batch]
      exporters: [jaeger, logging]

    metrics:
      receivers: [otlp]
      processors: [resource, batch]
      exporters: [prometheus, logging]
```

## 🚨 Alertas e Notificações

### AlertManager Configuration

```yaml
# config/alertmanager.yml
global:
  smtp_smarthost: 'localhost:587'
  smtp_from: 'alerts@company.com'

route:
  group_by: ['alertname', 'cluster', 'service']
  group_wait: 30s
  group_interval: 5m
  repeat_interval: 12h
  receiver: 'web.hook'
  routes:
  - match:
      severity: critical
    receiver: 'critical-alerts'
  - match:
      alertname: ComplianceViolation
    receiver: 'compliance-team'

receivers:
- name: 'web.hook'
  webhook_configs:
  - url: 'http://slack-webhook:3000/alerts'

- name: 'critical-alerts'
  email_configs:
  - to: 'sre-team@company.com'
    subject: 'CRITICAL: {{ .GroupLabels.alertname }}'
    body: |
      {{ range .Alerts }}
      Alert: {{ .Annotations.summary }}
      Description: {{ .Annotations.description }}
      {{ end }}

- name: 'compliance-team'
  email_configs:
  - to: 'dpo@company.com'
    subject: 'COMPLIANCE VIOLATION: {{ .GroupLabels.alertname }}'
```

### Slack Integration

```go
// Webhook para notificações Slack
func sendSlackAlert(alert Alert) error {
    payload := SlackPayload{
        Text: fmt.Sprintf("🚨 *%s*", alert.Summary),
        Attachments: []SlackAttachment{
            {
                Color: getColorBySeverity(alert.Severity),
                Fields: []SlackField{
                    {
                        Title: "Service",
                        Value: alert.Service,
                        Short: true,
                    },
                    {
                        Title: "Severity", 
                        Value: alert.Severity,
                        Short: true,
                    },
                    {
                        Title: "Description",
                        Value: alert.Description,
                        Short: false,
                    },
                },
            },
        },
    }
    
    return sendToSlack(payload)
}
```

## 📈 Performance Monitoring

### SLI/SLO Configuration

```yaml
# Service Level Objectives
slos:
  availability:
    target: 99.9%
    measurement: "sum(rate(http_requests_total{status!~'5..'}[5m])) / sum(rate(http_requests_total[5m]))"
    
  latency:
    target: 95%  # 95% of requests under 200ms
    measurement: "histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m])) < 0.2"
    
  error_rate:
    target: 99%  # Less than 1% error rate
    measurement: "sum(rate(http_requests_total{status!~'5..'}[5m])) / sum(rate(http_requests_total[5m]))"
```

### Custom Metrics

```go
// Business metrics
func recordBusinessMetrics(ctx context.Context, telemetry *observability.TelemetryService) {
    // Task metrics
    telemetry.IncrementRequestCounter(ctx, "POST", "/api/tasks", "201")
    telemetry.RecordRequestDuration(ctx, "POST", "/api/tasks", 125*time.Millisecond)
    
    // Processing metrics
    telemetry.RecordProcessingTime(ctx, "task_validation", 15*time.Millisecond)
    telemetry.RecordProcessingTime(ctx, "database_insert", 85*time.Millisecond)
    
    // Error tracking
    if err != nil {
        telemetry.IncrementErrorCounter(ctx, "database", "connection_timeout")
    }
    
    // Custom gauge
    meter := telemetry.GetMeter("business-metrics")
    queueSize, _ := meter.Int64UpDownCounter(
        "task_queue_size",
        metric.WithDescription("Number of tasks in processing queue"),
    )
    queueSize.Add(ctx, 1)
}
```

## 🔍 Debugging e Troubleshooting

### Log Correlation

```bash
# Buscar logs por trace ID
kubectl logs -f deployment/mcp-ultra | jq 'select(.trace_id == "4bf92f3577b34da6a3ce929d0e0e4736")'

# Buscar logs por user ID
kubectl logs -f deployment/mcp-ultra | jq 'select(.user_id == "user-789")'

# Buscar erros em um período
kubectl logs --since=1h deployment/mcp-ultra | jq 'select(.level == "error")'
```

### Trace Analysis

```bash
# Jaeger query by operation
curl "http://jaeger:16686/api/traces?service=mcp-ultra&operation=POST%20/api/tasks&limit=10"

# Find slow traces
curl "http://jaeger:16686/api/traces?service=mcp-ultra&minDuration=1s"
```

### Metrics Queries

```promql
# Top endpoints by traffic
topk(10, sum(rate(http_requests_total[5m])) by (endpoint))

# Slowest endpoints
topk(10, histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m])) by (endpoint))

# Error rate by endpoint
topk(10, sum(rate(http_requests_total{status=~"5.."}[5m])) by (endpoint))
```

---

**Atualizado em**: 2025-09-12  
**Versão**: 1.0.0  
**OpenTelemetry**: v1.21.0