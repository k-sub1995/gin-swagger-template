# Gin Swagger Template

This is a template project for building REST APIs using the Gin framework in Go lang.

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/k-sub1995/gin-swagger-template.git
cd gin-swagger-template
```

### 2. Run the application

```bash
docker compose up
```

## Usage

Once the container is running, you can access the following endpoints:

- **API Server**: `http://localhost:8080`
- **Ping Endpoint**: `http://localhost:8080/ping`

## API Documentation

API documentation is automatically generated from the source code comments.

- **Swagger UI**: `http://localhost:8080/swagger/index.html`
- **ReDoc**: `http://localhost:8080/redoc`

## Directory Strategy

```bash
.
├── docs/           # Contains Swagger API documentation
├── dto/            # Data Transfer Objects (DTOs)
├── handlers/       # HTTP handlers layer(controllers)
├── infra/          # Infrastructure layer
├── middlewares/    # Gin middleware functions
├── migrations/     # Database migration scripts
├── models/         # Domain models
├── repositories/   # Provides an abstraction layer
├── services/       # Business logic layer
```
