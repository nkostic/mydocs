# 📝 mydocs CLI

A console-based journaling utility written in Go to help you quickly generate daily Markdown notes and publish them to a centralized index.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Build Status](https://img.shields.io/badge/build-passing-green.svg)]()

## ✨ Features

- 🆕 **Create journal entries** for today or specific dates
- 📑 **Auto-generate home index** with links to all entries  
- 🎨 **Beautiful terminal UI** with colors and styling
- 🚀 **Fast and lightweight** - single binary with no dependencies
- 📅 **Smart date handling** with validation
- 🔄 **Duplicate prevention** - won't overwrite existing entries
- 📊 **Comprehensive testing** with benchmarks

## 🚀 Quick Start

### Build the application

```bash
make build
```

### Create a journal entry for today

```bash
./mydocs new
```

### Create a journal entry for a specific date

```bash
./mydocs new 2025-12-25
```

### Update the home index

```bash
./mydocs publish
```

### Show version information

```bash
./mydocs version
```

### Show help

```bash
./mydocs help
```

## 📁 Project Structure

```
mydocs/
├── cmd/mydocs/          # Application entry point
│   └── main.go
├── internal/            # Core application logic
│   ├── create.go        # Journal entry creation
│   ├── create_test.go   # Tests for create functionality
│   ├── publish.go       # Home index publishing
│   ├── publish_test.go  # Tests for publish functionality
│   ├── styles.go        # Terminal UI styling
│   └── version.go       # Version information
├── docs/                # Documentation
│   ├── CODE_STANDARDS.MD # Development standards
│   └── INSTRUCTIONS.MD   # Original project requirements
├── .editorconfig        # Editor configuration
├── .gitignore          # Git ignore rules
├── CONTRIBUTING.md     # Contribution guidelines
├── LICENSE             # MIT License
├── Makefile            # Build and development tasks
├── README.md           # This file
└── go.mod              # Go module definition
```

## 🛠 Development

### Available Make targets

```bash
make help              # Show available commands
make build             # Build the binary
make run               # Run the application
make test              # Run tests
make bench             # Run benchmarks
make lint              # Run linter
make fmt               # Format code
make vet               # Run go vet
make clean             # Remove built binary
make install           # Install to GOPATH/bin
```

### Running tests

```bash
make test
```

### Running benchmarks

```bash
make bench
```

## 📖 Usage Examples

### Creating entries

```bash
# Create entry for today
./mydocs new

# Create entry for Christmas
./mydocs new 2025-12-25
```

### Publishing to index

```bash
# Update home.md with all journal entries
./mydocs publish
```

This will create a `home.md` file with links to all your journal entries, sorted by date (newest first).

## 🧪 Testing

The project includes comprehensive tests for all core functionality:

- Unit tests for entry creation
- Unit tests for home publishing
- Benchmarks for performance testing
- Table-driven tests following Go best practices

Run the tests with:

```bash
go test ./...
```

## 🤝 Contributing

We welcome contributions! This is an open source project.

### Quick Start for Contributors

1. **Fork the repository** on GitHub
2. **Clone your fork**: `git clone https://github.com/YOUR_USERNAME/mydocs.git`
3. **Create a feature branch**: `git checkout -b feature/amazing-feature`
4. **Make your changes** following our [Code Standards](docs/CODE_STANDARDS.MD)
5. **Run tests**: `make test`
6. **Commit your changes**: `git commit -m 'feat: add amazing feature'`
7. **Push to your fork**: `git push origin feature/amazing-feature`
8. **Open a Pull Request** on GitHub

### Guidelines

- Follow the [Code Standards](docs/CODE_STANDARDS.MD)
- Read the full [Contributing Guide](CONTRIBUTING.md)
- Write tests for new features
- Use conventional commit messages
- Ensure all CI checks pass

### Development Setup

```bash
# Install dependencies
go mod tidy

# Build the application
make build

# Run all tests
make test

# Check code quality
make lint && make vet && make fmt
```

## 📚 Documentation

- **[Project Requirements](docs/INSTRUCTIONS.MD)** - Original specifications
- **[Code Standards](docs/CODE_STANDARDS.MD)** - Development guidelines  
- **[Contributing Guide](CONTRIBUTING.md)** - How to contribute
- **[License](LICENSE)** - MIT License terms

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Nenad Kostic**

## 🙏 Acknowledgments

This project follows the specifications in [docs/INSTRUCTIONS.MD](docs/INSTRUCTIONS.MD) and [docs/CODE_STANDARDS.MD](docs/CODE_STANDARDS.MD).

Built with ❤️ using:
- [Go](https://golang.org/) - The programming language
- [Charm](https://charm.sh/) - Beautiful terminal UI components
- [lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
