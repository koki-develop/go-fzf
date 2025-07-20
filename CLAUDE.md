# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

go-fzf is a Go library and CLI tool for creating fuzzy finders. It provides:
- A library (`github.com/koki-develop/go-fzf`) for building custom fuzzy finders
- A CLI tool (`gofzf`) that demonstrates the library's capabilities

## Development Commands

### Build
```bash
# Build the CLI tool
go build ./cmd/gofzf
```

### Test
```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./search_test.go

# Run a specific test
go test -run TestName
```

### Lint
```bash
# Run golangci-lint (used in CI)
golangci-lint run
```

## Architecture

### Core Components

1. **FZF struct** (`fzf.go`): Main entry point for the library
   - `New()`: Creates a new fuzzy finder instance with options
   - `Find()`: Launches the fuzzy finder and returns selected indexes
   - `ForceReload()`: Forces reload when hot reload is enabled

2. **Model** (`model.go`): Implements the Bubble Tea Model interface
   - Manages UI state, cursor position, selections
   - Handles window sizing and layout
   - Integrates with Bubble Tea for terminal UI

3. **Items** (`item.go`): Manages the list of searchable items
   - Supports both static slices and hot-reloadable pointer-to-slice
   - Provides thread-safe access for hot reload scenarios

4. **Search** (`search.go`): Implements fuzzy search algorithm
   - Case-sensitive/insensitive matching
   - Returns matched indexes for highlighting

5. **Options** (`option.go`): Configuration for fuzzy finder behavior
   - UI customization (prompts, cursors, styles)
   - Behavioral options (multiple selection, limits, case sensitivity)
   - Keymaps and preview windows

6. **Styles** (`styles.go`): Lipgloss styles for UI components
   - Customizable colors and formatting
   - Separate styles for different UI elements

### Key Design Patterns

- **Builder Pattern**: Options are configured through functional options
- **Model-View Pattern**: Uses Bubble Tea's Model interface for UI state management
- **Interface-based Design**: Items are accessed through reflection to support any slice type

### Dependencies

- **Bubble Tea**: Terminal UI framework
- **Lipgloss**: Styling and layout
- **Cobra**: CLI command framework (for gofzf)
- **testify**: Testing assertions

## Examples

The `/examples` directory contains multiple demonstrations of library usage:
- basic: Simple fuzzy finder
- multiple: Multi-selection
- hotreload: Dynamic item updates
- styles: Custom styling
- keymap: Custom key bindings
- preview-window: Preview pane functionality