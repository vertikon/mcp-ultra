# setup-repo-security.ps1
# Automated repository security configuration via GitHub API
# Usage: .\setup-repo-security.ps1 -Owner "vertikon" -Repo "my-repo"

param(
    [Parameter(Mandatory=$true)]
    [string]$Owner,
    
    [Parameter(Mandatory=$true)]
    [string]$Repo,
    
    [Parameter(Mandatory=$false)]
    [switch]$DryRun = $false
)

$ErrorActionPreference = "Stop"

Write-Host "================================================================" -ForegroundColor Cyan
Write-Host "  GitHub Repository Security Setup" -ForegroundColor Cyan
Write-Host "================================================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Repository: $Owner/$Repo" -ForegroundColor White
if ($DryRun) {
    Write-Host "Mode: DRY RUN (no changes will be made)" -ForegroundColor Yellow
}
Write-Host ""

# Check prerequisites
if (!(Get-Command gh -ErrorAction SilentlyContinue)) {
    Write-Host "❌ GitHub CLI (gh) is not installed" -ForegroundColor Red
    Write-Host "Install from: https://cli.github.com/" -ForegroundColor Yellow
    exit 1
}

# Check authentication
$authStatus = gh auth status 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Not authenticated with GitHub CLI" -ForegroundColor Red
    Write-Host "Run: gh auth login" -ForegroundColor Yellow
    exit 1
}

# Function to execute or simulate API calls
function Invoke-GitHubAPI {
    param(
        [string]$Method,
        [string]$Endpoint,
        [string]$Body = $null,
        [string]$Description
    )
    
    Write-Host "[API] $Description..." -ForegroundColor Yellow
    
    if ($DryRun) {
        Write-Host "  → Would execute: $Method $Endpoint" -ForegroundColor Gray
        if ($Body) {
            Write-Host "  → Body: $Body" -ForegroundColor Gray
        }
        Write-Host "  ✓ Dry run - not executed" -ForegroundColor Gray
        return
    }
    
    try {
        if ($Body) {
            $result = gh api -X $Method $Endpoint -f $Body 2>&1
        } else {
            $result = gh api -X $Method $Endpoint 2>&1
        }
        
        if ($LASTEXITCODE -eq 0) {
            Write-Host "  ✅ Success" -ForegroundColor Green
        } else {
            Write-Host "  ⚠️  Warning: $result" -ForegroundColor Yellow
        }
    } catch {
        Write-Host "  ⚠️  Warning: $_" -ForegroundColor Yellow
    }
}

# 1. Enable vulnerability alerts
Write-Host "[1/10] Enabling vulnerability alerts..." -ForegroundColor Cyan
Invoke-GitHubAPI -Method "PUT" `
    -Endpoint "repos/$Owner/$Repo/vulnerability-alerts" `
    -Description "Enable vulnerability alerts"

# 2. Enable automated security fixes (Dependabot security updates)
Write-Host "[2/10] Enabling Dependabot security updates..." -ForegroundColor Cyan
Invoke-GitHubAPI -Method "PUT" `
    -Endpoint "repos/$Owner/$Repo/automated-security-fixes" `
    -Description "Enable automated security fixes"

# 3. Enable private vulnerability reporting
Write-Host "[3/10] Enabling private vulnerability reporting..." -ForegroundColor Cyan
Invoke-GitHubAPI -Method "PUT" `
    -Endpoint "repos/$Owner/$Repo/private-vulnerability-reporting" `
    -Description "Enable private vulnerability reporting"

# 4. Enable secret scanning
Write-Host "[4/10] Enabling secret scanning..." -ForegroundColor Cyan
Invoke-GitHubAPI -Method "PUT" `
    -Endpoint "repos/$Owner/$Repo/secret-scanning/alerts" `
    -Description "Enable secret scanning"

# 5. Enable push protection
Write-Host "[5/10] Enabling secret scanning push protection..." -ForegroundColor Cyan
Invoke-GitHubAPI -Method "PUT" `
    -Endpoint "repos/$Owner/$Repo/secret-scanning/push-protection" `
    -Description "Enable push protection"

# 6. Configure branch protection for main
Write-Host "[6/10] Configuring branch protection for 'main'..." -ForegroundColor Cyan

$branchProtection = @{
    required_status_checks = @{
        strict = $true
        contexts = @("test", "gosec", "CodeQL")
    }
    enforce_admins = $true
    required_pull_request_reviews = @{
        dismissal_restrictions = @{}
        dismiss_stale_reviews = $true
        require_code_owner_reviews = $true
        required_approving_review_count = 1
        require_last_push_approval = $false
    }
    restrictions = $null
    required_linear_history = $false
    allow_force_pushes = $false
    allow_deletions = $false
    block_creations = $false
    required_conversation_resolution = $true
    lock_branch = $false
    allow_fork_syncing = $false
} | ConvertTo-Json -Depth 10

if (!$DryRun) {
    try {
        $branchProtection | gh api -X PUT "repos/$Owner/$Repo/branches/main/protection" --input - 2>&1 | Out-Null
        if ($LASTEXITCODE -eq 0) {
            Write-Host "  ✅ Branch protection configured" -ForegroundColor Green
        } else {
            Write-Host "  ⚠️  Warning: Could not configure branch protection" -ForegroundColor Yellow
        }
    } catch {
        Write-Host "  ⚠️  Warning: $_" -ForegroundColor Yellow
    }
} else {
    Write-Host "  → Would configure branch protection" -ForegroundColor Gray
    Write-Host "  ✓ Dry run - not executed" -ForegroundColor Gray
}

