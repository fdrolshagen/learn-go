build:
	@go build -o bin/goserve cmd/app/main.go

run: build
	@bin/goserve

test:
	@go test ./...