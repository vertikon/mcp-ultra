# Script de Correção Automatizada de Imports
# Substitui imports privados por mcp-ultra-fix

Write-Host "🔧 Iniciando correção de imports no template mcp-ultra..." -ForegroundColor Cyan

$replacements = @{
    'github.com/vertikon/mcp-ultra/internal/config' = 'github.com/vertikon/mcp-ultra-fix/pkg/config'
    'github.com/vertikon/mcp-ultra/internal/handlers' = 'github.com/vertikon/mcp-ultra-fix/pkg/handlers'
    'github.com/vertikon/mcp-ultra/internal/services' = 'github.com/vertikon/mcp-ultra-fix/pkg/services'
    'github.com/vertikon/mcp-ultra/internal/repository' = 'github.com/vertikon/mcp-ultra-fix/pkg/repository'
    'github.com/vertikon/mcp-ultra/internal/domain' = 'github.com/vertikon/mcp-ultra-fix/pkg/domain'
    'github.com/vertikon/mcp-ultra/internal/telemetry' = 'github.com/vertikon/mcp-ultra-fix/pkg/telemetry'
    'github.com/vertikon/mcp-ultra/internal/security' = 'github.com/vertikon/mcp-ultra-fix/pkg/security'
    'github.com/vertikon/mcp-ultra/internal/constants' = 'github.com/vertikon/mcp-ultra-fix/pkg/constants'
    'github.com/vertikon/mcp-ultra/internal/observability' = 'github.com/vertikon/mcp-ultra-fix/pkg/observability'
    'github.com/vertikon/mcp-ultra/internal/features' = 'github.com/vertikon/mcp-ultra-fix/pkg/features'
    'github.com/vertikon/mcp-ultra/internal/compliance' = 'github.com/vertikon/mcp-ultra-fix/pkg/compliance'
    'github.com/vertikon/mcp-ultra/internal/testhelpers' = 'github.com/vertikon/mcp-ultra-fix/pkg/testhelpers'
    'github.com/vertikon/mcp-ultra/internal/http' = 'github.com/vertikon/mcp-ultra-fix/pkg/http'
    'github.com/vertikon/mcp-ultra/internal/events' = 'github.com/vertikon/mcp-ultra-fix/pkg/nats'
    'github.com/vertikon/mcp-ultra/internal/cache' = 'github.com/vertikon/mcp-ultra-fix/pkg/cache'
    'github.com/vertikon/mcp-ultra/internal/ratelimit' = 'github.com/vertikon/mcp-ultra-fix/pkg/ratelimit'
    'github.com/vertikon/mcp-ultra/internal/lifecycle' = 'github.com/vertikon/mcp-ultra-fix/pkg/lifecycle'
    'github.com/vertikon/mcp-ultra/internal/metrics' = 'github.com/vertikon/mcp-ultra-fix/pkg/metrics'
    'github.com/vertikon/mcp-ultra/internal/tracing' = 'github.com/vertikon/mcp-ultra-fix/pkg/tracing'
    'github.com/vertikon/mcp-ultra/internal/ai' = 'github.com/vertikon/mcp-ultra-fix/pkg/ai'
    'github.com/vertikon/mcp-ultra/internal/repository/postgres' = 'github.com/vertikon/mcp-ultra-fix/pkg/postgres'
    'github.com/vertikon/mcp-ultra/internal/repository/redis' = 'github.com/vertikon/mcp-ultra-fix/pkg/redis'
}

$totalFiles = 0
$totalReplacements = 0

foreach ($old in $replacements.Keys) {
    $new = $replacements[$old]

    Write-Host "`n🔍 Procurando: $old" -ForegroundColor Yellow
    Write-Host "   ✅ Substituindo por: $new" -ForegroundColor Green

    $files = Get-ChildItem -Recurse -Filter *.go | Select-String -Pattern $old -List | Select-Object -ExpandProperty Path -Unique

    foreach ($file in $files) {
        $content = Get-Content $file -Raw
        $newContent = $content -replace [regex]::Escape($old), $new

        if ($content -ne $newContent) {
            Set-Content $file -Value $newContent -NoNewline
            $totalFiles++
            $totalReplacements++
            Write-Host "   📝 Corrigido: $file" -ForegroundColor Cyan
        }
    }
}

# Tratar test/mocks (remover imports)
Write-Host "`n🗑️  Removendo imports de test/mocks..." -ForegroundColor Yellow
$mockFiles = Get-ChildItem -Recurse -Filter *.go | Select-String -Pattern 'github.com/vertikon/mcp-ultra/test/mocks' -List | Select-Object -ExpandProperty Path -Unique

foreach ($file in $mockFiles) {
    Write-Host "   ⚠️  Arquivo com mock privado: $file" -ForegroundColor Red
    Write-Host "   💡 AÇÃO MANUAL: Criar mocks locais ou usar testify/gomock" -ForegroundColor Magenta
}

Write-Host "`n✅ Correção concluída!" -ForegroundColor Green
Write-Host "   📊 Total de arquivos corrigidos: $totalFiles" -ForegroundColor Cyan
Write-Host "   🔄 Total de substituições: $totalReplacements" -ForegroundColor Cyan

Write-Host "`n📦 Próximos passos:" -ForegroundColor Yellow
Write-Host "   1. go mod tidy" -ForegroundColor White
Write-Host "   2. go build ./..." -ForegroundColor White
Write-Host "   3. go test ./..." -ForegroundColor White
Write-Host "   4. Validar com enhanced_validator_v7" -ForegroundColor White
