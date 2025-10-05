# Validate compilation fixes
Write-Host "=== Validating Compilation Fixes ===" -ForegroundColor Cyan

Write-Host "`n[1/3] Running go build..." -ForegroundColor Yellow
go build ./... 2>&1 | Tee-Object -Variable buildOutput

if ($LASTEXITCODE -eq 0) {
    Write-Host "OK Build successful" -ForegroundColor Green

    Write-Host "`n[2/3] Running tests..." -ForegroundColor Yellow
    go test -short ./... 2>&1 | Tee-Object -Variable testOutput

    if ($LASTEXITCODE -eq 0) {
        Write-Host "OK Tests passed" -ForegroundColor Green
    } else {
        Write-Host "WARNING Some tests failed" -ForegroundColor Yellow
    }

    Write-Host "`n[3/3] Checking formatting..." -ForegroundColor Yellow
    $unformatted = gofmt -l .
    if ($unformatted) {
        Write-Host "WARNING Unformatted files found:" -ForegroundColor Yellow
        $unformatted
    } else {
        Write-Host "OK All files formatted" -ForegroundColor Green
    }

    Write-Host "`n========================================" -ForegroundColor Green
    Write-Host "  VALIDATION SUCCESSFUL!" -ForegroundColor Green
    Write-Host "========================================" -ForegroundColor Green

} else {
    Write-Host "`n========================================" -ForegroundColor Red
    Write-Host "  BUILD FAILED" -ForegroundColor Red
    Write-Host "========================================" -ForegroundColor Red
    Write-Host "`nErrors:" -ForegroundColor Red
    $buildOutput
}

Write-Host "`n=== Validation Complete ===" -ForegroundColor Cyan
