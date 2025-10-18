# Script to remove router_example.go that conflicts with the real router

Set-Location "E:\vertikon\business\SaaS\templates\mcp-ultra"

Write-Host "Removing router_example.go..." -ForegroundColor Yellow

# Check if file exists
if (Test-Path ".\internal\handlers\http\router_example.go") {
    # Try git rm first (if tracked)
    try {
        git rm ".\internal\handlers\http\router_example.go" 2>$null
        Write-Host "File removed via git rm" -ForegroundColor Green
    } catch {
        # If git rm fails, just delete the file
        Remove-Item ".\internal\handlers\http\router_example.go" -Force
        Write-Host "File deleted directly" -ForegroundColor Green
    }

    Write-Host "router_example.go has been removed successfully!" -ForegroundColor Green
} else {
    Write-Host "File not found - already removed?" -ForegroundColor Cyan
}

Write-Host "`nYou can now run the validation checks:" -ForegroundColor Cyan
Write-Host "  go build ./..." -ForegroundColor White
Write-Host "  golangci-lint run --out-format=tab" -ForegroundColor White
