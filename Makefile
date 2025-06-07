default: run

fmt:
	@go fmt ./...

run: fmt
	@go run main.go