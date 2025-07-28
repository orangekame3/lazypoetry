#!/bin/bash

# Release script for lazypoetry
# Usage: ./scripts/release.sh v1.0.0

set -e

if [ $# -eq 0 ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 v1.0.0"
    exit 1
fi

VERSION=$1

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Version must be in format v1.2.3"
    exit 1
fi

echo "🚀 Preparing release $VERSION"

# Check if we're on main branch
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ]; then
    echo "Error: Must be on main branch to release. Currently on: $CURRENT_BRANCH"
    exit 1
fi

# Check if working directory is clean
if [ -n "$(git status --porcelain)" ]; then
    echo "Error: Working directory is not clean. Please commit or stash changes."
    git status --short
    exit 1
fi

# Pull latest changes
echo "📥 Pulling latest changes..."
git pull origin main

# Run tests
echo "🧪 Running tests..."
go test ./...

# Build to ensure everything compiles
echo "🔨 Building..."
go build .

# Update version in documentation if needed
echo "📝 Version: $VERSION"

# Create and push tag
echo "🏷️  Creating tag $VERSION..."
git tag -a "$VERSION" -m "Release $VERSION"

echo "📤 Pushing tag to GitHub..."
git push origin "$VERSION"

echo "✅ Release $VERSION initiated!"
echo ""
echo "🎯 Next steps:"
echo "   1. GitHub Actions will automatically create the release"
echo "   2. Binaries will be built for multiple platforms"
echo "   3. Release notes will be generated from commits"
echo "   4. Homebrew formula will be updated (if tap exists)"
echo ""
echo "🔗 Check the release status at:"
echo "   https://github.com/orangekame3/lazypoetry/actions"