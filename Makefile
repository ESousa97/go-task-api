# Binary name
BINARY_NAME=apigotask

# Build targets
.PHONY: all build run test clean help

all: build

build: ## Build the application
	@echo "Building binary..."
	go build -o bin/$(BINARY_NAME) ./cmd/server

run: build ## Run the application
	@echo "Running application..."
	./bin/$(BINARY_NAME)

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	rm -f server.exe

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
