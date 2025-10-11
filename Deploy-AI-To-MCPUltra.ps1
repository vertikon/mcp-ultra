param(
  [Parameter(Mandatory=$true)][string]$ZipPath,
  [string]$TargetRoot = "E:\vertikon\business\SaaS\templates\mcp-ultra\templates\ai"
)

$ErrorActionPreference = "Stop"

function New-Log { param([string]$Msg) Write-Host "[AI-Update] $Msg" }

if (-not (Test-Path $ZipPath)) {
  throw "ZipPath '$ZipPath' não encontrado."
}

# Backup se já existir
if (Test-Path $TargetRoot) {
  $stamp = Get-Date -Format "yyyyMMdd-HHmmss"
  $backup = "$TargetRoot-backup-$stamp"
  New-Log "Backup do diretório existente em: $backup"
  Copy-Item -Path $TargetRoot -Destination $backup -Recurse -Force
} else {
  New-Log "Criando diretório alvo: $TargetRoot"
  New-Item -ItemType Directory -Path $TargetRoot | Out-Null
}

# Expand zip para temp
$temp = Join-Path $env:TEMP ("mcp-ultra-ai-update-" + [guid]::NewGuid().ToString())
New-Log "Extraindo ZIP para: $temp"
Expand-Archive -Path $ZipPath -DestinationPath $temp

# O ZIP contém raiz com ai/* e README
$srcAi = Join-Path $temp "ai"
$srcReadme = Join-Path $temp "README.md"

if (-not (Test-Path $srcAi)) { throw "Estrutura inesperada no ZIP: pasta 'ai' não encontrada." }

# Garantir subpastas
$folders = @("config","examples","nats-schemas","telemetry")
foreach ($f in $folders) {
  $p = Join-Path $TargetRoot $f
  if (-not (Test-Path $p)) { New-Item -ItemType Directory -Path $p | Out-Null }
}

# Copiar conteúdos
New-Log "Copiando arquivos de configuração..."
Copy-Item -Path (Join-Path $srcAi "config\*") -Destination (Join-Path $TargetRoot "config") -Recurse -Force
Copy-Item -Path (Join-Path $srcAi "examples\*") -Destination (Join-Path $TargetRoot "examples") -Recurse -Force
Copy-Item -Path (Join-Path $srcAi "nats-schemas\*") -Destination (Join-Path $TargetRoot "nats-schemas") -Recurse -Force
Copy-Item -Path (Join-Path $srcAi "telemetry\*") -Destination (Join-Path $TargetRoot "telemetry") -Recurse -Force
Copy-Item -Path (Join-Path $srcAi "feature_flags.json") -Destination (Join-Path $TargetRoot "feature_flags.json") -Force
Copy-Item -Path (Join-Path $srcAi "MIGRATION.md") -Destination (Join-Path $TargetRoot "MIGRATION.md") -Force

# README no nível superior do ai/
Copy-Item -Path $srcReadme -Destination (Join-Path $TargetRoot "README.md") -Force

# Sumário
New-Log "Estrutura final:"
Get-ChildItem -Recurse -File $TargetRoot | Select-Object FullName, Length | Format-Table -AutoSize

New-Log "Concluído com sucesso."
New-Log "Próximos passos:"
New-Log "1) Mesclar 'examples\.env.mcp.example' no .env.example do MCP."
New-Log "2) Atualizar Inventário usando 'examples\inventory-registry.example.json'."
New-Log "3) DRY-RUN: ENABLE_AI=true, PROVIDER_PRIMARY=local, AI_CANARY_PERCENT=0."
