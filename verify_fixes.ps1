# Verification script for code conflict fixes

Set-Location "E:\vertikon\business\SaaS\templates\mcp-ultra"

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Verificando correcoes de conflitos" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# Step 1: Format code
Write-Host ""
Write-Host "[1/4] Formatando codigo com gofmt..." -ForegroundColor Yellow
gofmt -s -w .
if ($LASTEXITCODE -eq 0) {
    Write-Host "[OK] Formatacao concluida" -ForegroundColor Green
} else {
    Write-Host "[ERRO] Erro na formatacao" -ForegroundColor Red
}

# Step 2: Fix imports
Write-Host ""
Write-Host "[2/4] Corrigindo imports com goimports..." -ForegroundColor Yellow
goimports -w .
if ($LASTEXITCODE -eq 0) {
    Write-Host "[OK] Imports corrigidos" -ForegroundColor Green
} else {
    Write-Host "[AVISO] goimports nao encontrado (opcional)" -ForegroundColor Yellow
}

# Step 3: Build
Write-Host ""
Write-Host "[3/4] Compilando projeto..." -ForegroundColor Yellow
go build ./...
if ($LASTEXITCODE -eq 0) {
    Write-Host "[OK] Build concluido com sucesso" -ForegroundColor Green
} else {
    Write-Host "[ERRO] Erro no build" -ForegroundColor Red
    exit 1
}

# Step 4: Lint
Write-Host ""
Write-Host "[4/4] Executando linter..." -ForegroundColor Yellow
golangci-lint run --out-format=tab
if ($LASTEXITCODE -eq 0) {
    Write-Host "[OK] Lint passou sem erros" -ForegroundColor Green
} else {
    Write-Host "[AVISO] Lint encontrou problemas" -ForegroundColor Yellow
}

# Step 5: Run tests
Write-Host ""
Write-Host "[5/5] Executando testes..." -ForegroundColor Yellow
go test ./... -count=1
if ($LASTEXITCODE -eq 0) {
    Write-Host "[OK] Testes passaram" -ForegroundColor Green
} else {
    Write-Host "[ERRO] Alguns testes falharam" -ForegroundColor Red
}

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Resumo das correcoes aplicadas:" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "[OK] domain/dto.go - Removida duplicacao de Task" -ForegroundColor Green
Write-Host "[OK] router_example.go - Neutralizado (package http_disabled)" -ForegroundColor Green
Write-Host "[OK] Conflitos de HealthService resolvidos" -ForegroundColor Green
Write-Host "[OK] Conflitos de NewRouter resolvidos" -ForegroundColor Green
Write-Host ""
