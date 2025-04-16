#!/bin/bash

# Check for folder bin and create it if it doesn't exist
if [ ! -d "bin" ]; then
    mkdir bin
fi



# Compile the validator for linux
GOOS=linux GOARCH=amd64 go build -o bin/validator-linux-amd64 validator/main.go

# Compile the validator for linux arm64
GOOS=linux GOARCH=arm64 go build -o bin/validator-linux-arm64 validator/main.go

# Compile the validator for darwin x86_64
GOOS=darwin GOARCH=amd64 go build -o bin/validator-darwin-amd64 validator/main.go

# Compile the validator for darwin arm64
GOOS=darwin GOARCH=arm64 go build -o bin/validator-darwin-arm64 validator/main.go

# Compile the validator for windows x86_64
GOOS=windows GOARCH=amd64 go build -o bin/validator.exe validator/main.go


# Check if the compilation was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to compile the validator"
    exit 1
fi