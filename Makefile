build:
	go build -o mydocs ./cmd/mydocs

run:
	go run ./cmd/mydocs

test:
	go test ./...

bench:
	go test -bench=. ./...

lint:
	golangci-lint run

fmt:
	gofmt -s -w .

vet:
	go vet ./...

clean:
	rm -f mydocs

install:
	go install ./cmd/mydocs

help:
	@echo "Available commands: build, run, test, bench, lint, fmt, vet, clean, install"

.PHONY: build run test bench lint fmt vet clean install help
