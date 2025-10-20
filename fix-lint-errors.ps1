# Script de Correção em Massa de Erros de Lint
# Aplica substituições automáticas para resolver depguard e revive

$projectRoot = "E:\vertikon\business\SaaS\templates\mcp-ultra"
Set-Location $projectRoot

Write-Host "🔧 Iniciando correções em massa..." -ForegroundColor Cyan

# 1. Corrigir imports UUID
Write-Host "`n📦 Corrigindo imports UUID..." -ForegroundColor Yellow
$files = @(
    "internal\handlers\http\task_handlers.go",
    "internal\handlers\http\router_test.go",
    "internal\compliance\framework.go",
    "test\component\task_service_test.go"
)

foreach ($file in $files) {
    if (Test-Path $file) {
        Write-Host "  • $file"
        $content = Get-Content $file -Raw

        # Substituir import
        $content = $content -replace '"github\.com/google/uuid"', '"github.com/vertikon/mcp-ultra/pkg/types"'

        # Substituir usos
        $content = $content -replace '\buuid\.UUID\b', 'types.UUID'
        $content = $content -replace '\buuid\.New\(\)', 'types.New()'
        $content = $content -replace '\buuid\.Nil\b', 'types.Nil'
        $content = $content -replace '\buuid\.Parse\(', 'types.Parse('
        $content = $content -replace '\buuid\.MustParse\(', 'types.MustParse('

        Set-Content $file $content -NoNewline
    }
}

# 2. Corrigir unused parameters (renomear para _)
Write-Host "`n🔧 Corrigindo unused parameters..." -ForegroundColor Yellow

# Mapeamento de arquivos e correções específicas
$unusedParams = @{
    "internal\handlers\health.go" = @(
        @{ Pattern = 'func \(h \*HealthHandler\) Live\(w http\.ResponseWriter, r \*http\.Request\)'; Replace = 'func (h *HealthHandler) Live(w http.ResponseWriter, _ *http.Request)' }
        @{ Pattern = 'func \(h \*HealthHandler\) Ready\(w http\.ResponseWriter, r \*http\.Request\)'; Replace = 'func (h *HealthHandler) Ready(w http.ResponseWriter, _ *http.Request)' }
        @{ Pattern = 'func \(h \*HealthHandler\) Health\(w http\.ResponseWriter, r \*http\.Request\)'; Replace = 'func (h *HealthHandler) Health(w http.ResponseWriter, _ *http.Request)' }
    )
    "internal\security\tls.go" = @(
        # ioutil também será corrigido aqui
    )
}

foreach ($file in $unusedParams.Keys) {
    if (Test-Path $file) {
        Write-Host "  • $file"
        $content = Get-Content $file -Raw

        foreach ($fix in $unusedParams[$file]) {
            $content = $content -replace $fix.Pattern, $fix.Replace
        }

        Set-Content $file $content -NoNewline
    }
}

# 3. Corrigir ioutil deprecated
Write-Host "`n📦 Corrigindo ioutil deprecated..." -ForegroundColor Yellow
$ioutilFile = "internal\security\tls.go"

if (Test-Path $ioutilFile) {
    Write-Host "  • $ioutilFile"
    $content = Get-Content $ioutilFile -Raw

    # Substituir import
    $content = $content -replace '"io/ioutil"', ''

    # Adicionar imports necessários se não existirem
    if ($content -notmatch '"io"') {
        $content = $content -replace '(import \()', '$1`n`t"io"'
    }
    if ($content -notmatch '"os"') {
        $content = $content -replace '(import \()', '$1`n`t"os"'
    }

    # Substituir funções
    $content = $content -replace '\bioutil\.ReadFile\(', 'os.ReadFile('
    $content = $content -replace '\bioutil\.WriteFile\(', 'os.WriteFile('
    $content = $content -replace '\bioutil\.ReadAll\(', 'io.ReadAll('
    $content = $content -replace '\bioutil\.NopCloser\(', 'io.NopCloser('

    Set-Content $ioutilFile $content -NoNewline
}

# 4. Corrigir SA9003 (empty branches)
Write-Host "`n🔧 Corrigindo empty branches..." -ForegroundColor Yellow
$emptyBranchFiles = @(
    "internal\compliance\framework.go",
    "internal\config\config.go"
)

foreach ($file in $emptyBranchFiles) {
    if (Test-Path $file) {
        Write-Host "  • $file"
        $content = Get-Content $file -Raw

        # Padrão: if err := something; err != nil { // comment }
        # Substituir por: _ = something
        $content = $content -replace 'if err := ([^;]+); err != nil \{\s*//[^\}]*\}', '_ = $1'

        Set-Content $file $content -NoNewline
    }
}

# 5. Executar gofmt e goimports
Write-Host "`n🎨 Formatando código..." -ForegroundColor Yellow
& gofmt -w .
& goimports -w .

Write-Host "`n✅ Correções aplicadas!" -ForegroundColor Green
Write-Host "`n📊 Executando validação..." -ForegroundColor Cyan

# 6. Executar linter para verificar
& golangci-lint run --out-format=colored

Write-Host "`n✅ Script concluído!" -ForegroundColor Green
