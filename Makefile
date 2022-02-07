lint:
	golangci-lint run

test:
	go test -v -race -cover ./...

PHONY: lint test
