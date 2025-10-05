# Script de correcao de compilacao
Write-Host "=== Iniciando correcoes de compilacao ===" -ForegroundColor Cyan

# 1. Remover arquivo conflitante
Write-Host "`n[1/4] Removendo arquivos conflitantes..." -ForegroundColor Yellow
if (Test-Path "test_constants.go") {
    Remove-Item "test_constants.go" -Force
    Write-Host "OK test_constants.go removido" -ForegroundColor Green
}

# 2. Executar go mod tidy
Write-Host "`n[2/4] Resolvendo dependencias (go mod tidy)..." -ForegroundColor Yellow
go mod tidy
if ($LASTEXITCODE -eq 0) {
    Write-Host "OK Dependencias resolvidas" -ForegroundColor Green
}
else {
    Write-Host "ERRO ao resolver dependencias" -ForegroundColor Red
    Write-Host "Tentando go get github.com/vertikon/mcp-ultra..." -ForegroundColor Yellow
    go get github.com/vertikon/mcp-ultra
    go mod tidy
}

# 3. Formatar codigo
Write-Host "`n[3/4] Formatando codigo (gofmt)..." -ForegroundColor Yellow
gofmt -w .
Write-Host "OK Codigo formatado" -ForegroundColor Green

# 4. Tentar compilar
Write-Host "`n[4/4] Compilando projeto..." -ForegroundColor Yellow
go build ./...
if ($LASTEXITCODE -eq 0) {
    Write-Host "`nOK COMPILACAO BEM-SUCEDIDA!" -ForegroundColor Green
}
else {
    Write-Host "`nERRO Compilacao falhou. Verifique os erros acima." -ForegroundColor Red
    Write-Host "`nSe houver erros de semconv, execute:" -ForegroundColor Yellow
    Write-Host "go get go.opentelemetry.io/otel/semconv/v1.21.0" -ForegroundColor Cyan
}

Write-Host "`n=== Script concluido ===" -ForegroundColor Cyan
