#!/usr/bin/env pwsh
# MCP Ultra - Script de Auditoria de Projeto
# ===========================================

$timestamp = Get-Date -Format "yyyy-MM-dd_HH-mm-ss"
$auditDir = "docs/audit"
$reportFile = "$auditDir/audit-report-$timestamp.json"

Write-Host "🔍 Iniciando Auditoria do Projeto MCP Ultra..." -ForegroundColor Cyan
Write-Host ""

# Estrutura do relatório
$report = @{
    timestamp = (Get-Date -Format "yyyy-MM-dd'T'HH:mm:ssK")
    project = "mcp-ultra"
    summary = @{}
    directories = @{}
    files = @{}
    dependencies = @{}
    candidates_for_cleanup = @()
    recommendations = @()
}

# 1. Análise de Diretórios
Write-Host "📁 Analisando estrutura de diretórios..." -ForegroundColor Yellow

$directories = @("internal", "pkg", "cmd", "docs", "test", "tests", "scripts", "deploy", "migrations", "api")

foreach ($dir in $directories) {
    if (Test-Path $dir) {
        $files = Get-ChildItem -Path $dir -Recurse -File -ErrorAction SilentlyContinue
        $totalSize = ($files | Measure-Object -Property Length -Sum).Sum
        $goFiles = ($files | Where-Object { $_.Extension -eq ".go" }).Count
        $testFiles = ($files | Where-Object { $_.Name -like "*_test.go" }).Count
        $mdFiles = ($files | Where-Object { $_.Extension -eq ".md" }).Count
        
        $report.directories[$dir] = @{
            total_files = $files.Count
            total_size_bytes = $totalSize
            total_size_mb = [math]::Round($totalSize / 1MB, 2)
            go_files = $goFiles
            test_files = $testFiles
            md_files = $mdFiles
        }
        
        Write-Host "  ✓ $dir : $($files.Count) files, $([math]::Round($totalSize / 1MB, 2)) MB" -ForegroundColor Green
    }
}

# 2. Análise de Arquivos Duplicados ou Obsoletos
Write-Host ""
Write-Host "🔍 Procurando arquivos obsoletos..." -ForegroundColor Yellow

$obsoletePatterns = @("*.old", "*.bak", "*.tmp", "*.backup", "*-old.*", "*_old.*")
$obsoleteFiles = @()

foreach ($pattern in $obsoletePatterns) {
    $found = Get-ChildItem -Path . -Filter $pattern -Recurse -File -ErrorAction SilentlyContinue
    $obsoleteFiles += $found
}

if ($obsoleteFiles.Count -gt 0) {
    Write-Host "  ⚠ Encontrados $($obsoleteFiles.Count) arquivos obsoletos:" -ForegroundColor Red
    foreach ($file in $obsoleteFiles) {
        Write-Host "    - $($file.FullName)" -ForegroundColor Gray
        $report.candidates_for_cleanup += @{
            path = $file.FullName.Replace((Get-Location).Path, ".")
            reason = "Arquivo obsoleto (padrão: $($file.Extension))"
            action = "remove"
            size_bytes = $file.Length
        }
    }
} else {
    Write-Host "  ✓ Nenhum arquivo obsoleto encontrado" -ForegroundColor Green
}

# 3. Análise de Documentação
Write-Host ""
Write-Host "📚 Analisando documentação..." -ForegroundColor Yellow

$docsGaps = Get-ChildItem -Path "docs/gaps" -File -ErrorAction SilentlyContinue
$docsMelhorias = Get-ChildItem -Path "docs/melhorias" -File -ErrorAction SilentlyContinue

$report.files["docs_gaps"] = @{
    count = $docsGaps.Count
    total_size_kb = [math]::Round(($docsGaps | Measure-Object -Property Length -Sum).Sum / 1KB, 2)
}

$report.files["docs_melhorias"] = @{
    count = $docsMelhorias.Count
    total_size_kb = [math]::Round(($docsMelhorias | Measure-Object -Property Length -Sum).Sum / 1KB, 2)
}

if ($docsGaps.Count -gt 50) {
    $report.recommendations += "Consolidar docs/gaps/ ($($docsGaps.Count) arquivos) em histórico único"
    Write-Host "  ⚠ docs/gaps/ tem $($docsGaps.Count) arquivos - considerar consolidação" -ForegroundColor Yellow
}

if ($docsMelhorias.Count -gt 30) {
    $report.recommendations += "Consolidar docs/melhorias/ ($($docsMelhorias.Count) arquivos) em histórico único"
    Write-Host "  ⚠ docs/melhorias/ tem $($docsMelhorias.Count) arquivos - considerar consolidação" -ForegroundColor Yellow
}

# 4. Análise de Scripts
Write-Host ""
Write-Host "📜 Analisando scripts..." -ForegroundColor Yellow

$scripts = Get-ChildItem -Path "scripts", "." -Include "*.ps1", "*.sh" -File -ErrorAction SilentlyContinue | 
    Where-Object { $_.Name -notlike "audit-*" }

