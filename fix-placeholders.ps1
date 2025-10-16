# Script para substituir MODULE_PATH por github.com/vertikon/mcp-ultra
$ModulePath = "github.com/vertikon/mcp-ultra"
$Root = "E:\vertikon\business\SaaS\templates\mcp-ultra"

Write-Host "Substituindo placeholders..." -ForegroundColor Cyan

$files = Get-ChildItem -Path $Root -Recurse -Include *.go -File | Where-Object {
    $_.FullName -notmatch '\\\.git\\' -and
    $_.FullName -notmatch '\\vendor\\' -and
    $_.FullName -notmatch '\\node_modules\\'
}

$count = 0
foreach ($file in $files) {
    $content = Get-Content -Path $file.FullName -Raw
    $newContent = $content -replace '\{\{MODULE_PATH\}\}', $ModulePath

    if ($content -ne $newContent) {
        Set-Content -Path $file.FullName -Value $newContent -NoNewline
        $count++
        Write-Host "  Fixed: $($file.Name)" -ForegroundColor Green
    }
}

Write-Host "`nTotal arquivos corrigidos: $count" -ForegroundColor Yellow
