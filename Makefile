.PHONY: build run dev test clean docker-build docker-up docker-down docker-logs swagger

# Build the application
build:
	go build -o bin/godago-rest-api ./cmd/main.go

# Run the application
run: build
	./bin/godago-rest-api

# Run in development mode with hot reload (requires air)
dev:
	air -c .air.toml || go run ./cmd/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Download dependencies
deps:
	go mod download
	go mod tidy

# Generate swagger docs (requires swag)
swagger:
	swag init -g cmd/main.go -o docs

# Docker commands
docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

# Production docker commands
docker-prod-up:
	docker-compose -f docker-compose.prod.yml up -d

docker-prod-down:
	docker-compose -f docker-compose.prod.yml down

docker-prod-logs:
	docker-compose -f docker-compose.prod.yml logs -f

# Database migrations
migrate-up:
	mysql -h 127.0.0.1 -u root -p < migrations/01_create_users_table.sql

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Help
help:
	@echo "Available commands:"
	@echo "  build           - Build the application"
	@echo "  run             - Build and run the application"
	@echo "  dev             - Run in development mode"
	@echo "  test            - Run tests"
	@echo "  test-coverage   - Run tests with coverage"
	@echo "  clean           - Clean build artifacts"
	@echo "  deps            - Download dependencies"
	@echo "  swagger         - Generate swagger documentation"
	@echo "  docker-build    - Build Docker image"
	@echo "  docker-up       - Start Docker containers"
	@echo "  docker-down     - Stop Docker containers"
	@echo "  docker-logs     - View Docker logs"
	@echo "  docker-prod-up  - Start production containers"
	@echo "  docker-prod-down- Stop production containers"
	@echo "  migrate-up      - Run database migrations"
	@echo "  fmt             - Format code"
	@echo "  lint            - Lint code"
