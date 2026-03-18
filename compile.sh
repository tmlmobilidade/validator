#!/bin/bash

# LDFLAGS for version injection at build time (from workflow or manual build)
LDFLAGS=""
if [ -n "$1" ]; then
    LDFLAGS="-ldflags \"-X main/config.Version=$1\""
fi

# Check for folder bin and create it if it doesn't exist
if [ ! -d "bin" ]; then
    mkdir bin
fi

cd validator

# Compile the validator for linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/validator-linux-amd64 ./main.go

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