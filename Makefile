.PHONY: build run

build:
	@echo "Building..."
	@go build -o bin/app ./cmd/app/main.go

run: build
	@echo "Running..."
	@./bin/app

test:
	@go test ./...

seed:
	@go run ./cmd/seed/seed.go
