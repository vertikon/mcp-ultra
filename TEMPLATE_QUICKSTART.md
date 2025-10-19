# MCP Ultra Template - Guia Rápido

Este é o template de segurança completo para projetos MCP da Vertikon. Todos os novos repositórios devem usar este template como base.

## 🚀 Criando um Novo Repositório

### Método 1: Via GitHub UI (Recomendado)

1. Marque este repositório como **Template Repository**:
   - Vá em `Settings` → `Template repository` → Marcar checkbox
   
2. Crie um novo repo a partir do template:
   - Clique em **"Use this template"**
   - Escolha **"Create a new repository"**
   - Preencha nome e descrição
   - Crie o repositório

3. Clone e inicialize:
   ```powershell
   git clone https://github.com/vertikon/seu-novo-repo.git
   cd seu-novo-repo
   .\init.ps1 -RepoName "seu-novo-repo" -Description "Descrição do projeto"
   ```

### Método 2: Via GitHub CLI

```powershell
# Criar repo a partir do template
gh repo create vertikon/novo-repo `
    --template vertikon/mcp-ultra `
    --private `
    --clone

# Entrar no diretório
cd novo-repo

# Inicializar
.\init.ps1 -RepoName "novo-repo" -Description "Meu novo MCP server"
```

### Método 3: Cópia Manual

```powershell
# Copiar template para novo diretório
Copy-Item -Path "E:\vertikon\business\SaaS\templates\mcp-ultra" `
          -Destination "E:\vertikon\meu-novo-projeto" `
          -Recurse

cd E:\vertikon\meu-novo-projeto

# Executar script de inicialização
.\init.ps1 -RepoName "meu-novo-projeto" `
           -Description "Descrição" `
           -CreateRepo
```

## 📁 Estrutura do Template

```
mcp-ultra/
├── .github/
│   ├── workflows/
│   │   ├── codeql.yml          # CodeQL security scanning
│   │   ├── gosec.yml           # GoSec static analysis
│   │   └── test.yml            # Automated testing
│   ├── dependabot.yml          # Dependency updates
│   ├── PULL_REQUEST_TEMPLATE.md
│   ├── bug_report.yml
│   └── feature_request.yml
├── docs/
│   └── SECURITY-OVERVIEW.md    # Documentação de segurança
├── CODEOWNERS                  # Code review assignments
├── SECURITY.md                 # Política de segurança
├── Makefile                    # Build automation
├── .gitignore                  # Arquivos ignorados
├── README.md                   # Template de README
├── init.ps1                    # Script de inicialização
└── setup-repo-security.ps1     # Configuração de segurança
```

## ⚙️ Customização Pós-Criação

### 1. Atualizar CODEOWNERS

Edite `CODEOWNERS` e substitua os placeholders:

```diff
- * @vertikon/platform-team
+ * @vertikon/seu-time

