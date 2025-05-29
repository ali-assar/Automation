# Makefile for Automation project

# Variables
BINARY_NAME = main
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOMOD = $(GOCMD) mod
PROJECT_DIR = /root/04Automation/backend"
CMD_DIR = $(PROJECT_DIR)/cmd/mainservice
OUTPUT_DIR = $(PROJECT_DIR)/bin

# Build the binary
build:
	@mkdir -p $(OUTPUT_DIR)
	@cd $(CMD_DIR) && CGO_ENABLED=0 GOOS=linux $(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) .

# Clean the binary and output directory
clean:
	@$(GOCLEAN)
	@rm -rf $(OUTPUT_DIR)

# Run tests
test:
	@$(GOTEST) -v ./...

# Download dependencies
deps:
	@cd $(PROJECT_DIR) && $(GOMOD) tidy
	@cd $(PROJECT_DIR) && $(GOMOD) download

# Default target
all: deps build

.PHONY: build clean test deps all