# 🤖 Agent Implementation Prompt

## Tarefa: Corrigir 46 Issues

**Prioridade:** CRITICAL | **Esforço:** LARGE (1-2 days) | **Impacto:** HIGH - Afeta segurança/estabilidade

## 📋 Contexto

Foram identificados 46 problemas que precisam ser corrigidos:

- critical: Hardcoded credentials detected
- critical: Hardcoded credentials detected
- critical: Hardcoded credentials detected
- critical: Hardcoded credentials detected
- critical: Hardcoded credentials detected
... e mais 41 issues


## 📂 Arquivos Afetados

- `configs\secrets\template.yaml`
- `configs\security.yaml`
- `deploy\docker\prometheus-dev.yml`
- `deploy\k8s\deployment.yaml`
- `deploy\k8s\rbac.yaml`
- `deploy\k8s\secrets.template.yaml`
- `deploy\k8s\secrets.yaml`
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
   Revisar os 46 issues encontrados e priorizar correções

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
