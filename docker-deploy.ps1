#!/usr/bin/env pwsh
# MCP Ultra - Script de Deploy Docker Automatizado
# ================================================

param(
    [Parameter(Mandatory=$false)]
    [ValidateSet('dev', 'prod', 'staging')]
    [string]$Environment = 'dev',
    
    [Parameter(Mandatory=$false)]
    [switch]$Build,
    
    [Parameter(Mandatory=$false)]
    [switch]$NoBuild,
    
    [Parameter(Mandatory=$false)]
    [switch]$Down,
    
    [Parameter(Mandatory=$false)]
    [switch]$Clean,
    
    [Parameter(Mandatory=$false)]
    [switch]$Logs,
    
    [Parameter(Mandatory=$false)]
    [switch]$Status
)

# Cores para output
$Green = "`e[32m"
$Yellow = "`e[33m"
$Red = "`e[31m"
$Blue = "`e[34m"
$Reset = "`e[0m"

function Write-ColoredOutput {
    param(
        [string]$Message,
        [string]$Color = $Reset
    )
    Write-Host "${Color}${Message}${Reset}"
}

function Write-Step {
    param([string]$Message)
    Write-ColoredOutput "===> $Message" $Blue
}

function Write-Success {
    param([string]$Message)
    Write-ColoredOutput "✓ $Message" $Green
}

function Write-Warning {
    param([string]$Message)
    Write-ColoredOutput "⚠ $Message" $Yellow
}

function Write-Error {
    param([string]$Message)
    Write-ColoredOutput "✗ $Message" $Red
}

# Banner
Write-Host @"
${Blue}
╔═══════════════════════════════════════════════════╗
║         MCP Ultra - Docker Deploy Tool           ║
║              Ambiente: $Environment                    ║
╚═══════════════════════════════════════════════════╝
${Reset}
"@

# Verificar pré-requisitos
Write-Step "Verificando pré-requisitos..."

# Verificar Docker
if (!(Get-Command docker -ErrorAction SilentlyContinue)) {
    Write-Error "Docker não encontrado. Por favor, instale o Docker Desktop."
    exit 1
}

# Verificar Docker Compose
if (!(Get-Command docker-compose -ErrorAction SilentlyContinue)) {
    Write-Warning "docker-compose não encontrado, tentando usar 'docker compose'..."
    $UseDockerCompose = $false
} else {
    $UseDockerCompose = $true
}

# Verificar se Docker está rodando
try {
    docker ps | Out-Null
    Write-Success "Docker está rodando"
} catch {
    Write-Error "Docker não está rodando. Por favor, inicie o Docker Desktop."
    exit 1
}

# Criar arquivo .env se não existir
if (!(Test-Path ".env")) {
    Write-Warning "Arquivo .env não encontrado."
    if (Test-Path "env.template") {
        Write-Step "Criando .env a partir de env.template..."
        Copy-Item "env.template" ".env"
        Write-Success "Arquivo .env criado. Revise as configurações antes de prosseguir."
        Write-Warning "IMPORTANTE: Altere as senhas padrão no arquivo .env!"
    } else {
        Write-Error "Template .env não encontrado. Crie um arquivo .env manualmente."
        exit 1
    }
}

# Comandos docker-compose
$DockerComposeCmd = if ($UseDockerCompose) { "docker-compose" } else { "docker compose" }

# Processar comandos
if ($Status) {
    Write-Step "Status dos containers..."
    & $DockerComposeCmd ps
    exit 0
}

if ($Logs) {
    Write-Step "Mostrando logs..."
    & $DockerComposeCmd logs -f mcp-ultra
    exit 0
}

if ($Down) {
    Write-Step "Parando containers..."
    & $DockerComposeCmd down
    Write-Success "Containers parados"
    exit 0
}

