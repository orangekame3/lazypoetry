# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`lazypoetry` is a simple CLI tool written in Go that generates random poems and allows browsing a built-in poetry library. It serves as a demonstration of the [vercheck](https://github.com/orangekame3/vercheck) library for automatic update notifications in CLI tools.

## Architecture

```
lazypoetry/
├── cmd/                    # CLI commands (Cobra-based)
│   ├── root.go            # Root command with vercheck integration
│   ├── generate.go        # Generate random poems
│   ├── list.go           # List library poems  
│   ├── show.go           # Show specific poems
│   └── version.go        # Version command
├── internal/poetry/       # Core poetry functionality
│   ├── library.go        # Embedded JSON poem library
│   ├── generate.go       # Haiku and limerick generators
│   └── poems.json        # Embedded poem data
└── main.go               # Entry point

```

## Development Commands

```bash
# Build and run
make build        # Build the binary
make run         # Build and run the application
go build .       # Alternative build command

# Testing and quality
make test        # Run tests
make coverage    # Run tests with coverage report
make fmt         # Format code
make lint        # Run linter (requires golangci-lint)
make ci          # Run all CI checks (fmt, lint, test)

# Dependencies
make deps        # Install and tidy dependencies
go mod tidy      # Tidy dependencies

# Release management
make install-goreleaser     # Install goreleaser
make release-test          # Test goreleaser configuration
make release VERSION=v1.0.0 # Create and push release tag
./scripts/release.sh v1.0.0 # Direct release script usage

# Usage examples
./lazypoetry generate                    # Generate haiku
./lazypoetry generate --form limerick    # Generate limerick
./lazypoetry list                        # List all poems
./lazypoetry show <id>                   # Show specific poem
./lazypoetry --no-update <command>       # Disable update check
```

## Key Dependencies

- `github.com/spf13/cobra` - CLI framework
- `github.com/orangekame3/vercheck` - Update notification system

## vercheck Integration

The tool demonstrates vercheck integration in `cmd/root.go`:
- Update check runs in `PersistentPreRunE` for all commands
- Uses 5-second timeout context
- Silent mode to avoid disrupting CLI experience
- Can be disabled with `--no-update` flag

## Release Process

The project uses automated releases with GitHub Actions and goreleaser:

1. **CI/CD Pipeline**: `.github/workflows/ci.yml` runs tests on push/PR
2. **Release Pipeline**: `.github/workflows/release.yml` triggers on git tags
3. **Multi-platform Builds**: Supports Linux, macOS, Windows (amd64, arm64, etc.)
4. **Version Management**: Version injected at build time via ldflags
5. **Homebrew Integration**: Can update homebrew tap automatically

### Release Steps:
```bash
# 1. Ensure main branch is clean and up-to-date
# 2. Run release script
make release VERSION=v1.0.0

# 3. GitHub Actions will:
#    - Build binaries for all platforms
#    - Create GitHub release with changelog
#    - Update homebrew tap (if configured)
```

## Code Conventions

- Use standard Go formatting (`gofmt`)
- Follow Go naming conventions
- Embed static assets using `//go:embed`
- Handle errors gracefully without disrupting user experience
- Use Cobra for consistent CLI patterns
- Version is injected at build time, use `var Version = "dev"` pattern