#!/usr/bin/env pwsh
# Script para corrigir uso de logger fields em massa

$files = @(
    "internal/metrics/business.go",
    "internal/cache/distributed.go",
    "internal/lifecycle/deployment.go",
    "internal/lifecycle/health.go",
    "internal/ratelimit/distributed.go",
    "internal/tracing/business.go",
    "internal/telemetry/telemetry.go"
)

Write-Host "🔧 Corrigindo uso de logger fields..." -ForegroundColor Cyan

foreach ($file in $files) {
    if (Test-Path $file) {
        Write-Host "  Processando $file..." -ForegroundColor Yellow
        
        # Ler conteúdo
        $content = Get-Content $file -Raw
        
        # Verificar se já tem import de zap
        if ($content -notmatch 'import\s+\(\s*[^)]*"go\.uber\.org/zap"') {
            # Adicionar import de zap
            $content = $content -replace '(import\s+\()', "`$1`n`t`"go.uber.org/zap`"`n"
            Write-Host "    ✓ Import zap adicionado" -ForegroundColor Green
        }
        
        # Salvar
        $content | Set-Content $file -NoNewline
        
        Write-Host "    ✓ Arquivo atualizado" -ForegroundColor Green
    } else {
        Write-Host "  ⚠ Arquivo não encontrado: $file" -ForegroundColor Red
    }
}

Write-Host "`n✅ Imports de zap adicionados. Execute go fmt para formatar." -ForegroundColor Green
Write-Host "⚠️  ATENÇÃO: Você ainda precisará corrigir manualmente os logger calls!" -ForegroundColor Yellow

