# MCP Ultra Template - Guia de Uso

## 🎯 Propósito

Este é o **template base** para criação de novos microserviços na arquitetura Vertikon.
Ele usa placeholders {{MODULE_PATH}} que são substituídos durante o bootstrap.

## 🚀 Criando um novo serviço (semente)

### 1. Clone o template

\\\powershell
Copy-Item -Recurse -Force \
  "E:\vertikon\business\SaaS\templates\mcp-ultra" \
  "E:\vertikon\.ecosistema-vertikon\NeuraLead\waba\meu-servico"
\\\

### 2. Execute o bootstrap

\\\powershell
cd "E:\vertikon\.ecosistema-vertikon\NeuraLead\waba\meu-servico"
.\scripts\bootstrap.ps1 github.com/vertikon/meu-servico
\\\

### 3. Valide

\\\powershell
go mod tidy
go build ./...
go test ./...
\\\

## 📦 Dependências

### ✅ Permitidas (via mcp-ultra-fix)

\\\go
import (
  "github.com/vertikon/mcp-ultra-fix/pkg/logger"
  "github.com/vertikon/mcp-ultra-fix/pkg/version"
  "github.com/vertikon/mcp-ultra-fix/pkg/config"
  // ... outros pacotes do fix
)
\\\

### ❌ Proibidas (privadas)

\\\go
// NUNCA use:
import "github.com/vertikon/mcp-ultra/internal/..."
import "github.com/vertikon/mcp-ultra/test/mocks"
\\\

### ✅ Imports internos (após bootstrap)

\\\go
import (
  "github.com/vertikon/meu-servico/internal/config"
  "github.com/vertikon/meu-servico/internal/handlers"
  "github.com/vertikon/meu-servico/test/mocks"
)
\\\

## 🧪 Mocks

Cada projeto mantém seus próprios mocks em 	est/mocks/.

**Opção 1: testify**
\\\ash
go get github.com/stretchr/testify/mock
\\\

**Opção 2: gomock**
\\\ash
go install github.com/golang/mock/mockgen@latest
go generate ./...
\\\

## 🔧 Manutenção do Template

### Atualizar placeholders (raro)

\\\powershell
.\scripts\prepare-template.ps1
\\\

### Adicionar nova dependência compartilhada

1. Adicione ao mcp-ultra-fix (não ao template!)
2. Publique nova versão do fix
3. Use no template via import do fix

## ✅ Checklist de Qualidade

- [ ] go mod tidy sem erros
- [ ] go build ./... compila
- [ ] go test ./... passa
- [ ] Nenhum import para github.com/vertikon/mcp-ultra/...
- [ ] Apenas imports {{MODULE_PATH}}/... ou mcp-ultra-fix/pkg/...
- [ ] Mocks em 	est/mocks/ (local)

## 📞 Suporte

Dúvidas? Consulte o time de arquitetura.
