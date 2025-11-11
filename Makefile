# Makefile for mcpgo

.PHONY: all build run test lint clean

# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_RUN=$(GO_CMD) run
GO_TEST=$(GO_CMD) test
GO_FMT=$(GO_CMD) fmt
GO_LINT=golangci-lint run

# Binary name
BINARY_NAME=mcpgo

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	$(GO_BUILD) -o bin/$(BINARY_NAME) ./cmd/mcpgo

run:
	@echo "Running $(BINARY_NAME)..."
	$(GO_RUN) ./cmd/mcpgo

test:
	@echo "Running tests..."
	$(GO_TEST) -v ./...

lint:
	@echo "Linting code..."
	$(GO_FMT) ./...
	# $(GO_LINT) ./... # Uncomment when golangci-lint is configured

clean:
	@echo "Cleaning up..."
	@rm -f bin/$(BINARY_NAME)
