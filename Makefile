build:
	@go build -o bin/clean-architecture cmd/api/main.go

test:
	@go test -v ./...

run:
	@go run cmd/api/main.go