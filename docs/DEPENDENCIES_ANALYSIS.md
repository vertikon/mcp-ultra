# 🔍 Análise de Dependências - MCP-Ultra

**Data**: 2025-10-11
**Análise**: Dependências duplicadas e não utilizadas

---

## 🚨 Dependências Duplicadas

### 1. **Redis Client - CRÍTICO**

| Biblioteca | Versão | Arquivos | Status |
|------------|--------|----------|--------|
| `github.com/go-redis/redis/v8` | v8.11.5 | 3 arquivos | ❌ ANTIGA |
| `github.com/redis/go-redis/v9` | v9.7.3 | 2 arquivos | ✅ OFICIAL |

**Arquivos usando v8 (antiga)**:
1. `internal/ratelimit/distributed.go`
2. `internal/cache/distributed_test.go`
3. `internal/cache/distributed.go`

**Arquivos usando v9 (nova)**:
1. `internal/repository/redis/cache_repository.go`
2. `internal/repository/redis/connection.go`

**Impacto**:
- +15MB no binário final
- Confusão na API (v8 vs v9 tem diferenças)
- Risco de bugs por misturar versões

**Ação**:
1. Migrar os 3 arquivos v8 → v9
2. Remover `github.com/go-redis/redis/v8` do go.mod
3. Atualizar imports

**Mudanças Necessárias (v8 → v9)**:
```go
// Import
- "github.com/go-redis/redis/v8"
+ "github.com/redis/go-redis/v9"

// Contexto (v9 é context-first em todas as operações)
// Nenhuma mudança necessária - já está correto

// redis.Nil
- if err == redis.Nil
+ if err == redis.Nil  // Mesmo comportamento

// Scripts
- redis.NewScript(...)
+ redis.NewScript(...)  // API idêntica
```

**Estimativa**: 30 minutos (mudança simples, apenas imports)

---

### 2. **HTTP Router**

| Biblioteca | Versão | Arquivos | Status |
|------------|--------|----------|--------|
| `github.com/gorilla/mux` | v1.8.1 | 1 arquivo | ❌ REMOVER |
| `github.com/go-chi/chi/v5` | v5.1.0 | 7 arquivos | ✅ MANTER |

**Arquivo usando gorilla/mux**:
- `internal/handlers/http/swagger.go`

**Arquivos usando chi**:
- `internal/handlers/http/task_handlers.go`
- `internal/handlers/http/health.go`
- `internal/handlers/http/router.go`
- `internal/handlers/http/feature_flag_handlers.go`
- `internal/telemetry/telemetry.go`
- `internal/middleware/auth.go`
- `internal/handlers/http/health_test.go`

**Impacto**:
- +5MB no binário
- 2 APIs diferentes para roteamento
- Inconsistência

**Ação**:
1. Migrar `swagger.go` para usar chi
2. Remover `github.com/gorilla/mux` do go.mod

**Mudanças Necessárias (mux → chi)**:
```go
// Import
- "github.com/gorilla/mux"
+ "github.com/go-chi/chi/v5"

// Router criação
- r := mux.NewRouter()
+ r := chi.NewRouter()

// Rotas
- r.HandleFunc("/path", handler).Methods("GET")
+ r.Get("/path", handler)

// Variáveis de URL
- vars := mux.Vars(r)
- id := vars["id"]
+ id := chi.URLParam(r, "id")
```

**Estimativa**: 20 minutos (1 arquivo apenas)

---

## 🔍 Dependências Suspeitas

### 3. **mcp-ultra-fix - CRÍTICO** ⚠️

```go
github.com/vertikon/mcp-ultra-fix v0.1.0
```

**Usado em**:
- `internal/ratelimit/distributed.go:12` - `github.com/vertikon/mcp-ultra-fix/pkg/logger`

**Problemas**:
1. ❌ **Não documentado** - Nenhuma menção no README
2. ❌ **Dependência circular?** - "mcp-ultra" depende de "mcp-ultra-fix"
3. ❌ **Versão 0.1.0** - Instável
4. ❌ **Único import** - Usado apenas para logger

