.PHONY: all test lint build clean fmt verify

all: verify build

fmt:
	gofmt -s -w .
	gofumpt -w .

lint:
	golangci-lint run ./...

test:
	go test -v -race -cover ./...

build:
	go build -o bin/platform ./cmd/platform/main.go

verify: fmt lint test

clean:
	rm -rf bin/
