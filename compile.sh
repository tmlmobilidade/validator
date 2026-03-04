#!/bin/bash

# Modify the version number in the validator/services/cli.go file if a version number is provided
if [ "$1" ]; then
    if [[ "$OSTYPE" == "darwin"* ]]; then
        sed -i '' 's/const version = "0.0.0"/const version = "'"$1"'"/' validator/services/cli.go
    else
        sed -i 's/const version = "0.0.0"/const version = "'"$1"'"/' validator/services/cli.go
    fi
fi

# Check for folder bin and create it if it doesn't exist
if [ ! -d "bin" ]; then
    mkdir bin
fi

cd validator

# Compile the validator for linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build go build -o ../bin/validator-linux-amd64 ./main.go

# Compile the validator for linux arm64
GOOS=linux GOARCH=arm64 go build -o ../bin/validator-linux-arm64 ./main.go

# Compile the validator for darwin x86_64
GOOS=darwin GOARCH=amd64 go build -o ../bin/validator-darwin-amd64 ./main.go

# Compile the validator for darwin arm64
GOOS=darwin GOARCH=arm64 go build -o ../bin/validator-darwin-arm64 ./main.go

# Compile the validator for windows x86_64
GOOS=windows GOARCH=amd64 go build -o ../bin/validator.exe ./main.go

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