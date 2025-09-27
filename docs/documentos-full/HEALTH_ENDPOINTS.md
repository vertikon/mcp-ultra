# Health Check Endpoints

## 📋 Visão Geral

O sistema MCP Ultra implementa endpoints de health check abrangentes para monitoramento de aplicações, seguindo as melhores práticas do Kubernetes e microserviços.

## 🏥 Endpoints Disponíveis

### `/health` - Status Detalhado
**Método**: `GET`  
**Descrição**: Retorna status completo com métricas detalhadas de todos os componentes.

```bash
curl http://localhost:8080/health
```

**Resposta de Exemplo**:
```json
{
  "status": "healthy",
  "version": "1.0.0",
  "timestamp": "2025-09-12T18:50:58Z",
  "uptime": "2h45m30s",
  "environment": "production",
  "checks": {
    "database": {
      "name": "postgresql",
      "status": "healthy",
      "message": "Database postgresql is healthy",
      "duration": "15ms",
      "timestamp": "2025-09-12T18:50:58Z",
      "metadata": {
        "connection_pool_size": 10,
        "active_connections": 3
      }
    },
    "redis": {
      "name": "redis", 
      "status": "healthy",
      "message": "Redis is healthy",
      "duration": "5ms",
      "timestamp": "2025-09-12T18:50:58Z",
      "metadata": {
        "connected_clients": 2,
        "used_memory": "1.2MB"
      }
    },
    "nats": {
      "name": "nats",
      "status": "healthy", 
      "message": "NATS is connected",
      "duration": "3ms",
      "timestamp": "2025-09-12T18:50:58Z",
      "metadata": {
        "server_id": "nats-1",
        "connections": 1
      }
    }
  },
  "system": {
    "go_version": "go1.21.0",
    "goroutines": 42,
    "cpu_count": 8,
    "memory": {
      "alloc_bytes": 15728640,
      "total_alloc_bytes": 67108864,
      "sys_bytes": 25165824,
      "gc_count": 5,
      "last_gc": "2025-09-12T18:48:30Z",
      "gc_pause_total_ns": 1500000
    }
  }
}
```

### `/healthz` - Health Check Simples
**Método**: `GET`  
**Descrição**: Health check simples compatível com Kubernetes.

```bash
curl http://localhost:8080/healthz
```

**Respostas**:
- **200 OK**: `OK` (sistema saudável)
- **503 Service Unavailable**: `Service Unavailable` (sistema não saudável)

### `/ready` / `/readyz` - Readiness Check
**Método**: `GET`  
**Descrição**: Verifica se o serviço está pronto para receber tráfego.

```bash
curl http://localhost:8080/ready
curl http://localhost:8080/readyz
```

**Critérios de Prontidão**:
- ✅ Conexão com banco de dados funcional
- ✅ Cache Redis disponível  
- ✅ Message broker NATS conectado

**Respostas**:
- **200 OK**: `Ready` (pronto para tráfego)
- **503 Service Unavailable**: `Service Not Ready` (não pronto)

### `/live` / `/livez` - Liveness Check
**Método**: `GET`  
**Descrição**: Verifica se o serviço está vivo e operacional.

```bash
curl http://localhost:8080/live
curl http://localhost:8080/livez
```

**Resposta**:
- **200 OK**: `Alive` (sempre retorna se o processo estiver rodando)

### `/status` - Status Abrangente
**Método**: `GET`  
**Descrição**: Status detalhado com informações de request e trace.

```bash
curl -H "X-Request-ID: req-123" http://localhost:8080/status
```

**Resposta de Exemplo**:
```json
{
  "status": "healthy",
  "version": "1.0.0",
  "timestamp": "2025-09-12T18:50:58Z",
  "uptime": "2h45m30s",
  "environment": "production",
  "checks": { /* ... mesmos dados do /health ... */ },
  "system": { /* ... informações do sistema ... */ },
  "request_id": "req-123",
  "trace_id": "4bf92f3577b34da6a3ce929d0e0e4736"
}
```

## 📊 Status Types

### Health Status
```go
type HealthStatus string

const (
    StatusHealthy   HealthStatus = "healthy"   // Todos os componentes funcionando
    StatusDegraded  HealthStatus = "degraded"  // Alguns problemas, mas operacional
    StatusUnhealthy HealthStatus = "unhealthy" // Problemas críticos
)
```

### HTTP Status Codes
| Health Status | HTTP Code | Descrição |
|---------------|-----------|-----------|
| `healthy` | 200 | Todos os componentes saudáveis |
| `degraded` | 200 | Operacional com degradação |
| `unhealthy` | 503 | Serviço indisponível |

## ⚙️ Configuração

### Health Checkers Registrados

```go
// Database health checker
healthService.RegisterChecker("database", 
    httphandlers.NewDatabaseHealthChecker("postgresql", func(ctx context.Context) error {
        return db.PingContext(ctx)
    }))

// Redis health checker  
healthService.RegisterChecker("redis",
    httphandlers.NewRedisHealthChecker(func(ctx context.Context) error {
        return redis.Ping(cacheClient)
    }))

// NATS health checker
healthService.RegisterChecker("nats",
    httphandlers.NewNATSHealthChecker(func() bool {
        return eventBus.IsConnected()
    }))
```