if ($Clean) {
    Write-Warning "ATENÇÃO: Isso irá remover todos os containers, volumes e imagens!"
    $confirmation = Read-Host "Tem certeza? (yes/no)"
    if ($confirmation -eq "yes") {
        Write-Step "Limpando ambiente Docker..."
        & $DockerComposeCmd down -v --rmi all
        Write-Success "Ambiente limpo"
    } else {
        Write-ColoredOutput "Operação cancelada" $Yellow
    }
    exit 0
}

# Deploy normal
Write-Step "Iniciando deploy do MCP Ultra..."

# Parar containers existentes
Write-Step "Parando containers existentes..."
& $DockerComposeCmd down

if ($Build -or !$NoBuild) {
    Write-Step "Compilando aplicação..."
    
    # Limpar build anterior
    if (Test-Path "mcp-ultra.exe") {
        Remove-Item "mcp-ultra.exe" -Force
    }
    
    Write-Step "Building Docker images..."
    & $DockerComposeCmd build --no-cache mcp-ultra
    
    if ($LASTEXITCODE -ne 0) {
        Write-Error "Falha ao buildar imagem Docker"
        exit 1
    }
    
    Write-Success "Build concluído com sucesso"
}

# Subir containers
Write-Step "Iniciando containers..."
& $DockerComposeCmd up -d

if ($LASTEXITCODE -ne 0) {
    Write-Error "Falha ao iniciar containers"
    exit 1
}

# Aguardar serviços ficarem prontos
Write-Step "Aguardando serviços iniciarem..."
Start-Sleep -Seconds 10

# Verificar health dos containers
Write-Step "Verificando saúde dos containers..."
$containers = @("postgres", "redis", "nats", "mcp-ultra")
$allHealthy = $true

foreach ($container in $containers) {
    $health = docker inspect --format='{{.State.Health.Status}}' "mcp-ultra-${container}-1" 2>$null
    if (!$health) {
        $health = docker inspect --format='{{.State.Status}}' "mcp-ultra-${container}-1" 2>$null
    }
    
    if ($health -eq "healthy" -or $health -eq "running") {
        Write-Success "${container}: OK"
    } else {
        Write-Warning "${container}: $health"
        $allHealthy = $false
    }
}

Write-Host ""
Write-ColoredOutput "═══════════════════════════════════════════════════" $Blue
Write-ColoredOutput "           Deploy Status" $Blue
Write-ColoredOutput "═══════════════════════════════════════════════════" $Blue

if ($allHealthy) {
    Write-Success "Todos os serviços estão rodando!"
} else {
    Write-Warning "Alguns serviços podem não estar completamente prontos."
    Write-ColoredOutput "Execute '$DockerComposeCmd logs' para mais detalhes." $Yellow
}

Write-Host ""
Write-ColoredOutput "Serviços disponíveis:" $Blue
Write-Host "  • MCP Ultra API:     http://localhost:9655"
Write-Host "  • Metrics:           http://localhost:9656/metrics"
Write-Host "  • Health Check:      http://localhost:9655/healthz"
Write-Host "  • PostgreSQL:        localhost:5432"
Write-Host "  • Redis:             localhost:6379"
Write-Host "  • NATS:              localhost:4222"
Write-Host "  • NATS Monitoring:   http://localhost:8222"
Write-Host "  • Jaeger UI:         http://localhost:16686"
Write-Host "  • Prometheus:        http://localhost:9090"
Write-Host "  • Grafana:           http://localhost:3000"

Write-Host ""
Write-ColoredOutput "Comandos úteis:" $Yellow
Write-Host "  • Ver logs:          $DockerComposeCmd logs -f mcp-ultra"
Write-Host "  • Parar:             $DockerComposeCmd down"
Write-Host "  • Restart:           $DockerComposeCmd restart mcp-ultra"
Write-Host "  • Shell:             $DockerComposeCmd exec mcp-ultra sh"
Write-Host "  • Status:            $DockerComposeCmd ps"

Write-Host ""
Write-Success "Deploy concluído com sucesso!"
Write-Host ""

