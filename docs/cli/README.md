# Using as a CLI

🚧 WIP 🚧

- [Installation](#installation)
- [Usage](#usage)

## Installation

### Homebrew

```console
$ brew install koki-develop/tap/gofzf
```

### `go install`

```console
$ go install github.com/koki-develop/go-fzf/cmd/gofzf@latest
```

### Releases

Download the binary from the [releases page](https://github.com/koki-develop/go-fzf/releases/latest).

## Usage

```console
$ gofzf --help
Usage:
  gofzf [flags]

Flags:
      --limit int                     maximum number of items to select (default 1)
      --no-limit                      unlimited number of items to select
      --prompt string                  (default "> ")
      --cursor string                  (default "> ")
      --selected-prefix string         (default "● ")
      --unselected-prefix string       (default "◯ ")
      --input-placeholder string       (default "Filter...")
      --cursor-fg string               (default "#00ADD8")
      --cursor-bg string
      --cursor-bold
      --cursor-blink
      --cursor-italic
      --cursor-strike
      --cursor-underline
      --cursor-faint
      --cursorline-fg string
      --cursorline-bg string
      --cursorline-bold                (default true)
      --cursorline-blink
      --cursorline-italic
      --cursorline-strke
      --cursorline-underline
      --cursorline-faint
      --selected-prefix-fg string      (default "#00ADD8")
      --selected-prefix-bg string
      --selected-prefix-bold
      --selected-prefix-blink
      --selected-prefix-italic
      --selected-prefix-strke
      --selected-prefix-underline
      --selected-prefix-faint
      --unselected-prefix-fg string
      --unselected-prefix-bg string
      --unselected-prefix-bold
      --unselected-prefix-blink
      --unselected-prefix-italic
      --unselected-prefix-strke
      --unselected-prefix-underline
      --unselected-prefix-faint        (default true)
      --matches-fg string              (default "#00ADD8")
      --matches-bg string
      --matches-bold
      --matches-blink
      --matches-italic
      --matches-strke
      --matches-underline
      --matches-faint
  -h, --help                          help for gofzf
  -v, --version                       version for gofzf
```
