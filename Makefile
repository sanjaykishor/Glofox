.PHONY: build run clean test lint docker help

BINARY_NAME=glofox
BUILD_DIR=./bin
VERSION=0.1.0
MAIN_PATH=./cmd/glofox
GOFLAGS=-ldflags "-X main.Version=$(VERSION)"

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete!"

run:
	@echo "Starting Glofox API server..."
	@go run $(MAIN_PATH)

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete!"

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Running linter..."
	@golangci-lint run ./...

docker:
	@echo "Building Docker image..."
	@docker build -t glofox:$(VERSION) .
	@echo "Docker image built!"

build-all: build-linux build-mac build-windows

build-linux:
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux $(MAIN_PATH)

build-mac:
	@echo "Building for macOS..."
	@GOOS=darwin GOARCH=amd64 go build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-mac $(MAIN_PATH)

build-windows:
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows.exe $(MAIN_PATH)

help:
	@echo "Glofox - Fitness Studio Management API"
	@echo ""
	@echo "Usage:"
	@echo "  make build             - Build the application"
	@echo "  make run               - Run the application"
	@echo "  make clean             - Clean build artifacts"
	@echo "  make test              - Run tests"
	@echo "  make lint              - Run linter"
	@echo "  make docker            - Build Docker image"
	@echo "  make build-all         - Build for all platforms"
	@echo "  make help              - Show this help"
