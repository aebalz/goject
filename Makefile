.PHONY: all tidy swagger run dev

all: tidy swagger run

tidy:
	@echo "Running go mod tidy..."
	go mod tidy

swagger:
	@echo "Generating Swagger docs..."
	swag init -g cmd/myapp/main.go

run:
	@echo "Running the application..."
	go run cmd/myapp/main.go

dev:
	@echo "Watching for changes..."
	dogo