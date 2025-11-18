# Security Policy

## Supported Versions

We release patches for security vulnerabilities. Which versions are eligible for receiving such patches depends on the CVSS v3.0 Rating:

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take the security of MCP Ultra seriously. If you have discovered a security vulnerability in this project, please report it to us as described below.

### Reporting Process

1. **DO NOT** open a public GitHub issue if the vulnerability is security-related
2. Email us at: security@vertikon.com
3. Include the following information:
   - Type of issue (e.g., buffer overflow, SQL injection, cross-site scripting, etc.)
   - Full paths of source file(s) related to the manifestation of the issue
   - Location of affected source code (tag/branch/commit or direct URL)
   - Step-by-step instructions to reproduce the issue
   - Proof-of-concept or exploit code (if possible)
   - Impact of the issue, including how an attacker might exploit it

### Response Timeline

- **Initial Response**: Within 48 hours
- **Status Update**: Within 5 business days
- **Resolution Target**:
  - Critical: 7 days
  - High: 14 days
  - Medium: 30 days
  - Low: 90 days

## Security Measures

### Code Security

- **Static Analysis**: All code undergoes static security analysis using:
  - CodeQL for comprehensive SAST
  - gosec for Go-specific security issues
  - golangci-lint with security linters enabled

- **Dependency Management**:
  - Regular dependency updates
  - Automated vulnerability scanning with Grype and Trivy
  - SBOM generation for supply chain transparency

- **Secret Detection**:
  - Gitleaks integration in CI/CD
  - Pre-commit hooks for local detection
  - Regular repository scanning

### Container Security

- **Base Images**:
  - Using distroless or minimal Alpine images
  - Regular base image updates
  - No root user execution

- **Runtime Security**:
  - Read-only root filesystem
  - Non-root user execution
  - Dropped capabilities
  - Security contexts enforced

### Infrastructure Security

- **Kubernetes Policies**:
  - Network policies enforced
  - Pod Security Standards applied
  - RBAC with least privilege
  - Resource limits mandatory

- **TLS/Encryption**:
  - TLS 1.3 minimum
  - Mutual TLS for service-to-service
  - Encryption at rest for sensitive data

### Authentication & Authorization

- **JWT Implementation**:
  - RS256 signing algorithm
  - Short token lifetime (15 minutes)
  - Refresh token rotation
  - JTI for token revocation

- **RBAC**:
  - Fine-grained permissions
  - Default deny policy
  - Regular permission audits

## Security Checklist

Before deploying to production, ensure:

- [ ] All security scans pass in CI/CD
- [ ] No HIGH or CRITICAL vulnerabilities
- [ ] Secrets are properly managed (not in code)
- [ ] Container runs as non-root
- [ ] Network policies are configured
- [ ] TLS is properly configured
- [ ] Rate limiting is enabled
- [ ] Audit logging is configured
- [ ] Monitoring and alerting are active
- [ ] Backup and recovery tested

## Compliance

This project aims to comply with:

- OWASP Top 10
- CIS Docker Benchmark
- CIS Kubernetes Benchmark
- NIST Cybersecurity Framework
- SOC2 Type II requirements

## Security Tools

### Required for Development

```bash
# Install security tools
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install github.com/google/go-licenses@latest
brew install gitleaks trivy grype conftest
```

### Running Security Checks Locally

```bash
# Run all security validations
./validate-security.sh

# Individual checks
gitleaks detect --config=.gitleaks.toml
gosec -severity high ./...
grype dir:. --fail-on high
trivy image mcp-ultra:latest
```

## Security Contacts

- Security Team: security@vertikon.com
- Security Updates: https://github.com/vertikon/mcp-ultra/security/advisories

## Acknowledgments

We appreciate responsible disclosure and will acknowledge security researchers who report valid vulnerabilities.

---

Last Updated: 2024
Version: 1.0