# 🔧 Docker Port Conflict - Solução Rápida

## ❌ Problema Encontrado

```
Error response from daemon: failed to set up container networking:
driver failed programming external connectivity on endpoint mcp-ultra-nats-1:
Bind for 0.0.0.0:4222 failed: port is already allocated
```

**Causa**: A porta 4222 (NATS) já está sendo usada por outro processo ou container.

---

## ✅ Soluções (Escolha Uma)

### Solução 1: Parar Container/Processo Existente (Recomendado)

#### Opção A: Parar containers Docker antigos

```powershell
# Ver todos os containers (incluindo parados)
docker ps -a

# Parar container que está usando a porta
docker stop <container-name-or-id>

# Ou parar todos os containers
docker stop $(docker ps -q)

# Remover containers parados
docker container prune -f

# Tentar novamente
docker-compose up -d
```

#### Opção B: Identificar processo Windows na porta

```powershell
# PowerShell: Encontrar PID usando porta 4222
netstat -ano | findstr :4222

# Exemplo de output:
#   TCP    0.0.0.0:4222    0.0.0.0:0    LISTENING    1234
#                                                     ^^^^ PID

# Ver qual processo é
tasklist | findstr <PID>

# Parar processo (se não for crítico)
Stop-Process -Id <PID> -Force

# Tentar novamente
docker-compose up -d
```

### Solução 2: Alterar Portas no docker-compose.yml

Se você não pode parar o processo que está usando 4222, altere as portas:

```yaml
# docker-compose.yml
services:
  nats:
    image: nats:2.10-alpine
    ports:
      - "14222:4222"  # ← Mudar de 4222:4222 para 14222:4222
      - "18222:8222"  # ← Mudar de 8222:8222 para 18222:8222
    # ... resto da config
```

**IMPORTANTE**: Se alterar a porta NATS, também atualize `.env`:
```env
# .env
NATS_URL=nats://nats:4222  # ← NÃO MUDE! Isso é interno ao container
```

A porta interna (4222) permanece a mesma, apenas a porta externa muda (14222).

### Solução 3: Subir Apenas Alguns Serviços

Se NATS não for crítico para teste inicial:

```powershell
# Subir apenas app + banco + redis (sem NATS)
docker-compose up -d postgres redis mcp-ultra

# Ver status
docker-compose ps

# Verificar logs
docker-compose logs -f mcp-ultra
```

### Solução 4: Usar Docker Compose Override

Criar `docker-compose.override.yml`:

```yaml
# docker-compose.override.yml
version: '3.8'

services:
  nats:
    ports:
      - "14222:4222"
      - "18222:8222"
```

Este arquivo é automaticamente lido pelo docker-compose e sobrescreve as portas.

---

## 🔍 Verificar Portas Antes de Subir

```powershell
# PowerShell: Verificar todas as portas que o docker-compose usa
netstat -ano | findstr ":5432 :6379 :4222 :9655 :9656 :9090 :3000 :16686"

# Se alguma porta aparecer, identifique o processo:
netstat -ano | findstr :<PORTA>
```

**Portas Usadas pelo MCP Ultra**:
- 5432 - PostgreSQL
- 6379 - Redis
- 4222 - NATS (messaging)
- 8222 - NATS (monitoring)
- 9655 - MCP Ultra (API)
- 9656 - MCP Ultra (metrics)
- 9090 - Prometheus
- 3000 - Grafana
- 16686 - Jaeger

---

## ✅ Solução Aplicada (Passo a Passo)

### 1. Limpar Ambiente Docker

```powershell
# Parar todos os containers
docker-compose down

# Remover containers antigos
docker container prune -f

# Remover networks não usadas
docker network prune -f
```

### 2. Verificar Portas

```powershell
# Verificar se 4222 está livre
netstat -ano | findstr :4222

# Se retornar vazio = porta livre ✅
# Se retornar algo = porta ocupada ❌
```

