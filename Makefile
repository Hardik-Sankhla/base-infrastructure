.PHONY: all test lint build clean

all: lint test build

test:
	go test -v -race -cover ./...

lint:
	golangci-lint run ./...

build:
	go build -o bin/platform ./cmd/platform/main.go

clean:
	rm -rf bin/
