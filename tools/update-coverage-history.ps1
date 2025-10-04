param(
  [string]$CoverageFuncPath = "coverage_func.txt",
  [string]$CsvPath = "docs\coverage_history.csv",
  [string]$MdPath  = "docs\coverage_history.md",
  [string]$BadgePath = "docs\badges\coverage.svg",
  [string]$Version = $env:REL_VERSION,   # pode vir do workflow (v35/v36), senão inferimos
  [int]$Window = 7                       # janela para média/mediana/tendência
)

$ErrorActionPreference = "Stop"

if (-not (Test-Path $CoverageFuncPath)) {
  throw "Não achei $CoverageFuncPath. Gere antes com: go tool cover -func=coverage.out > coverage_func.txt"
}

# --- 1) Extrair cobertura do coverage_func.txt ---
$raw = Get-Content $CoverageFuncPath -Raw
$match = [regex]::Match($raw, 'total:\s+\(statements\)\s+([0-9]+\.[0-9]+)%')
if (-not $match.Success) { throw "Não consegui extrair 'total' de $CoverageFuncPath" }
$cov = [double]$match.Groups[1].Value

# --- 2) Determinar versão e timestamp ---
if (-not $Version -or $Version.Trim() -eq "") {
  $tag = (git describe --tags --always 2>$null)
  $Version = if ($tag) { $tag } else { "local" }
}
# Fortaleza (UTC-3) simplificado sem DST: hora local -3 a partir do UTC
$now = (Get-Date).ToUniversalTime().AddHours(-3)
$ts  = $now.ToString("yyyy-MM-dd HH:mm:ss")
$line = "{0},{1},{2}" -f $ts, $Version, ("{0:n1}" -f $cov)

# --- 3) CSV: criar cabeçalho se preciso e anexar linha ---
New-Item -ItemType Directory -Force -Path (Split-Path $CsvPath) | Out-Null
if (-not (Test-Path $CsvPath)) {
  "timestamp,version,coverage" | Out-File -Encoding UTF8 $CsvPath
}
Add-Content -Path $CsvPath -Value $line

# --- 4) Carregar série e calcular estatísticas ---
$rows = (Get-Content $CsvPath | Select-Object -Skip 1) | ForEach-Object {
  $p = $_.Split(',')
  [pscustomobject]@{ timestamp=$p[0]; version=$p[1]; coverage=[double]$p[2] }
}
$totalCount = $rows.Count
$prev = if ($totalCount -ge 2) { $rows[$totalCount-2].coverage } else { $cov }
$delta = [math]::Round($cov - $prev, 1)

$windowRows = @($rows | Select-Object -Last $Window)
$wCount = $windowRows.Count
$wAvg = if ($wCount -gt 0) { [math]::Round(($windowRows | Measure-Object -Property coverage -Average).Average, 1) } else { $cov }
$wSorted = $windowRows.coverage | Sort-Object
$wMedian = if ($wCount -gt 0) {
  if ($wCount % 2 -eq 1) { [math]::Round($wSorted[ [int]([math]::Floor($wCount/2)) ], 1) }
  else { [math]::Round( ($wSorted[$wCount/2 - 1] + $wSorted[$wCount/2]) / 2, 1) }
} else { $cov }
$wMin = if ($wCount -gt 0) { [math]::Round($wSorted[0],1) } else { $cov }
$wMax = if ($wCount -gt 0) { [math]::Round($wSorted[-1],1) } else { $cov }
$wDelta = if ($wCount -gt 1) { [math]::Round($windowRows[-1].coverage - $windowRows[0].coverage, 1) } else { 0.0 }

# Tendência (setinha): threshold 0.1pp
function Get-TrendArrow([double]$d) {
  if ($d -gt 0.1) { return "↗" }
  elseif ($d -lt -0.1) { return "↘" }
  else { return "→" }
}
$arrow = Get-TrendArrow $delta

# --- 5) Gerar badge SVG ---
function Get-BadgeColor([double]$p) {
  if ($p -ge 90) { return "#4c1" }         # brightgreen
  elseif ($p -ge 80) { return "#97CA00" }  # greenish
  elseif ($p -ge 70) { return "#a4a61d" }  # yellowgreen
  elseif ($p -ge 50) { return "#dfb317" }  # yellow
  elseif ($p -ge 20) { return "#fe7d37" }  # orange
  else { return "#e05d44" }                # red
}
$color = Get-BadgeColor $cov
$label = "coverage"
$value = ("{0:n1}%" -f $cov)

# badge simples (compatível com GitHub render)
$svg = @"
<svg xmlns="http://www.w3.org/2000/svg" width="150" height="20" role="img" aria-label="$label: $value">
  <linearGradient id="s" x2="0" y2="100%">
    <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
    <stop offset="1" stop-opacity=".1"/>
  </linearGradient>
  <mask id="m"><rect width="150" height="20" rx="3" fill="#fff"/></mask>
  <g mask="url(#m)">
    <rect width="70" height="20" fill="#555"/>
    <rect x="70" width="80" height="20" fill="$color"/>
    <rect width="150" height="20" fill="url(#s)"/>
  </g>
  <g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" font-size="11">
    <text x="35" y="14">$label</text>
    <text x="110" y="14">$value</text>
  </g>
</svg>
"@

New-Item -ItemType Directory -Force -Path (Split-Path $BadgePath) | Out-Null
$svg | Out-File -Encoding utf8 $BadgePath

# --- 6) Gerar Markdown com dashboard + tabela (últimos 30) ---
$last30 = $rows | Select-Object -Last 30
$tbl = "| Data/Hora (Fortaleza) | Versão | Coverage (%) |\n|---|---|---:|\n"
foreach ($r in $last30) {
  $tbl += "| {0} | {1} | {2:n1} |\n" -f $r.timestamp, $r.version, $r.coverage
}

$dash = @"
# Histórico de Coverage

**Atual:** {0:n1}% {1} (Δ {2:+0.0;-0.0;0.0} pp vs. anterior)
**Janela {3}:** média {4:n1}%, mediana {5:n1}% (melhor {6:n1}% · pior {7:n1}%, Δ janela {8:+0.0;-0.0;0.0} pp)
**Execuções registradas:** {9}

![coverage](badges/coverage.svg)

{10}
"@ -f $cov, $arrow, $delta, $Window, $wAvg, $wMedian, $wMax, $wMin, $wDelta, $totalCount, $tbl

New-Item -ItemType Directory -Force -Path (Split-Path $MdPath) | Out-Null
$dash | Out-File -Encoding UTF8 $MdPath

Write-Host ("Histórico atualizado ✓  {0} | {1} | {2}  (Δ {3:+0.0;-0.0;0.0} pp)" -f $CsvPath, $MdPath, $BadgePath, $delta)
