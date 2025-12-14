# Godago REST API

A simple REST API built with Go, Gin, MySQL, and Swagger/OpenAPI.

## Tech Stack

- **Go 1.21** - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM library
- **MySQL 8.0** - Database
- **Swagger** - API documentation
- **Docker** - Containerization

## Project Structure

```
godago-rest-api/
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── config/           # Database configuration
│   ├── dto/              # Data Transfer Objects
│   ├── errors/           # Error handling
│   ├── handlers/         # HTTP handlers
│   ├── models/           # Database models
│   ├── routes/           # Route definitions
│   └── services/         # Business logic
├── docs/                 # Swagger documentation
├── migrations/           # Database migrations
├── scripts/              # Utility scripts
├── docker-compose.yml    # Docker compose for development
├── docker-compose.prod.yml # Docker compose for production
├── Dockerfile
├── Makefile
└── .env.example
```

## Getting Started

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- MySQL 8.0 (if running locally without Docker)

### Environment Variables

Copy `.env.example` to `.env` and configure:

```bash
cp .env.example .env
```

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | MySQL connection string | `root:@tcp(127.0.0.1:3306)/golang?...` |
| `DB_USER` | Database user | `gouser` |
| `DB_PASSWORD` | Database password | `devpassword` |
| `DB_NAME` | Database name | `golang` |
| `DB_PORT` | Database port | `3306` |
| `HOST` | Server host | `127.0.0.1` |
| `PORT` | Server port | `8080` |
| `GIN_MODE` | Gin mode (debug/release) | `debug` |

### Running with Docker (Recommended)

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Running Locally

```bash
# Install dependencies
make deps

# Run the application
make run

# Or run in development mode
make dev
```

### Running in Production

```bash
docker-compose -f docker-compose.prod.yml up -d
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/swagger/*` | Swagger UI |
| `POST` | `/api/users` | Create user |
| `GET` | `/api/users` | Get all users |
| `GET` | `/api/users/:id` | Get user by ID |
| `PUT` | `/api/users/:id` | Update user |
| `DELETE` | `/api/users/:id` | Delete user |

## Swagger Documentation

Access Swagger UI at:

```
http://<host>:<port>/swagger/index.html
```

### Regenerate Swagger Docs

```bash
make swagger
```

## Available Make Commands

```bash
make build          # Build the application
make run            # Build and run
make dev            # Run in development mode
make test           # Run tests
make test-coverage  # Run tests with coverage
make clean          # Clean build artifacts
make deps           # Download dependencies
make swagger        # Generate swagger docs
make docker-build   # Build Docker image
make docker-up      # Start Docker containers
make docker-down    # Stop Docker containers
make docker-logs    # View Docker logs
make docker-prod-up # Start production containers
make fmt            # Format code
make lint           # Lint code
make help           # Show all commands
```

## API Examples

### Create User

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

### Get All Users

```bash
curl http://localhost:8080/api/users
```

### Get User by ID

```bash
curl http://localhost:8080/api/users/1
```

### Update User

```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "email": "jane@example.com"}'
```

### Delete User

```bash
curl -X DELETE http://localhost:8080/api/users/1
```

## License

MIT License - see [LICENSE](LICENSE) file for details.
