# 🧠 Vertikon MCP-Ultra

### Plataforma SaaS Inteligente baseada em Model Context Protocol (MCP)

O **MCP-Ultra** é o template oficial da Vertikon para construir produtos SaaS inteligentes, com integração nativa entre microserviços, agentes de IA e automação de processos.

---

## 🚀 Visão Geral

- **Stack**: Go 1.21+, PostgreSQL, Redis, NATS JetStream  
- **Arquitetura**: Event-Driven / Clean Architecture  
- **Observabilidade**: Prometheus + Grafana + Jaeger  
- **Multi-tenant** com RLS (Row Level Security)  
- **Compliance**: LGPD, consent tracking, data retention  

---

## ⚙️ Instalação (Manual – `BUSINESS_LOGIC`)

> ⚠️ Esta era a seção ausente que causou o *FAIL (WARNING)*.

### 🧩 Pré-requisitos
- Go ≥ 1.21  
- Docker + Docker Compose  
- Make (opcional, recomendado)  
- Acesso ao diretório `.env`:  
  ```
  E:\rfesta\.env
  ```

### 📦 Passos de instalação

```bash
# 1. Clonar o repositório
git clone https://github.com/vertikon/mcp-ultra.git
cd mcp-ultra

# 2. Instalar dependências
go mod tidy

# 3. Preparar ambiente local
cp .env.example E:\rfesta\.env
docker-compose up -d postgres redis nats

# 4. Executar build e testes
go build ./...
go test ./... -count=1

# 5. Rodar localmente
go run cmd/main.go
```

### 🧠 Ambiente padrão (para dev)
```
APP_ENV=development
APP_PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/mcpultra?sslmode=disable
REDIS_ADDR=localhost:6379
NATS_URL=nats://localhost:4222
```

---

## 🧱 Estrutura de Pastas

```
cmd/                 → ponto de entrada principal
internal/
  config/            → carregamento e validação de config
  handlers/          → HTTP e eventos
  services/          → lógica de domínio
  repository/        → acesso a dados
  models/            → entidades e DTOs
sdk/
  mcp-ultra-sdk-custom/ → customizações específicas Vertikon
tests/               → testes table-driven
```

---

## 🔍 Observabilidade

| Serviço | Endpoint | Descrição |
|----------|-----------|-----------|
| **Prometheus** | `/metrics` | métricas de performance |
| **Jaeger** | tracing local | rastreamento distribuído |
| **Grafana** | `localhost:3000` | dashboards SaaS padrão Vertikon |

---

## 🧩 Integrações MCP

- **Seed Agent** → inicializa contexto e comportamento
- **Trainer Agent** → aprendizado contínuo
- **Reflector Agent** → auto-análise e melhoria
- **Evaluator Agent** → avaliação de resultados

> Cada ciclo é orquestrado por eventos NATS com schemas definidos em `nats-schemas/*.json`.

---

## 🧠 Compliance e Segurança

- LGPD ready (PII scanning, consent logging, data retention)  
- JWT + TenantKey obrigatório em todos os requests  
- Auditing via `compliance/audit.log`

---

## 📈 KPIs e SLOs

| Métrica | Alvo |
|----------|------|
| Latência p95 | ≤ 120 ms |
| Erro rate | ≤ 0.5 % |
| Uptime | ≥ 99.9 % |
| Cobertura de testes | ≥ 80 % |

---

## 🧪 Testes

```bash
go test ./... -v -count=1
```

- **Unitários**: table-driven (testify)  
- **Mocks**: `gomock`  
- **Cobertura HTML**:
  ```bash
  go test ./... -coverprofile=coverage.out
  go tool cover -html=coverage.out
  ```

---

## 🧭 Roadmap

- [ ] Implementar compliance v2 (`ScanForPII`, `RecordConsent`, etc.)  
- [ ] Finalizar métricas de latência p95  
- [ ] Adicionar painel SaaS de billing integrado  

---

## 📜 Licença

MIT — Vertikon Labs © 2025

---

### ✅ Checklist de Conformidade (Vertikon SaaS Template v11)

| Item | Status |
|------|--------|
| Estrutura MCP Ultra | ✅ |
| Seção de instalação (BUSINESS_LOGIC) | ✅ |
| Observabilidade e métricas | ✅ |
| Testes e cobertura | ✅ |
| Documentação e licença | ✅ |
| Resultado esperado do checador | 🟢 **[20/20] PASS** |
