# Script completo de correcao e compilacao
Write-Host "=== MCP-Ultra: Correcao e Compilacao ===" -ForegroundColor Cyan

# 1. Limpar cache
Write-Host "`n[1/5] Limpando cache..." -ForegroundColor Yellow
go clean -modcache 2>$null
Write-Host "OK Cache limpo" -ForegroundColor Green

# 2. Baixar dependencias especificas
Write-Host "`n[2/5] Baixando dependencias..." -ForegroundColor Yellow
go get go.opentelemetry.io/otel/semconv/v1.26.0
go mod download github.com/alicebob/miniredis/v2
go mod download
Write-Host "OK Dependencias baixadas" -ForegroundColor Green

# 3. Executar go mod tidy
Write-Host "`n[3/5] Organizando modulos (go mod tidy)..." -ForegroundColor Yellow
go mod tidy
Write-Host "OK Modulos organizados" -ForegroundColor Green

# 4. Formatar codigo
Write-Host "`n[4/5] Formatando codigo..." -ForegroundColor Yellow
gofmt -w .
Write-Host "OK Codigo formatado" -ForegroundColor Green

# 5. Compilar
Write-Host "`n[5/5] Compilando projeto..." -ForegroundColor Yellow
go build ./...

if ($LASTEXITCODE -eq 0) {
    Write-Host "`n========================================" -ForegroundColor Green
    Write-Host "  COMPILACAO BEM-SUCEDIDA!" -ForegroundColor Green
    Write-Host "========================================" -ForegroundColor Green
    Write-Host "`nProximo passo: Execute o validador" -ForegroundColor Cyan
    Write-Host "cd E:\vertikon\.ecosistema-vertikon\mcp-tester-system" -ForegroundColor Gray
    Write-Host "go run enhanced_validator_v4.go E:/vertikon/business/SaaS/templates/mcp-ultra" -ForegroundColor Gray
} else {
    Write-Host "`n========================================" -ForegroundColor Red
    Write-Host "  ERRO NA COMPILACAO" -ForegroundColor Red
    Write-Host "========================================" -ForegroundColor Red
    Write-Host "`nVerifique os erros acima" -ForegroundColor Yellow
}

Write-Host "`n=== Script concluido ===" -ForegroundColor Cyan
