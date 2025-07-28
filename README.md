# lazypoetry

A simple CLI tool that generates random poems and lets you browse a small collection of poetry. Perfect for when you need a quick dose of creativity or inspiration!

This project serves as a demonstration of the [vercheck](https://github.com/orangekame3/vercheck) library, which helps CLI tools check for newer versions on GitHub.

## Features

- 🎭 Generate random haiku and limerick poems
- 📚 Browse a built-in poetry library
- 🔔 Automatic update notifications via vercheck
- ⚡ Lightweight and fast

## Installation

```bash
go install github.com/orangekame3/lazypoetry@latest
```

## Usage

### Generate a random poem

```bash
# Generate a haiku (default)
lazypoetry generate

# Generate a limerick
lazypoetry generate --form limerick
```

### Browse the poetry library

```bash
# List all poems
lazypoetry list

# Show a specific poem by ID
lazypoetry show 1
```

### Other commands

```bash
# Check version (and get update notifications)
lazypoetry version

# Disable update checks
lazypoetry --no-update generate
```

## Development

### Prerequisites

- Go 1.21 or higher
- [goreleaser](https://goreleaser.com/) (for releases)

### Building

```bash
# Build the project
make build

# Run the application
make run

# Run tests
make test

# Format and lint
make fmt
make lint

# Run all CI checks
make ci
```

### Releasing

This project uses automated releases via GitHub Actions and goreleaser:

```bash
# Install goreleaser
make install-goreleaser

# Test release configuration
make release-test

# Create a new release (pushes tag and triggers GitHub Actions)
make release VERSION=v1.0.0
```

The release process:
1. Creates and pushes a git tag
2. GitHub Actions automatically builds binaries for multiple platforms
3. Creates a GitHub release with auto-generated changelog
4. Updates Homebrew tap (if configured)

## vercheck Integration

This project demonstrates how to integrate [vercheck](https://github.com/orangekame3/vercheck) into a CLI tool:

```go
vercheck.CheckWithContext(ctx, vercheck.Options{
    CurrentVersion: Version,
    RepoOwner:      repoOwner,
    RepoName:       repoName,
    Silent:         true,
})
```

The update check runs automatically on every command (unless `--no-update` is specified) and will notify users when a newer version is available.

## License

MIT License