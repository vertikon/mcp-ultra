# Compilacao forcada com limpeza total
Write-Host "=== Compilacao Forcada MCP-Ultra ===" -ForegroundColor Cyan

Write-Host "`n[1/6] Limpando cache do Go..." -ForegroundColor Yellow
Remove-Item -Recurse -Force -ErrorAction SilentlyContinue "$env:GOPATH\pkg\mod\cache"
go clean -cache -modcache -testcache
Write-Host "OK Cache limpo" -ForegroundColor Green

Write-Host "`n[2/6] Removendo go.sum..." -ForegroundColor Yellow
Remove-Item -Force -ErrorAction SilentlyContinue go.sum
Write-Host "OK go.sum removido" -ForegroundColor Green

Write-Host "`n[3/6] Baixando modulos..." -ForegroundColor Yellow
go mod download
Write-Host "OK Modulos baixados" -ForegroundColor Green

Write-Host "`n[4/6] Executando go mod tidy..." -ForegroundColor Yellow
go mod tidy
Write-Host "OK Modulos organizados" -ForegroundColor Green

Write-Host "`n[5/6] Formatando codigo..." -ForegroundColor Yellow
gofmt -w .
Write-Host "OK Codigo formatado" -ForegroundColor Green

Write-Host "`n[6/6] Compilando..." -ForegroundColor Yellow
go build -a ./...

if ($LASTEXITCODE -eq 0) {
    Write-Host "`n========================================" -ForegroundColor Green
    Write-Host "  COMPILACAO BEM-SUCEDIDA!" -ForegroundColor Green
    Write-Host "========================================" -ForegroundColor Green
} else {
    Write-Host "`n========================================" -ForegroundColor Red
    Write-Host "  ERRO NA COMPILACAO" -ForegroundColor Red
    Write-Host "========================================" -ForegroundColor Red
}

Write-Host "`n=== Concluido ===" -ForegroundColor Cyan
