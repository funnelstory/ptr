.PHONY: test lint

test:
	go test ./...

lint:
	golangci-lint run

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out