# Contributing to Go Task API

First off, thank you for considering contributing to Go Task API! It's people like you who make the open-source community such an amazing place to learn, inspire, and create.

## Development Setup

### Prerequisites
- **Go 1.22+**
- **Docker & Docker Compose**
- **Make**

### Local Environment
1. Fork the repository and clone it locally.
2. Run `docker-compose up -d` to start the PostgreSQL database.
3. Use `make build` to ensure everything compiles.
4. Run `make test` to verify the codebase.

## Formatting and Style
We follow the standard Go formatting rules. Please run `gofmt -s` before submitting any PR. We also value clean, documented code using godoc comments.

## Pull Request Process
1. Create a new branch for your feature or bugfix: `git checkout -b feature/my-new-feature`.
2. Ensure all tests pass: `make test`.
3. Fill out the PR template with a clear description of the changes.
4. Once approved, your PR will be merged into the `main` branch.

## Commit Conventions
We follow [Conventional Commits](https://www.conventionalcommits.org/):
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `style:` for formatting
- `refactor:` for code restructuring
