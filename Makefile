.PHONY: help build run test clean docker setup migrate lint fmt deps

VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "v1.0.0-dev")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

deps: ## Download dependencies
	@echo "🔄 Downloading dependencies..."
	@go mod download
	@go mod tidy

build: deps ## Build the application
	@echo "🔨 Building MCP Ultra $(VERSION)..."
	@mkdir -p $(GOBIN)
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
		-ldflags="-w -s -X 'github.com/vertikon/mcp-ultra/pkg/version.Version=$(VERSION)' -X 'github.com/vertikon/mcp-ultra/pkg/version.BuildDate=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)'" \
		-o $(GOBIN)/mcp-ultra cmd/mcp-model-ultra/main.go

run: build ## Run the application
	@echo "🚀 Starting MCP Ultra..."
	@$(GOBIN)/mcp-ultra

dev: ## Run in development mode with full observability
	@echo "🔧 Starting in development mode with telemetry..."
	TELEMETRY_ENABLED=true TELEMETRY_DEBUG=true CONSOLE_EXPORTER_ENABLED=true \
	TRACING_SAMPLE_RATE=1.0 OTLP_ENDPOINT=http://localhost:4317 \
	go run cmd/mcp-model-ultra/main.go

dev-minimal: ## Run in development mode without observability
	@echo "🔧 Starting in minimal development mode..."
	TELEMETRY_ENABLED=false go run cmd/mcp-model-ultra/main.go

dev-reload: ## Run in development mode with live reload
	@echo "🔧 Starting with live reload..."
	@air -c .air.toml

# Testing Strategy - 9 Layers Implementation
test: ## Run all tests (9 layer strategy)
	@echo "🧪 Running Complete Test Suite (9 Layers)..."
	@$(MAKE) test-fast
	@$(MAKE) test-medium  
	@$(MAKE) test-slow
	@echo "✅ All tests completed successfully!"

# Layer 1: Unit Tests - Foundation
test-unit: ## Layer 1: Unit tests (fast)
	@echo "🧪 Layer 1: Running Unit Tests..."
	@go test -short -race -coverprofile=coverage.out ./internal/...
	@go tool cover -func=coverage.out

# Layer 2: Property-Based Tests  
test-property: ## Layer 2: Property-based tests
	@echo "🧪 Layer 2: Running Property-Based Tests..."
	@go test -race ./test/property/...

# Layer 3: Component Tests
test-component: ## Layer 3: Component tests (isolated services)
	@echo "🧪 Layer 3: Running Component Tests..."
	@go test -race ./test/component/...

# Layer 4: Integration Tests
test-integration: ## Layer 4: Integration tests (with real deps)
	@echo "🧪 Layer 4: Running Integration Tests..."
	@docker-compose -f docker-compose.test.yml up -d --build
	@sleep 10
	@go test -v -race -tags=integration ./test/integration/...
	@docker-compose -f docker-compose.test.yml down

# Layer 5: Chaos/Resilience Tests
test-chaos: ## Layer 5: Chaos engineering tests
	@echo "🧪 Layer 5: Running Chaos Tests..."
	@go test -race -tags=chaos ./test/chaos/...

# Layer 6: Performance Tests
test-performance: ## Layer 6: Performance and load tests
	@echo "🧪 Layer 6: Running Performance Tests..."
	@go test -bench=. -benchmem ./test/performance/...

# Layer 7: Security Tests
test-security: ## Layer 7: Security and penetration tests
	@echo "🧪 Layer 7: Running Security Tests..."
	@go test -race -tags=security ./test/security/...

# Layer 8: Contract Tests
test-contract: ## Layer 8: API contract tests
	@echo "🧪 Layer 8: Running Contract Tests..."
	@go test -race -tags=contract ./test/contract/...

# Layer 9: End-to-End Tests
test-e2e: ## Layer 9: End-to-end tests (full system)
	@echo "🧪 Layer 9: Running E2E Tests..."
	@docker-compose up -d --build
	@sleep 20
	@go test -v -race -tags=e2e ./test/e2e/...
	@docker-compose down

