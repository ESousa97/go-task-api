<div align="center">
  <h1>Go Task API</h1>
  <p>Minimalist and high-performance task management API built strictly with the Go standard library.</p>

  <img src="assets/github-go.png" alt="Go Task API Banner" width="600px">

  <br>

[![CI/CD](https://github.com/ESousa97/go-task-api/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/ESousa97/go-task-api/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ESousa97/go-task-api)](https://goreportcard.com/report/github.com/ESousa97/go-task-api)
[![Go Reference](https://pkg.go.dev/badge/github.com/ESousa97/go-task-api.svg)](https://pkg.go.dev/github.com/ESousa97/go-task-api)
[![License: MIT](https://img.shields.io/github/license/ESousa97/go-task-api)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ESousa97/go-task-api)](https://github.com/ESousa97/go-task-api)
[![Last Commit](https://img.shields.io/github/last-commit/ESousa97/go-task-api)](https://github.com/ESousa97/go-task-api/commits/main)

</div>

---

**Go Task API** is a high-performance RESTful service designed to manage task lifecycles with zero external framework dependencies. It implements architectural patterns like **Repository** and **Dependency Injection** to ensure modularity, testability, and a low blast radius (Failure domain isolation) in multi-provider environments (Memory and PostgreSQL).

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

> View full documentation at [pkg.go.dev/github.com/ESousa97/go-task-api](https://pkg.go.dev/github.com/ESousa97/go-task-api).

## Configuration

| Variable | Description | Type | Default |
| --- | --- | --- | --- |
| `APP_PORT` | Port where the server listens | int | 8080 |
| `DB_URL` | PostgreSQL connection string | string | postgres://postgres:password@localhost:5433/taskdb?sslmode=disable |

## Roadmap

Acompanhe as etapas de evolução do projeto:

- [x] **Phase 1: Foundation** — In-Memory persistence core.
- [x] **Phase 2: Persistence** — PostgreSQL integration with Docker.
- [x] **Phase 3: Security** — Middleware implementation (Auth & Recovery).
- [x] **Phase 4: Patterns** — Repository Pattern & Dependency Injection.
- [x] **Phase 5: Governance** — CI/CD, Professional Documentation and Badges.
- [ ] **Phase 6: Observability** — Structured logging and health checks.

## Contributing

Interessado em colaborar? Veja nosso [CONTRIBUTING.md](CONTRIBUTING.md) para padrões de código e processo de PR.

## License

Este projeto está licenciado sob a **MIT License** — veja o arquivo [LICENSE](LICENSE) para detalhes.

<div align="center">

## Autor

**Enoque Sousa**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/enoque-sousa-bb89aa168/)
[![GitHub](https://img.shields.io/badge/GitHub-100000?style=flat&logo=github&logoColor=white)](https://github.com/ESousa97)
[![Portfolio](https://img.shields.io/badge/Portfolio-FF5722?style=flat&logo=target&logoColor=white)](https://enoquesousa.vercel.app)

**[⬆ Voltar ao topo](#go-task-api)**

Feito com ❤️ por [Enoque Sousa](https://github.com/ESousa97)

**Status do Projeto:** Ativo — Em constante atualização

</div>
