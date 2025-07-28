# Makefile for lazypoetry

# Variables
BINARY_NAME=lazypoetry
VERSION=v0.1.0

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	go build -o $(BINARY_NAME) .

# Run the application
.PHONY: run
run: build
	./$(BINARY_NAME)

# Clean build artifacts
.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)

# Run tests
.PHONY: test
test:
	go test -v ./...

# Run tests with coverage
.PHONY: coverage
coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Install dependencies
.PHONY: deps
deps:
	go mod download
	go mod tidy

# Lint the code (requires golangci-lint)
.PHONY: lint
lint:
	golangci-lint run

# Format the code
.PHONY: fmt
fmt:
	go fmt ./...

# Run all CI checks
.PHONY: ci
ci: fmt lint test

# Release management
.PHONY: release
release:
	@echo "Usage: make release VERSION=v1.0.0"
	@if [ -z "$(VERSION)" ]; then echo "ERROR: VERSION is required"; exit 1; fi
	./scripts/release.sh $(VERSION)

# Test goreleaser configuration
.PHONY: release-test
release-test:
	goreleaser release --snapshot --clean

# Install goreleaser
.PHONY: install-goreleaser
install-goreleaser:
	go install github.com/goreleaser/goreleaser@latest

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build              - Build the application"
	@echo "  run                - Build and run the application"
	@echo "  clean              - Clean build artifacts"
	@echo "  test               - Run tests"
	@echo "  coverage           - Run tests with coverage report"
	@echo "  deps               - Install and tidy dependencies"
	@echo "  lint               - Run linter (requires golangci-lint)"
	@echo "  fmt                - Format code"
	@echo "  ci                 - Run all CI checks (fmt, lint, test)"
	@echo "  release VERSION=X  - Create and push a release tag"
	@echo "  release-test       - Test goreleaser configuration"
	@echo "  install-goreleaser - Install goreleaser"
	@echo "  help               - Show this help message"