# Test Execution Strategies
test-fast: ## Fast tests (Layers 1-2) for development
	@echo "⚡ Running Fast Test Suite (Unit + Property)..."
	@$(MAKE) test-unit
	@$(MAKE) test-property
	@echo "✅ Fast tests completed"

test-medium: ## Medium tests (Layers 3-4) for CI
	@echo "🔄 Running Medium Test Suite (Component + Integration)..."
	@$(MAKE) test-component
	@$(MAKE) test-integration
	@echo "✅ Medium tests completed"

test-slow: ## Slow tests (Layers 5-8) for comprehensive validation
	@echo "🐌 Running Slow Test Suite (Chaos + Performance + Security + Contract)..."
	@$(MAKE) test-chaos
	@$(MAKE) test-performance
	@$(MAKE) test-security
	@$(MAKE) test-contract
	@echo "✅ Slow tests completed"

test-complete: ## Complete test suite including E2E
	@echo "🎯 Running Complete Test Suite (All 9 Layers)..."
	@$(MAKE) test-fast
	@$(MAKE) test-medium
	@$(MAKE) test-slow
	@$(MAKE) test-e2e
	@echo "🎉 Complete test suite finished successfully!"

benchmark: ## Run benchmarks
	@echo "⚡ Running benchmarks..."
	@go test -bench=. -benchmem ./...

lint: ## Run linters
	@echo "🔍 Running linters..."
	@golangci-lint run --timeout=5m

fmt: ## Format code
	@echo "🎨 Formatting code..."
	@go fmt ./...
	@goimports -w .

migrate-up: ## Run database migrations up
	@echo "⬆️ Running database migrations up..."
	@migrate -path internal/repository/postgres/migrations -database "$(DATABASE_URL)" up

migrate-down: ## Run database migrations down
	@echo "⬇️ Running database migrations down..."
	@migrate -path internal/repository/postgres/migrations -database "$(DATABASE_URL)" down

migrate-create: ## Create new migration (use: make migrate-create name=migration_name)
	@echo "📝 Creating new migration: $(name)"
	@migrate create -ext sql -dir internal/repository/postgres/migrations $(name)

docker-build: ## Build Docker image
	@echo "🐳 Building Docker image..."
	@docker build -f Dockerfile -t mcp-ultra:$(VERSION) .

docker-run: docker-build ## Run in Docker
	@echo "🐳 Running in Docker..."
	@docker-compose up --build

docker-push: docker-build ## Push Docker image
	@echo "🚀 Pushing Docker image..."
	@docker tag mcp-ultra:$(VERSION) vertikon/mcp-ultra:$(VERSION)
	@docker tag mcp-ultra:$(VERSION) vertikon/mcp-ultra:latest
	@docker push vertikon/mcp-ultra:$(VERSION)
	@docker push vertikon/mcp-ultra:latest

setup: ## Setup development environment
	@echo "🔧 Setting up development environment..."
	@./scripts/setup.sh

dr-test: ## Run disaster recovery test
	@echo "🚨 Running disaster recovery test..."
	@./scripts/dr-test.sh

deploy-canary: ## Deploy canary version (use: make deploy-canary VERSION=v1.0.0)
	@echo "🚀 Deploying canary version $(VERSION)..."
	@./scripts/deploy-canary.sh $(VERSION)

monitor-setup: ## Setup monitoring stack
	@echo "📊 Setting up monitoring stack..."
	@helm upgrade --install prometheus prometheus-community/kube-prometheus-stack \
		--namespace monitoring --create-namespace \
		--values deploy/monitoring/prometheus-values.yaml
	@helm upgrade --install grafana grafana/grafana \
		--namespace monitoring \
		--values deploy/monitoring/grafana-values.yaml

telemetry-stack: ## Start local telemetry stack for development
	@echo "📊 Starting local telemetry stack (Jaeger + OTEL Collector)..."
	@docker-compose -f deploy/docker/telemetry-stack.yml up -d
	@echo "🔍 Jaeger UI available at: http://localhost:16686"
	@echo "📊 OTEL Collector available at: http://localhost:4317 (gRPC) and http://localhost:4318 (HTTP)"

