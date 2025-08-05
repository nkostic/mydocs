# 🧰 CODE_STANDARDS.md — Project Standards for `mydocs` CLI

This document outlines code structure, design choices, testing expectations, and development conventions for the `mydocs` Go CLI app.

---

## 📐 Architecture & CLI Design

- Use idiomatic Go structure:
  - `cmd/mydocs/main.go`: entry point
  - `internal/` folder for core logic (`create.go`, `publish.go`, `utils.go`)
- All commands should follow a standard signature:
  ```go
  func Run(args []string) error
  ```
- Graceful error handling and clear exit codes.

---

## 🎨 CLI UI Styling

- Use [`charmbracelet/bubbletea`](https://github.com/charmbracelet/bubbletea) for rich terminal UI (TUI) where applicable.
- Follow charmbracelet ecosystem tools for coloring, prompts, etc.
- Keep `help` and feedback output readable and friendly.

---

## 🧪 Testing Requirements

- All commands must have unit tests using Go’s testing framework.
- Follow table-driven tests where appropriate.
- Aim for testable functions — avoid hardcoded `os.Args`/`os.Exit` inside logic.

```go
func CreateEntry(date string) error
func PublishHome() error
```

---

## 📈 Benchmarking

- Add `*_test.go` files with `BenchmarkXxx` functions for:
  - File scanning
  - Markdown generation
  - Sorting logic
- Example:
  ```go
  func BenchmarkPublishHome(b *testing.B) { ... }
  ```

Use:

```bash
go test -bench=.
```

---

## 📦 Dependency Management

- Use Go modules (`go.mod`)
- Keep dependencies minimal and purposeful
- Prefer standard library unless justified

---

## 🧹 Linting, Formatting & Static Analysis

- Use `golangci-lint`:
  ```bash
  golangci-lint run
  ```
- Format all code using:
  ```bash
  gofmt -s -w .
  ```
- Run `go vet` regularly.

---

## 🛠 Makefile Tasks

Include a `Makefile` with the following targets:

```makefile
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

help:
	@echo "Available commands: build, run, test, bench, lint, fmt, vet"
```

---

## 🔁 Commit Conventions (Optional)

Follow [Conventional Commits](https://www.conventionalcommits.org/) if this grows:

- `feat:` New feature
- `fix:` Bug fix
- `refactor:` Code change without behavior change
- `test:` Testing changes
- `chore:` Linting, tooling, deps

---

## 🧪 Example Test File Template

```go
func TestCreateEntry(t *testing.T) {
    tests := []struct {
        name string
        date string
        wantErr bool
    }{
        {"valid date", "2025-12-25", false},
        {"invalid date", "invalid-date", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := CreateEntry(tt.date)
            if (err != nil) != tt.wantErr {
                t.Errorf("got error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

---

## 📚 References

- https://github.com/charmbracelet/bubbletea
- https://golang.org/pkg/testing/
- https://github.com/golangci/golangci-lint
- https://golang.org/cmd/vet/
