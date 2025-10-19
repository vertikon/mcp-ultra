# 🐳 MCP Ultra - Docker Quick Start

[![Docker](https://img.shields.io/badge/docker-ready-blue.svg)](https://www.docker.com/)
[![Score](https://img.shields.io/badge/validation-100%25-brightgreen.svg)]()
[![Go](https://img.shields.io/badge/go-1.24-00ADD8.svg)](https://golang.org/)

---

## 🚀 Quick Start (3 Comandos)

```bash
# 1. Copiar configuração
cp .env.example .env

# 2. Subir stack completo
docker-compose up -d

# 3. Verificar health
curl http://localhost:9655/healthz
```

**Pronto!** Acesse:
- **API**: http://localhost:9655
- **Metrics**: http://localhost:9656/metrics
- **Grafana**: http://localhost:3000 (admin/password do .env)
- **Prometheus**: http://localhost:9090
- **Jaeger**: http://localhost:16686

---

## 📦 Build Manual

### Opção 1: Script PowerShell (Windows)

```powershell
# Build simples
.\docker-build.ps1

# Build com tag específica
.\docker-build.ps1 -Tag v1.0.0

# Build sem cache
.\docker-build.ps1 -NoBuildCache

# Build e push para registry
.\docker-build.ps1 -Tag v1.0.0 -Push -Registry vertikon
```

### Opção 2: Script Bash (Linux/Mac)

```bash
# Dar permissão de execução
chmod +x docker-build.sh

# Build simples
./docker-build.sh

# Build com tag
./docker-build.sh --tag v1.0.0

# Build e push
./docker-build.sh --tag v1.0.0 --push --registry vertikon
```

### Opção 3: Docker Command (Manual)

```bash
# Build
docker build -t mcp-ultra:latest .

# Build sem cache
docker build --no-cache -t mcp-ultra:latest .

# Verificar imagem
docker images mcp-ultra
```

---

## 🏃 Executar Container

### Standalone (Sem docker-compose)

```bash
# Run básico
docker run -d \
  --name mcp-ultra \
  -p 9655:9655 \
  -p 9656:9656 \
  -e LOG_LEVEL=info \
  mcp-ultra:latest

# Ver logs
docker logs -f mcp-ultra

# Parar
docker stop mcp-ultra

# Remover
docker rm mcp-ultra
```

### Com Docker Compose (Recomendado)

```bash
# Subir todos os serviços
docker-compose up -d

# Ver logs
docker-compose logs -f mcp-ultra

# Parar tudo
docker-compose down

# Parar e remover volumes
docker-compose down -v

# Restart
docker-compose restart mcp-ultra

# Rebuild e restart
docker-compose up -d --build
```

---

## 🔍 Validação

### Health Checks

```bash
# Liveness (está vivo?)
curl http://localhost:9655/livez

# Readiness (pronto?)
curl http://localhost:9655/readyz

# Health completo
curl http://localhost:9655/health

# Métricas Prometheus
curl http://localhost:9656/metrics
```

### Status dos Containers

```bash
# Ver containers rodando
docker-compose ps

# Ver logs de todos os serviços
docker-compose logs -f

# Ver logs de um serviço específico
docker-compose logs -f mcp-ultra

# Ver recursos usados
docker stats
```

---

## 📊 Monitoramento

### Grafana

**URL**: http://localhost:3000
**Login**: admin / (senha do .env)

**Dashboards Incluídos**:
- MCP Ultra Application Metrics
- Go Runtime Metrics
- Database Metrics
- Cache Metrics

### Prometheus

**URL**: http://localhost:9090

**Queries Úteis**:
```promql
# Request rate
rate(http_requests_total[5m])

# Error rate
rate(http_requests_total{status=~"5.."}[5m])

# Latency p99
histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m]))
```

### Jaeger

**URL**: http://localhost:16686

- Tracing distribuído
- Análise de latência
- Service dependency graph

---

## 🛠️ Troubleshooting

### Container não inicia

```bash
# Ver logs de erro
docker logs mcp-ultra

# Ver todas as variáveis de ambiente
docker exec mcp-ultra env

# Verificar conectividade com banco
docker exec mcp-ultra ping postgres
```

### Build falha

```bash
# Build com output verbose
DOCKER_BUILDKIT=1 docker build --progress=plain -t mcp-ultra:latest .

# Limpar cache
docker builder prune -a
```

### Performance ruim

```bash
# Ver recursos usados
docker stats mcp-ultra

# Aumentar limites (editar docker-compose.yml)
services:
  mcp-ultra:
    deploy:
      resources:
        limits:
          cpus: '2.0'
          memory: 2G
```

---

## 🔐 Produção

### Antes do Deploy

- [ ] Alterar senhas em `.env`
- [ ] Gerar JWT_SECRET seguro
- [ ] Gerar ENCRYPTION_MASTER_KEY
- [ ] Configurar SSL/TLS
- [ ] Configurar backups de volumes
- [ ] Configurar monitoramento externo
- [ ] Configurar alertas

### Gerar Secrets

```bash
# JWT Secret (256-bit)
openssl rand -base64 64

# Encryption Key (256-bit)
openssl rand -base64 32

# Database Password
openssl rand -base64 24

# API Key
uuidgen
```

### Push para Registry

```bash
# Docker Hub
docker login
docker tag mcp-ultra:latest vertikon/mcp-ultra:latest
docker push vertikon/mcp-ultra:latest

# Private Registry
docker tag mcp-ultra:latest registry.exemplo.com/mcp-ultra:latest
docker push registry.exemplo.com/mcp-ultra:latest
```

---

## 📁 Estrutura de Arquivos

```
mcp-ultra/
├── Dockerfile              # Multi-stage Dockerfile otimizado
├── .dockerignore           # Arquivos excluídos do build
├── docker-compose.yml      # Stack completo (app + deps + monitoring)
├── docker-build.ps1        # Script de build (Windows)
├── docker-build.sh         # Script de build (Linux/Mac)
├── .env.example            # Template de configuração
└── docs/
    ├── DOCKER_DEPLOYMENT.md    # Documentação completa
    └── DOCKER_README.md        # Este arquivo
```

---

## 🏗️ Arquitetura

```
┌─────────────────────────────────────────────────────────┐
│                    MCP Ultra Stack                      │
└─────────────────────────────────────────────────────────┘

┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│  MCP Ultra   │  │  Prometheus  │  │   Grafana    │
│  :9655/9656  │  │    :9090     │  │    :3000     │
└──────┬───────┘  └──────┬───────┘  └──────┬───────┘
       │                 │                  │
       └─────────────────┴──────────────────┘
                        │
       ┌────────────────┴────────────────┐
       │                                 │
┌──────▼──────┐  ┌──────▼──────┐  ┌─────▼──────┐
│ PostgreSQL  │  │    Redis    │  │    NATS    │
│   :5432     │  │    :6379    │  │   :4222    │
└─────────────┘  └─────────────┘  └────────────┘
```

---

## 🎯 Especificações Técnicas

### Imagem Docker

- **Base**: `alpine:latest` (~5MB)
- **Go Version**: 1.24.0
- **Compilação**: CGO_ENABLED=0 (static binary)
- **Otimizações**: -trimpath, -ldflags "-w -s"
- **User**: non-root (appuser:1000)
- **Tamanho Final**: ~15-20MB

### Dependências

| Serviço | Versão | Porta | Descrição |
|---------|--------|-------|-----------|
| PostgreSQL | 16-alpine | 5432 | Database |
| Redis | 7-alpine | 6379 | Cache |
| NATS | 2.10-alpine | 4222 | Messaging |
| Jaeger | latest | 16686 | Tracing |
| Prometheus | latest | 9090 | Metrics |
| Grafana | latest | 3000 | Dashboards |

---

## 📚 Mais Informações

- **Documentação Completa**: [`docs/DOCKER_DEPLOYMENT.md`](docs/DOCKER_DEPLOYMENT.md)
- **Validação 100/100**: [`docs/LINTING_LOOP_ANALYSIS.md`](docs/LINTING_LOOP_ANALYSIS.md)
- **Análise de Regressão**: [`docs/REGRESSION_ANALYSIS_CURSOR.md`](docs/REGRESSION_ANALYSIS_CURSOR.md)
- **Configurações Env**: [`.env.example`](.env.example)
- **Main README**: [`README.md`](README.md)

---

## 🤝 Suporte

**Issues**: Problemas? Abra uma issue com logs:
```bash
docker-compose logs --tail=100 > debug.log
```

**Logs**: Sempre inclua os logs em bug reports:
```bash
docker logs mcp-ultra 2>&1 | grep ERROR
```

---

## ✅ Status

- ✅ **Validação**: 100/100
- ✅ **Build**: OK
- ✅ **Testes**: Passando
- ✅ **Linter**: Limpo
- ✅ **Docker**: Ready
- ✅ **Production**: Ready

---

**Versão**: v1.0.0
**Data**: 2025-10-19
**Autor**: Vertikon Team + Claude Code
**Licença**: MIT

🚀 **Happy Dockering!**
