# init.ps1
# Initialize new MCP repository with security settings
# Usage: .\init.ps1 -RepoName "my-new-repo" -Description "My new MCP server"

param(
    [Parameter(Mandatory=$true)]
    [string]$RepoName,
    
    [Parameter(Mandatory=$false)]
    [string]$Description = "MCP Server",
    
    [Parameter(Mandatory=$false)]
    [string]$GoVersion = "1.22",
    
    [Parameter(Mandatory=$false)]
    [switch]$EnableGitHubActions = $true,
    
    [Parameter(Mandatory=$false)]
    [switch]$CreateRepo = $false
)

$ErrorActionPreference = "Stop"

Write-Host "================================================================" -ForegroundColor Cyan
Write-Host "  MCP Repository Initialization" -ForegroundColor Cyan
Write-Host "================================================================" -ForegroundColor Cyan
Write-Host ""

# 1. Verify prerequisites
Write-Host "[1/8] Checking prerequisites..." -ForegroundColor Yellow
$prerequisites = @("go", "git", "gh")
foreach ($cmd in $prerequisites) {
    if (!(Get-Command $cmd -ErrorAction SilentlyContinue)) {
        Write-Host "❌ $cmd is not installed or not in PATH" -ForegroundColor Red
        exit 1
    }
}
Write-Host "✅ All prerequisites found" -ForegroundColor Green

# 2. Initialize Go module
Write-Host "[2/8] Initializing Go module..." -ForegroundColor Yellow
$modulePath = "github.com/vertikon/$RepoName"
go mod init $modulePath
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Failed to initialize Go module" -ForegroundColor Red
    exit 1
}
Write-Host "✅ Go module initialized: $modulePath" -ForegroundColor Green

# 3. Create directory structure
Write-Host "[3/8] Creating directory structure..." -ForegroundColor Yellow
$dirs = @(
    "cmd",
    "internal/handlers",
    "internal/security",
    "internal/config",
    "pkg",
    "test",
    "docs",
    "deploy/k8s",
    "deploy/docker"
)
foreach ($dir in $dirs) {
    New-Item -ItemType Directory -Force -Path $dir | Out-Null
}
Write-Host "✅ Directory structure created" -ForegroundColor Green

# 4. Update README with repo-specific info
Write-Host "[4/8] Customizing README..." -ForegroundColor Yellow
if (Test-Path "README.md") {
    $readme = Get-Content "README.md" -Raw
    $readme = $readme -replace '\{REPO_NAME\}', $RepoName
    $readme = $readme -replace '\{DESCRIPTION\}', $Description
    $readme = $readme -replace '\{GO_VERSION\}', $GoVersion
    Set-Content "README.md" -Value $readme
    Write-Host "✅ README customized" -ForegroundColor Green
} else {
    Write-Host "⚠️  README.md not found, skipping..." -ForegroundColor Yellow
}

# 5. Initialize Git repository
Write-Host "[5/8] Initializing Git repository..." -ForegroundColor Yellow
if (!(Test-Path ".git")) {
    git init
    git branch -m main
    Write-Host "✅ Git repository initialized" -ForegroundColor Green
} else {
    Write-Host "⚠️  Git repository already exists" -ForegroundColor Yellow
}

# 6. Create initial commit
Write-Host "[6/8] Creating initial commit..." -ForegroundColor Yellow
git add .
git commit -m "feat: initialize repository from template

- Setup project structure
- Configure security workflows
- Add standard documentation

🤖 Generated from mcp-ultra template"
Write-Host "✅ Initial commit created" -ForegroundColor Green

# 7. Create GitHub repository (if requested)
if ($CreateRepo) {
    Write-Host "[7/8] Creating GitHub repository..." -ForegroundColor Yellow
    gh repo create "vertikon/$RepoName" `
        --private `
        --description $Description `
        --source . `
        --push
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ GitHub repository created and pushed" -ForegroundColor Green
    } else {
        Write-Host "❌ Failed to create GitHub repository" -ForegroundColor Red
        exit 1
    }
} else {
    Write-Host "[7/8] Skipping GitHub repository creation" -ForegroundColor Yellow
}

# 8. Enable GitHub security features (if repo was created)
if ($CreateRepo -and $EnableGitHubActions) {
    Write-Host "[8/8] Configuring GitHub security features..." -ForegroundColor Yellow
    
    # Enable vulnerability alerts
    gh api -X PUT "repos/vertikon/$RepoName/vulnerability-alerts" | Out-Null
    
    # Enable Dependabot security updates
    gh api -X PUT "repos/vertikon/$RepoName/automated-security-fixes" | Out-Null
    
    Write-Host "✅ Security features enabled" -ForegroundColor Green
} else {
    Write-Host "[8/8] Skipping security feature configuration" -ForegroundColor Yellow
}

# Summary
Write-Host ""
Write-Host "================================================================" -ForegroundColor Cyan
Write-Host "  Initialization Complete!" -ForegroundColor Cyan
Write-Host "================================================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Repository: $RepoName" -ForegroundColor White
Write-Host "Module: $modulePath" -ForegroundColor White
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Yellow
Write-Host "  1. Review and customize .github/workflows/*.yml" -ForegroundColor White
Write-Host "  2. Update SECURITY.md with contact information" -ForegroundColor White
Write-Host "  3. Configure CODEOWNERS with your team names" -ForegroundColor White
Write-Host "  4. Run: make deps" -ForegroundColor White
Write-Host "  5. Run: make test" -ForegroundColor White
Write-Host "  6. Start coding!" -ForegroundColor White
Write-Host ""

if (!$CreateRepo) {
    Write-Host "To create the GitHub repository later, run:" -ForegroundColor Yellow
    Write-Host "  gh repo create vertikon/$RepoName --private --source . --push" -ForegroundColor White
    Write-Host ""
}

Write-Host "Documentation:" -ForegroundColor Yellow
Write-Host "  - README.md" -ForegroundColor White
Write-Host "  - SECURITY.md" -ForegroundColor White
Write-Host "  - docs/SECURITY-OVERVIEW.md" -ForegroundColor White
Write-Host ""
Write-Host "Happy coding! 🚀" -ForegroundColor Green