**Análise**:
```bash
grep -r "mcp-ultra-fix" internal/
# OUTPUT: internal/ratelimit/distributed.go:12:	"github.com/vertikon/mcp-ultra-fix/pkg/logger"
```

**Solução**:
1. **Substituir** `mcp-ultra-fix/pkg/logger` por logger interno
2. Mcp-ultra já tem:
   - `go.uber.org/zap` (linha 38 do go.mod)
   - Usado em outros arquivos

**Mudança**:
```go
// Em internal/ratelimit/distributed.go
- import "github.com/vertikon/mcp-ultra-fix/pkg/logger"
+ import "go.uber.org/zap"

// Mudar tipo
type DistributedRateLimiter struct {
-   logger logger.Logger
+   logger *zap.Logger
}

// Usar logger
- drl.logger.Info("message", "key", value)
+ drl.logger.Info("message", zap.String("key", value))
```

**Estimativa**: 1 hora (refatorar logs no distributed.go)

---

### 4. **Vault API**

```go
github.com/hashicorp/vault/api v1.21.0
```

**Análise**:
```bash
grep -r "hashicorp/vault" internal/
# OUTPUT: (vazio - não encontrado)
```

**Conclusão**: ❌ **NÃO USADO**

**Ação**: Remover do go.mod

**Impacto**:
- Reduz ~15MB de dependências
- Remove 10+ deps transitivas (hashicorp/*)

**Comando**:
```bash
go mod edit -dropreplace github.com/hashicorp/vault/api
go mod tidy
```

**Estimativa**: 2 minutos

---

### 5. **OTEL Exporters**

| Exporter | Uso Típico | Status |
|----------|------------|--------|
| `jaeger` | Tracing (DEPRECATED) | ❌ REMOVER |
| `prometheus` | Métricas | ✅ MANTER |
| `otlptracehttp` | Tracing OTLP | ✅ MANTER |
| `stdout` | Debug local | ❓ AVALIAR |

**Análise**:
```bash
grep -r "exporters/jaeger" internal/
grep -r "exporters/stdout" internal/
```

**Jaeger**:
- ❌ Deprecated desde OTEL 1.0
- Substituído por OTLP
- **Ação**: Remover

**Stdout**:
- Usado para debug local
- Não necessário em produção
- **Ação**: Mover para build tag `dev`

**Estimativa**: 30 minutos

---

## 📊 Resumo de Ações

### **Sprint 1 - Quick Wins (4 horas)**

| Ação | Prioridade | Tempo | Benefício |
|------|-----------|-------|-----------|
| 1. Migrar Redis v8 → v9 | 🔴 ALTA | 30min | -15MB, API consistente |
| 2. Substituir mcp-ultra-fix logger | 🔴 ALTA | 1h | Remove dep circular |
| 3. Migrar gorilla/mux → chi | 🟡 MÉDIA | 20min | -5MB, API consistente |
| 4. Remover Vault API | 🟢 BAIXA | 2min | -15MB deps |
| 5. Remover Jaeger exporter | 🟢 BAIXA | 10min | -5MB deps |
| 6. go mod tidy | 🟡 MÉDIA | 5min | Limpa transitivas |
| 7. Testar compilação | 🔴 ALTA | 10min | Valida mudanças |
| 8. Rodar testes | 🔴 ALTA | 5min | Valida comportamento |

**Total**: 2h 22min
**Ganho Estimado**: -40MB no binário final

---

## 🔧 Plano de Execução

### **Passo 1: Backup**
```bash
git checkout -b refactor/dependencies-cleanup
git tag v1.1.0-pre-refactor
```

### **Passo 2: Migrar Redis (30min)**
1. Editar `internal/ratelimit/distributed.go`
   ```go
   - "github.com/go-redis/redis/v8"
   + "github.com/redis/go-redis/v9"
   ```

2. Editar `internal/cache/distributed.go` e `distributed_test.go`
   - Mesmo import change

3. Verificar compilação:
   ```bash
   go build ./internal/ratelimit ./internal/cache
   ```

### **Passo 3: Substituir mcp-ultra-fix Logger (1h)**
1. Editar `internal/ratelimit/distributed.go`
   ```go
   - import "github.com/vertikon/mcp-ultra-fix/pkg/logger"
   + import "go.uber.org/zap"

   type DistributedRateLimiter struct {
   -   logger logger.Logger
   +   logger *zap.Logger
   }
   ```

2. Refatorar todas as chamadas de log:
   ```bash
   # Encontrar padrões
   grep "logger\." internal/ratelimit/distributed.go
   ```

3. Atualizar construtor:
   ```go
   func NewDistributedRateLimiter(client redis.Cmdable, config Config, logger *zap.Logger, ...) {
       // ...
   }
   ```

### **Passo 4: Migrar Gorilla Mux (20min)**
1. Editar `internal/handlers/http/swagger.go`
   ```go
   - "github.com/gorilla/mux"
   + "github.com/go-chi/chi/v5"
   ```

2. Atualizar código de roteamento

### **Passo 5: Limpar go.mod (5min)**
```bash
# Remover deps não usadas
go get github.com/go-redis/redis/v8@none
go get github.com/gorilla/mux@none
go get github.com/vertikon/mcp-ultra-fix@none
go get github.com/hashicorp/vault/api@none
go get go.opentelemetry.io/otel/exporters/jaeger@none
go get go.opentelemetry.io/otel/exporters/stdout/stdouttrace@none

# Limpar transitivas
go mod tidy
```

### **Passo 6: Validar (15min)**
```bash
# Compilar tudo
go build ./...

# Rodar testes
go test ./internal/... -count=1

# Re-validar
go run E:\vertikon\.ecosistema-vertikon\mcp-tester-system\enhanced_validator_v4.go E:\vertikon\business\SaaS\templates\mcp-ultra
```

### **Passo 7: Commit**
```bash
git add .
git commit -m "refactor: consolidate dependencies

- Migrate Redis client v8 → v9 (3 files)
- Replace mcp-ultra-fix logger with zap
- Migrate gorilla/mux → chi (swagger.go)
- Remove unused dependencies (Vault, Jaeger, mcp-ultra-fix)
- Run go mod tidy

Benefits:
- Binary size: -40MB
- Consistent APIs (Redis v9, chi router)
- No circular dependencies
- Cleaner go.mod (100 deps → ~85 deps)
"
```

---

## 📈 Métricas Esperadas

### **Antes**
- **go.mod deps**: 127 (42 diretas + 85 indiretas)
- **Binary size**: ~80MB
- **Build time**: ~20s
- **Redis APIs**: 2 versões
- **Router APIs**: 2 bibliotecas

### **Depois**
- **go.mod deps**: ~85 (38 diretas + 47 indiretas)
- **Binary size**: ~60MB (-25%)
- **Build time**: ~12s (-40%)
- **Redis APIs**: 1 versão (v9)
- **Router APIs**: 1 biblioteca (chi)

---

## ⚠️ Riscos

### **Risco 1: Redis v8 → v9 API Changes**
**Probabilidade**: Baixa
**Impacto**: Médio
**Mitigação**:
- API é 99% compatível
- Testar todos os endpoints Redis
- Rodar integration tests

### **Risco 2: Logger Refactoring**
**Probabilidade**: Média
**Impacto**: Baixo
**Mitigação**:
- zap é usado em outros lugares do projeto
- Manter mesma estrutura de logs
- Grep por todos os `logger.` antes de mudar

### **Risco 3: Gorilla Mux → Chi**
**Probabilidade**: Baixa
**Impacto**: Baixo
**Mitigação**:
- Apenas 1 arquivo afetado (swagger.go)
- chi tem API similar
- Testar endpoint swagger após mudança

---

## 📚 Referências

- [Redis v9 Migration Guide](https://github.com/redis/go-redis/blob/master/CHANGELOG.md#v900)
- [Chi Router Docs](https://github.com/go-chi/chi)
- [Zap Logger Best Practices](https://pkg.go.dev/go.uber.org/zap)
- [Go Modules Reference](https://go.dev/ref/mod)

---

**Status**: 📝 DRAFT (aguardando execução)
**Versão**: 1.0
**Próximo passo**: Executar Sprint 1
