# Project variables
MODULE_NAME := piscine
BUILD_DIR := sudoko
EXECUTABLE := $(BUILD_DIR)/$(MODULE_NAME)

# Files
MAIN_FILE := main/main.go
UTILS := utils.go

# Go tools
GO := go
GOFMT := gofmt
GOTEST := go test
GOLINT := golint

# Default target
all: build

# Build the executable
build:
	@echo "Building the project..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(EXECUTABLE) $(MAIN_FILE)
	@echo "Build complete: $(EXECUTABLE)"

# Run the application
run: build
	@echo "Running the application..."
	./$(EXECUTABLE) "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79"

# Format code
fmt:
	@echo "Formatting Go code..."
	$(GOFMT) -w .
	@echo "Code formatted."

# Lint code
lint:
	@echo "Running linter..."
	$(GOLINT) ./...

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

# Display help
help:
	@echo "Makefile commands:"
	@echo "  all      - Build the project (default)"
	@echo "  build    - Build the project"
	@echo "  run      - Build and run the project with sample input"
	@echo "  fmt      - Format the Go code"
	@echo "  lint     - Run Go lint"
	@echo "  test     - Run tests"
	@echo "  clean    - Remove build artifacts"
	@echo "  help     - Display this help message"

