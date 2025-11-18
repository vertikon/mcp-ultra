# 🤖 Agent Implementation Prompt

## Tarefa: Corrigir 257 Issues

**Prioridade:** CRITICAL | **Esforço:** LARGE (1-2 days) | **Impacto:** HIGH - Afeta segurança/estabilidade

## 📋 Contexto

Foram identificados 257 problemas que precisam ser corrigidos:

- medium: Memory exhaustion in multipart form parsing in net/http
- high: HTTP/2 CONTINUATION flood in net/http
- critical: Hardcoded credentials detected
- critical: Hardcoded credentials detected
- critical: Hardcoded credentials detected
... e mais 252 issues


## 📂 Arquivos Afetados

- `config\telemetry.yaml`
- `configs\security.yaml`
- `deploy\docker\prometheus-dev.yml`
- `deploy\k8s\deployment.yaml`
- `deploy\k8s\rbac.yaml`
- `gitleaks\cmd\generate\config\rules\azure.go`
- `gitleaks\cmd\generate\config\rules\cloudflare.go`
- `gitleaks\cmd\generate\config\rules\facebook.go`
- `gitleaks\cmd\generate\config\rules\flyio.go`
- `gitleaks\cmd\generate\config\rules\gcp.go`
- `gitleaks\cmd\generate\config\rules\generic.go`
- `gitleaks\cmd\generate\config\rules\grafana.go`
- `gitleaks\cmd\generate\config\rules\hashicorp.go`
- `gitleaks\cmd\generate\config\rules\heroku.go`
- `gitleaks\cmd\generate\config\rules\hubspot.go`
- `gitleaks\cmd\generate\config\rules\huggingface.go`
- `gitleaks\cmd\generate\config\rules\kubernetes.go`
- `gitleaks\cmd\generate\config\rules\mailchimp.go`
- `gitleaks\cmd\generate\config\rules\nuget.go`
- `gitleaks\cmd\generate\config\rules\octopusdeploy.go`
- `gitleaks\cmd\generate\config\rules\okta.go`
- `gitleaks\cmd\generate\config\rules\openshift.go`
- `gitleaks\cmd\generate\config\rules\prefect.go`
- `gitleaks\cmd\generate\config\rules\privateai.go`
- `gitleaks\cmd\generate\config\rules\readme.go`
- `gitleaks\cmd\generate\config\rules\scalingo.go`
- `gitleaks\cmd\generate\config\rules\sentry.go`
- `gitleaks\cmd\generate\config\rules\snyk.go`
- `gitleaks\cmd\generate\config\rules\sonar.go`
- `gitleaks\cmd\generate\config\rules\sumologic.go`
- `gitleaks\cmd\generate\config\rules\telegram.go`
- `gitleaks\cmd\generate\config\utils\generate.go`
- `gitleaks\cmd\generate\config\utils\generate_test.go`
- `gitleaks\config\allowlist_test.go`
- `gitleaks\config\config_test.go`
- `gitleaks\detect\baseline_test.go`
- `gitleaks\detect\codec\decoder_test.go`
- `gitleaks\detect\detect_test.go`
- `gitleaks\detect\reader_test.go`
- `gitleaks\report\csv_test.go`
- `gitleaks\report\finding.go`
- `gitleaks\report\finding_test.go`
- `gitleaks\report\json_test.go`
- `gitleaks\report\junit_test.go`
- `gitleaks\report\sarif_test.go`
- `gitleaks\report\template_test.go`
- `go.mod`
- `internal\compliance\audit_logger.go`
- `internal\features\flags.go`
- `internal\features\manager_test.go`
- `internal\grpc\server\system_server.go`
- `internal\lifecycle\deployment.go`
- `internal\repository\postgres\task_repository.go`
- `test\integration\database_integration_test.go`
- `test\property\task_properties_test.go`

## 🔧 Passos de Implementação

1. **Analisar issues identificados**
   Revisar os 257 issues encontrados e priorizar correções

2. **Implementar correções**
   Aplicar as correções necessárias seguindo as soluções sugeridas

3. **Adicionar testes**
   Criar testes para validar as correções implementadas

4. **Validar mudanças**
   Executar testes e validação completa

## 📝 Exemplos de Código

## 🧪 Guia de Testes


### Testes Gerais

1. **Executar Todos os Testes**
   ```bash
   go test ./... -v -cover
   ```

2. **Lint**
   ```bash
   golangci-lint run
   ```


## ✅ Critérios de Aceitação

Após implementar as mudanças:

1. Execute os testes: `go test ./...`
2. Valide o código: `golangci-lint run`
3. Execute o validador: `make validate`
4. Confirme que os issues foram resolvidos

## 📚 Referências

- [Go Best Practices](https://golang.org/doc/effective_go)

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