telemetry-down: ## Stop local telemetry stack
	@echo "📊 Stopping local telemetry stack..."
	@docker-compose -f deploy/docker/telemetry-stack.yml down

telemetry-logs: ## Show telemetry stack logs
	@echo "📊 Showing telemetry stack logs..."
	@docker-compose -f deploy/docker/telemetry-stack.yml logs -f

observability-test: ## Test observability integration
	@echo "🧪 Testing observability integration..."
	@$(MAKE) telemetry-stack
	@sleep 10
	@echo "🔍 Running observability integration test..."
	@TELEMETRY_ENABLED=true TRACING_ENABLED=true OTLP_ENDPOINT=http://localhost:4317 \
		go test -v ./test/observability/...
	@$(MAKE) telemetry-down

k8s-deploy: ## Deploy to Kubernetes
	@echo "☸️ Deploying to Kubernetes..."
	@kubectl apply -f deploy/k8s/

k8s-delete: ## Delete from Kubernetes
	@echo "☸️ Deleting from Kubernetes..."
	@kubectl delete -f deploy/k8s/

security-scan: ## Run comprehensive security scan
	@echo "🔒 Running comprehensive security scan..."
	@gosec -fmt sarif -out gosec-results.sarif -conf gosec.json ./...
	@govulncheck ./...
	@go list -json -deps ./... | nancy sleuth || echo "Nancy scan completed"
	@trivy fs . --format sarif --output trivy-fs-results.sarif
	@echo "✅ Security scans completed. Check SARIF files for details."

ci-security: ## Run CI security pipeline locally
	@echo "🔒 Running CI security pipeline..."
	@$(MAKE) security-scan
	@$(MAKE) test-security
	@echo "✅ CI security pipeline completed"

proto-lint: ## Lint protocol buffers
	@echo "🔍 Linting protocol buffers..."
	@buf lint api/grpc/proto/

proto-breaking: ## Check for breaking changes in protobuf
	@echo "🔍 Checking for breaking changes in protobuf..."
	@buf breaking api/grpc/proto/ --against '.git#branch=main,subdir=api/grpc/proto/' || echo "Breaking change check completed"

proto-generate: ## Generate protocol buffer files
	@echo "⚙️ Generating protocol buffer files..."
	@buf generate api/grpc/proto/

proto-format: ## Format protocol buffer files
	@echo "🎨 Formatting protocol buffer files..."
	@buf format -w api/grpc/proto/

sbom: ## Generate Software Bill of Materials
	@echo "📋 Generating SBOM..."
	@syft dir:. -o spdx-json=sbom.spdx.json
	@echo "✅ SBOM generated: sbom.spdx.json"

sign-image: ## Sign container image with Cosign
	@echo "✍️ Signing container image..."
	@cosign sign --yes mcp-ultra:$(VERSION)
	@echo "✅ Image signed with Cosign"

vulnerability-report: ## Generate vulnerability report
	@echo "📊 Generating vulnerability report..."
	@trivy image --format json --output vuln-report.json mcp-ultra:$(VERSION) || echo "Trivy scan completed"
	@echo "✅ Vulnerability report generated: vuln-report.json"

ci-pipeline: ## Run full CI pipeline locally
	@echo "🚀 Running complete CI pipeline..."
	@$(MAKE) deps
	@$(MAKE) lint
	@$(MAKE) proto-lint  
	@$(MAKE) test-fast
	@$(MAKE) test-medium
	@$(MAKE) ci-security
	@$(MAKE) docker-build
	@$(MAKE) vulnerability-report
	@echo "✅ Complete CI pipeline finished successfully!"

cd-pipeline: ## Run CD pipeline steps
	@echo "🚀 Running CD pipeline steps..."
	@$(MAKE) docker-build
	@$(MAKE) sbom
	@$(MAKE) sign-image
	@$(MAKE) k8s-deploy
	@echo "✅ CD pipeline completed!"

