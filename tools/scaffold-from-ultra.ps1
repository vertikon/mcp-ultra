param(
  [Parameter(Mandatory=$true)][string]$NewPath,
  [Parameter(Mandatory=$true)][string]$NewModule
)

$ErrorActionPreference = "Stop"

Write-Host ""
Write-Host "🌱 Scaffold from mcp-ultra" -ForegroundColor Cyan
Write-Host "──────────────────────────" -ForegroundColor DarkCyan
Write-Host ""

# 1) Copiar (sem .git)
Write-Host "📦 Copiando estrutura..." -ForegroundColor Yellow
robocopy . $NewPath /E /XD .git .github\workflows\_cache 1> $null
if ($LASTEXITCODE -gt 3) { throw "Falha ao copiar ($LASTEXITCODE)" }

# 2) Limpar VCS/artefatos
Write-Host "🧹 Limpando artefatos..." -ForegroundColor Yellow
Get-ChildItem -Path $NewPath -Recurse -Force -Include '.git','coverage.out','coverage.html','coverage_func.txt' |
  Remove-Item -Recurse -Force -ErrorAction SilentlyContinue

# 3) Ajustar module
Write-Host "📝 Ajustando go.mod..." -ForegroundColor Yellow
$gomod = Join-Path $NewPath 'go.mod'
if (Test-Path $gomod) {
  (Get-Content $gomod) -replace '^module\s+.*', "module $NewModule" | Set-Content $gomod -Encoding UTF8
}

# 4) Reescrever imports do módulo antigo
Write-Host "🔄 Reescrevendo imports..." -ForegroundColor Yellow
$old = 'github.com/vertikon/mcp-ultra'
Get-ChildItem -Path $NewPath -Recurse -Include *.go,*.md |
  ForEach-Object {
    (Get-Content $_.FullName) -replace [regex]::Escape($old), $NewModule | Set-Content $_.FullName -Encoding UTF8
  }

# 5) go mod tidy + smoke test
Write-Host "⚙️  Rodando go mod tidy..." -ForegroundColor Yellow
Push-Location $NewPath
try {
  go mod tidy

  Write-Host "🧪 Smoke test..." -ForegroundColor Yellow
  go test ./... -count=1

  Write-Host ""
  Write-Host "✅ Scaffold criado com sucesso!" -ForegroundColor Green
  Write-Host ""
  Write-Host "📂 Local: $NewPath" -ForegroundColor White
  Write-Host "📦 Module: $NewModule" -ForegroundColor White
  Write-Host ""

} catch {
  Write-Host ""
  Write-Host "⚠️  Erro durante setup: $_" -ForegroundColor Red
  Write-Host ""
  throw
} finally {
  Pop-Location
}