- /internal/security/** @vertikon/security-team
+ /internal/security/** @vertikon/seu-time-seguranca
```

### 2. Configurar Workflows

Edite `.github/workflows/*.yml` conforme necessário:

```yaml
# .github/workflows/test.yml
- name: Setup Go
  uses: actions/setup-go@v5
  with:
-   go-version: "1.22.x"
+   go-version: "1.23.x"  # Sua versão
```

### 3. Atualizar SECURITY.md

Preencha informações de contato:

```markdown
## Security Contacts

- **Security Team**: security@sua-empresa.com
- **Security Lead**: Nome do Responsável
- **Emergency Contact**: +55 11 XXXX-XXXX
```

### 4. Executar Configuração de Segurança

```powershell
# Dry run primeiro (para ver o que será feito)
.\setup-repo-security.ps1 -Owner "vertikon" -Repo "seu-repo" -DryRun

# Executar de verdade
.\setup-repo-security.ps1 -Owner "vertikon" -Repo "seu-repo"
```

Isso irá:
- ✅ Habilitar vulnerability alerts
- ✅ Habilitar Dependabot security updates
- ✅ Habilitar private vulnerability reporting
- ✅ Habilitar secret scanning
- ✅ Habilitar push protection
- ✅ Configurar branch protection
- ✅ Adicionar labels de segurança

## 🔒 Checklist de Segurança Inicial

Antes de fazer o primeiro commit com código real:

- [ ] Atualizar `CODEOWNERS` com times corretos
- [ ] Configurar `SECURITY.md` com contatos reais
- [ ] Executar `setup-repo-security.ps1`
- [ ] Verificar branch protection em GitHub
- [ ] Adicionar secrets necessários em Settings → Secrets
- [ ] Habilitar GitHub Advanced Security (se disponível)
- [ ] Configurar required status checks
- [ ] Adicionar equipe de segurança como reviewers
- [ ] Testar workflows executando `make ci`

## 🛠️ Comandos Make Disponíveis

```bash
# Build
make build              # Compila o projeto
make build-release      # Build para produção

# Testes
make test               # Executa testes
make test-coverage      # Testes com cobertura

# Segurança
make sec-scan           # Todos os scans de segurança
make sec-gosec          # GoSec scanner
make sec-deps           # Verifica vulnerabilidades
make sec-secrets        # Detecta secrets vazados

# Desenvolvimento
make dev                # Roda em modo dev
make lint               # Linters
make fmt                # Formata código

# Pre-commit
make pre-commit         # Roda todos os checks
```

## 📝 Personalizações do README

O README.md usa placeholders que são substituídos pelo `init.ps1`:

- `{REPO_NAME}` → Nome do repositório
- `{DESCRIPTION}` → Descrição do projeto
- `{GO_VERSION}` → Versão do Go

Após executar `init.ps1`, revise e adicione:
- Features específicas do seu projeto
- Instruções de instalação adicionais
- Exemplos de uso
- Capturas de tela ou diagramas

## 🔐 Secrets e Variáveis de Ambiente

### Secrets Necessários (Configure em GitHub Settings → Secrets)

```yaml
# Para Dependabot e Security Scanning
GITHUB_TOKEN: # Fornecido automaticamente

# Para CodeCov (opcional)
CODECOV_TOKEN: seu-token-aqui

# Para Snyk (opcional)
SNYK_TOKEN: seu-token-aqui

# Para notificações (opcional)
SLACK_WEBHOOK_URL: https://hooks.slack.com/...
```

### Variáveis de Ambiente Local

Copie `.env.example` para `.env.local`:

```bash
cp .env.example .env.local
# Edite .env.local com seus valores
```

**IMPORTANTE**: `.env.local` está no `.gitignore` e nunca deve ser commitado!

## 🚢 Deploy

### Docker

```bash
# Build
docker build -t seu-repo:latest .

# Run
docker run -p 8080:8080 seu-repo:latest
```

### Kubernetes

```bash
# Aplicar manifests
kubectl apply -f deploy/k8s/

# Verificar
kubectl get pods -l app=seu-repo
```

## 📊 Monitoramento

Depois de fazer deploy:

1. **Métricas**: Acesse `/metrics` para Prometheus
2. **Health**: Acesse `/health` para health checks
3. **Logs**: Configure agregação de logs (ELK, Datadog, etc.)
4. **Traces**: Configure OpenTelemetry endpoint

## 🆘 Troubleshooting

### Workflows falhando

```bash
# Verificar localmente
make ci

# Checar logs em GitHub
gh run list --repo vertikon/seu-repo
gh run view <run-id> --log
```

### GoSec reportando falsos positivos

Adicione `// #nosec` com justificativa:

```go
// #nosec G204 - Command arguments are validated
cmd := exec.Command(tool, args...)
```

### Dependabot PRs demais

Ajuste `.github/dependabot.yml`:

```yaml
- package-ecosystem: "gomod"
  schedule:
    interval: "monthly"  # Era "weekly"
  open-pull-requests-limit: 5  # Era 10
```

## 📚 Recursos Adicionais

### Documentação

- [SECURITY.md](SECURITY.md) - Política de segurança
- [docs/SECURITY-OVERVIEW.md](docs/SECURITY-OVERVIEW.md) - Arquitetura de segurança
- [GitHub Security Best Practices](https://docs.github.com/en/code-security)

### Ferramentas

- [GoSec](https://github.com/securego/gosec) - Security scanner
- [CodeQL](https://codeql.github.com/) - Code analysis
- [Dependabot](https://docs.github.com/en/code-security/dependabot) - Dependency updates

### Compliance

- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [CIS Benchmarks](https://www.cisecurity.org/cis-benchmarks/)
- [NIST Framework](https://www.nist.gov/cyberframework)

## 🤝 Suporte

Problemas com o template? Abra uma issue ou entre em contato:

- **Platform Team**: platform@vertikon.com
- **Security Team**: security@vertikon.com

---

**Template Version**: 1.0.0  
**Last Updated**: 2025-01-19  
**Maintained by**: Vertikon Platform Team
