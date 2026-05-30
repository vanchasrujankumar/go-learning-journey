.PHONY: help build run test clean docker-up docker-down logs

PROJECT_NAME=go-api
BIN_NAME=server
DOCKER_COMPOSE=docker-compose

help: ## Display this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building $(PROJECT_NAME)..."
	@go build -o bin/$(BIN_NAME) ./cmd/server
	@echo "✓ Build complete"

run: docker-up ## Run the application
	@echo "Starting API server..."
	@sleep 5
	@go run ./cmd/server/main.go

test: ## Run tests
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...
	@echo "✓ Tests complete"

coverage: test ## Generate coverage report
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Coverage report generated (coverage.html)"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "✓ Clean complete"

docker-up: ## Start Docker services (MongoDB, Kafka, etc.)
	@echo "Starting Docker services..."
	@$(DOCKER_COMPOSE) up -d
	@echo "✓ Docker services started"
	@$(DOCKER_COMPOSE) ps

docker-down: ## Stop Docker services
	@echo "Stopping Docker services..."
	@$(DOCKER_COMPOSE) down
	@echo "✓ Docker services stopped"

logs: ## View Docker logs
	@$(DOCKER_COMPOSE) logs -f

fmt: ## Format code
	@go fmt ./...
	@echo "✓ Code formatted"

lint: ## Run linter
	@golangci-lint run ./...

vet: ## Run go vet
	@go vet ./...

deps: ## Download dependencies
	@go mod download
	@go mod tidy
	@echo "✓ Dependencies downloaded"

mongo-shell: ## Connect to MongoDB shell
	@$(DOCKER_COMPOSE) exec mongo mongosh

kafka-topics: ## List Kafka topics
	@$(DOCKER_COMPOSE) exec kafka kafka-topics --bootstrap-server localhost:9092 --list

kafka-describe: ## Describe Kafka topics
	@$(DOCKER_COMPOSE) exec kafka kafka-topics --bootstrap-server localhost:9092 --describe

# API Testing
api-create-user: ## Create a test user (requires running API)
	@curl -X POST http://localhost:8080/api/v1/users \
		-H "Content-Type: application/json" \
		-d '{"name":"John Doe","email":"john@example.com","password":"password123"}'

api-list-users: ## List all users
	@curl -X GET http://localhost:8080/api/v1/users

api-health: ## Check API health
	@curl -X GET http://localhost:8080/health

all: clean build test ## Run all checks

.DEFAULT_GOAL := help
