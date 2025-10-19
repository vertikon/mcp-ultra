# Análise da evolução dos gaps reports
$reports = Get-ChildItem 'E:\vertikon\business\SaaS\templates\mcp-ultra\docs\gaps' -Filter 'gaps-report-*.json' | Sort-Object Name

Write-Host "`n=== EVOLUÇÃO DOS ERROS ===" -ForegroundColor Cyan
Write-Host "Report                          | Score | Erros | Hora" -ForegroundColor Yellow
Write-Host "-" * 70

foreach ($report in $reports) {
    try {
        $json = Get-Content $report.FullName -Raw | ConvertFrom-Json
        $erros = "N/A"

        if ($json.Critical) {
            $errorObj = $json.Critical | Where-Object { $_.Type -eq "Erros não tratados" }
            if ($errorObj) {
                $erros = $errorObj.Description
            }
        }

        $name = $report.Name -replace 'gaps-report-', '' -replace '.json', ''
        $time = $report.LastWriteTime.ToString("dd/MM HH:mm")

        Write-Host ("{0,-32} | {1,5}% | {2,-20} | {3}" -f $name, $json.Score, $erros, $time)
    }
    catch {
        Write-Host "Erro ao processar $($report.Name)" -ForegroundColor Red
    }
}

Write-Host "`n" -ForegroundColor Cyan
