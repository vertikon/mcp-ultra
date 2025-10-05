# 🤖 Agent Implementation Prompt

## Tarefa: Remover Secrets Hardcoded

**Prioridade:** CRITICAL | **Esforço:** SMALL (1-2h) | **Impacto:** HIGH - Afeta segurança/estabilidade

## 📋 Contexto

Foram identificados 2 problemas que precisam ser corrigidos:

- critical: JWT Token
- critical: JWT Token


## 📂 Arquivos Afetados

- `smart_validation_report.json`
- `test_constants.go`

## 🔧 Passos de Implementação

1. **Identificar todos os secrets hardcoded**
   Listar todos os secrets que precisam ser movidos para variáveis de ambiente

   ```bash
   # Buscar secrets no código
   grep -r "password" --include="*.go" .
   grep -r "api_key" --include="*.go" .
   grep -r "secret" --include="*.go" .
   ```

2. **Criar arquivo .env.example**
   Criar template de variáveis de ambiente sem valores sensíveis

   ```bash
   # .env.example
   DATABASE_URL=postgresql://user:password@localhost:5432/dbname
   API_KEY=your_api_key_here
   JWT_SECRET=your_jwt_secret_here
   SMTP_PASSWORD=your_smtp_password_here
   ```

3. **Criar estrutura de configuração**
   Implementar carregamento de configurações de ambiente

   ```go
   package config
   
   import (
       "github.com/kelseyhightower/envconfig"
   )
   
   type Config struct {
       DatabaseURL  string `envconfig:"DATABASE_URL" required:"true"`
       APIKey       string `envconfig:"API_KEY" required:"true"`
       JWTSecret    string `envconfig:"JWT_SECRET" required:"true"`
       SMTPPassword string `envconfig:"SMTP_PASSWORD" required:"true"`
   }
   
   func Load() (*Config, error) {
       var cfg Config
       if err := envconfig.Process("", &cfg); err != nil {
           return nil, err
       }
       return &cfg, nil
   }
   ```

4. **Substituir secrets hardcoded por configuração**
   Remover secrets do código e usar configuração

   ```go
   // ❌ ANTES (Inseguro)
   db, err := sql.Open("postgres", "postgresql://user:mypassword123@localhost:5432/db")
   
   // ✅ DEPOIS (Seguro)
   cfg, err := config.Load()
   if err != nil {
       log.Fatal(err)
   }
   db, err := sql.Open("postgres", cfg.DatabaseURL)
   ```

5. **Adicionar .env ao .gitignore**
   Garantir que secrets não sejam commitados

   ```bash
   # .gitignore
   .env
   .env.local
   *.pem
   *.key
   config/secrets.yaml
   ```

## 📝 Exemplos de Código

## 🧪 Guia de Testes


### Testes de Segurança

1. **Scan de Secrets**
   ```bash
   gitleaks detect --source=.
   ```

2. **Verificar Variáveis de Ambiente**
   ```bash
   grep -r "password\|secret\|key" --include="*.go" .
   ```

3. **Testar Configuração**
   ```bash
   go test ./internal/config/... -v
   ```


## ✅ Critérios de Aceitação

Após implementar as mudanças:

1. Execute os testes: `go test ./...`
2. Valide o código: `golangci-lint run`
3. Execute o validador: `make validate`
4. Confirme que os issues foram resolvidos

## 📚 Referências

- [OWASP - Sensitive Data Exposure](https://owasp.org/www-project-top-ten/2017/A3_2017-Sensitive_Data_Exposure)
- [12 Factor App - Config](https://12factor.net/config)

## 🎯 Checklist de Implementação

- [ ] Ler e entender o contexto completo
- [ ] Analisar os arquivos afetados
- [ ] Implementar as mudanças conforme os exemplos
- [ ] Adicionar/atualizar testes
- [ ] Executar testes localmente
- [ ] Validar com linter
- [ ] Executar validação completa
- [ ] Documentar mudanças no commit

---

**Nota para o Agente:** Este prompt foi gerado automaticamente pelo MCP Tester System. Siga as instruções cuidadosamente e valide cada etapa antes de prosseguir.
