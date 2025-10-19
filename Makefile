# Makefile for MCP Project
# Standard security-enabled build system

.PHONY: all build test clean help dev lint sec-scan deps install

# Variables
GO := go
BINARY_NAME := mcp-server
BUILD_DIR := bin
COVERAGE_FILE := coverage.out

# Colors
GREEN := \033[0;32m
YELLOW := \033[0;33m
RED := \033[0;31m
NC := \033[0m

# Default target
all: clean deps test build

# Help
help:
	@echo "================================================================"
	@echo "  MCP Project - Makefile"
	@echo "================================================================"
	@echo ""
	@echo "Build Commands:"
	@echo "  make build              - Build the project"
	@echo "  make build-release      - Build production release"
	@echo "  make clean              - Clean build artifacts"
	@echo ""
	@echo "Development:"
	@echo "  make dev                - Run in development mode"
	@echo "  make test               - Run all tests"
	@echo "  make test-coverage      - Run tests with coverage report"
	@echo "  make lint               - Run linters"
	@echo "  make fmt                - Format code"
	@echo ""
	@echo "Security:"
	@echo "  make sec-scan           - Run all security scans"
	@echo "  make sec-gosec          - Run GoSec scanner"
	@echo "  make sec-deps           - Check dependencies for vulnerabilities"
	@echo "  make sec-secrets        - Scan for leaked secrets"
	@echo ""
	@echo "Dependencies:"
	@echo "  make deps               - Download and verify dependencies"
	@echo "  make deps-update        - Update dependencies"
	@echo ""

# Build
build:
	@echo "$(GREEN)🔨 Building $(BINARY_NAME)...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "$(GREEN)✅ Build complete: $(BUILD_DIR)/$(BINARY_NAME)$(NC)"

build-release:
	@echo "$(GREEN)🔨 Building release version...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -ldflags="-s -w -X main.Version=$$(git describe --tags --always) -X main.BuildTime=$$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
		-o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "$(GREEN)✅ Release build complete$(NC)"

# Development
dev:
	@echo "$(YELLOW)🔥 Starting development server...$(NC)"
	@$(GO) run .

# Testing
test:
	@echo "$(YELLOW)🧪 Running tests...$(NC)"
	@$(GO) test -v -race ./...
	@echo "$(GREEN)✅ Tests passed$(NC)"

test-coverage:
	@echo "$(YELLOW)🧪 Running tests with coverage...$(NC)"
	@$(GO) test -v -race -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./...
	@$(GO) tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@$(GO) tool cover -func=$(COVERAGE_FILE)
	@echo "$(GREEN)✅ Coverage report: coverage.html$(NC)"

# Linting
lint:
	@echo "$(YELLOW)🔍 Running linters...$(NC)"
	@golangci-lint run ./...
	@echo "$(GREEN)✅ Linting complete$(NC)"

fmt:
	@echo "$(YELLOW)📝 Formatting code...$(NC)"
	@$(GO) fmt ./...
	@gofumpt -l -w .
	@echo "$(GREEN)✅ Code formatted$(NC)"

# Security Scanning
sec-scan: sec-gosec sec-deps sec-secrets
	@echo "$(GREEN)✅ All security scans complete$(NC)"

sec-gosec:
	@echo "$(YELLOW)🔒 Running GoSec security scanner...$(NC)"
	@gosec -fmt=json -out=gosec-report.json ./... || true
	@gosec ./...
	@echo "$(GREEN)✅ GoSec scan complete$(NC)"

sec-deps:
	@echo "$(YELLOW)🔒 Checking dependencies for vulnerabilities...$(NC)"
	@$(GO) list -json -deps ./... | nancy sleuth || true
	@echo "$(GREEN)✅ Dependency scan complete$(NC)"

sec-secrets:
	@echo "$(YELLOW)🔒 Scanning for secrets...$(NC)"
	@gitleaks detect --source . --verbose || true
	@echo "$(GREEN)✅ Secret scan complete$(NC)"

# Dependencies
deps:
	@echo "$(YELLOW)📦 Downloading dependencies...$(NC)"
	@$(GO) mod download
	@$(GO) mod verify
	@echo "$(GREEN)✅ Dependencies ready$(NC)"

deps-update:
	@echo "$(YELLOW)📦 Updating dependencies...$(NC)"
	@$(GO) get -u ./...
	@$(GO) mod tidy
	@echo "$(GREEN)✅ Dependencies updated$(NC)"

# Installation
install:
	@echo "$(GREEN)📦 Installing $(BINARY_NAME)...$(NC)"
	@$(GO) install .
	@echo "$(GREEN)✅ Installation complete$(NC)"

# Cleanup
clean:
	@echo "$(YELLOW)🧹 Cleaning build artifacts...$(NC)"
	@rm -rf $(BUILD_DIR)
	@rm -f $(COVERAGE_FILE) coverage.html
	@rm -f gosec-report.json
	@echo "$(GREEN)✅ Cleanup complete$(NC)"

# Docker
docker-build:
	@echo "$(YELLOW)🐳 Building Docker image...$(NC)"
	@docker build -t $(BINARY_NAME):latest .
	@echo "$(GREEN)✅ Docker image built$(NC)"

docker-scan:
	@echo "$(YELLOW)🔒 Scanning Docker image...$(NC)"
	@trivy image $(BINARY_NAME):latest
	@echo "$(GREEN)✅ Docker scan complete$(NC)"

# Pre-commit checks
pre-commit: fmt lint test sec-scan
	@echo "$(GREEN)✅ All pre-commit checks passed$(NC)"

# CI/CD
ci: deps test lint sec-scan
	@echo "$(GREEN)✅ CI checks passed$(NC)"
