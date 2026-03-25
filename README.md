# Shopping Event-Driven System

A modular event-driven shopping platform built with Go, featuring microservices architecture with gRPC, REST APIs, PostgreSQL persistence, and NATS JetStream for event streaming.

## 📋 Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Key Technologies](#key-technologies)
- [Repository Structure](#repository-structure)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Project Structure](#project-structure)
- [Database & Migrations](#database--migrations)
- [API Documentation](#api-documentation)
- [Code Generation](#code-generation)
- [Development](#development)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Overview

This repository implements a comprehensive shopping platform with separate bounded contexts (Customer, Product, Payment, Order) organized as independent service modules. Each module is built using Domain-Driven Design (DDD) principles and communicates through event-driven patterns.

### Core Features

- **Microservices Architecture**: Customer, Product, Payment, and Order services
- **Multiple Communication Patterns**: gRPC for service-to-service, REST for client APIs
- **Event-Driven Integration**: NATS JetStream for asynchronous event streaming
- **Persistence Layer**: PostgreSQL with structured schemas
- **API Documentation**: Auto-generated Swagger/OpenAPI specifications
- **Developer Experience**: Embedded Swagger UI, Docker Compose setup

## Architecture

### Service Modules

1. **Customer Service** - Customer registration and management
2. **Product Service** - Product catalog management with pricing
3. **Payment Service** - Invoice creation and payment processing
4. **Order Service** - Order orchestration and checkout workflows

### Infrastructure Components

- **PostgreSQL**: Persistent data store with separate schemas per service
- **NATS (JetStream)**: Event streaming for inter-service communication
- **gRPC + grpc-gateway**: Type-safe RPC and REST endpoint generation
- **Swagger UI**: Interactive API documentation (embedded)

## Key Technologies

| Technology              | Purpose                              |
| ----------------------- | ------------------------------------ |
| Go                      | Primary programming language         |
| gRPC                    | Service-to-service communication     |
| grpc-gateway            | REST API generation from protobuf    |
| Protocol Buffers        | Service interface definitions        |
| PostgreSQL 12           | Relational data persistence          |
| NATS 2.9                | Event streaming (JetStream)          |
| buf                     | Protobuf linting and code generation |
| Docker & Docker Compose | Containerization and orchestration   |
| Swagger UI              | API documentation UI                 |

## Repository Structure

```
shopping-event-driven-system/
├── cmd/
│   └── shopping/              # Application entry point
│       ├── application.go      # App initialization
│       └── monolith.go         # Monolith deployment mode
├── customer/                  # Customer service module
│   ├── customerspb/           # Generated protobuf code
│   ├── domain/                # DDD domain layer
│   ├── internal/application/  # Application services
│   ├── infra/                 # Infrastructure (repos, handlers)
│   ├── usecase/               # Business logic
│   ├── buf.yaml               # buf configuration
│   └── *.proto                # Protocol buffer definitions
├── product/                   # Product service module (similar structure)
├── payment/                   # Payment service module (similar structure)
├── order/                     # Order service module (similar structure)
├── internal/
│   ├── am/                    # Event messaging adapters
│   ├── ddd/                   # DDD helpers and abstractions
│   ├── es/                    # Event store implementations
│   ├── jetstream/             # NATS JetStream integration
│   ├── registry/              # Protocol serialization registry
│   ├── web/                   # Embedded web assets (Swagger UI)
│   └── tools.go               # Build tools reference
├── scripts/                   # Database migration SQL files
├── docker/
│   ├── database/              # PostgreSQL init scripts
│   ├── config_nats.conf       # NATS JetStream configuration
│   ├── .env                   # Environment variables
│   └── Dockerfile             # Application container image
├── docker-compose.yaml        # Local development stack
├── docs/                      # Generated OpenAPI specs
└── go.mod                     # Go module dependencies
```

## Prerequisites

- **Go** >= 1.20
- **Docker** & **Docker Compose** (for local development)
- **buf** (recommended for protobuf management): `https://docs.buf.build/installation`
- **PostgreSQL** client tools (optional, for direct database access)

### Optional Tools

- **protoc** and plugins (if using manual protobuf compilation)
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
  ```

## Quick Start

### 1. Start Docker Infrastructure

```bash
docker-compose up -d
```

This starts:

- **PostgreSQL** (port 5432): Main database
- **pgAdmin** (port 83): Database UI
- **NATS** (ports 4222, 8222): Event streaming server

### 2. Verify Connectivity

```bash
# Check PostgreSQL
docker-compose exec postgres psql -U postgres -c "SELECT 1;"

# Check NATS
nc -zv localhost 4222
```

### 3. Build & Run Application

```bash
# Download dependencies
go mod download

# Build
go build -o shopping ./cmd/shopping

# Run
./shopping
```

Or directly:

```bash
go run ./cmd/shopping
```

### 4. Access Services

- **Swagger UI**: http://localhost:8080/swagger-ui/ (REST API documentation)
- **gRPC Services**: localhost:8085
- **Customers**: http://localhost:8080/customers/
- **Products**: http://localhost:8080/products/
- **Payments**: http://localhost:8080/payments/
- **Orders**: http://localhost:8080/orders/

## Project Structure

### Module Organization (e.g., Customer Service)

```
customer/
├── domain/                 # Domain models, value objects, aggregates
├── infra/                  # Repositories, persistence, queries
├── internal/
│   ├── application/
│   │   └── router/
│   │       ├── grpc/       # gRPC handlers
│   │       └── rest/       # REST handlers (generated by grpc-gateway)
│   ├── logging/            # Structured logging
│   ├── handlers/           # Business logic handlers
│   └── middleware/         # Request/response interceptors
├── usecase/                # Application use cases / orchestration
├── customerspb/            # Generated protobuf code
│   ├── customer.pb.go      # Generated message types
│   ├── customer_grpc.pb.go # Generated service stubs
│   └── customer_grpc.gw.pb.go # Generated REST gateway
├── *.proto                 # Service interface definitions
├── buf.yaml                # buf configuration
└── buf.gen.yaml            # buf code generation config
```

### Shared Internal Components

```
internal/
├── am/                     # Aggregate-Message adapter for event handling
├── ddd/                    # Value, Entity, Aggregate abstractions
├── es/                     # Event store implementations
├── jetstream/              # NATS JetStream subscriber/publisher
├── registry/               # Protocol serialization registry
├── web/                    # Embedded Swagger UI static assets
└── container/              # Dependency injection container
```

## Database & Migrations

### Structure

SQL migration files are organized in `scripts/` with naming convention: `{sequence}_{SERVICE}.{direction}.sql`

Example files:

- `1_CUSTOMER.up.sql` - Create customer schema
- `1_CUSTOMER.down.sql` - Drop customer schema
- `2_PRODUCT.up.sql` - Create product schema
- etc.

### Docker Initialization

On first run, Docker initializes the database using scripts in `docker/database/`:

- `1_create_db.sh` - Creates shopping database and user
- `2_create_customers_schema.sh` - Customer schema setup
- `3_create_products_schema.sh` - Product schema setup
- etc.

### Manual Execution

If you need to re-apply migrations:

```bash
# Connect to database
docker-compose exec postgres psql -U postgres -d shopping

# Run migration SQL
\i /path/to/scripts/1_CUSTOMER.up.sql
```

## API Documentation

### Swagger UI

Access auto-generated API documentation at `http://localhost:8080/swagger-ui/`

The Swagger UI aggregates OpenAPI specifications from all service modules:

- Customers API
- Products API
- Payments API
- Orders API

### REST Endpoints

Each module exposes REST endpoints via gRPC-gateway:

```bash
# Customer Service Examples
GET    /customers                    # List customers
POST   /customers                    # Create customer
GET    /customers/{id}               # Get customer by ID

# Product Service Examples
GET    /products                     # List products
POST   /products                     # Create product
GET    /products/{id}                # Get product by ID
DELETE /products/{id}                # Delete product

# Payment Service Examples
POST   /payments/invoices            # Create invoice
POST   /payments/invoices/{id}/pay   # Pay invoice
POST   /payments/invoices/{id}/cancel # Cancel invoice

# Order Service Examples
GET    /orders                       # List orders
POST   /orders                       # Create order
GET    /orders/{id}                  # Get order
POST   /orders/{id}/items            # Add item to order
POST   /orders/{id}/complete         # Complete order
POST   /orders/{id}/cancel           # Cancel order
```

### gRPC Endpoints

Services are also exposed via gRPC on port 8085:

```bash
# Example: Customer service
grpcurl -plaintext \
  -d '{"name":"John","email":"john@example.com","sms_number":"+1234567890"}' \
  localhost:8085 customerspb.CustomersService/RegisterCustomer
```

## Code Generation

### Using buf (Recommended)

buf simplifies protobuf compilation with versioning and plugin management.

#### Installation

```bash
# Install buf (macOS)
brew install bufbuild/buf/buf

# Or visit https://docs.buf.build/installation
```

#### Generate Code

```bash
# Generate code for a specific module
cd customer
buf generate

# Or for all modules at once (from root)
for module in customer product payment order; do
  (cd "$module" && buf generate)
done
```

Generated files are placed in module directories (e.g., `customer/customerspb/`)

#### buf Configuration Files

- `buf.yaml` - Workspace and linting config
- `buf.gen.yaml` - Code generation rules
- `buf.lock` - Dependency lock file

### Manual protoc (If Not Using buf)

```bash
# Install plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# Generate (example for customer service)
cd customer
protoc --go_out=. --go-grpc_out=. \
  --grpc-gateway_out=. --openapiv2_out=docs \
  customerspb/customer.proto
```

## Development

### Local Environment Setup

1. **Start containers**:

   ```bash
   docker-compose up -d
   ```

2. **Generate protobuf code** (if modified):

   ```bash
   cd customer && buf generate
   cd ../product && buf generate
   # ... repeat for other modules
   ```

3. **Run application**:

   ```bash
   go run ./cmd/shopping
   ```

4. **Access Swagger UI**:
   ```
   http://localhost:8080/swagger-ui/
   ```

### Common Development Tasks

#### Viewing Logs

```bash
# Application logs
docker-compose logs shopping

# Database logs
docker-compose logs postgres

# NATS logs
docker-compose logs nats

# Follow logs in real-time
docker-compose logs -f shopping
```

#### Database Access

```bash
# Direct database access
docker-compose exec postgres psql -U postgres -d shopping

# Using pgAdmin UI
# URL: http://localhost:83
# Email: giangle2198@gmail.com
# Password: 1
```

#### NATS JetStream Inspection

```bash
# Check NATS streams
nc localhost 4222  # Direct connection test

# NATS monitor UI (if enabled)
# http://localhost:8222/
```

## Testing

### Run Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific package tests
go test -v ./customer/...

# Run with coverage
go test -cover ./...
```

### Protobuf Linting

```bash
# Lint protobuf files (using buf)
cd customer && buf lint
```

## Contributing

### Code Changes

1. **Create feature branch**: `git checkout -b feature/my-feature`
2. **Make changes**: Modify code as needed
3. **Update protobuf** (if needed):
   - Modify `.proto` files
   - Run `buf generate`
4. **Run tests**: `go test ./...`
5. **Commit**: `git commit -m "feat: description"`
6. **Push & create PR**: `git push origin feature/my-feature`

### Pull Request Guidelines

- Include descriptive commit messages
- Update protobuf files if service contracts change
- Ensure all tests pass
- Follow Go code style conventions (go fmt)
- Update documentation for API changes

## Useful Commands

### Docker Management

```bash
# Start all services
docker-compose up -d

# Stop all services
docker-compose down

# View running containers
docker-compose ps

# View logs for all services
docker-compose logs -f

# Stop and remove volumes (full cleanup)
docker-compose down -v
```

### Building & Running

```bash
# Build binary
go build -o shopping ./cmd/shopping

# Run binary
./shopping

# Run with output
go run ./cmd/shopping

# Build for specific OS
GOOS=linux GOARCH=amd64 go build -o shopping ./cmd/shopping
```

### Database Operations

```bash
# Connect to database
docker-compose exec postgres psql -U postgres -d shopping

# List tables
\dt

# List all databases
\l

# Describe table
\d table_name

# Exit psql
\q
```

### Protobuf Operations

```bash
# Lint all .proto files in module
cd customer && buf lint

# Format protobuf files
cd customer && buf format -w

# Generate code (creates pb.go files)
cd customer && buf generate

# Check for breaking changes
buf breaking --against .git#branch=main
```

## Troubleshooting

### Port Already in Use

```bash
# Find and kill process using port
lsof -i :8080
kill -9 <PID>

# Or change ports in docker-compose.yaml
```

### Database Connection Issues

```bash
# Verify PostgreSQL is running
docker-compose ps postgres

# Check connection string in code matches docker-compose
# Should be: postgresql://shopping_user:shopping_pass@postgres:5432/shopping
```

### NATS Connection Issues

```bash
# Verify NATS is running
docker-compose ps nats

# Check NATS configuration
docker-compose exec nats cat /config/config_nats.conf
```

### Protobuf Generation Issues

```bash
# Clean generated files
find . -name "*.pb.go" -delete

# Regenerate
buf generate
```

---

## 📄 License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

---

For questions or support, please open an issue on the repository.
