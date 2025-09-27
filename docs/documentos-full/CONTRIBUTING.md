# Contributing to MCP Ultra

🎉 Thank you for considering contributing to MCP Ultra! We're excited to have you as part of our community.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Contributing Guidelines](#contributing-guidelines)
- [Testing Strategy](#testing-strategy)
- [Security Guidelines](#security-guidelines)
- [Pull Request Process](#pull-request-process)
- [Release Process](#release-process)

## 📜 Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/1/code_of_conduct/). By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

## 🚀 Getting Started

### Prerequisites

- **Go 1.22+**
- **Docker & Docker Compose**
- **Make** (for build automation)
- **Git** (obviously!)

### Quick Setup

```bash
# 1. Fork and clone the repository
git clone https://github.com/your-username/mcp-ultra.git
cd mcp-ultra

# 2. Set up environment
cp config/.env.example .env
make install-tools

# 3. Start dependencies
docker-compose up -d postgres redis nats

# 4. Run tests to ensure everything works
make test-fast

# 5. Start the application
make dev
```

## 🛠 Development Setup

### Environment Configuration

1. **Copy the environment template:**
   ```bash
   cp config/.env.example .env
   ```

2. **Edit `.env` with your local settings:**
   ```bash
   # Database
   POSTGRES_PASSWORD=your_dev_password
   
   # JWT Secret (generate a strong one!)
   JWT_SECRET=your_super_secure_jwt_secret_key_minimum_32_characters
   
   # Development settings
   LOG_LEVEL=debug
   DEBUG_MODE=true
   ```

### Development Tools

Install required development tools:
```bash
make install-tools
```

This installs:
- `golangci-lint` - Code linting
- `air` - Live reload for development
- `migrate` - Database migrations
- `buf` - Protocol buffer management
- `govulncheck` - Vulnerability scanning

### IDE Setup

#### VS Code
1. Install the Go extension
2. Add this to your `.vscode/settings.json`:
   ```json
   {
     "go.lintTool": "golangci-lint",
     "go.testTimeout": "10m",
     "go.coverOnTestPackage": true,
     "go.testFlags": ["-race", "-count=1"],
     "files.exclude": {
       "**/.git": true,
       "**/bin": true,
       "**/coverage.out": true
     }
   }
   ```

#### GoLand/IntelliJ
- Enable Go modules support
- Configure golangci-lint as external tool
- Set up run configurations for different make targets

## 🤝 Contributing Guidelines

### Types of Contributions

We welcome various types of contributions:

- 🐛 **Bug Reports**: Help us identify and fix issues
- 🚀 **Feature Requests**: Suggest new functionality
- 💡 **Enhancements**: Improve existing features
- 📖 **Documentation**: Improve docs, examples, tutorials
- 🧪 **Tests**: Add missing tests, improve test coverage
- 🔒 **Security**: Report security vulnerabilities (privately)
- 🎨 **UI/UX**: Improve user experience

### Before You Start

1. **Check existing issues** to avoid duplicate work
2. **Discuss major changes** in an issue first
3. **Follow the coding standards** described below
4. **Ensure compatibility** with our supported Go versions

### Coding Standards

#### Go Code Style
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` and `goimports` for formatting
- Write godoc comments for public functions
- Keep functions small and focused
- Use meaningful variable and function names

#### Project Structure
- Follow Clean Architecture principles
- Separate concerns: domain, repository, service, handler layers
- Place new features in appropriate packages
- Maintain consistent import ordering

#### Error Handling
```go
// ✅ Good: Wrap errors with context
if err := someOperation(); err != nil {
    return fmt.Errorf("failed to perform operation: %w", err)
}

// ❌ Bad: Return raw errors
if err := someOperation(); err != nil {
    return err
}
```

#### Logging
```go
// ✅ Good: Structured logging with context
logger.Info("User authenticated", 
    zap.String("user_id", userID),
    zap.String("method", "jwt"),
    zap.Duration("duration", authDuration))

// ❌ Bad: Plain string logging
log.Printf("User %s authenticated with JWT in %v", userID, authDuration)
```

#### Security Guidelines
- Never hardcode secrets or passwords
- Use environment variables for configuration
- Validate all inputs
- Use prepared statements for SQL queries
- Implement proper authentication and authorization
- Follow principle of least privilege

### Documentation Standards

- Keep README.md up to date
- Document new features in CHANGELOG.md
- Add inline comments for complex logic
- Include examples in godoc comments
- Update API documentation when changing endpoints

## 🧪 Testing Strategy

MCP Ultra implements a comprehensive 9-layer testing strategy:

### Layer 1: Unit Tests (Fast)
```bash
make test-unit
```
- Test individual functions and methods
- Mock external dependencies
- Aim for >80% code coverage
- Run in <30 seconds

### Layer 2: Property-Based Tests
```bash
make test-property
```
- Test with generated random inputs
- Verify invariants and properties
- Catch edge cases

### Layer 3: Component Tests
```bash
make test-component
```
- Test individual services in isolation
- Use test doubles for dependencies
- Verify service contracts

### Layer 4: Integration Tests
```bash
make test-integration
```
- Test with real dependencies (database, cache)
- Verify data flow between components
- Use Docker containers for dependencies

### Layer 5: Chaos Tests
```bash
make test-chaos
```
- Test system resilience
- Simulate network failures, timeouts
- Verify graceful degradation

### Layer 6: Performance Tests
```bash
make test-performance
```
- Benchmark critical paths
- Load testing with concurrent requests
- Memory usage analysis

### Layer 7: Security Tests
```bash
make test-security
```
- Authentication and authorization tests
- Input validation testing
- Vulnerability scanning

### Layer 8: Contract Tests
```bash
make test-contract
```
- API contract verification
- Schema validation
- Backward compatibility

### Layer 9: End-to-End Tests
```bash
make test-e2e
```
- Full system integration
- User journey testing
- Complete Docker environment

### Test Development Guidelines

#### Writing Good Tests
```go
func TestUserService_CreateUser(t *testing.T) {
    // Arrange
    ctx := context.Background()
    userRepo := &mockUserRepository{}
    service := NewUserService(userRepo)
    
    // Act
    user, err := service.CreateUser(ctx, "john@example.com", "password123")
    
    // Assert
    assert.NoError(t, err)
    assert.NotEmpty(t, user.ID)
    assert.Equal(t, "john@example.com", user.Email)
}
```

#### Test Naming Convention
- `TestFunctionName_Scenario_ExpectedResult`
- Use table-driven tests for multiple scenarios
- Include both positive and negative test cases

#### Test Organization
```
test/
├── unit/           # Layer 1: Unit tests
├── property/       # Layer 2: Property-based tests
├── component/      # Layer 3: Component tests
├── integration/    # Layer 4: Integration tests
├── chaos/          # Layer 5: Chaos engineering tests
├── performance/    # Layer 6: Performance tests
├── security/       # Layer 7: Security tests
├── contract/       # Layer 8: Contract tests
└── e2e/           # Layer 9: End-to-end tests
```

## 🔒 Security Guidelines

### Reporting Security Issues

**⚠️ IMPORTANT**: Do not open public GitHub issues for security vulnerabilities.

Instead, please email security issues to: **security@vertikon.com**

Include:
- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Any suggested fixes

### Security Best Practices

- Keep dependencies updated
- Follow OWASP guidelines
- Use static analysis tools
- Implement proper input validation
- Use secure communication (HTTPS/TLS)
- Apply principle of least privilege
- Regular security audits

## 🔄 Pull Request Process

### Before Submitting

1. **Create a feature branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes:**
   - Write clean, tested code
   - Follow coding standards
   - Update documentation

3. **Run the full test suite:**
   ```bash
   make ci-pipeline
   ```

4. **Commit with conventional commits:**
   ```bash
   git commit -m "feat: add user authentication middleware"
   ```

### Commit Message Format

We follow [Conventional Commits](https://conventionalcommits.org/):

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

**Types:**
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that don't affect code meaning
- `refactor`: Code changes that neither fix a bug nor add a feature
- `perf`: Performance improvements
- `test`: Adding missing tests
- `chore`: Changes to build process or auxiliary tools

**Examples:**
```
feat(auth): add JWT authentication middleware
fix(health): resolve database connection health check
docs(api): update OpenAPI specification
test(auth): add unit tests for JWT validation
```

### PR Checklist

- [ ] 🧪 All tests pass (`make test-complete`)
- [ ] 🔍 Code linting passes (`make lint`)
- [ ] 📖 Documentation updated
- [ ] 🔒 Security considerations addressed
- [ ] 🎯 Feature flag added (if applicable)
- [ ] 📊 Metrics/observability added
- [ ] 🐳 Docker build succeeds
- [ ] ☸️ Kubernetes manifests updated (if needed)

### Review Process

1. **Automated Checks**: GitHub Actions will run automatically
2. **Code Review**: At least one maintainer must approve
3. **Security Review**: Required for security-related changes
4. **Testing**: All CI checks must pass
5. **Documentation**: Verify docs are updated

### Merging

- Use "Squash and merge" for feature branches
- Use "Create a merge commit" for release branches
- Delete feature branch after merge

## 🚀 Release Process

### Versioning

We use [Semantic Versioning](https://semver.org/):
- `MAJOR.MINOR.PATCH`
- `MAJOR`: Breaking changes
- `MINOR`: New features (backward compatible)
- `PATCH`: Bug fixes (backward compatible)

### Release Workflow

1. **Create release branch:**
   ```bash
   git checkout -b release/v1.2.0
   ```

2. **Update version and changelog:**
   - Update `CHANGELOG.md`
   - Bump version in relevant files

3. **Run release checks:**
   ```bash
   make release-build
   make ci-pipeline
   ```

4. **Create PR to main:**
   - Title: `Release v1.2.0`
   - Include changelog in description

5. **After merge, create GitHub release:**
   - Tag: `v1.2.0`
   - Include release notes
   - Attach release artifacts

## 🎯 Areas We Need Help

### High Priority
- [ ] **Performance optimization** - Database queries, caching strategies
- [ ] **Documentation** - More examples, tutorials, API guides
- [ ] **Testing** - Increase test coverage, add more integration tests
- [ ] **Monitoring** - Enhanced dashboards, alerting rules

### Medium Priority
- [ ] **New features** - GraphQL support, WebAssembly plugins
- [ ] **DevOps** - Helm charts, Terraform modules
- [ ] **Security** - Advanced RBAC, audit logging
- [ ] **Compliance** - GDPR/LGPD automation

### Good First Issues
Look for issues labeled `good first issue` - these are perfect for new contributors!

## 📞 Getting Help

### Communication Channels
- **GitHub Issues**: Bug reports, feature requests
- **GitHub Discussions**: Questions, ideas, community chat
- **Discord**: Real-time chat with maintainers and community
- **Email**: Direct contact for sensitive issues

### Documentation
- **README.md**: Getting started guide
- **API Docs**: OpenAPI/Swagger documentation
- **Architecture Docs**: System design and patterns
- **Runbooks**: Operational procedures

## 🙏 Recognition

We value all contributions! Contributors will be:
- Listed in our README
- Mentioned in release notes
- Invited to maintainer discussions (for regular contributors)
- Given special Discord roles

## 📄 License

By contributing to MCP Ultra, you agree that your contributions will be licensed under the MIT License.

---

**Happy coding! 🚀**

Together, we're building the future of enterprise microservices, one commit at a time.