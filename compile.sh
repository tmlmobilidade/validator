#!/bin/bash

# Version for ldflags injection (from workflow or manual build)
VERSION="${1:-}"

# Check for folder bin and create it if it doesn't exist
if [ ! -d "bin" ]; then
    mkdir bin
fi

cd validator

build_binary() {
    local goos=$1 goarch=$2 output=$3
    if [ -n "$VERSION" ]; then
        CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -ldflags "-X main/services.Version=$VERSION" -o "$output" ./main.go
    else
        CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -o "$output" ./main.go
    fi
}

# Compile the validator for linux arm64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ../bin/validator-linux-arm64 ./main.go

# Compile the validator for darwin x86_64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../bin/validator-darwin-amd64 ./main.go

# Compile the validator for darwin arm64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ../bin/validator-darwin-arm64 ./main.go

# Compile the validator for windows x86_64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../bin/validator.exe ./main.go

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