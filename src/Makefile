# Makefile for Technical Debt Record (TDR) Generator Tool

# Variables
BINARY_NAME=generate-td
GO_FILES=generate-td.go

# Default target
.PHONY: all
all: build

# Initialize Go module if not already initialized
.PHONY: init
init:
	@echo "Initializing Go module..."
	go mod init technical_debt_generator || echo "Go module already initialized."
	@echo "Go module initialization complete."

# Download and add dependencies
.PHONY: deps
deps: init
	@echo "Adding required dependencies..."
	go get github.com/phpdave11/gofpdf
	go get github.com/xuri/excelize/v2
	@echo "Dependencies added successfully."

# Tidy up the Go module (optional)
.PHONY: tidy
tidy: deps
	@echo "Tidying up Go modules..."
	go mod tidy
	@echo "Go modules tidied successfully."

# Build the executable
.PHONY: build
build: deps
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(GO_FILES)
	@echo "Build completed. Executable: $(BINARY_NAME)"

# Clean the executable and Go module files
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	rm -f go.mod go.sum
	@echo "Clean completed."

# Run the tool (optional)
.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

# Test the application
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...
	@echo "Tests completed."
