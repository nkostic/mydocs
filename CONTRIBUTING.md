# Contributing to mydocs

Thank you for your interest in contributing to mydocs! This document provides guidelines and information for contributors.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Testing](#testing)
- [Submitting Changes](#submitting-changes)
- [Code Style](#code-style)
- [Project Structure](#project-structure)

## 🤝 Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct:

- Be respectful and inclusive
- Welcome newcomers and help them learn
- Focus on constructive feedback
- Respect differing viewpoints and experiences

## 🚀 Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/mydocs.git
   cd mydocs
   ```
3. **Add the upstream remote**:
   ```bash
   git remote add upstream https://github.com/nkostic/mydocs.git
   ```

## 🛠 Development Setup

### Prerequisites

- **Go 1.21+** - [Download & Install Go](https://golang.org/dl/)
- **Make** - For running build tasks
- **Git** - For version control

### Setup

1. **Install dependencies**:

   ```bash
   go mod tidy
   ```

2. **Build the application**:

   ```bash
   make build
   ```

3. **Run tests**:

   ```bash
   make test
   ```

4. **Check code quality**:
   ```bash
   make lint
   make vet
   make fmt
   ```

## 🔧 Making Changes

1. **Create a feature branch**:

   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following our [Code Standards](docs/CODE_STANDARDS.MD)

3. **Write or update tests** for your changes

4. **Test your changes**:

   ```bash
   make test
   make bench
   ```

5. **Ensure code quality**:
   ```bash
   make lint
   make fmt
   make vet
   ```

## 🧪 Testing

We use comprehensive testing practices:

- **Unit tests**: Test individual functions
- **Table-driven tests**: Test multiple scenarios
- **Benchmarks**: Performance testing
- **Integration tests**: Test component interactions

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
go test -cover ./...

# Run benchmarks
make bench

# Run specific test
go test -run TestCreateEntry ./internal
```

### Writing Tests

Follow our testing patterns:

```go
func TestYourFunction(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {"valid case", "input", "expected", false},
        {"error case", "bad", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := YourFunction(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("got error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("got = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## 📤 Submitting Changes

1. **Commit your changes**:

   ```bash
   git add .
   git commit -m "feat: add new feature description"
   ```

2. **Push to your fork**:

   ```bash
   git push origin feature/your-feature-name
   ```

3. **Create a Pull Request** on GitHub with:
   - Clear title and description
   - Reference any related issues
   - Include screenshots if UI changes
   - Ensure CI passes

### Commit Message Format

We follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation changes
- `style:` Code style changes (formatting, etc.)
- `refactor:` Code refactoring
- `test:` Adding or updating tests
- `chore:` Maintenance tasks

## 🎨 Code Style

We follow strict code standards defined in [docs/CODE_STANDARDS.MD](docs/CODE_STANDARDS.MD):

### Go Code Style

- Use `gofmt` for formatting
- Follow Go idioms and best practices
- Use meaningful variable and function names
- Include doc comments for exported functions
- Handle errors appropriately

### Terminal UI

- Use `charmbracelet/lipgloss` for styling
- Maintain consistent color scheme
- Provide clear visual feedback
- Use emojis appropriately for UX

### Example

```go
// CreateEntry creates a new journal entry for the specified date.
// If date is empty, uses today's date.
func CreateEntry(date string) error {
    var targetDate time.Time
    var err error

    if date == "" {
        targetDate = time.Now()
    } else {
        targetDate, err = time.Parse("2006-01-02", date)
        if err != nil {
            return fmt.Errorf("invalid date format. Please use YYYY-MM-DD format")
        }
    }

    // ... rest of implementation
}
```

## 📁 Project Structure

```
mydocs/
├── cmd/mydocs/          # Application entry point
│   └── main.go
├── internal/            # Core application logic
│   ├── create.go        # Journal creation
│   ├── publish.go       # Home publishing
│   ├── styles.go        # Terminal styling
│   └── version.go       # Version information
├── docs/                # Documentation
│   ├── CODE_STANDARDS.MD
│   └── INSTRUCTIONS.MD
├── .editorconfig        # Editor configuration
├── .gitignore          # Git ignore rules
├── CONTRIBUTING.md     # This file
├── LICENSE             # MIT License
├── Makefile           # Build automation
├── README.md          # Project overview
└── go.mod             # Go module definition
```

## 🐛 Reporting Issues

When reporting issues, please include:

- **Description**: Clear description of the problem
- **Steps to reproduce**: Exact steps to trigger the issue
- **Expected behavior**: What should happen
- **Actual behavior**: What actually happens
- **Environment**: OS, Go version, etc.
- **Logs**: Any relevant error messages

## 💡 Feature Requests

We welcome feature requests! Please:

- Check existing issues first
- Provide clear use case
- Explain expected behavior
- Consider implementation impact

## 📞 Getting Help

- **GitHub Issues**: For bugs and feature requests
- **GitHub Discussions**: For questions and general discussion
- **Email**: [contact info if available]

## 🏆 Recognition

Contributors will be recognized in:

- GitHub contributors list
- Release notes for significant contributions
- Special thanks in README for major features

Thank you for contributing to mydocs! 🙏
