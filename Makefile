# Description: Makefile for building and running the API
# define go as the language
GO = go

build:
	@echo "Building..."
	$(GO) build -o bin/ ./cmd/api

run-build: build
	@echo "Running..."
	@./bin/api

run:
	@echo "Running..."
	$(GO) run ./cmd/api
	
clean:
	@echo "Cleaning..."
	rm -rf bin/