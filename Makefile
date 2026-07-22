.PHONY: all test race lint build clean fmt verify

all: verify build

fmt:
	gofmt -s -w .
	gofumpt -extra -w .

lint:
	golangci-lint run ./...

test:
	go test -v -cover ./...

race:
	go test -v -race -cover ./...

build:
	go build -o bin/platform ./cmd/platform/main.go

verify:
	go fmt ./...
	gofumpt -extra -w .
	golangci-lint run ./...
	go test -race ./...
	go test ./...
	npm run docs:build

clean:
	rm -rf bin/