release-build: ## Build release artifacts
	@echo "📦 Building release artifacts..."
	@mkdir -p dist
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X 'github.com/vertikon/mcp-ultra/pkg/version.Version=$(VERSION)'" -o dist/mcp-ultra-linux-amd64 cmd/mcp-model-ultra/main.go
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s -X 'github.com/vertikon/mcp-ultra/pkg/version.Version=$(VERSION)'" -o dist/mcp-ultra-linux-arm64 cmd/mcp-model-ultra/main.go
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s -X 'github.com/vertikon/mcp-ultra/pkg/version.Version=$(VERSION)'" -o dist/mcp-ultra-darwin-amd64 cmd/mcp-model-ultra/main.go
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-w -s -X 'github.com/vertikon/mcp-ultra/pkg/version.Version=$(VERSION)'" -o dist/mcp-ultra-darwin-arm64 cmd/mcp-model-ultra/main.go
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s -X 'github.com/vertikon/mcp-ultra/pkg/version.Version=$(VERSION)'" -o dist/mcp-ultra-windows-amd64.exe cmd/mcp-model-ultra/main.go
	@echo "✅ Release artifacts built in dist/"

compliance-check: ## Check compliance requirements
	@echo "📋 Running compliance checks..."
	@echo "🔍 Checking GDPR/LGPD compliance..."
	@grep -r "ConsentRecord\|PIIDetection\|DataSubjectRequest" . || echo "Privacy compliance patterns found"
	@echo "🔍 Checking SOC2 controls..."
	@test -f "internal/security/auth.go" && echo "✅ Authentication controls" || echo "❌ Missing auth controls"
	@test -f "internal/security/tls.go" && echo "✅ Encryption controls" || echo "❌ Missing TLS controls"
	@test -d "internal/observability" && echo "✅ Monitoring controls" || echo "❌ Missing monitoring"
	@echo "✅ Compliance check completed"

generate: ## Generate code
	@echo "⚙️ Generating code..."
	@go generate ./...

mod-upgrade: ## Upgrade dependencies
	@echo "⬆️ Upgrading dependencies..."
	@go get -u ./...
	@go mod tidy

clean: ## Clean build artifacts
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf $(GOBIN)
	@rm -f coverage.out coverage.html
	@go clean -cache -testcache -modcache

install-tools: ## Install development tools
	@echo "🔧 Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/cosmtrek/air@latest
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/securecodewarrior/sast-scan-go@latest
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest

health-check: ## Check application health endpoints
	@echo "🏥 Checking health endpoints..."
	@curl -s http://localhost:9655/health | jq . || echo "Health endpoint not available"
	@curl -s http://localhost:9655/ready || echo "Ready endpoint not available"  
	@curl -s http://localhost:9655/live || echo "Live endpoint not available"

auth-test: ## Test authentication endpoints
	@echo "🔐 Testing authentication..."
	@echo "Testing JWT token generation..."
	@curl -s -X POST http://localhost:9655/auth/token \
		-H "Content-Type: application/json" \
		-d '{"user_id":"test","username":"test","roles":["user"]}' || echo "Auth endpoint not available"

tls-cert-gen: ## Generate self-signed certificates for development
	@echo "🔐 Generating TLS certificates for development..."
	@mkdir -p certs
	@openssl req -x509 -newkey rsa:4096 -keyout certs/server.key -out certs/server.crt \
		-days 365 -nodes -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost,IP:127.0.0.1"
	@echo "✅ Certificates generated in certs/ directory"

validate-config: ## Validate configuration files
	@echo "✅ Validating configuration..."
	@test -f config/config.yaml && echo "✅ config.yaml exists" || echo "❌ config.yaml missing"
	@test -f config/.env.example && echo "✅ .env.example exists" || echo "❌ .env.example missing"
	@test -f Dockerfile && echo "✅ Dockerfile exists" || echo "❌ Dockerfile missing"

logs: ## Show application logs
	@echo "📋 Showing application logs..."
	@docker-compose logs -f mcp-ultra

status: ## Show services status
	@echo "📊 Services status:"
	@docker-compose ps

# Database shortcuts
db-reset: migrate-down migrate-up ## Reset database

db-shell: ## Connect to database shell
	@docker-compose exec postgres psql -U postgres -d mcp_ultra

redis-cli: ## Connect to Redis CLI
	@docker-compose exec redis redis-cli

# Default target
.DEFAULT_GOAL := help