$unusedScripts = @()
foreach ($script in $scripts) {
    # Verificar se o script é referenciado no Makefile ou em outros scripts
    $makefileRefs = Select-String -Path "Makefile" -Pattern $script.Name -ErrorAction SilentlyContinue
    $scriptRefs = Select-String -Path "scripts/*" -Pattern $script.Name -ErrorAction SilentlyContinue
    
    if (-not $makefileRefs -and -not $scriptRefs -and $script.Name -like "*old*") {
        $unusedScripts += $script
        Write-Host "  ⚠ Script possivelmente não usado: $($script.Name)" -ForegroundColor Yellow
    }
}

$report.files["scripts_total"] = $scripts.Count
$report.files["scripts_potentially_unused"] = $unusedScripts.Count

# 5. Análise de Binários e Arquivos de Build
Write-Host ""
Write-Host "🔨 Procurando binários e arquivos de build..." -ForegroundColor Yellow

$buildArtifacts = Get-ChildItem -Path "." -Include "*.exe", "*.test", "*.out", "*.prof" -File -ErrorAction SilentlyContinue |
    Where-Object { $_.Name -ne "mcp-ultra.exe" }

foreach ($artifact in $buildArtifacts) {
    Write-Host "  ⚠ Artefato de build encontrado: $($artifact.Name)" -ForegroundColor Yellow
    $report.candidates_for_cleanup += @{
        path = $artifact.FullName.Replace((Get-Location).Path, ".")
        reason = "Artefato de build temporário"
        action = "remove"
        size_bytes = $artifact.Length
    }
}

# 6. Análise de Dependências no go.mod
Write-Host ""
Write-Host "📦 Analisando dependências..." -ForegroundColor Yellow

if (Test-Path "go.mod") {
    $goModContent = Get-Content "go.mod"
    $requires = $goModContent | Select-String "^\s+github.com/" | Measure-Object
    $replaces = $goModContent | Select-String "^replace " | Measure-Object
    
    $report.dependencies["total_requires"] = $requires.Count
    $report.dependencies["total_replaces"] = $replaces.Count
    
    Write-Host "  ✓ Total de dependências require: $($requires.Count)" -ForegroundColor Green
    Write-Host "  ✓ Total de replace directives: $($replaces.Count)" -ForegroundColor Green
    
    # Procurar por replace com caminhos locais
    $localReplaces = $goModContent | Select-String "^replace .* => [A-Z]:" -ErrorAction SilentlyContinue
    if ($localReplaces) {
        Write-Host "  ⚠ Encontrado(s) replace com caminho local (problema para Docker):" -ForegroundColor Red
        foreach ($replace in $localReplaces) {
            Write-Host "    - $($replace.Line)" -ForegroundColor Gray
        }
        $report.recommendations += "Remover ou comentar replace directives com caminhos locais antes do build Docker"
    }
}

# 7. Resumo Final
Write-Host ""
Write-Host "📊 Gerando Resumo..." -ForegroundColor Cyan

$totalFiles = (Get-ChildItem -Path . -Recurse -File -ErrorAction SilentlyContinue).Count
$totalGoFiles = (Get-ChildItem -Path . -Include "*.go" -Recurse -File -ErrorAction SilentlyContinue).Count
$totalTestFiles = (Get-ChildItem -Path . -Include "*_test.go" -Recurse -File -ErrorAction SilentlyContinue).Count

$report.summary = @{
    total_files = $totalFiles
    total_go_files = $totalGoFiles
    total_test_files = $totalTestFiles
    total_obsolete_files = $obsoleteFiles.Count
    total_build_artifacts = $buildArtifacts.Count
    total_cleanup_candidates = $report.candidates_for_cleanup.Count
    total_recommendations = $report.recommendations.Count
}

# Salvar relatório
$reportJson = $report | ConvertTo-Json -Depth 10
$reportJson | Out-File -FilePath $reportFile -Encoding utf8

Write-Host ""
Write-Host "═══════════════════════════════════════════════════════" -ForegroundColor Cyan
Write-Host "                 RESUMO DA AUDITORIA" -ForegroundColor Cyan
Write-Host "═══════════════════════════════════════════════════════" -ForegroundColor Cyan
Write-Host ""
Write-Host "Total de arquivos no projeto: $totalFiles" -ForegroundColor White
Write-Host "  - Arquivos Go: $totalGoFiles" -ForegroundColor White
Write-Host "  - Arquivos de teste: $totalTestFiles" -ForegroundColor White
Write-Host ""
Write-Host "Candidatos para limpeza: $($report.candidates_for_cleanup.Count)" -ForegroundColor $(if ($report.candidates_for_cleanup.Count -gt 0) { "Yellow" } else { "Green" })
Write-Host "  - Arquivos obsoletos: $($obsoleteFiles.Count)" -ForegroundColor Yellow
Write-Host "  - Artefatos de build: $($buildArtifacts.Count)" -ForegroundColor Yellow
Write-Host ""
Write-Host "Recomendações: $($report.recommendations.Count)" -ForegroundColor $(if ($report.recommendations.Count -gt 0) { "Yellow" } else { "Green" })
foreach ($rec in $report.recommendations) {
    Write-Host "  • $rec" -ForegroundColor Yellow
}
Write-Host ""
Write-Host "✅ Relatório salvo em: $reportFile" -ForegroundColor Green
Write-Host ""

