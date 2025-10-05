#!/bin/bash
# MCP Ultra - Local Security Validation Script
# =============================================
# This script runs the same security checks as the CI/CD pipeline locally
# Usage: ./validate-security.sh

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
PROJECT_NAME="mcp-ultra"
GO_VERSION="1.22"

# Functions
print_header() {
    echo -e "\n${BLUE}========================================${NC}"
    echo -e "${BLUE}$1${NC}"
    echo -e "${BLUE}========================================${NC}\n"
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

print_error() {
    echo -e "${RED}✗ $1${NC}"
}

check_command() {
    if command -v "$1" &> /dev/null; then
        print_success "$1 is installed"
        return 0
    else
        print_warning "$1 is not installed"
        return 1
    fi
}

# Main execution
print_header "MCP Ultra Security Validation"
echo "Starting security checks for $PROJECT_NAME"
echo "Timestamp: $(date)"

# Check prerequisites
print_header "Checking Prerequisites"

MISSING_TOOLS=()
check_command "go" || MISSING_TOOLS+=("go")
check_command "docker" || MISSING_TOOLS+=("docker")
check_command "git" || MISSING_TOOLS+=("git")

# Optional tools (warn but don't fail)
check_command "gitleaks" || print_warning "Install gitleaks: https://github.com/gitleaks/gitleaks"
check_command "trivy" || print_warning "Install trivy: brew install aquasecurity/trivy/trivy"
check_command "grype" || print_warning "Install grype: curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sh"
check_command "gosec" || print_warning "Install gosec: go install github.com/securego/gosec/v2/cmd/gosec@latest"
check_command "golangci-lint" || print_warning "Install golangci-lint: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

if [ ${#MISSING_TOOLS[@]} -gt 0 ]; then
    print_error "Missing required tools: ${MISSING_TOOLS[*]}"
    echo "Please install missing tools and try again"
    exit 1
fi

# 1. Go Module Verification
print_header "1. Go Module Verification"
if [ -f "go.mod" ]; then
    go mod verify && print_success "Go modules verified"
    go mod tidy && print_success "Go modules tidied"

    # Check for outdated dependencies
    echo -e "\n${YELLOW}Checking for available updates:${NC}"
    go list -m -u all | grep -E '\[.*\]' || echo "All dependencies up to date"
else
    print_error "go.mod not found"
fi

# 2. Go Security Analysis
print_header "2. Go Security Analysis"

# Run go vet
echo "Running go vet..."
go vet ./... && print_success "go vet passed" || print_error "go vet failed"

# Run gosec if available
if command -v gosec &> /dev/null; then
    echo "Running gosec..."
    gosec -fmt json -out gosec-report.json -severity high ./... || true
    if [ -f "gosec-report.json" ]; then
        ISSUES=$(jq '.Issues | length' gosec-report.json)
        if [ "$ISSUES" -eq 0 ]; then
            print_success "No high/critical security issues found by gosec"
        else
            print_warning "Found $ISSUES security issues (see gosec-report.json)"
        fi
    fi
fi

# Run golangci-lint if available
if command -v golangci-lint &> /dev/null; then
    echo "Running golangci-lint..."
    golangci-lint run --timeout=5m || print_warning "golangci-lint found issues"
fi

# 3. Secret Detection
print_header "3. Secret Detection"
if command -v gitleaks &> /dev/null; then
    echo "Running gitleaks..."
    if [ -f ".gitleaks.toml" ]; then
        gitleaks detect --config=.gitleaks.toml --report-format json --report-path gitleaks-report.json || true
    else
        gitleaks detect --report-format json --report-path gitleaks-report.json || true
    fi

    if [ -f "gitleaks-report.json" ]; then
        LEAKS=$(jq '. | length' gitleaks-report.json)
        if [ "$LEAKS" -eq 0 ]; then
            print_success "No secrets detected"
        else
            print_error "Found $LEAKS potential secrets (see gitleaks-report.json)"
        fi
    fi
else
    # Fallback: basic pattern matching
    echo "Running basic secret detection..."
    PATTERNS=("password.*=" "api[_-]?key.*=" "secret.*=" "token.*=" "private[_-]?key")
    for pattern in "${PATTERNS[@]}"; do
        if grep -rI --exclude-dir=.git --exclude-dir=vendor --exclude="*.md" -i "$pattern" . 2>/dev/null; then
            print_warning "Found potential secrets matching pattern: $pattern"
        fi
    done
fi

# 4. Dependency Vulnerability Scan
print_header "4. Dependency Vulnerability Scan"

# Generate SBOM if syft is available
if command -v syft &> /dev/null; then
    echo "Generating SBOM..."
    syft dir:. -o cyclonedx-json > sbom.json
    print_success "SBOM generated (sbom.json)"
fi

# Run vulnerability scan
if command -v grype &> /dev/null; then
    echo "Running Grype vulnerability scan..."
    if [ -f "sbom.json" ]; then
        grype sbom:sbom.json --fail-on high --add-cpes-if-none -o json > grype-report.json || true
    else
        grype dir:. --fail-on high -o json > grype-report.json || true
    fi

    if [ -f "grype-report.json" ]; then
        HIGH_VULNS=$(jq '[.matches[] | select(.vulnerability.severity == "High")] | length' grype-report.json)
        CRITICAL_VULNS=$(jq '[.matches[] | select(.vulnerability.severity == "Critical")] | length' grype-report.json)

        if [ "$CRITICAL_VULNS" -gt 0 ]; then
            print_error "Found $CRITICAL_VULNS CRITICAL vulnerabilities"
        fi
        if [ "$HIGH_VULNS" -gt 0 ]; then
            print_warning "Found $HIGH_VULNS HIGH vulnerabilities"
        fi
        if [ "$CRITICAL_VULNS" -eq 0 ] && [ "$HIGH_VULNS" -eq 0 ]; then
            print_success "No high/critical vulnerabilities found"
        fi
    fi
fi

# 5. Docker Security
print_header "5. Docker Security"

# Build Docker image
if [ -f "Dockerfile" ] || [ -f "Dockerfile.secure" ]; then
    DOCKERFILE="Dockerfile"
    [ -f "Dockerfile.secure" ] && DOCKERFILE="Dockerfile.secure"

    echo "Building Docker image..."
    docker build -t "${PROJECT_NAME}:security-test" -f "$DOCKERFILE" . || print_error "Docker build failed"

    # Run Trivy if available
    if command -v trivy &> /dev/null; then
        echo "Running Trivy container scan..."
        trivy image --severity HIGH,CRITICAL --exit-code 0 "${PROJECT_NAME}:security-test" > trivy-report.txt
        print_success "Trivy scan completed (see trivy-report.txt)"
    fi

    # Check if container runs as non-root
    echo "Checking container user..."
    USER=$(docker run --rm --entrypoint sh "${PROJECT_NAME}:security-test" -c "whoami" 2>/dev/null || echo "root")
    if [ "$USER" != "root" ]; then
        print_success "Container runs as non-root user: $USER"
    else
        print_error "Container runs as root!"
    fi
fi

# 6. Kubernetes Manifest Validation
print_header "6. Kubernetes Manifest Validation"

if [ -d "deploy/k8s" ] && command -v conftest &> /dev/null; then
    echo "Validating Kubernetes manifests with OPA..."
    conftest test deploy/k8s/*.yaml -p policy/ || print_warning "Some policies failed"
elif [ -d "deploy/k8s" ]; then
    print_warning "Conftest not installed - skipping K8s validation"
    echo "Install with: brew install conftest"
fi

# 7. License Check
print_header "7. License Compliance"

if command -v go-licenses &> /dev/null; then
    echo "Checking licenses..."
    go-licenses check ./... --disallowed_types=forbidden,restricted || print_warning "License issues detected"
else
    print_warning "go-licenses not installed"
    echo "Install with: go install github.com/google/go-licenses@latest"
fi

# Summary
print_header "Security Validation Summary"

echo "Reports generated:"
[ -f "gosec-report.json" ] && echo "  - gosec-report.json"
[ -f "gitleaks-report.json" ] && echo "  - gitleaks-report.json"
[ -f "grype-report.json" ] && echo "  - grype-report.json"
[ -f "trivy-report.txt" ] && echo "  - trivy-report.txt"
[ -f "sbom.json" ] && echo "  - sbom.json"

echo -e "\n${GREEN}Security validation completed!${NC}"
echo "Review any warnings or errors above before pushing to CI/CD"

# Cleanup (optional)
read -p "Clean up report files? (y/N) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    rm -f gosec-report.json gitleaks-report.json grype-report.json trivy-report.txt sbom.json
    print_success "Report files cleaned up"
fi

exit 0