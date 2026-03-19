#!/bin/bash

# Version for ldflags injection (from workflow or manual build)
# -ldflags "-X main/services.Version=$VERSION" replaces the default "0.0.0" at compile time
VERSION="${1:-}"

if [ -z "$VERSION" ]; then
  echo "WARNING: No version passed - binary will show v0.0.0. Usage: compile.sh 20250319.1500.00"
fi

# Check for folder bin and create it if it doesn't exist
if [ ! -d "bin" ]; then
    mkdir bin
fi

cd validator

build_binary() {
    local goos=$1 goarch=$2 output=$3
    if [ -n "$VERSION" ]; then
        echo "  Building $output with version=$VERSION"
        CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -ldflags "-X main/services.Version=$VERSION" -o "$output" ./main.go
    else
        CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -o "$output" ./main.go
    fi
}

# Compile the validator for each platform
build_binary linux amd64 ../bin/validator-linux-amd64
build_binary linux arm64 ../bin/validator-linux-arm64
build_binary darwin amd64 ../bin/validator-darwin-amd64
build_binary darwin arm64 ../bin/validator-darwin-arm64
build_binary windows amd64 ../bin/validator.exe

# Allow all users to execute the validator
chmod +x ../bin/validator-linux-amd64
chmod +x ../bin/validator-linux-arm64
chmod +x ../bin/validator-darwin-amd64
chmod +x ../bin/validator-darwin-arm64
chmod +x ../bin/validator.exe

# Check if the compilation was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to compile the validator"
    exit 1
fi

cd ..