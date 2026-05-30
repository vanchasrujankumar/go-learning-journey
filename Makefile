.PHONY: help build run test clean fmt lint docs docker-up docker-down install-tools

# Variables
GO := go
GOFMT := gofmt
GOLINT := golangci-lint
DOCKER_COMPOSE := docker-compose

# Colors for output
BLUE := \033[0;34m
GREEN := \033[0;32m
YELLOW := \033[0;33m
RED := \033[0;31m
NC := \033[0m # No Color

help: ## Display this help message
	@echo "$(BLUE)Go Learning Journey - Available Commands$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(GREEN)%-15s$(NC) %s\n", $$1, $$2}'
	@echo ""
	@echo "$(YELLOW)Usage:$(NC)"
	@echo "  make [command]"
	@echo "  make build module=01-basics"
	@echo "  make run example=01-hello-world"

build: ## Build all Go programs
	@echo "$(BLUE)Building Go programs...$(NC)"
	@find . -name "*.go" -type f ! -path "./vendor/*" ! -path "./.git/*" -exec dirname {} \; | sort -u | while read dir; do \
		if [ -f "$$dir/go.mod" ]; then \
			echo "$(GREEN)Building $$dir$(NC)"; \
			cd $$dir && $(GO) build -o bin/app . || echo "$(RED)Failed to build $$dir$(NC)"; \
			cd - > /dev/null; \
		fi; \
	done
	@echo "$(GREEN)Build complete!$(NC)"

run: ## Run a specific example (make run example=01-hello-world module=01-basics)
	@if [ -z "$(module)" ] || [ -z "$(example)" ]; then \
		echo "$(RED)Error: Please specify module and example$(NC)"; \
		echo "$(YELLOW)Usage: make run module=01-basics example=01-hello-world$(NC)"; \
		exit 1; \
	fi
	@if [ ! -f "$(module)/examples/$(example).go" ]; then \
		echo "$(RED)Error: File not found: $(module)/examples/$(example).go$(NC)"; \
		exit 1; \
	fi
	@echo "$(BLUE)Running $(example) from $(module)...$(NC)"
	@cd $(module) && $(GO) run examples/$(example).go

test: ## Run all tests
	@echo "$(BLUE)Running tests...$(NC)"
	@$(GO) test -v -race -cover ./...
	@echo "$(GREEN)Tests complete!$(NC)"

coverage: ## Generate test coverage report
	@echo "$(BLUE)Generating coverage report...$(NC)"
	@$(GO) test -cover -coverprofile=coverage.out ./...
	@$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report saved to coverage.html$(NC)"

bench: ## Run benchmarks
	@echo "$(BLUE)Running benchmarks...$(NC)"
	@$(GO) test -bench=. -benchmem ./...

fmt: ## Format all Go code
	@echo "$(BLUE)Formatting Go code...$(NC)"
	@$(GO) fmt ./...
	@echo "$(GREEN)Formatting complete!$(NC)"

lint: ## Run linter
	@echo "$(BLUE)Running linter...$(NC)"
	@if ! command -v $(GOLINT) &> /dev/null; then \
		echo "$(YELLOW)Installing golangci-lint...$(NC)"; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@$(GOLINT) run ./...
	@echo "$(GREEN)Linting complete!$(NC)"

vet: ## Run go vet
	@echo "$(BLUE)Running go vet...$(NC)"
	@$(GO) vet ./...
	@echo "$(GREEN)Vet check complete!$(NC)"

clean: ## Remove build artifacts and cache
	@echo "$(BLUE)Cleaning up...$(NC)"
	@find . -name "bin" -type d -exec rm -rf {} + 2>/dev/null || true
	@find . -name "dist" -type d -exec rm -rf {} + 2>/dev/null || true
	@$(GO) clean -cache -modcache -testcache
	@rm -f coverage.out coverage.html
	@echo "$(GREEN)Clean complete!$(NC)"

install-tools: ## Install required development tools
	@echo "$(BLUE)Installing development tools...$(NC)"
	@echo "$(YELLOW)Installing golangci-lint...$(NC)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "$(YELLOW)Installing goimports...$(NC)"
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "$(YELLOW)Installing golines...$(NC)"
	@go install github.com/segmentio/golines@latest
	@echo "$(GREEN)Tools installed!$(NC)"

docs: ## Generate documentation
	@echo "$(BLUE)Generating documentation...$(NC)"
	@echo "$(YELLOW)Documentation is in docs/ directory$(NC)"
	@ls -la docs/

docker-up: ## Start Docker containers
	@echo "$(BLUE)Starting Docker containers...$(NC)"
	@$(DOCKER_COMPOSE) up -d
	@echo "$(GREEN)Containers started!$(NC)"
	@$(DOCKER_COMPOSE) ps

docker-down: ## Stop Docker containers
	@echo "$(BLUE)Stopping Docker containers...$(NC)"
	@$(DOCKER_COMPOSE) down
	@echo "$(GREEN)Containers stopped!$(NC)"

docker-logs: ## View Docker container logs
	@$(DOCKER_COMPOSE) logs -f

docker-build: ## Build Docker image
	@echo "$(BLUE)Building Docker image...$(NC)"
	@$(DOCKER_COMPOSE) build
	@echo "$(GREEN)Docker image built!$(NC)"

git-status: ## Show git status
	@echo "$(BLUE)Git Status$(NC)"
	@git status

git-log: ## Show recent commits
	@echo "$(BLUE)Recent Commits$(NC)"
	@git log --oneline -10

setup: install-tools ## Setup development environment
	@echo "$(GREEN)Development environment setup complete!$(NC)"
	@echo "$(YELLOW)Next steps:$(NC)"
	@echo "1. Read the README.md"
	@echo "2. Start with Module 01: make run module=01-basics example=01-hello-world"
	@echo "3. Check the CONTRIBUTING.md for guidelines"

all: clean fmt lint test ## Run all checks
	@echo "$(GREEN)All checks complete!$(NC)"

# Module helpers
module-info: ## Show module structure
	@echo "$(BLUE)Available Modules:$(NC)"
	@ls -d */ | grep -E "^[0-9]{2}-" | sort

examples: ## List all examples
	@echo "$(BLUE)Available Examples:$(NC)"
	@find . -path "*/examples/*.go" -type f | sort

.DEFAULT_GOAL := help
