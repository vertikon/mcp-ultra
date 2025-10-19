#!/usr/bin/env pwsh
# Script para corrigir todos os calls de logger em massa

$files = @(
    "internal/cache/distributed.go",
    "internal/lifecycle/health.go",
    "internal/metrics/business.go",
    "internal/ratelimit/distributed.go",
    "internal/tracing/business.go"
)

Write-Host "🔧 Corrigindo logger calls..." -ForegroundColor Cyan

foreach ($file in $files) {
    if (!(Test-Path $file)) {
        Write-Host "  ⚠ Arquivo não encontrado: $file" -ForegroundColor Red
        continue
    }
    
    Write-Host "`n  Processando $file..." -ForegroundColor Yellow
    
    $content = Get-Content $file -Raw
    
    # Pattern 1: String fields - "key", value → zap.String("key", value)
    $content = $content -replace '(logger\.(Info|Error|Warn|Debug)\([^)]+)\s+"([^"]+)",\s+([a-zA-Z][a-zA-Z0-9_.()]*)\s*([,)])', '${1} zap.String("${3}", ${4})${5}'
    
    # Pattern 2: Bool fields
    $content = $content -replace '(logger\.(Info|Error|Warn|Debug)\([^)]+)\s+"([^"]+)",\s+(true|false|[a-zA-Z][a-zA-Z0-9_.()]*\.Is[A-Z][a-zA-Z]*\(\)|[a-zA-Z][a-zA-Z0-9_.()]*Enabled)\s*([,)])', '${1} zap.Bool("${3}", ${4})${5}'
    
    # Pattern 3: Int64 fields
    $content = $content -replace '(logger\.(Info|Error|Warn|Debug)\([^)]+)\s+"([^"]+)",\s+([a-zA-Z][a-zA-Z0-9_.()]*Limit|[a-zA-Z][a-zA-Z0-9_.()]*Count)\s*([,)])', '${1} zap.Int64("${3}", ${4})${5}'
    
    # Pattern 4: Duration fields
    $content = $content -replace '(logger\.(Info|Error|Warn|Debug)\([^)]+)\s+"([^"]+)",\s+([a-zA-Z][a-zA-Z0-9_.()]*Window|[a-zA-Z][a-zA-Z0-9_.()]*Interval|[a-zA-Z][a-zA-Z0-9_.()]*Timeout|[a-zA-Z][a-zA-Z0-9_.()]*Duration)\s*([,)])', '${1} zap.Duration("${3}", ${4})${5}'
    
    # Pattern 5: Error fields - "error", err → zap.Error(err)
    $content = $content -replace '(logger\.(Info|Error|Warn|Debug)\([^)]+)\s+"error",\s+([a-zA-Z][a-zA-Z0-9_]*)\s*([,)])', '${1} zap.Error(${3})${4}'
    
    # Pattern 6: Float64 fields
    $content = $content -replace '(logger\.(Info|Error|Warn|Debug)\([^)]+)\s+"([^"]+)",\s+([a-zA-Z][a-zA-Z0-9_.()]*Rate|[a-zA-Z][a-zA-Z0-9_.()]*Value|[a-zA-Z][a-zA-Z0-9_.()]*Threshold)\s*([,)])', '${1} zap.Float64("${3}", ${4})${5}'
    
    # Salvar
    $content | Set-Content $file -NoNewline
    
    Write-Host "    ✓ Arquivo atualizado" -ForegroundColor Green
}

Write-Host "`n✅ Correções aplicadas. Executando go fmt..." -ForegroundColor Green
go fmt ./...

Write-Host "`n📊 Testando build..." -ForegroundColor Cyan
$buildResult = go build ./... 2>&1
if ($LASTEXITCODE -eq 0) {
    Write-Host "✅ BUILD SUCESSO!" -ForegroundColor Green
} else {
    Write-Host "⚠️ BUILD com erros (verificar manualmente):" -ForegroundColor Yellow
    $buildResult | Select-Object -First 20
}

