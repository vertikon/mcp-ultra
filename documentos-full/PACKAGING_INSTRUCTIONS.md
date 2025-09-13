# 📦 MCP Ultra - Packaging Instructions

## Overview

This guide provides comprehensive instructions for packaging the MCP Ultra application for different environments including Docker containers, Kubernetes deployments, Helm charts, binary distributions, and cloud-native deployments. The packaging system supports multiple architectures and deployment strategies.

## Table of Contents

1. [Docker Packaging](#docker-packaging)
2. [Kubernetes Deployment](#kubernetes-deployment)
3. [Helm Charts](#helm-charts)
4. [Binary Distribution](#binary-distribution)
5. [Cloud Provider Packaging](#cloud-provider-packaging)
6. [CI/CD Integration](#cicd-integration)
7. [Security Considerations](#security-considerations)
8. [Troubleshooting](#troubleshooting)

## Docker Packaging

### 1. Multi-Stage Docker Build

#### Production-Ready Dockerfile
```dockerfile
# MCP Ultra Multi-stage Dockerfile
# Build stage
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git make ca-certificates tzdata

WORKDIR /build

# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -X 'github.com/vertikon/mcp-ultra/pkg/version.Version=${VERSION:-v1.0.0}' -X 'github.com/vertikon/mcp-ultra/pkg/version.GitCommit=$(git rev-parse HEAD 2>/dev/null || echo 'unknown')' -X 'github.com/vertikon/mcp-ultra/pkg/version.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)'" \
    -o mcp-ultra cmd/mcp-model-ultra/main.go

# Runtime stage
FROM alpine:3.19

# Install required packages and create user
RUN apk add --no-cache ca-certificates tzdata wget curl && \
    adduser -D -g '' appuser

WORKDIR /app

# Copy binary and configuration
COPY --from=builder /build/mcp-ultra .
COPY --from=builder /build/config ./config
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set ownership
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose ports
EXPOSE 9655 9656

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:9655/healthz || exit 1

# Set entrypoint
ENTRYPOINT ["./mcp-ultra"]
```

### 2. Multi-Architecture Builds

#### Build for Multiple Platforms
```bash
#!/bin/bash
# scripts/docker-build-multiarch.sh

set -e

IMAGE_NAME="ghcr.io/vertikon/mcp-ultra"
VERSION=${1:-latest}
PLATFORMS="linux/amd64,linux/arm64"

echo "🐳 Building multi-architecture Docker images..."

# Create and use buildx builder
docker buildx create --name mcp-ultra-builder --use || true
docker buildx inspect --bootstrap

# Build and push multi-architecture images
docker buildx build \
    --platform $PLATFORMS \
    --tag $IMAGE_NAME:$VERSION \
    --tag $IMAGE_NAME:latest \
    --build-arg VERSION=$VERSION \
    --build-arg BUILD_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
    --build-arg VCS_REF=$(git rev-parse HEAD) \
    --push \
    .

echo "✅ Multi-architecture images built and pushed successfully"

# Clean up builder
docker buildx rm mcp-ultra-builder
```

### 3. Development Docker Compose

#### Complete Development Environment
```yaml
# docker-compose.yml
version: '3.8'

services:
  mcp-ultra:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "9655:9655"  # HTTP API
      - "9656:9656"  # Metrics
    environment:
      - APP_ENV=development
      - LOG_LEVEL=debug
      - DATABASE_URL=postgres://postgres:password@postgres:5432/mcp_ultra?sslmode=disable
      - REDIS_URL=redis://redis:6379/0
      - NATS_URL=nats://nats:4222
      - OTEL_EXPORTER_OTLP_TRACES_ENDPOINT=http://jaeger:4318/v1/traces
      - OTEL_EXPORTER_PROMETHEUS_PORT=9656
    depends_on:
      postgres:
        condition: service_healthy
      redis:
        condition: service_healthy
      nats:
        condition: service_started
    networks:
      - mcp-network
    restart: unless-stopped

  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: mcp_ultra
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./deploy/database/init.sql:/docker-entrypoint-initdb.d/init.sql
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - mcp-network
    restart: unless-stopped

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
      - ./deploy/redis/redis.conf:/usr/local/etc/redis/redis.conf
    command: redis-server /usr/local/etc/redis/redis.conf
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - mcp-network
    restart: unless-stopped

  nats:
    image: nats:2-alpine
    ports:
      - "4222:4222"
      - "8222:8222"  # HTTP monitoring
    command: ["-js", "-m", "8222"]
    networks:
      - mcp-network
    restart: unless-stopped

  # Observability Stack
  prometheus:
    image: prom/prometheus:latest
    ports:
      - "9090:9090"
    volumes:
      - ./deploy/monitoring/prometheus.yml:/etc/prometheus/prometheus.yml
      - prometheus_data:/prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--web.console.libraries=/etc/prometheus/console_libraries'
      - '--web.console.templates=/etc/prometheus/consoles'
      - '--web.enable-lifecycle'
    networks:
      - mcp-network
    restart: unless-stopped

  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    environment:
      GF_SECURITY_ADMIN_PASSWORD: admin
    volumes:
      - grafana_data:/var/lib/grafana
      - ./deploy/monitoring/grafana/dashboards:/etc/grafana/provisioning/dashboards
      - ./deploy/monitoring/grafana/datasources:/etc/grafana/provisioning/datasources
    networks:
      - mcp-network
    restart: unless-stopped

  jaeger:
    image: jaegertracing/all-in-one:latest
    ports:
      - "16686:16686"  # Jaeger UI
      - "14268:14268"  # Jaeger collector
      - "4317:4317"    # OTLP gRPC
      - "4318:4318"    # OTLP HTTP
    environment:
      COLLECTOR_OTLP_ENABLED: true
    networks:
      - mcp-network
    restart: unless-stopped

volumes:
  postgres_data:
  redis_data:
  prometheus_data:
  grafana_data:

networks:
  mcp-network:
    driver: bridge
```

## Kubernetes Deployment

### 1. Base Kubernetes Manifests

#### Deployment Configuration
```yaml
# deploy/k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mcp-ultra
  labels:
    app: mcp-ultra
    version: v1.0.0
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  selector:
    matchLabels:
      app: mcp-ultra
  template:
    metadata:
      labels:
        app: mcp-ultra
        version: v1.0.0
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "9656"
        prometheus.io/path: "/metrics"
    spec:
      serviceAccountName: mcp-ultra
      securityContext:
        runAsNonRoot: true
        runAsUser: 1000
        fsGroup: 1000
      containers:
      - name: mcp-ultra
        image: ghcr.io/vertikon/mcp-ultra:v1.0.0
        imagePullPolicy: Always
        ports:
        - name: http
          containerPort: 9655
          protocol: TCP
        - name: metrics
          containerPort: 9656
          protocol: TCP
        env:
        - name: APP_ENV
          value: "production"
        - name: LOG_LEVEL
          value: "info"
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: mcp-ultra-secrets
              key: database-url
        - name: REDIS_URL
          valueFrom:
            secretKeyRef:
              name: mcp-ultra-secrets
              key: redis-url
        - name: NATS_URL
          valueFrom:
            secretKeyRef:
              name: mcp-ultra-secrets
              key: nats-url
        - name: JWT_SECRET
          valueFrom:
            secretKeyRef:
              name: mcp-ultra-secrets
              key: jwt-secret
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 500m
            memory: 512Mi
        livenessProbe:
          httpGet:
            path: /live
            port: http
          initialDelaySeconds: 30
          periodSeconds: 10
          timeoutSeconds: 5
          failureThreshold: 3
        readinessProbe:
          httpGet:
            path: /ready
            port: http
          initialDelaySeconds: 5
          periodSeconds: 5
          timeoutSeconds: 3
          failureThreshold: 3
        startupProbe:
          httpGet:
            path: /healthz
            port: http
          initialDelaySeconds: 10
          periodSeconds: 10
          timeoutSeconds: 3
          failureThreshold: 30
        volumeMounts:
        - name: config
          mountPath: /app/config
          readOnly: true
        - name: tls-certs
          mountPath: /etc/ssl/certs/mcp-ultra
          readOnly: true
      volumes:
      - name: config
        configMap:
          name: mcp-ultra-config
      - name: tls-certs
        secret:
          secretName: mcp-ultra-tls
          defaultMode: 0400
---
apiVersion: v1
kind: Service
metadata:
  name: mcp-ultra
  labels:
    app: mcp-ultra
spec:
  type: ClusterIP
  ports:
  - port: 80
    targetPort: http
    protocol: TCP
    name: http
  - port: 9656
    targetPort: metrics
    protocol: TCP
    name: metrics
  selector:
    app: mcp-ultra
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: mcp-ultra
  labels:
    app: mcp-ultra
automountServiceAccountToken: true
```

#### ConfigMap and Secrets
```yaml
# deploy/k8s/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mcp-ultra-config
  labels:
    app: mcp-ultra
data:
  config.yaml: |
    server:
      host: "0.0.0.0"
      port: 9655
      read_timeout: 30s
      write_timeout: 30s
      shutdown_timeout: 30s
    
    metrics:
      port: 9656
      path: "/metrics"
      
    observability:
      tracing:
        enabled: true
        sampler_ratio: 0.01
        exporter: "otlp"
      metrics:
        enabled: true
        namespace: "mcp_ultra"
        
    security:
      tls:
        enabled: true
        cert_file: "/etc/ssl/certs/mcp-ultra/tls.crt"
        key_file: "/etc/ssl/certs/mcp-ultra/tls.key"
      auth:
        jwt:
          expiration: "24h"
          refresh_expiration: "168h"
        rate_limiting:
          requests: 1000
          window: "1h"
          burst: 100
          
    compliance:
      lgpd:
        enabled: true
        data_retention_days: 2555
        anonymization_enabled: true
        
---
apiVersion: v1
kind: Secret
metadata:
  name: mcp-ultra-secrets
  labels:
    app: mcp-ultra
type: Opaque
data:
  database-url: cG9zdGdyZXM6Ly9wb3N0Z3JlczpwYXNzd29yZEBwb3N0Z3JlczppNTQzMi9tY3BfdWx0cmE/c3NsbW9kZT1kaXNhYmxl
  redis-url: cmVkaXM6Ly9yZWRpczoyNjM3OS8w
  nats-url: bmF0czovL25hdHM6NDIyMg==
  jwt-secret: c3VwZXJfc2VjcmV0X2p3dF9rZXlfaGVyZQ==
  api-key: YXBpX2tleV9mb3JfYXV0aGVudGljYXRpb24=
```

### 2. Advanced Kubernetes Resources

#### Horizontal Pod Autoscaler
```yaml
# deploy/k8s/hpa.yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: mcp-ultra-hpa
  labels:
    app: mcp-ultra
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: mcp-ultra
  minReplicas: 3
  maxReplicas: 50
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
  - type: Pods
    pods:
      metric:
        name: http_requests_per_second
      target:
        type: AverageValue
        averageValue: "100"
  behavior:
    scaleUp:
      stabilizationWindowSeconds: 60
      policies:
      - type: Percent
        value: 50
        periodSeconds: 60
      - type: Pods
        value: 2
        periodSeconds: 60
    scaleDown:
      stabilizationWindowSeconds: 300
      policies:
      - type: Percent
        value: 10
        periodSeconds: 60
```

#### Network Policy
```yaml
# deploy/k8s/network-policy.yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: mcp-ultra-network-policy
  labels:
    app: mcp-ultra
spec:
  podSelector:
    matchLabels:
      app: mcp-ultra
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          name: ingress-nginx
    - namespaceSelector:
        matchLabels:
          name: monitoring
    ports:
    - protocol: TCP
      port: 9655
    - protocol: TCP
      port: 9656
  egress:
  # Allow DNS
  - to: []
    ports:
    - protocol: UDP
      port: 53
    - protocol: TCP
      port: 53
  # Allow database access
  - to:
    - namespaceSelector:
        matchLabels:
          name: database
    ports:
    - protocol: TCP
      port: 5432
  # Allow Redis access
  - to:
    - namespaceSelector:
        matchLabels:
          name: cache
    ports:
    - protocol: TCP
      port: 6379
  # Allow NATS access
  - to:
    - namespaceSelector:
        matchLabels:
          name: messaging
    ports:
    - protocol: TCP
      port: 4222
  # Allow external HTTP/HTTPS
  - to: []
    ports:
    - protocol: TCP
      port: 80
    - protocol: TCP
      port: 443
```

#### Pod Disruption Budget
```yaml
# deploy/k8s/pdb.yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: mcp-ultra-pdb
  labels:
    app: mcp-ultra
spec:
  minAvailable: 2
  selector:
    matchLabels:
      app: mcp-ultra
```

## Helm Charts

### 1. Helm Chart Structure

#### Chart Metadata
```yaml
# deploy/helm/mcp-ultra/Chart.yaml
apiVersion: v2
name: mcp-ultra
description: Enterprise-grade Go microservice template with health checks, observability, and compliance
type: application
version: 1.0.0
appVersion: "1.0.0"
home: https://github.com/vertikon/mcp-ultra
sources:
- https://github.com/vertikon/mcp-ultra
maintainers:
- name: Vertikon Team
  email: devops@vertikon.com
keywords:
- microservice
- go
- health-checks
- observability
- compliance
- enterprise
dependencies:
- name: postgresql
  version: "12.x.x"
  repository: https://charts.bitnami.com/bitnami
  condition: postgresql.enabled
- name: redis
  version: "17.x.x"
  repository: https://charts.bitnami.com/bitnami
  condition: redis.enabled
- name: prometheus
  version: "23.x.x"
  repository: https://prometheus-community.github.io/helm-charts
  condition: monitoring.prometheus.enabled
```

#### Values Template
```yaml
# deploy/helm/mcp-ultra/values.yaml
# Global configuration
global:
  imageRegistry: "ghcr.io"
  imagePullSecrets: []

# Application configuration
image:
  repository: vertikon/mcp-ultra
  tag: "v1.0.0"
  pullPolicy: Always

# Deployment configuration
replicaCount: 3

strategy:
  type: RollingUpdate
  rollingUpdate:
    maxSurge: 1
    maxUnavailable: 0

# Service configuration
service:
  type: ClusterIP
  port: 80
  targetPort: http
  annotations: {}

# Ingress configuration
ingress:
  enabled: true
  className: "nginx"
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
    cert-manager.io/cluster-issuer: "letsencrypt-prod"
  hosts:
  - host: mcp-ultra.example.com
    paths:
    - path: /
      pathType: Prefix
  tls:
  - secretName: mcp-ultra-tls
    hosts:
    - mcp-ultra.example.com

# Resource configuration
resources:
  requests:
    cpu: 100m
    memory: 128Mi
  limits:
    cpu: 500m
    memory: 512Mi

# Autoscaling configuration
autoscaling:
  enabled: true
  minReplicas: 3
  maxReplicas: 50
  targetCPUUtilizationPercentage: 70
  targetMemoryUtilizationPercentage: 80
  behavior:
    scaleUp:
      stabilizationWindowSeconds: 60
      policies:
      - type: Percent
        value: 50
        periodSeconds: 60
    scaleDown:
      stabilizationWindowSeconds: 300
      policies:
      - type: Percent
        value: 10
        periodSeconds: 60

# Pod Disruption Budget
podDisruptionBudget:
  enabled: true
  minAvailable: 2

# Security configuration
security:
  podSecurityContext:
    runAsNonRoot: true
    runAsUser: 1000
    fsGroup: 1000
  securityContext:
    allowPrivilegeEscalation: false
    readOnlyRootFilesystem: true
    capabilities:
      drop:
      - ALL

# Network Policy
networkPolicy:
  enabled: true
  policyTypes:
  - Ingress
  - Egress

# Configuration
config:
  server:
    host: "0.0.0.0"
    port: 9655
    readTimeout: "30s"
    writeTimeout: "30s"
    shutdownTimeout: "30s"
  
  observability:
    tracing:
      enabled: true
      samplerRatio: 0.01
      exporter: "otlp"
    metrics:
      enabled: true
      namespace: "mcp_ultra"
  
  security:
    tls:
      enabled: true
    auth:
      jwt:
        expiration: "24h"
        refreshExpiration: "168h"
      rateLimiting:
        requests: 1000
        window: "1h"
        burst: 100
  
  compliance:
    lgpd:
      enabled: true
      dataRetentionDays: 2555
      anonymizationEnabled: true

# Secrets
secrets:
  databaseUrl: "postgres://postgres:password@postgres:5432/mcp_ultra?sslmode=disable"
  redisUrl: "redis://redis:6379/0"
  natsUrl: "nats://nats:4222"
  jwtSecret: "super_secret_jwt_key_here"
  apiKey: "api_key_for_authentication"

# Dependencies
postgresql:
  enabled: true
  auth:
    postgresPassword: "password"
    database: "mcp_ultra"
  primary:
    persistence:
      enabled: true
      size: 10Gi

redis:
  enabled: true
  auth:
    enabled: false
  master:
    persistence:
      enabled: true
      size: 5Gi

# Monitoring
monitoring:
  prometheus:
    enabled: true
    serviceMonitor:
      enabled: true
      interval: 30s
      path: /metrics
      port: metrics
  grafana:
    enabled: true
    dashboards:
      enabled: true

# Logging
logging:
  level: "info"
  format: "json"
```

#### Deployment Template
```yaml
# deploy/helm/mcp-ultra/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "mcp-ultra.fullname" . }}
  labels:
    {{- include "mcp-ultra.labels" . | nindent 4 }}
spec:
  {{- if not .Values.autoscaling.enabled }}
  replicas: {{ .Values.replicaCount }}
  {{- end }}
  strategy:
    {{- toYaml .Values.strategy | nindent 4 }}
  selector:
    matchLabels:
      {{- include "mcp-ultra.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      annotations:
        checksum/config: {{ include (print $.Template.BasePath "/configmap.yaml") . | sha256sum }}
        checksum/secret: {{ include (print $.Template.BasePath "/secret.yaml") . | sha256sum }}
        {{- with .Values.podAnnotations }}
        {{- toYaml . | nindent 8 }}
        {{- end }}
      labels:
        {{- include "mcp-ultra.selectorLabels" . | nindent 8 }}
    spec:
      {{- with .Values.global.imagePullSecrets }}
      imagePullSecrets:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      serviceAccountName: {{ include "mcp-ultra.serviceAccountName" . }}
      securityContext:
        {{- toYaml .Values.security.podSecurityContext | nindent 8 }}
      containers:
        - name: {{ .Chart.Name }}
          securityContext:
            {{- toYaml .Values.security.securityContext | nindent 12 }}
          image: "{{ .Values.global.imageRegistry | default "ghcr.io" }}/{{ .Values.image.repository }}:{{ .Values.image.tag | default .Chart.AppVersion }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          ports:
            - name: http
              containerPort: {{ .Values.config.server.port }}
              protocol: TCP
            - name: metrics
              containerPort: 9656
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /live
              port: http
            initialDelaySeconds: 30
            periodSeconds: 10
            timeoutSeconds: 5
            failureThreshold: 3
          readinessProbe:
            httpGet:
              path: /ready
              port: http
            initialDelaySeconds: 5
            periodSeconds: 5
            timeoutSeconds: 3
            failureThreshold: 3
          startupProbe:
            httpGet:
              path: /healthz
              port: http
            initialDelaySeconds: 10
            periodSeconds: 10
            timeoutSeconds: 3
            failureThreshold: 30
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
          env:
            - name: APP_ENV
              value: "production"
            - name: LOG_LEVEL
              value: {{ .Values.logging.level | quote }}
            - name: LOG_FORMAT
              value: {{ .Values.logging.format | quote }}
            - name: DATABASE_URL
              valueFrom:
                secretKeyRef:
                  name: {{ include "mcp-ultra.fullname" . }}-secrets
                  key: database-url
            - name: REDIS_URL
              valueFrom:
                secretKeyRef:
                  name: {{ include "mcp-ultra.fullname" . }}-secrets
                  key: redis-url
            - name: NATS_URL
              valueFrom:
                secretKeyRef:
                  name: {{ include "mcp-ultra.fullname" . }}-secrets
                  key: nats-url
            - name: JWT_SECRET
              valueFrom:
                secretKeyRef:
                  name: {{ include "mcp-ultra.fullname" . }}-secrets
                  key: jwt-secret
            - name: API_KEY
              valueFrom:
                secretKeyRef:
                  name: {{ include "mcp-ultra.fullname" . }}-secrets
                  key: api-key
          volumeMounts:
            - name: config
              mountPath: /app/config
              readOnly: true
            {{- if .Values.config.security.tls.enabled }}
            - name: tls-certs
              mountPath: /etc/ssl/certs/mcp-ultra
              readOnly: true
            {{- end }}
      volumes:
        - name: config
          configMap:
            name: {{ include "mcp-ultra.fullname" . }}-config
        {{- if .Values.config.security.tls.enabled }}
        - name: tls-certs
          secret:
            secretName: {{ include "mcp-ultra.fullname" . }}-tls
            defaultMode: 0400
        {{- end }}
      {{- with .Values.nodeSelector }}
      nodeSelector:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.affinity }}
      affinity:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.tolerations }}
      tolerations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
```

### 2. Helm Packaging Script

```bash
#!/bin/bash
# scripts/helm-package.sh

set -e

CHART_DIR="deploy/helm/mcp-ultra"
PACKAGE_DIR="dist/helm"
VERSION=${1:-$(grep '^version:' $CHART_DIR/Chart.yaml | awk '{print $2}')}

echo "📦 Packaging Helm chart version $VERSION..."

# Create package directory
mkdir -p $PACKAGE_DIR

# Update dependencies
cd $CHART_DIR
helm dependency update

# Lint the chart
echo "🔍 Linting Helm chart..."
helm lint .

# Package the chart
echo "📦 Creating Helm package..."
helm package . --destination ../../$PACKAGE_DIR

# Create index
cd ../../$PACKAGE_DIR
helm repo index . --url https://github.com/vertikon/mcp-ultra/releases/download

echo "✅ Helm chart packaged successfully"
echo "📁 Package location: $PACKAGE_DIR/mcp-ultra-$VERSION.tgz"
```

## Binary Distribution

### 1. Multi-Platform Binary Build

#### Build Script for All Platforms
```bash
#!/bin/bash
# scripts/build-binaries.sh

set -e

VERSION=${1:-$(git describe --tags --always)}
BUILD_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_COMMIT=$(git rev-parse HEAD)

PLATFORMS=(
    "linux/amd64"
    "linux/arm64"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
)

DIST_DIR="dist/binaries"
mkdir -p $DIST_DIR

echo "🏗️ Building binaries for version $VERSION..."

for platform in "${PLATFORMS[@]}"; do
    GOOS=${platform%/*}
    GOARCH=${platform#*/}
    
    echo "Building for $GOOS/$GOARCH..."
    
    output_name="mcp-ultra-$VERSION-$GOOS-$GOARCH"
    if [ "$GOOS" = "windows" ]; then
        output_name+=".exe"
    fi
    
    env GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build \
        -ldflags="-w -s -X 'github.com/vertikon/mcp-ultra/pkg/version.Version=$VERSION' -X 'github.com/vertikon/mcp-ultra/pkg/version.GitCommit=$GIT_COMMIT' -X 'github.com/vertikon/mcp-ultra/pkg/version.BuildDate=$BUILD_DATE'" \
        -o $DIST_DIR/$output_name \
        cmd/mcp-model-ultra/main.go
        
    # Create archive
    cd $DIST_DIR
    if [ "$GOOS" = "windows" ]; then
        zip -q mcp-ultra-$VERSION-$GOOS-$GOARCH.zip $output_name
    else
        tar -czf mcp-ultra-$VERSION-$GOOS-$GOARCH.tar.gz $output_name
    fi
    rm $output_name
    cd - > /dev/null
done

# Generate checksums
cd $DIST_DIR
sha256sum * > checksums.txt
cd - > /dev/null

echo "✅ Binary builds completed"
echo "📁 Binaries available in: $DIST_DIR"
```

### 2. Release Packaging

#### Complete Release Package Script
```bash
#!/bin/bash
# scripts/create-release-package.sh

set -e

VERSION=${1:-v1.0.0}
PACKAGE_NAME="mcp-ultra-$VERSION"
TEMP_DIR="/tmp/mcp-ultra-release"
DIST_DIR="dist"

echo "📦 Creating release package for $VERSION..."

# Clean up previous builds
rm -rf $TEMP_DIR $DIST_DIR
mkdir -p $TEMP_DIR $DIST_DIR

# Copy source files
echo "📋 Copying source files..."
rsync -av --progress . $TEMP_DIR/$PACKAGE_NAME \
    --exclude='.git' \
    --exclude='bin/' \
    --exclude='dist/' \
    --exclude='coverage.out' \
    --exclude='*.test' \
    --exclude='.env' \
    --exclude='*.key' \
    --exclude='*.crt' \
    --exclude='*.pem' \
    --exclude='node_modules/' \
    --exclude='vendor/' \
    --exclude='*.log' \
    --exclude='*.sarif'

# Create documentation package
echo "📚 Creating documentation package..."
mkdir -p $TEMP_DIR/$PACKAGE_NAME/docs/release
cp documentos-full/*.md $TEMP_DIR/$PACKAGE_NAME/docs/release/

# Create quick start script
cat << 'EOF' > $TEMP_DIR/$PACKAGE_NAME/quickstart.sh
#!/bin/bash
set -e

echo "🚀 MCP Ultra Quick Start"
echo "======================="

# Check prerequisites
command -v go >/dev/null 2>&1 || { echo "Go is required but not installed. Aborting." >&2; exit 1; }
command -v docker >/dev/null 2>&1 || { echo "Docker is required but not installed. Aborting." >&2; exit 1; }

# Setup environment
echo "📋 Setting up environment..."
cp config/.env.example .env
echo "✅ Environment file created (.env)"

# Install dependencies
echo "📦 Installing dependencies..."
go mod download
echo "✅ Dependencies installed"

# Start infrastructure
echo "🏗️ Starting infrastructure..."
docker-compose up -d postgres redis nats
sleep 10
echo "✅ Infrastructure started"

# Run tests
echo "🧪 Running tests..."
go test ./... -short
echo "✅ Tests passed"

# Build application
echo "🏗️ Building application..."
go build -o bin/mcp-ultra cmd/mcp-model-ultra/main.go
echo "✅ Application built"

# Start application
echo "🚀 Starting application..."
echo "Visit http://localhost:9655/healthz to check health"
echo "Visit http://localhost:9655/metrics for metrics"
echo ""
echo "Press Ctrl+C to stop"
./bin/mcp-ultra
EOF

chmod +x $TEMP_DIR/$PACKAGE_NAME/quickstart.sh

# Create README for release
cat << EOF > $TEMP_DIR/$PACKAGE_NAME/RELEASE_NOTES.md
# MCP Ultra $VERSION Release

## 🎉 What's New

This release includes:
- ✅ Enterprise-grade Go microservice template
- ✅ Comprehensive health check endpoints
- ✅ 95%+ test coverage with multiple testing strategies
- ✅ OpenTelemetry observability integration
- ✅ LGPD/GDPR compliance framework
- ✅ Production-ready Docker containers
- ✅ Kubernetes deployment manifests
- ✅ Helm charts for easy deployment

## 🚀 Quick Start

1. Extract the archive:
   \`\`\`bash
   tar -xzf mcp-ultra-$VERSION.tar.gz
   cd mcp-ultra-$VERSION
   \`\`\`

2. Run quick start:
   \`\`\`bash
   ./quickstart.sh
   \`\`\`

## 📚 Documentation

- [GitHub Setup Guide](docs/release/GITHUB_SETUP.md)
- [Lifecycle Management](docs/release/LIFECYCLE_MANAGEMENT.md)
- [Packaging Instructions](docs/release/PACKAGING_INSTRUCTIONS.md)
- [Production Deployment](docs/release/SEND_READY.md)

## 🆘 Support

- GitHub Issues: https://github.com/vertikon/mcp-ultra/issues
- Documentation: https://github.com/vertikon/mcp-ultra
- Email: support@vertikon.com

---

**Enterprise microservice template ready for production use!** ✨
EOF

# Create archives
echo "🗜️ Creating archives..."
cd $TEMP_DIR

# Create tar.gz archive
tar -czf $PACKAGE_NAME.tar.gz $PACKAGE_NAME/
mv $PACKAGE_NAME.tar.gz $(pwd)/../$DIST_DIR/

# Create zip archive
zip -r -q $PACKAGE_NAME.zip $PACKAGE_NAME/
mv $PACKAGE_NAME.zip $(pwd)/../$DIST_DIR/

cd - > /dev/null

# Generate checksums
cd $DIST_DIR
sha256sum *.tar.gz *.zip > checksums.txt
cd - > /dev/null

# Clean up
rm -rf $TEMP_DIR

echo "✅ Release package created successfully"
echo "📁 Packages available in: $DIST_DIR"
echo "📋 Files created:"
ls -la $DIST_DIR/
```

## Cloud Provider Packaging

### 1. AWS Lambda Packaging

#### Lambda Deployment Package
```bash
#!/bin/bash
# scripts/lambda-package.sh

set -e

FUNCTION_NAME="mcp-ultra"
REGION=${AWS_REGION:-us-east-1}
LAMBDA_DIR="dist/lambda"

echo "🚀 Creating AWS Lambda package..."

mkdir -p $LAMBDA_DIR

# Build for Lambda
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
    -ldflags="-w -s" \
    -o $LAMBDA_DIR/main \
    cmd/lambda/main.go

# Create deployment package
cd $LAMBDA_DIR
zip -r mcp-ultra-lambda.zip main
cd - > /dev/null

echo "✅ Lambda package created: $LAMBDA_DIR/mcp-ultra-lambda.zip"
```

### 2. Google Cloud Run Packaging

#### Cloud Run Deployment
```yaml
# deploy/cloud-run/service.yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: mcp-ultra
  annotations:
    run.googleapis.com/ingress: all
    run.googleapis.com/execution-environment: gen2
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/minScale: "1"
        autoscaling.knative.dev/maxScale: "100"
        run.googleapis.com/cpu-throttling: "false"
    spec:
      containerConcurrency: 1000
      timeoutSeconds: 300
      serviceAccountName: mcp-ultra@PROJECT_ID.iam.gserviceaccount.com
      containers:
      - image: gcr.io/PROJECT_ID/mcp-ultra:latest
        ports:
        - name: http1
          containerPort: 9655
        env:
        - name: APP_ENV
          value: "production"
        - name: PORT
          value: "9655"
        resources:
          limits:
            cpu: "1"
            memory: "512Mi"
        livenessProbe:
          httpGet:
            path: /live
            port: 9655
          initialDelaySeconds: 10
          periodSeconds: 10
        startupProbe:
          httpGet:
            path: /healthz
            port: 9655
          initialDelaySeconds: 10
          periodSeconds: 5
          failureThreshold: 10
```

### 3. Azure Container Apps

#### Container App Configuration
```yaml
# deploy/azure/containerapp.yaml
apiVersion: 2022-03-01
location: East US
type: Microsoft.App/containerApps
properties:
  managedEnvironmentId: /subscriptions/SUBSCRIPTION_ID/resourceGroups/RG_NAME/providers/Microsoft.App/managedEnvironments/ENV_NAME
  configuration:
    ingress:
      external: true
      targetPort: 9655
      allowInsecure: false
    secrets:
    - name: database-url
      value: "postgres://..."
    - name: jwt-secret
      value: "secret-value"
    registries:
    - server: ghcr.io
      username: username
      passwordSecretRef: registry-password
  template:
    containers:
    - image: ghcr.io/vertikon/mcp-ultra:latest
      name: mcp-ultra
      env:
      - name: APP_ENV
        value: "production"
      - name: DATABASE_URL
        secretRef: database-url
      - name: JWT_SECRET
        secretRef: jwt-secret
      resources:
        cpu: 0.5
        memory: 1Gi
      probes:
      - type: liveness
        httpGet:
          path: /live
          port: 9655
        initialDelaySeconds: 10
        periodSeconds: 10
      - type: readiness
        httpGet:
          path: /ready
          port: 9655
        initialDelaySeconds: 5
        periodSeconds: 5
    scale:
      minReplicas: 1
      maxReplicas: 10
      rules:
      - name: http-rule
        http:
          metadata:
            concurrentRequests: "100"
```

## CI/CD Integration

### 1. GitHub Actions Packaging Workflow

```yaml
# .github/workflows/package-and-release.yml
name: Package and Release

on:
  push:
    tags:
      - 'v*'
  workflow_dispatch:
    inputs:
      version:
        description: 'Version to package'
        required: true
        default: 'v1.0.0'

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  build-binaries:
    name: Build Binaries
    runs-on: ubuntu-latest
    strategy:
      matrix:
        include:
        - goos: linux
          goarch: amd64
        - goos: linux
          goarch: arm64
        - goos: darwin
          goarch: amd64
        - goos: darwin
          goarch: arm64
        - goos: windows
          goarch: amd64
    
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      
    - name: Setup Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.22'
        
    - name: Build binary
      env:
        GOOS: ${{ matrix.goos }}
        GOARCH: ${{ matrix.goarch }}
        CGO_ENABLED: 0
      run: |
        VERSION=${GITHUB_REF_NAME:-${{ github.event.inputs.version }}}
        BUILD_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)
        GIT_COMMIT=${{ github.sha }}
        
        OUTPUT_NAME="mcp-ultra-$VERSION-$GOOS-$GOARCH"
        if [ "$GOOS" = "windows" ]; then
          OUTPUT_NAME="${OUTPUT_NAME}.exe"
        fi
        
        go build \
          -ldflags="-w -s -X 'github.com/vertikon/mcp-ultra/pkg/version.Version=$VERSION' -X 'github.com/vertikon/mcp-ultra/pkg/version.GitCommit=$GIT_COMMIT' -X 'github.com/vertikon/mcp-ultra/pkg/version.BuildDate=$BUILD_DATE'" \
          -o $OUTPUT_NAME \
          cmd/mcp-model-ultra/main.go
          
    - name: Upload binary
      uses: actions/upload-artifact@v3
      with:
        name: binaries
        path: mcp-ultra-*
        retention-days: 30

  build-container:
    name: Build Container
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write
    
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      
    - name: Setup Docker Buildx
      uses: docker/setup-buildx-action@v3
      
    - name: Login to Container Registry
      uses: docker/login-action@v3
      with:
        registry: ${{ env.REGISTRY }}
        username: ${{ github.actor }}
        password: ${{ secrets.GITHUB_TOKEN }}
        
    - name: Extract metadata
      id: meta
      uses: docker/metadata-action@v5
      with:
        images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
        tags: |
          type=ref,event=tag
          type=raw,value=latest,enable={{is_default_branch}}
          
    - name: Build and push
      uses: docker/build-push-action@v5
      with:
        context: .
        platforms: linux/amd64,linux/arm64
        push: true
        tags: ${{ steps.meta.outputs.tags }}
        labels: ${{ steps.meta.outputs.labels }}
        cache-from: type=gha
        cache-to: type=gha,mode=max

  package-helm:
    name: Package Helm Chart
    runs-on: ubuntu-latest
    
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      
    - name: Setup Helm
      uses: azure/setup-helm@v3
      with:
        version: 'latest'
        
    - name: Package Helm chart
      run: |
        cd deploy/helm/mcp-ultra
        helm dependency update
        helm lint .
        helm package . --destination ../../../dist/helm
        
    - name: Upload Helm chart
      uses: actions/upload-artifact@v3
      with:
        name: helm-chart
        path: dist/helm/*.tgz
        retention-days: 30

  create-release:
    name: Create Release
    runs-on: ubuntu-latest
    needs: [build-binaries, build-container, package-helm]
    if: startsWith(github.ref, 'refs/tags/')
    permissions:
      contents: write
    
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      
    - name: Download artifacts
      uses: actions/download-artifact@v3
      with:
        path: dist/
        
    - name: Create release package
      run: |
        ./scripts/create-release-package.sh ${{ github.ref_name }}
        
    - name: Generate checksums
      run: |
        cd dist
        find . -type f -name "*.tar.gz" -o -name "*.zip" -o -name "*.tgz" | xargs sha256sum > checksums.txt
        
    - name: Create Release
      uses: softprops/action-gh-release@v1
      with:
        tag_name: ${{ github.ref_name }}
        name: MCP Ultra ${{ github.ref_name }}
        body_path: RELEASE_NOTES.md
        files: |
          dist/**/*
        draft: false
        prerelease: false
        generate_release_notes: true
```

## Security Considerations

### 1. Container Security

#### Security Scanning
```bash
#!/bin/bash
# scripts/security-scan-container.sh

set -e

IMAGE_NAME="ghcr.io/vertikon/mcp-ultra:latest"

echo "🔒 Running container security scans..."

# Trivy vulnerability scan
echo "Running Trivy scan..."
trivy image --severity HIGH,CRITICAL --format table $IMAGE_NAME

# Docker Scout scan
echo "Running Docker Scout scan..."
docker scout cves $IMAGE_NAME

# Container structure test
echo "Running container structure tests..."
container-structure-test test --image $IMAGE_NAME --config deploy/tests/container-structure-test.yaml

echo "✅ Container security scans completed"
```

### 2. Supply Chain Security

#### SLSA Provenance Generation
```yaml
# .github/workflows/slsa-provenance.yml
name: SLSA Provenance

on:
  push:
    tags:
      - 'v*'

permissions:
  actions: read
  id-token: write
  contents: write

jobs:
  build:
    outputs:
      hashes: ${{ steps.hash.outputs.hashes }}
    runs-on: ubuntu-latest
    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      
    - name: Setup Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.22'
        
    - name: Build binary
      run: |
        go build -o mcp-ultra cmd/mcp-model-ultra/main.go
        
    - name: Generate hashes
      shell: bash
      id: hash
      run: |
        echo "hashes=$(sha256sum mcp-ultra | base64 -w0)" >> "$GITHUB_OUTPUT"
        
    - name: Upload binary
      uses: actions/upload-artifact@v3
      with:
        name: mcp-ultra
        path: mcp-ultra

  provenance:
    needs: [build]
    permissions:
      actions: read
      id-token: write
      contents: write
    uses: slsa-framework/slsa-github-generator/.github/workflows/generator_generic_slsa3.yml@v1.4.0
    with:
      base64-subjects: "${{ needs.build.outputs.hashes }}"
      upload-assets: true
```

## Troubleshooting

### 1. Common Issues

#### Docker Build Issues
```bash
# Problem: Docker build fails with permission issues
# Solution: Ensure proper user permissions in Dockerfile
USER appuser
WORKDIR /app
RUN chown -R appuser:appuser /app

# Problem: Multi-arch build fails
# Solution: Use docker buildx with proper platform specification
docker buildx build --platform linux/amd64,linux/arm64 .

# Problem: Container security scan failures
# Solution: Update base image and dependencies
FROM alpine:3.19  # Use latest secure base image
RUN apk add --no-cache ca-certificates  # Only essential packages
```

#### Kubernetes Deployment Issues
```bash
# Problem: Pods fail to start
kubectl describe pod mcp-ultra-xxx
kubectl logs mcp-ultra-xxx

# Problem: Health checks fail
kubectl exec -it mcp-ultra-xxx -- curl localhost:9655/healthz

# Problem: Resource limits too low
kubectl top pod mcp-ultra-xxx
# Increase resource requests/limits in deployment
```

#### Helm Chart Issues
```bash
# Problem: Helm install fails
helm lint deploy/helm/mcp-ultra
helm template deploy/helm/mcp-ultra --debug

# Problem: Dependencies not resolved
helm dependency update deploy/helm/mcp-ultra

# Problem: Values not applied correctly
helm get values mcp-ultra-release
```

### 2. Debugging Scripts

#### Container Debug Script
```bash
#!/bin/bash
# scripts/debug-container.sh

IMAGE_NAME=${1:-"ghcr.io/vertikon/mcp-ultra:latest"}

echo "🐛 Debugging container: $IMAGE_NAME"

# Check image layers
echo "📋 Image layers:"
docker history $IMAGE_NAME

# Run container in debug mode
echo "🚀 Running container in debug mode..."
docker run -it --rm \
    --entrypoint /bin/sh \
    -p 9655:9655 \
    -p 9656:9656 \
    $IMAGE_NAME

# Check container security
echo "🔒 Checking container security..."
docker run --rm -it \
    -v /var/run/docker.sock:/var/run/docker.sock \
    aquasec/trivy image $IMAGE_NAME
```

## Conclusion

This packaging guide provides comprehensive instructions for packaging MCP Ultra for various deployment environments. Key takeaways:

### ✅ **Comprehensive Packaging Support**
- **Docker Containers**: Multi-stage, multi-architecture builds with security hardening
- **Kubernetes**: Complete manifests with advanced features (HPA, PDB, Network Policies)
- **Helm Charts**: Production-ready charts with configurable values and dependencies
- **Binary Distribution**: Cross-platform binaries with automated packaging

### ✅ **Cloud-Native Ready**
- **AWS Lambda**: Serverless deployment packages
- **Google Cloud Run**: Container-based serverless deployment
- **Azure Container Apps**: Managed container platform deployment
- **Multi-cloud Support**: Platform-agnostic packaging strategies

### ✅ **Enterprise Features**
- **Security Scanning**: Container vulnerability scanning and compliance
- **Supply Chain Security**: SLSA provenance generation and verification
- **CI/CD Integration**: Automated packaging and release workflows
- **Monitoring Integration**: Built-in observability and health checks

### ✅ **Production Excellence**
- **Zero-Downtime Deployments**: Rolling updates and health checks
- **Auto-Scaling**: HPA configuration and resource optimization
- **High Availability**: Multi-replica deployments and pod disruption budgets
- **Security Hardening**: Non-root containers and security contexts

**This guide ensures MCP Ultra can be packaged and deployed across any modern infrastructure with enterprise-grade reliability and security.**

---

**📋 Next Steps:**
1. Choose the appropriate packaging method for your deployment environment
2. Customize the configurations based on your specific requirements
3. Set up CI/CD pipelines using the provided workflows
4. Implement monitoring and alerting for your deployments
5. Test disaster recovery and rollback procedures

**Your enterprise microservice is ready for any deployment scenario!** ✨