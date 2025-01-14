# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=backend-in-go
BINARY_UNIX=$(BINARY_NAME)_unix

# All target
all: test build

# Build the project
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run tests
test: 
	$(GOTEST) -v ./...

# Clean the build directory
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Install dependencies
deps:
	$(GOGET) -u ./...

# Cross compilation for Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

.PHONY: all build clean test deps build-linux