build:
	@echo "Building..."
	@go build -o bin/app ./cmd/...

run: build
	@echo "Running..."
	@./bin/app
