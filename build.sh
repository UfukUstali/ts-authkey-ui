#!/bin/sh

echo "Building Tailscale Auth Key UI for multiple platforms..."

mkdir -p dist

echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o dist/ts-authkey-ui-windows-amd64.exe

echo "Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o dist/ts-authkey-ui-macos-amd64

echo "Building for macOS (arm64)..."
GOOS=darwin GOARCH=arm64 go build -o dist/ts-authkey-ui-macos-arm64

echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o dist/ts-authkey-ui-linux-amd64

echo "Building for Linux (arm64)..."
GOOS=linux GOARCH=arm64 go build -o dist/ts-authkey-ui-linux-arm64

echo "Build complete! Binaries are in the dist/ directory."
ls -la dist/
