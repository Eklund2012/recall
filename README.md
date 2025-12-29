# recall

A terminal-based flashcard management app built with Go.

## Tech Stack

- **Cobra** - CLI framework
- **Bubble Tea** - Terminal UI framework
- **Lipgloss** - Terminal styling

## Features

- Add flashcards (front/back)
- List flashcards in an interactive TUI
- Edit flashcards from the list view
- Delete flashcards from the list view
- Persistent JSON storage
- Status messages in the TUI

## Installation

```bash
# Build the binary
go build -o recall.exe main.go

# Or use make
make build

# Run tests
go test ./...
```

## Usage

```bash
# Add a new flashcard
./recall add

# List all flashcards (interactive view with edit/delete)
./recall list

# Start review session of flashcards
./recall review
```

## Project Structure
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

## Roadmap
- y/n confirmation for deleting in recall list 

## License

[MIT](LICENSE) © 2025 David Eklund