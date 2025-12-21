# recall

A simple command-line study helper to manage flashcards and review sessions.
Build, add, and review study cards right from your terminal.

## Getting Started

These instructions will get you a copy of the project running locally for development or testing.

## Installation
You can either build or run the CLI directly:

```bash
# Run the application without building
go run main.go

# Build the binary
go build -o recall main.go

# Run tests
go test ./...
```

## Usage
```bash
# Add a new card
./recall add

# Review cards interactively
./recall review
```

## Makefile

```bash
make build

./recall.exe
```

## Project Structure (overview)
```
recall/
├── cmd/        # Cobra commands
├── internal/   # TUI & card logic
├── main.go     # CLI entrypoint
├── go.mod
├── go.sum
├── tools.go    # Tool dependencies (Cobra CLI)
└── LICENSE
```

## License

[MIT](LICENSE) © 2025 David Eklund