### 3. Subir Novamente

```powershell
# Subir stack completo
docker-compose up -d

# Ver logs em tempo real
docker-compose logs -f

# Verificar health
curl http://localhost:9655/healthz
```

---

## 🎯 Teste Rápido (Sem NATS)

Se você só quer testar o MCP Ultra rapidamente sem NATS:

```powershell
# Criar .env mínimo
cp .env.example .env

# Subir apenas essenciais
docker-compose up -d postgres redis

# Aguardar banco ficar pronto (10-15 segundos)
timeout /t 15

# Subir aplicação
docker-compose up -d mcp-ultra

# Testar
curl http://localhost:9655/healthz
curl http://localhost:9656/metrics
```

---

## 🔧 Fix Permanente

Se você sempre tem conflito de porta, edite `docker-compose.yml`:

```yaml
# docker-compose.yml
services:
  nats:
    image: nats:2.10-alpine
    ports:
      - "14222:4222"  # Porta externa diferente
      - "18222:8222"
    # ... resto permanece igual
```

**Commit a mudança**:
```bash
git add docker-compose.yml
git commit -m "fix: change NATS external ports to avoid conflicts"
git push
```

---

## 📊 Status Esperado Após Fix

```powershell
docker-compose ps
```

**Output Esperado**:
```
NAME                        STATUS              PORTS
mcp-ultra-grafana-1         Up (healthy)        0.0.0.0:3000->3000/tcp
mcp-ultra-jaeger-1          Up                  0.0.0.0:16686->16686/tcp, ...
mcp-ultra-mcp-ultra-1       Up (healthy)        0.0.0.0:9655-9656->9655-9656/tcp
mcp-ultra-nats-1            Up (healthy)        0.0.0.0:4222->4222/tcp
mcp-ultra-postgres-1        Up (healthy)        0.0.0.0:5432->5432/tcp
mcp-ultra-prometheus-1      Up                  0.0.0.0:9090->9090/tcp
mcp-ultra-redis-1           Up (healthy)        0.0.0.0:6379->6379/tcp
```

Todos devem estar com status **"Up (healthy)"**.

---

## 🚨 Troubleshooting

### Container não fica "healthy"

```powershell
# Ver logs do container com problema
docker logs mcp-ultra-nats-1

# Ver health check details
docker inspect mcp-ultra-nats-1 | findstr -i health
```

### Porta ainda ocupada após docker stop

```powershell
# Encontrar processo
netstat -ano | findstr :4222

# Matar processo (CUIDADO!)
Stop-Process -Id <PID> -Force
```

### Permissão negada

```powershell
# Executar PowerShell como Administrador
# Ou adicionar seu usuário ao grupo docker-users
```

---

## ✅ Checklist de Validação

Após aplicar fix:

- [ ] `docker-compose down` executado sem erros
- [ ] Portas verificadas e livres
- [ ] `docker-compose up -d` executado sem erros
- [ ] `docker-compose ps` mostra todos containers "Up"
- [ ] Health checks passando (aguardar 30s)
- [ ] `curl http://localhost:9655/healthz` retorna OK
- [ ] `curl http://localhost:9656/metrics` retorna métricas
- [ ] Grafana acessível em http://localhost:3000
- [ ] Prometheus acessível em http://localhost:9090

---

## 📞 Próximos Passos

1. ✅ **Fix aplicado**: Portas liberadas, containers rodando
2. ✅ **Validação**: Health checks passando
3. ⏭️ **Testes**: Testar API endpoints
4. ⏭️ **Monitoramento**: Configurar Grafana dashboards
5. ⏭️ **Deploy**: Push para registry (opcional)

---

**Documento gerado**: 2025-10-19
**Autor**: Claude Code
**Status**: Ready para uso

Se o problema persistir, execute:
```powershell
# Logs completos
docker-compose logs --tail=100 > docker-error.log

# Anexe docker-error.log ao issue/ticket
```
