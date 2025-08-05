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

install-global:
	sudo cp mydocs /usr/local/bin/mydocs
	sudo chmod +x /usr/local/bin/mydocs

uninstall:
	rm -f $(GOPATH)/bin/mydocs /usr/local/bin/mydocs

help:
	@echo "Available commands:"
	@echo "  build          Build the binary"
	@echo "  run            Run the application"
	@echo "  test           Run tests"
	@echo "  bench          Run benchmarks"
	@echo "  lint           Run linter"
	@echo "  fmt            Format code"
	@echo "  vet            Run go vet"
	@echo "  clean          Remove built binary"
	@echo "  install        Install to GOPATH/bin (requires Go env)"
	@echo "  install-global Install to /usr/local/bin (system-wide)"
	@echo "  uninstall      Remove installed binary"

.PHONY: build run test bench lint fmt vet clean install install-global uninstall help
