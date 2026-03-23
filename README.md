# Go Task API

> Minimalist and high-performance task management API built strictly with the Go standard library.

[![CI/CD](https://github.com/ESousa97/go-task-api/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/ESousa97/go-task-api/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ESousa97/go-task-api)](https://goreportcard.com/report/github.com/ESousa97/go-task-api)
[![Go Reference](https://pkg.go.dev/badge/github.com/ESousa97/go-task-api.svg)](https://pkg.go.dev/github.com/ESousa97/go-task-api)
[![License: MIT](https://img.shields.io/github/license/ESousa97/go-task-api)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ESousa97/go-task-api)](https://github.com/ESousa97/go-task-api)
[![Last Commit](https://img.shields.io/github/last-commit/ESousa97/go-task-api)](https://github.com/ESousa97/go-task-api)

---

**Go Task API** is a high-performance RESTful service designed to manage task lifecycles with zero external framework dependencies. It implements architectural patterns like Repository and Dependency Injection to ensure modularity, testability, and a low blast radius in multi-provider environments (Memory and PostgreSQL).

## Project Showcase

### API Interaction (curl)

**Create a Task:**
```bash
curl -i -X POST \
  -H "X-API-Key: secret-key" \
  -H "Content-Type: application/json" \
  -d '{"title": "Documentation Refactor", "description": "Implement high-quality docs", "status": "doing"}' \
  http://localhost:8080/tasks
```

**List All Tasks:**
```bash
curl -i -H "X-API-Key: secret-key" http://localhost:8080/tasks
```

## Tech Stack

| Technology | Role |
| --- | --- |
| **Go 1.22+** | Core runtime and standard library (`net/http`) |
| **PostgreSQL** | Persistent data storage |
| **Docker** | Containerization and infrastructure orchestration |
| **GitHub Actions** | Automated CI/CD pipeline |

## Prerequisites

- **Go >= 1.22**
- **Docker & Docker Compose** (for PostgreSQL)
- **Make** (optional, but recommended)

## Installation and Usage

### From Source

```bash
# 1. Clone the repository
git clone https://github.com/ESousa97/go-task-api.git
cd go-task-api

# 2. Start the database
docker-compose up -d

# 3. Build and Run
make run
```

## Makefile Targets

| Target | Description |
| --- | --- |
| `make build` | Compiles the project into `bin/go-task-api` |
| `make run` | Builds and executes the server |
| `make test` | Runs all unit and integration tests |
| `make clean` | Removes build artifacts and binaries |
| `make help` | Displays all available commands |

## Architecture

The project follows a modular structure separated by logical responsibilities, ensuring that business rules remain decoupled from infrastructure details.

```mermaid
graph TD
    Client[HTTP Client] --> Mux[net.ServeMux]
    Mux --> Handler[internal/handler]
    Handler --> RepoInterface[internal/repository Interface]
    RepoInterface --> Postgres[internal/repository/PostgreSQL]
    RepoInterface --> Memory[internal/repository/Memory]
    
    subgraph "Domain Layer"
    Task[internal/domain/Task]
    end
```

## API Reference

| Method | Endpoint | Description | Auth Required |
| --- | --- | --- | --- |
| `GET` | `/tasks` | List all available tasks | Yes |
| `POST` | `/tasks` | Create a new task entry | Yes |
| `GET` | `/tasks/{id}` | Retrieve a specific task | Yes |
| `PUT` | `/tasks/{id}` | Update an existing task | Yes |
| `DELETE` | `/tasks/{id}` | Remove a task | Yes |

> View full documentation at [pkg.go.dev](https://pkg.go.dev/github.com/ESousa97/go-task-api).

## Configuration

| Variable | Description | Type | Default |
| --- | --- | --- | --- |
| `APP_PORT` | Port where the server listens | int | 8080 |
| `DB_URL` | PostgreSQL connection string | string | postgres://postgres:password@localhost:5433/taskdb?sslmode=disable |

## Roadmap

- [x] **Phase 1: Basic CRUD** — In-Memory persistence core.
- [x] **Phase 2: Persistence** — PostgreSQL integration with Docker.
- [x] **Phase 3: Security** — Middleware implementation (Auth & Recovery).
- [x] **Phase 4: Patterns** — Repository Pattern & Dependency Injection.
- [x] **Phase 5: Automation** — GitHub Actions & Makefile integration.
- [ ] **Phase 6: Observability** — Structured logging and health checks.

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines.

## License

Distributed under the **MIT License**. See `LICENSE` for more information.

## Author

**Enoque Sousa**
- [Portfolio](https://enoquesousa.vercel.app)
- [GitHub](https://github.com/ESousa97)
- [LinkedIn](https://www.linkedin.com/in/enoque-sousa-bb89aa168/)

---
[⬆ return to top](#go-task-api)
