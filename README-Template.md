# %PROJECT_NAME%

> Serviço gerado a partir do **mcp-ultra** (template canônico).
> Status: ✅ pronto para CI/CD, coverage e observabilidade "de fábrica".

## 🚀 Quickstart

```bash
# Rodar testes rápidos
go test ./internal/handlers ./tests/integration ./tests/smoke -count=1

# Coverage + HTML
go test ./internal/handlers ./tests/integration ./tests/smoke -coverpkg=./... -coverprofile=coverage.out
go tool cover -func coverage.out > coverage_func.txt
go tool cover -html coverage.out > coverage.html
```

## 📈 Coverage

Histórico (se habilitado): [`docs/coverage_history.md`](docs/coverage_history.md)
Badge (gerado no CI): ![coverage](docs/badges/coverage.svg)

## 🧪 Scripts úteis

- `tools\update-coverage-history.ps1` → atualiza histórico/csv/badge
- `tools\coverage-badge-check.ps1` → confere consistência do badge
- `tools\latency-stability-check.ps1` → mede latência (opcional)
- `tools\scaffold-from-ultra.ps1` → cria um novo serviço a partir do template

## 🛠️ CI (resumo)

1. Roda testes do core estável
2. Gera `coverage.out` / `coverage.html` / `coverage_func.txt`
3. Se houver scripts: atualiza `docs/coverage_history.*` e `docs/badges/coverage.svg`
4. Faz commit automático desses artefatos apenas na branch `main`

---

**Gerado via mcp-ultra Template** ✨
