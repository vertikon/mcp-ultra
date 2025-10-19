# MCP Ultra - Docker Build Script (PowerShell)
# Versão: 1.0.0
# Data: 2025-10-19

param(
    [string]$Tag = "latest",
    [switch]$NoBuildCache,
    [switch]$Push,
    [string]$Registry = "vertikon"
)

Write-Host "🚀 MCP Ultra - Docker Build Script" -ForegroundColor Cyan
Write-Host "=================================" -ForegroundColor Cyan
Write-Host ""

# Verificar se Docker está rodando
Write-Host "🔍 Verificando Docker..." -ForegroundColor Yellow
try {
    $dockerVersion = docker --version
    Write-Host "✅ Docker encontrado: $dockerVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Docker não encontrado ou não está rodando!" -ForegroundColor Red
    Write-Host "   Instale Docker Desktop: https://www.docker.com/products/docker-desktop/" -ForegroundColor Red
    exit 1
}

# Verificar se estamos no diretório correto
if (-not (Test-Path "Dockerfile")) {
    Write-Host "❌ Dockerfile não encontrado!" -ForegroundColor Red
    Write-Host "   Execute este script no diretório raiz do projeto." -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "📦 Configuração do Build:" -ForegroundColor Yellow
Write-Host "   Tag: $Tag" -ForegroundColor White
Write-Host "   No Cache: $NoBuildCache" -ForegroundColor White
Write-Host "   Push: $Push" -ForegroundColor White
if ($Push) {
    Write-Host "   Registry: $Registry" -ForegroundColor White
}
Write-Host ""

# Build command
$buildCmd = "docker build -t mcp-ultra:$Tag"
if ($NoBuildCache) {
    $buildCmd += " --no-cache"
}
$buildCmd += " ."

Write-Host "🔨 Executando build..." -ForegroundColor Yellow
Write-Host "   Comando: $buildCmd" -ForegroundColor Gray
Write-Host ""

# Execute build
$startTime = Get-Date
Invoke-Expression $buildCmd
$buildExitCode = $LASTEXITCODE
$endTime = Get-Date
$duration = $endTime - $startTime

if ($buildExitCode -ne 0) {
    Write-Host ""
    Write-Host "❌ Build falhou!" -ForegroundColor Red
    exit $buildExitCode
}

Write-Host ""
Write-Host "✅ Build concluído com sucesso!" -ForegroundColor Green
Write-Host "   Tempo: $($duration.TotalSeconds) segundos" -ForegroundColor Gray
Write-Host ""

# Mostrar informações da imagem
Write-Host "📊 Informações da Imagem:" -ForegroundColor Yellow
docker images mcp-ultra:$Tag
Write-Host ""

# Tag e Push se solicitado
if ($Push) {
    $fullTag = "$Registry/mcp-ultra:$Tag"

    Write-Host "🏷️  Tagging imagem: $fullTag" -ForegroundColor Yellow
    docker tag mcp-ultra:$Tag $fullTag

    if ($LASTEXITCODE -ne 0) {
        Write-Host "❌ Tag falhou!" -ForegroundColor Red
        exit $LASTEXITCODE
    }

    Write-Host "📤 Pushing para registry..." -ForegroundColor Yellow
    docker push $fullTag

    if ($LASTEXITCODE -ne 0) {
        Write-Host "❌ Push falhou!" -ForegroundColor Red
        Write-Host "   Certifique-se de estar logado: docker login" -ForegroundColor Red
        exit $LASTEXITCODE
    }

    Write-Host "✅ Push concluído!" -ForegroundColor Green
}

Write-Host ""
Write-Host "🎉 Processo concluído!" -ForegroundColor Green
Write-Host ""
Write-Host "📝 Próximos passos:" -ForegroundColor Cyan
Write-Host "   1. Testar localmente:" -ForegroundColor White
Write-Host "      docker run -d -p 9655:9655 -p 9656:9656 mcp-ultra:$Tag" -ForegroundColor Gray
Write-Host ""
Write-Host "   2. Ou subir stack completo:" -ForegroundColor White
Write-Host "      docker-compose up -d" -ForegroundColor Gray
Write-Host ""
Write-Host "   3. Verificar health:" -ForegroundColor White
Write-Host "      curl http://localhost:9655/healthz" -ForegroundColor Gray
Write-Host ""
Write-Host "   4. Ver logs:" -ForegroundColor White
Write-Host "      docker logs -f mcp-ultra" -ForegroundColor Gray
Write-Host ""

# Exemplos de uso
Write-Host "💡 Exemplos de uso deste script:" -ForegroundColor Cyan
Write-Host "   .\docker-build.ps1                          # Build com tag 'latest'" -ForegroundColor Gray
Write-Host "   .\docker-build.ps1 -Tag v1.0.0              # Build com tag específica" -ForegroundColor Gray
Write-Host "   .\docker-build.ps1 -NoBuildCache            # Build sem cache" -ForegroundColor Gray
Write-Host "   .\docker-build.ps1 -Tag v1.0.0 -Push        # Build e push para registry" -ForegroundColor Gray
Write-Host "   .\docker-build.ps1 -Push -Registry myrepo   # Push para registry customizado" -ForegroundColor Gray
Write-Host ""
