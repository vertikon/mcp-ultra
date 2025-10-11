# 🚀 MCP Ultra - Template Completo para Automação GitHub

[![MCP](https://img.shields.io/badge/MCP-Ultra-blue)](https://modelcontextprotocol.io)
[![GitHub](https://img.shields.io/badge/GitHub-Automation-green)](https://github.com)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8)](https://golang.org)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.0+-3178C6)](https://typescript.com)

**MCP Ultra** é um template completo para automação de repositórios GitHub usando o protocolo MCP (Model Context Protocol), desenvolvido pela equipe Vertikon.

## 🎯 **O que este template resolve**

✅ **Criação automática de repositórios GitHub**
✅ **Automação completa de commits e push** 
✅ **Integração MCP Server <-> GitHub API**
✅ **Gerenciamento de diretórios locais**
✅ **Scripts de setup automático**
✅ **Pipeline de testes end-to-end**

## 🏗️ **Arquitetura**

```
mcp-ultra/
├── 🤖 mcp-server/           # Servidor MCP com GitHub integration
├── 🔧 automation/           # Ferramentas de automação Go/PowerShell  
├── 📋 scripts/              # Scripts de setup e deployment
├── 🧪 testing/              # Testes e validação
├── 📚 docs/                 # Documentação completa
└── 🐳 deploy/               # Configurações de deployment
```

## 📦 Installation

### Prerequisites

- **Go**: 1.21+ ([download](https://golang.org/dl/))
- **PostgreSQL**: 14+ ([download](https://www.postgresql.org/download/))
- **NATS**: 2.10+ with JetStream ([download](https://nats.io/download/))
- **Redis**: 7.0+ (optional, for caching) ([download](https://redis.io/download))
- **Docker**: 20+ (optional, for containerized deployment) ([download](https://www.docker.com/get-started))

### Quick Install

```bash
# Clone the repository
git clone https://github.com/vertikon/mcp-ultra.git
cd mcp-ultra

# Install dependencies
go mod download
go mod tidy

# Build the project
go build ./...

# Run database migrations
psql -U postgres -d mcp_ultra -f migrations/0001_baseline.sql

# Configure environment variables
cp .env.example .env
# Edit .env with your configuration

# Start the server
go run ./cmd/server
```

### Docker Installation

```bash
# Build the Docker image
docker build -t mcp-ultra:latest .

# Run with docker-compose (includes PostgreSQL, NATS, Redis)
docker-compose up -d

# Check health
curl http://localhost:9655/healthz
```

### Configuration

Create a `.env` file in the project root:

```env
# Server Configuration
SERVER_PORT=9655
SERVER_HOST=0.0.0.0

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_secure_password
DB_NAME=mcp_ultra
DB_SSL_MODE=disable

# NATS Configuration
NATS_URL=nats://localhost:4222
NATS_CLUSTER_ID=mcp-ultra-cluster

# Redis Configuration (optional)
REDIS_URL=redis://localhost:6379
REDIS_DB=0

# JWT Configuration
JWT_SECRET=your_jwt_secret_here
JWT_ISSUER=mcp-ultra
JWT_EXPIRY=24h

# AI Configuration (opt-in)
ENABLE_AI=false
AI_ROUTER_MODE=balanced
AI_OPENAI_KEY=your_openai_key
AI_QWEN_KEY=your_qwen_key

# Feature Flags
ENABLE_METRICS=true
ENABLE_TRACING=true
LOG_LEVEL=info
```

### Verify Installation

```bash
# Run tests
go test ./...

# Check code coverage
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run linter
golangci-lint run

# Format code
go fmt ./...
```

## 🏃 Usage

### Development Mode

```bash
# Start with hot reload (requires air)
air

# Or run directly
go run ./cmd/server

# Access the API
curl http://localhost:9655/healthz
```

### Production Mode

```bash
# Build optimized binary
CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o bin/server ./cmd/server

# Run the binary
./bin/server

# Or use systemd service
sudo systemctl start mcp-ultra
sudo systemctl enable mcp-ultra
```

## 📊 Test & Coverage

![coverage](docs/badges/coverage.svg)

Veja o histórico detalhado em [`docs/coverage_history.md`](docs/coverage_history.md).

```bash
# Run tests
go test ./internal/handlers ./tests/integration ./tests/smoke -count=1

# Generate coverage
go test ./internal/handlers ./tests/integration ./tests/smoke -coverpkg=./... -coverprofile=coverage.out
go tool cover -func coverage.out > coverage_func.txt
go tool cover -html coverage.out > coverage.html
```

## 📡 NATS

Complete documentation in [docs/NATS.md](docs/NATS.md).

## 🚀 **Quick Start**

### 1️⃣ **Setup Inicial**
```powershell
# Clone o template
git clone https://github.com/vertikon/mcp-ultra.git
cd mcp-ultra

# Execute setup automático
.\scripts\setup-complete.ps1 -GithubToken "ghp_seu_token_aqui"
```

### 2️⃣ **Teste Pipeline**
```powershell
# Teste completo do pipeline
.\testing\test-complete-pipeline.ps1
```

### 3️⃣ **Usar MCP Server**
```bash
# Iniciar servidor MCP
npm start

# Criar repositório via MCP
echo '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"create_repository","arguments":{"name":"meu-repo"}}}' | node dist/index.js
```

## 📦 **Componentes Principais**

### 🤖 **MCP Server** (`mcp-server/`)
- **6 ferramentas GitHub**: create_repository, create_issue, create_pull_request, search_code, list_workflow_runs, get_repo_stats
- **4 recursos**: repositories, readme, issues, pulls  
- **Cache inteligente** com TTL configurável
- **Validação Zod** para todos os inputs
- **Tratamento robusto de erros**

### 🔧 **Automation Tools** (`automation/`)
- **AutoCommit Go**: Criação de diretórios + Git automation
- **PowerShell Scripts**: Setup e configuração Windows
- **Token Management**: Configuração segura de credenciais

### 🧪 **Testing Suite** (`testing/`)
- **Pipeline completo**: MCP → GitHub → Local → Commit → Push
- **Validação de permissões** 
- **Cleanup automático**
- **Relatórios detalhados**

## 🛠️ **Ferramentas MCP Disponíveis**

| Ferramenta | Descrição | Parâmetros |
|------------|-----------|------------|
| `create_repository` | Cria repositório GitHub | name, description, private, auto_init |
| `create_issue` | Cria issue no GitHub | repo, title, body, labels, assignees |
| `create_pull_request` | Cria pull request | repo, title, body, head, base |
| `search_code` | Busca código nos repos | query, repo, language, path |
| `list_workflow_runs` | Lista GitHub Actions | repo, workflow, branch, status |
| `get_repo_stats` | Estatísticas do repo | repo |

## 🔐 **Configuração de Segurança**

### **GitHub Token** (Obrigatório)
Crie um Personal Access Token com as seguintes permissões:
- ✅ `repo` (acesso completo a repositórios)
- ✅ `read:org` (ler organização) 
- ✅ `read:user` (ler perfil usuário)
- ✅ `workflow` (GitHub Actions)

### **Configuração Segura**
```bash
# Via script interativo
.\scripts\configure-github-token.ps1

# Ou definir manualmente no .env
GITHUB_TOKEN=ghp_seu_token_aqui
GITHUB_ORG=vertikon
GITHUB_DEFAULT_REPO=ecosystem
```

## 📋 **Casos de Uso**

### **🔨 Desenvolvimento**
```bash
# Criar novo repositório para projeto
mcp-tool create_repository --name "meu-projeto" --description "Novo projeto" --private false

# Automatizar commits durante desenvolvimento  
autocommit commit meu-projeto
```

### **🏗️ DevOps**
```bash
# Monitorar status de GitHub Actions
mcp-tool list_workflow_runs --repo "vertikon/ecosystem" --status "failed"

# Criar issues automaticamente para falhas
mcp-tool create_issue --repo "vertikon/ecosystem" --title "Build falhou" --body "Detalhes..."
```

### **🔍 Auditoria**
```bash  
# Buscar padrões de código problemáticos
mcp-tool search_code --query "TODO:" --language "go" 

# Obter estatísticas de repositórios
mcp-tool get_repo_stats --repo "vertikon/ecosystem"
```

## 🎨 **Customização**

### **Configurar Organização**
```bash
# No arquivo .env
GITHUB_ORG=sua_organizacao
GITHUB_DEFAULT_REPO=seu_repo_principal
```

### **Personalizar Automação**
```go
// Em automation/autocommit.go
config := Config{
    BasePath: "C:\\seus\\projetos",
    CommitMsg: "🚀 Deploy automático",
    Branch:   "develop",
}
```

### **Adicionar Ferramentas MCP**
```typescript
// Em mcp-server/src/index.ts
server.setRequestHandler(CallToolRequestSchema, async (request) => {
  switch (request.params.name) {
    case 'sua_nova_ferramenta':
      // Implementar lógica
      return { content: [{ type: 'text', text: 'Resultado' }] };
  }
});
```

## 🧪 **Testes**

### **Teste Unitário**
```bash
cd mcp-server && npm test
cd automation && go test ./...
```

### **Teste de Integração**
```powershell
.\testing\test-complete-pipeline.ps1 -TestRepoName "teste-$(Get-Date -Format 'yyyyMMdd')"
```

### **Validação de Setup**
```powershell
.\scripts\validate-setup.ps1
```

## 📚 **Documentação**

- 📖 [**Guia de Execução**](./docs/GUIA-EXECUCAO.md) - Tutorial passo a passo
- 🔧 [**Setup Avançado**](./docs/SETUP-AVANCADO.md) - Configurações detalhadas  
- 🐛 [**Troubleshooting**](./docs/TROUBLESHOOTING.md) - Solução de problemas
- 🏗️ [**Arquitetura**](./docs/ARQUITETURA.md) - Visão técnica detalhada
- 🔐 [**Segurança**](./docs/SEGURANCA.md) - Boas práticas de segurança

## 🤝 **Contribuição**

1. Fork o projeto
2. Crie branch para feature (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova funcionalidade'`)
4. Push para branch (`git push origin feature/nova-funcionalidade`)  
5. Abra Pull Request

## 📄 **Licença**

Este projeto está licenciado sob MIT License - veja [LICENSE](LICENSE) para detalhes.

## 🆘 **Suporte**

- 📧 **Email**: suporte@vertikon.com
- 💬 **Discord**: [Vertikon Community](https://discord.gg/vertikon)
- 📋 **Issues**: [GitHub Issues](https://github.com/vertikon/mcp-ultra/issues)
- 📖 **Wiki**: [Documentação Completa](https://github.com/vertikon/mcp-ultra/wiki)

## 🏆 **Créditos**

Desenvolvido com ❤️ pela equipe **Vertikon**

- 🤖 **MCP Protocol**: [Anthropic](https://modelcontextprotocol.io)
- 🐙 **GitHub API**: [GitHub](https://docs.github.com/en/rest)
- 🟢 **Node.js**: [Node.js Foundation](https://nodejs.org)
- 🐹 **Go**: [Google](https://golang.org)

---

⭐ **Se este template foi útil, considere dar uma estrela no GitHub!**