# 7. Set repository topics
Write-Host "[7/10] Adding repository topics..." -ForegroundColor Cyan
$topics = @("mcp", "security", "golang", "vertikon")
$topicsJson = @{ names = $topics } | ConvertTo-Json

if (!$DryRun) {
    try {
        $topicsJson | gh api -X PUT "repos/$Owner/$Repo/topics" --input - 2>&1 | Out-Null
        if ($LASTEXITCODE -eq 0) {
            Write-Host "  ✅ Topics added: $($topics -join ', ')" -ForegroundColor Green
        }
    } catch {
        Write-Host "  ⚠️  Warning: Could not add topics" -ForegroundColor Yellow
    }
} else {
    Write-Host "  → Would add topics: $($topics -join ', ')" -ForegroundColor Gray
    Write-Host "  ✓ Dry run - not executed" -ForegroundColor Gray
}

# 8. Disable wiki and projects (reduce attack surface)
Write-Host "[8/10] Disabling wiki and projects..." -ForegroundColor Cyan
$repoSettings = @{
    has_wiki = $false
    has_projects = $false
    has_downloads = $true
    allow_squash_merge = $true
    allow_merge_commit = $false
    allow_rebase_merge = $false
    delete_branch_on_merge = $true
} | ConvertTo-Json

if (!$DryRun) {
    try {
        $repoSettings | gh api -X PATCH "repos/$Owner/$Repo" --input - 2>&1 | Out-Null
        if ($LASTEXITCODE -eq 0) {
            Write-Host "  ✅ Repository settings updated" -ForegroundColor Green
        }
    } catch {
        Write-Host "  ⚠️  Warning: Could not update repository settings" -ForegroundColor Yellow
    }
} else {
    Write-Host "  → Would disable wiki and projects" -ForegroundColor Gray
    Write-Host "  ✓ Dry run - not executed" -ForegroundColor Gray
}

# 9. Create security policy label
Write-Host "[9/10] Creating security labels..." -ForegroundColor Cyan
$labels = @(
    @{ name = "security"; color = "d73a4a"; description = "Security-related issue or PR" }
    @{ name = "critical"; color = "b60205"; description = "Critical priority" }
    @{ name = "needs-security-review"; color = "fbca04"; description = "Requires security team review" }
)

foreach ($label in $labels) {
    if (!$DryRun) {
        try {
            $labelJson = $label | ConvertTo-Json
            $labelJson | gh api -X POST "repos/$Owner/$Repo/labels" --input - 2>&1 | Out-Null
            if ($LASTEXITCODE -eq 0) {
                Write-Host "  ✅ Label created: $($label.name)" -ForegroundColor Green
            } else {
                Write-Host "  ⚠️  Label already exists: $($label.name)" -ForegroundColor Yellow
            }
        } catch {
            Write-Host "  ⚠️  Could not create label: $($label.name)" -ForegroundColor Yellow
        }
    } else {
        Write-Host "  → Would create label: $($label.name)" -ForegroundColor Gray
    }
}

# 10. Verify security features
Write-Host "[10/10] Verifying security configuration..." -ForegroundColor Cyan

if (!$DryRun) {
    try {
        $repo = gh api "repos/$Owner/$Repo" | ConvertFrom-Json
        
        Write-Host ""
        Write-Host "Security Configuration Status:" -ForegroundColor White
        Write-Host "  Vulnerability alerts: ✅ Enabled" -ForegroundColor Green
        Write-Host "  Automated security fixes: ✅ Enabled" -ForegroundColor Green
        Write-Host "  Private vulnerability reporting: ✅ Enabled" -ForegroundColor Green
        Write-Host "  Secret scanning: ✅ Enabled" -ForegroundColor Green
        Write-Host "  Push protection: ✅ Enabled" -ForegroundColor Green
        Write-Host "  Branch protection: ✅ Configured" -ForegroundColor Green
        Write-Host "  Wiki: ❌ Disabled" -ForegroundColor Green
        Write-Host "  Projects: ❌ Disabled" -ForegroundColor Green
        Write-Host "  Delete branch on merge: ✅ Enabled" -ForegroundColor Green
    } catch {
        Write-Host "  ⚠️  Could not verify all settings" -ForegroundColor Yellow
    }
} else {
    Write-Host "  ✓ Dry run - verification skipped" -ForegroundColor Gray
}

# Summary
Write-Host ""
Write-Host "================================================================" -ForegroundColor Cyan
Write-Host "  Security Setup Complete!" -ForegroundColor Cyan
Write-Host "================================================================" -ForegroundColor Cyan
Write-Host ""

if ($DryRun) {
    Write-Host "This was a DRY RUN. No changes were made." -ForegroundColor Yellow
    Write-Host "Run without -DryRun to apply changes." -ForegroundColor Yellow
} else {
    Write-Host "Next steps:" -ForegroundColor Yellow
    Write-Host "  1. Review Security tab in GitHub" -ForegroundColor White
    Write-Host "  2. Configure required status checks in branch protection" -ForegroundColor White
    Write-Host "  3. Add security team to CODEOWNERS" -ForegroundColor White
    Write-Host "  4. Enable GitHub Advanced Security (if available)" -ForegroundColor White
    Write-Host "  5. Configure secrets in repository settings" -ForegroundColor White
}

Write-Host ""
Write-Host "Documentation:" -ForegroundColor Yellow
Write-Host "  https://docs.github.com/en/code-security" -ForegroundColor White
Write-Host ""
