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
  -h, --help        help for gofzf
  -l, --limit int   maximum number of items to select (default 1)
      --no-limit    unlimited number of items to select
  -v, --version     version for gofzf
```
