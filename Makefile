.PHONY: help build run test clean db-up db-down dev

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the Go application
	cd backend && go build -o bin/server main.go

run: ## Run the application
	cd backend && go run main.go

test: ## Run tests
	cd backend && go test ./...

clean: ## Clean build artifacts
	cd backend && rm -rf bin/

db-up: ## Start PostgreSQL database
	docker-compose up -d postgres

db-down: ## Stop PostgreSQL database
	docker-compose down

dev: db-up ## Start development environment
	@echo "Starting development environment..."
	@echo "Database starting up..."
	@sleep 5
	@echo "Starting Go server..."
	cd backend && go run main.go

install: ## Install dependencies
	cd backend && go mod tidy

frontend: ## Serve frontend files
	cd public && python3 -m http.server 3000

# Docker commands
docker-build: ## Build Docker image
	docker build -t nomado-houses .

docker-run: ## Run Docker container
	docker run -p 8080:8080 --env-file backend/.env nomado-houses

docker-up: ## Start application with Docker Compose
	docker-compose up -d

docker-down: ## Stop Docker Compose services
	docker-compose down

docker-logs: ## View Docker Compose logs
	docker-compose logs -f

docker-prod: ## Start production environment
	docker-compose -f docker-compose.prod.yml up -d

docker-prod-down: ## Stop production environment
	docker-compose -f docker-compose.prod.yml down

docker-clean: ## Clean Docker images and containers
	docker-compose down -v
	docker system prune -f
