#!/bin/bash

# Container Security Check Script
# This script performs comprehensive security checks on Docker images

set -euo pipefail

# Configuration
IMAGE_NAME="${1:-mcp-ultra}"
IMAGE_TAG="${2:-latest}"
FULL_IMAGE_NAME="${IMAGE_NAME}:${IMAGE_TAG}"
REPORT_DIR="security-reports"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Create report directory
mkdir -p "${REPORT_DIR}"

echo -e "${BLUE}🔒 Starting comprehensive security check for ${FULL_IMAGE_NAME}${NC}"
echo "=================================================="

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to print section header
print_section() {
    echo -e "\n${BLUE}📋 $1${NC}"
    echo "----------------------------------------"
}

# Function to print success
print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

# Function to print warning
print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

# Function to print error
print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Check required tools
print_section "Checking Required Security Tools"

REQUIRED_TOOLS=("docker" "trivy" "syft" "grype" "cosign")
MISSING_TOOLS=()

for tool in "${REQUIRED_TOOLS[@]}"; do
    if command_exists "$tool"; then
        print_success "$tool is installed"
    else
        print_error "$tool is not installed"
        MISSING_TOOLS+=("$tool")
    fi
done

if [ ${#MISSING_TOOLS[@]} -ne 0 ]; then
    print_error "Missing required tools: ${MISSING_TOOLS[*]}"
    echo "Please install missing tools and re-run the script"
    exit 1
fi

# 1. Image Vulnerability Scanning with Trivy
print_section "1. Vulnerability Scanning (Trivy)"

echo "Scanning for HIGH and CRITICAL vulnerabilities..."
trivy image \
    --format json \
    --output "${REPORT_DIR}/trivy-report-${TIMESTAMP}.json" \
    "${FULL_IMAGE_NAME}"

trivy image \
    --format table \
    --severity HIGH,CRITICAL \
    "${FULL_IMAGE_NAME}" | tee "${REPORT_DIR}/trivy-summary-${TIMESTAMP}.txt"

CRITICAL_COUNT=$(trivy image --format json "${FULL_IMAGE_NAME}" | jq '.Results[]?.Vulnerabilities[]? | select(.Severity=="CRITICAL") | .VulnerabilityID' | wc -l)
HIGH_COUNT=$(trivy image --format json "${FULL_IMAGE_NAME}" | jq '.Results[]?.Vulnerabilities[]? | select(.Severity=="HIGH") | .VulnerabilityID' | wc -l)

if [ "$CRITICAL_COUNT" -gt 0 ]; then
    print_error "Found $CRITICAL_COUNT CRITICAL vulnerabilities"
else
    print_success "No CRITICAL vulnerabilities found"
fi

if [ "$HIGH_COUNT" -gt 0 ]; then
    print_warning "Found $HIGH_COUNT HIGH vulnerabilities"
else
    print_success "No HIGH vulnerabilities found"
fi

# 2. Software Bill of Materials (SBOM) Generation
print_section "2. Software Bill of Materials (SBOM)"

echo "Generating SBOM with Syft..."
syft "${FULL_IMAGE_NAME}" -o json > "${REPORT_DIR}/sbom-${TIMESTAMP}.json"
syft "${FULL_IMAGE_NAME}" -o spdx-json > "${REPORT_DIR}/sbom-spdx-${TIMESTAMP}.json"
syft "${FULL_IMAGE_NAME}" -o table > "${REPORT_DIR}/sbom-table-${TIMESTAMP}.txt"

PACKAGE_COUNT=$(syft "${FULL_IMAGE_NAME}" -o json | jq '.artifacts | length')
print_success "SBOM generated with $PACKAGE_COUNT packages"

# 3. Vulnerability Scanning with Grype
print_section "3. Additional Vulnerability Scanning (Grype)"

echo "Running Grype vulnerability scan..."
grype "${FULL_IMAGE_NAME}" \
    -o json > "${REPORT_DIR}/grype-report-${TIMESTAMP}.json"

grype "${FULL_IMAGE_NAME}" \
    -o table > "${REPORT_DIR}/grype-summary-${TIMESTAMP}.txt"

GRYPE_CRITICAL=$(grype "${FULL_IMAGE_NAME}" -o json | jq '.matches[]? | select(.vulnerability.severity=="Critical") | .vulnerability.id' | wc -l)
GRYPE_HIGH=$(grype "${FULL_IMAGE_NAME}" -o json | jq '.matches[]? | select(.vulnerability.severity=="High") | .vulnerability.id' | wc -l)

print_success "Grype scan completed: $GRYPE_CRITICAL Critical, $GRYPE_HIGH High vulnerabilities"

# 4. Image Configuration Security Check
print_section "4. Image Configuration Security"

echo "Inspecting image configuration..."
docker inspect "${FULL_IMAGE_NAME}" > "${REPORT_DIR}/image-config-${TIMESTAMP}.json"

# Check for security best practices
CONFIG_ISSUES=()

# Check if running as root
USER_ID=$(docker inspect "${FULL_IMAGE_NAME}" | jq -r '.[0].Config.User // "root"')
if [[ "$USER_ID" == "root" || "$USER_ID" == "" || "$USER_ID" == "0" ]]; then
    CONFIG_ISSUES+=("Running as root user")
    print_warning "Image runs as root user"
else
    print_success "Image runs as non-root user ($USER_ID)"
fi

# Check exposed ports
EXPOSED_PORTS=$(docker inspect "${FULL_IMAGE_NAME}" | jq -r '.[0].Config.ExposedPorts // {} | keys[]' 2>/dev/null || echo "")
if [[ -n "$EXPOSED_PORTS" ]]; then
    print_success "Exposed ports: $(echo "$EXPOSED_PORTS" | tr '\n' ' ')"
    # Check for privileged ports
    while IFS= read -r port; do
        port_num=$(echo "$port" | cut -d'/' -f1)
        if [[ "$port_num" -lt 1024 ]]; then
            CONFIG_ISSUES+=("Privileged port exposed: $port")
            print_warning "Privileged port exposed: $port"
        fi
    done <<< "$EXPOSED_PORTS"
else
    print_success "No ports explicitly exposed"
fi

# Check for secrets in environment variables
ENV_VARS=$(docker inspect "${FULL_IMAGE_NAME}" | jq -r '.[0].Config.Env[]?' 2>/dev/null || echo "")
SECRET_KEYWORDS=("PASSWORD" "SECRET" "KEY" "TOKEN" "API_KEY" "PRIVATE")
while IFS= read -r env_var; do
    for keyword in "${SECRET_KEYWORDS[@]}"; do
        if echo "$env_var" | grep -qi "$keyword"; then
            CONFIG_ISSUES+=("Potential secret in environment variable: $env_var")
            print_warning "Potential secret in environment variable: $(echo "$env_var" | cut -d'=' -f1)"
        fi
    done
done <<< "$ENV_VARS"

# 5. Image Signing Verification (if using Cosign)
print_section "5. Image Signature Verification"

if command_exists cosign; then
    echo "Checking for image signatures..."
    if cosign verify "${FULL_IMAGE_NAME}" 2>/dev/null; then
        print_success "Image signature verified"
    else
        print_warning "Image is not signed or signature verification failed"
        echo "Consider signing your images with: cosign sign ${FULL_IMAGE_NAME}"
    fi
else
    print_warning "Cosign not available, skipping signature verification"
fi

# 6. Base Image Analysis
print_section "6. Base Image Analysis"

# Extract base image information
BASE_IMAGE=$(docker history "${FULL_IMAGE_NAME}" --format "table {{.CreatedBy}}" --no-trunc | grep -E "FROM|COPY --from" | tail -1 | sed 's/.*FROM \([^ ]*\).*/\1/' || echo "unknown")
if [[ "$BASE_IMAGE" != "unknown" && "$BASE_IMAGE" != "" ]]; then
    print_success "Base image detected: $BASE_IMAGE"
    
    # Check if base image is minimal
    if echo "$BASE_IMAGE" | grep -qi "scratch\|distroless\|alpine"; then
        print_success "Using minimal base image"
    else
        CONFIG_ISSUES+=("Not using minimal base image")
        print_warning "Consider using minimal base image (scratch, distroless, alpine)"
    fi
else
    print_warning "Could not determine base image"
fi

# 7. Layer Analysis
print_section "7. Layer Analysis"

echo "Analyzing image layers..."
LAYER_COUNT=$(docker history "${FULL_IMAGE_NAME}" --format "table {{.ID}}" | wc -l)
IMAGE_SIZE=$(docker inspect "${FULL_IMAGE_NAME}" | jq -r '.[0].Size')
IMAGE_SIZE_MB=$((IMAGE_SIZE / 1024 / 1024))

print_success "Image has $LAYER_COUNT layers"
print_success "Image size: ${IMAGE_SIZE_MB}MB"

if [[ "$LAYER_COUNT" -gt 20 ]]; then
    CONFIG_ISSUES+=("High number of layers: $LAYER_COUNT")
    print_warning "High number of layers ($LAYER_COUNT) - consider optimizing Dockerfile"
fi

if [[ "$IMAGE_SIZE_MB" -gt 500 ]]; then
    CONFIG_ISSUES+=("Large image size: ${IMAGE_SIZE_MB}MB")
    print_warning "Large image size (${IMAGE_SIZE_MB}MB) - consider optimization"
fi

# 8. Generate Security Report
print_section "8. Security Report Generation"

REPORT_FILE="${REPORT_DIR}/security-report-${TIMESTAMP}.json"

cat > "$REPORT_FILE" << EOF
{
  "image": "${FULL_IMAGE_NAME}",
  "scan_timestamp": "$(date -Iseconds)",
  "vulnerabilities": {
    "trivy_critical": $CRITICAL_COUNT,
    "trivy_high": $HIGH_COUNT,
    "grype_critical": $GRYPE_CRITICAL,
    "grype_high": $GRYPE_HIGH
  },
  "configuration": {
    "user": "$USER_ID",
    "layer_count": $LAYER_COUNT,
    "size_mb": $IMAGE_SIZE_MB,
    "base_image": "$BASE_IMAGE"
  },
  "issues": [
EOF

# Add configuration issues
for ((i=0; i<${#CONFIG_ISSUES[@]}; i++)); do
    echo "    \"${CONFIG_ISSUES[i]}\"" >> "$REPORT_FILE"
    if [[ $i -lt $((${#CONFIG_ISSUES[@]} - 1)) ]]; then
        echo "," >> "$REPORT_FILE"
    fi
done

cat >> "$REPORT_FILE" << EOF
  ],
  "recommendations": [
    "Fix all CRITICAL and HIGH vulnerabilities",
    "Use minimal base images (scratch, distroless, alpine)",
    "Run as non-root user",
    "Sign images with cosign",
    "Regularly update base images and dependencies",
    "Minimize image layers and size",
    "Avoid secrets in environment variables",
    "Use multi-stage builds to reduce attack surface"
  ]
}
EOF

print_success "Security report generated: $REPORT_FILE"

# 9. Security Score Calculation
print_section "9. Security Score"

SECURITY_SCORE=100
SCORE_DEDUCTIONS=0

# Deduct points for vulnerabilities
SCORE_DEDUCTIONS=$((SCORE_DEDUCTIONS + CRITICAL_COUNT * 20 + HIGH_COUNT * 5))

# Deduct points for configuration issues
SCORE_DEDUCTIONS=$((SCORE_DEDUCTIONS + ${#CONFIG_ISSUES[@]} * 10))

SECURITY_SCORE=$((SECURITY_SCORE - SCORE_DEDUCTIONS))
if [[ $SECURITY_SCORE -lt 0 ]]; then
    SECURITY_SCORE=0
fi

echo "Security Score: $SECURITY_SCORE/100"

if [[ $SECURITY_SCORE -ge 90 ]]; then
    print_success "Excellent security posture!"
elif [[ $SECURITY_SCORE -ge 70 ]]; then
    print_warning "Good security posture with room for improvement"
elif [[ $SECURITY_SCORE -ge 50 ]]; then
    print_warning "Moderate security posture - address critical issues"
else
    print_error "Poor security posture - immediate attention required"
fi

# 10. Summary and Recommendations
print_section "10. Summary"

echo "Security Check Complete!"
echo "========================"
echo "Image: $FULL_IMAGE_NAME"
echo "Security Score: $SECURITY_SCORE/100"
echo "Critical Vulnerabilities: $CRITICAL_COUNT"
echo "High Vulnerabilities: $HIGH_COUNT"
echo "Configuration Issues: ${#CONFIG_ISSUES[@]}"
echo ""
echo "Reports generated in: $REPORT_DIR/"
echo "- Trivy report: trivy-report-${TIMESTAMP}.json"
echo "- SBOM: sbom-${TIMESTAMP}.json"
echo "- Grype report: grype-report-${TIMESTAMP}.json"
echo "- Security summary: security-report-${TIMESTAMP}.json"

# Exit with appropriate code
if [[ $CRITICAL_COUNT -gt 0 ]]; then
    print_error "Build should fail due to CRITICAL vulnerabilities"
    exit 1
elif [[ $SECURITY_SCORE -lt 70 ]]; then
    print_warning "Build warning due to security score below 70"
    exit 2
else
    print_success "Security check passed!"
    exit 0
fi