### Customização de Health Checkers

```go
// Custom health checker
type CustomHealthChecker struct {
    name    string
    checker func(context.Context) error
}

func (c *CustomHealthChecker) Check(ctx context.Context) HealthCheck {
    start := time.Now()
    check := HealthCheck{
        Name:      c.name,
        Timestamp: start,
    }

    if err := c.checker(ctx); err != nil {
        check.Status = StatusUnhealthy
        check.Error = err.Error()
        check.Message = fmt.Sprintf("Service %s is unhealthy", c.name)
    } else {
        check.Status = StatusHealthy
        check.Message = fmt.Sprintf("Service %s is healthy", c.name)
    }

    check.Duration = time.Since(start)
    return check
}

// Registro
healthService.RegisterChecker("my-service", &CustomHealthChecker{
    name: "my-service",
    checker: func(ctx context.Context) error {
        // Lógica de verificação customizada
        return nil
    },
})
```

## 🐳 Kubernetes Integration

### Deployment Configuration

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mcp-ultra
spec:
  template:
    spec:
      containers:
      - name: mcp-ultra
        image: vertikon/mcp-ultra:latest
        ports:
        - containerPort: 8080
        livenessProbe:
          httpGet:
            path: /livez
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
          timeoutSeconds: 5
          failureThreshold: 3
        readinessProbe:
          httpGet:
            path: /readyz
            port: 8080
          initialDelaySeconds: 10
          periodSeconds: 5
          timeoutSeconds: 3
          failureThreshold: 2
        startupProbe:
          httpGet:
            path: /healthz
            port: 8080
          initialDelaySeconds: 10
          periodSeconds: 10
          timeoutSeconds: 5
          failureThreshold: 10
```

### Service Monitor (Prometheus)

```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: mcp-ultra-health
spec:
  selector:
    matchLabels:
      app: mcp-ultra
  endpoints:
  - port: http
    path: /health
    interval: 30s
    scrapeTimeout: 10s
```

## 📈 Monitoramento

### Métricas Disponíveis

As health checks expõem métricas via Prometheus:

```
# Health check duration
health_check_duration_seconds{service="database",status="healthy"} 0.015
health_check_duration_seconds{service="redis",status="healthy"} 0.005
health_check_duration_seconds{service="nats",status="healthy"} 0.003

# Health check status
health_check_status{service="database"} 1  # 1=healthy, 0=unhealthy
health_check_status{service="redis"} 1
health_check_status{service="nats"} 1

# System metrics
system_goroutines 42
system_memory_alloc_bytes 15728640
system_gc_count 5
```

### Alertas Sugeridos

```yaml
# Prometheus alerting rules
groups:
- name: mcp-ultra-health
  rules:
  - alert: ServiceUnhealthy
    expr: health_check_status == 0
    for: 1m
    labels:
      severity: critical
    annotations:
      summary: "MCP Ultra service {{ $labels.service }} is unhealthy"
      description: "Service {{ $labels.service }} has been unhealthy for more than 1 minute"

  - alert: HighLatencyHealthCheck
    expr: health_check_duration_seconds > 1
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "High latency on health check for {{ $labels.service }}"
      description: "Health check for {{ $labels.service }} is taking {{ $value }}s"
```

## 🔧 Troubleshooting

### Common Issues

#### 1. Database Connection Issues
```bash
curl http://localhost:8080/health | jq '.checks.database'
```

**Possíveis soluções**:
- Verificar string de conexão
- Validar credenciais do banco
- Testar conectividade de rede

#### 2. Redis Connection Issues
```bash
curl http://localhost:8080/health | jq '.checks.redis'
```

**Possíveis soluções**:
- Verificar se Redis está rodando
- Validar configuração de host/porta
- Testar autenticação se configurada

#### 3. NATS Connection Issues
```bash
curl http://localhost:8080/health | jq '.checks.nats'
```

**Possíveis soluções**:
- Verificar se NATS server está rodando
- Validar URL de conexão
- Testar conectividade de rede

### Health Check Logs

```json
{
  "level": "info",
  "timestamp": "2025-09-12T18:50:58Z",
  "message": "Health check completed",
  "status": "healthy",
  "duration": "23ms",
  "checks": 3,
  "service": "mcp-ultra"
}
```

## 🎯 Best Practices

### 1. Timeout Configuration
- Health checks devem ter timeout curto (< 5s)
- Implementar circuit breakers para dependências externas

### 2. Dependency Isolation
- Falhas em dependências não críticas não devem afetar liveness
- Usar diferentes criticidades para readiness vs liveness

### 3. Graceful Degradation  
- Retornar `degraded` quando possível ao invés de `unhealthy`
- Manter funcionalidades essenciais disponíveis

### 4. Monitoring Integration
- Expor métricas detalhadas para observabilidade
- Configurar alertas baseados em SLOs

---

**Atualizado em**: 2025-09-12  
**Versão**: 1.0.0