# recall

A terminal-based flashcard management app built with Go. Organize your flashcards into decks, review them interactively, and manage them directly from your terminal.

## Tech Stack

- **Cobra** - CLI framework
- **Bubble Tea** - Terminal UI framework
- **Lipgloss** - Terminal styling

## Features

- Deck management: create, select, and list multiple decks
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

### Deck Commands
```bash
# Show all deck commands
./recall deck

# Create a new deck
./recall deck create --name "Biology"

# Select an active deck
./recall deck select --name "Biology"

# List all decks
./recall deck ls
```

### Flashcard commands
```bash
# Add a new flashcard to the active deck
./recall add

# List all flashcards in the active deck (interactive view with edit/delete)
./recall list

# Start a review session for the active deck
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
- Add confirmation prompt (y/n) for deleting cards in list view
- Export/import decks as JSON
- Multiple-choice flashcard generation with AI
- Review statistics and spaced repetition support
- Add database support
- Delete/edit decks
- bubble tea / lipgloss for deck ls
- Shuffle cards
- Implement the Go color package

## License

[MIT](LICENSE) **© 2025 David Eklund**