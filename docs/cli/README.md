# Using as a CLI

## Contents

- [Installation](TODO)
- [Usage](TODO)
  - [Basic](TODO)
  - [Select Multiple](TODO)
  - [Customize Styles](TODO)

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

Download the binary from the [release page](https://github.com/koki-develop/go-fzf/releases/latest).

## Usage

### Basic

Running `gofzf` without arguments will recursively fuzzy search for files from the current working directory.

![](./basic.gif)

You can also pass items to search for from stdin, separated by a new line.

![](./basic-stdin.gif)

### Select Multiple

The `--limit` flag allows you to set the number of items that can be selected.  
Items can be selected with the Tab key.

![](./limit.gif)

Setting the `--no-limit` flag allows unlimited item selection.

![](./no-limit.gif)

### Customize Styles

The `gofzf` CLI allows for various visual customizations using flags.

- [Prompt](#prompt)
- [Cursor](#cursor)
- [Cursor Line](#cursor-line)
- [Prefix of selected/unselected items](#prefix-of-selectedunselected-items)
- [Placeholder for input](#placeholder-for-input)
- [Matched characters](#matched-characters)

#### Prompt

| Flag       | Default | Description    |
| ---------- | ------- | -------------- |
| `--prompt` | `"> "`  | Prompt string. |

#### Cursor

| Flag                 | Default     | Description                 |
| -------------------- | ----------- | --------------------------- |
| `--cursor`           | `"> "`      | Cursor string.              |
| `--cursor-fg`        | `"#00ADD8"` | Foreground color of cursor. |
| `--cursor-bg`        | N/A         | Background color of cursor. |
| `--cursor-bold`      | `false`     | Bold cursor.                |
| `--cursor-blink`     | `false`     | Blink cursor.               |
| `--cursor-italic`    | `false`     | Italicize cursor.           |
| `--cursor-strike`    | `false`     | Strkethrough cursor.        |
| `--cursor-underline` | `false`     | Underline cursor.           |
| `--cursor-faint`     | `false`     | Faint cursor.               |

#### Cursor Line

| Flag                     | Default | Description                      |
| ------------------------ | ------- | -------------------------------- |
| `--cursorline-fg`        | N/A     | Foreground color of cursor line. |
| `--cursorline-bg`        | N/A     | Background color of cursor line. |
| `--cursorline-bold`      | `true`  | Bold cursor line.                |
| `--cursorline-blink`     | `false` | Blink cursor line.               |
| `--cursorline-italic`    | `false` | Italicize cursor line.           |
| `--cursorline-strike`    | `false` | Strkethrough cursor line.        |
| `--cursorline-underline` | `false` | Underline cursor line.           |
| `--cursorline-faint`     | `false` | Faint cursor line.               |

#### Prefix of selected/unselected items

| Flag                          | Default     | Description                                   |
| ----------------------------- | ----------- | --------------------------------------------- |
| `--selected-prefix`           | `"●"`       | Prefix of selected items.                     |
| `--selected-prefix-fg`        | `"#00ADD8"` | Foreground color of prefix of selected items. |
| `--selected-prefix-bg`        | N/A         | Background color of prefix of selected items. |
| `--selected-prefix-bold`      | `false`     | Bold prefix of selected items.                |
| `--selected-prefix-blink`     | `false`     | Blink prefix of selected items.               |
| `--selected-prefix-italic`    | `false`     | Italicize prefix of selected items.           |
| `--selected-prefix-strike`    | `false`     | Strkethrough prefix of selected items.        |
| `--selected-prefix-underline` | `false`     | Underline prefix of selected items.           |
| `--selected-prefix-faint`     | `false`     | Faint prefix of selected items.               |

| Flag                            | Default | DescriptioDescription                           |
| ------------------------------- | ------- | ----------------------------------------------- |
| `--unselected-prefix`           | `"◯"`   | Prefix of unselected items.                     |
| `--unselected-prefix-fg`        | N/A     | Foreground color of prefix of unselected items. |
| `--unselected-prefix-bg`        | N/A     | Background color of prefix of unselected items. |
| `--unselected-prefix-bold`      | `false` | Bold prefix of unselected items.                |
| `--unselected-prefix-blink`     | `false` | Blink prefix of unselected items.               |
| `--unselected-prefix-italic`    | `false` | Italicize prefix of unselected items.           |
| `--unselected-prefix-strike`    | `false` | Strkethrough prefix of unselected items.        |
| `--unselected-prefix-underline` | `false` | Underline prefix of unselected items.           |
| `--unselected-prefix-faint`     | `true`  | Faint prefix of unselected items.               |

#### Placeholder for input

| Flag                  | Default       | Description            |
| --------------------- | ------------- | ---------------------- |
| `--input-placeholder` | `"Filter..."` | Placeholder for input. |

#### Matched characters

| Flag                  | Default     | Description                             |
| --------------------- | ----------- | --------------------------------------- |
| `--matches-fg`        | `"#00ADD8"` | Foreground color of matched characters. |
| `--matches-bg`        | N/A         | Background color of matched characters. |
| `--matches-bold`      | `false`     | Bold matched characters.                |
| `--matches-blink`     | `false`     | Blink matched characters.               |
| `--matches-italic`    | `false`     | Italicize matched characters.           |
| `--matches-strike`    | `false`     | Strikethrough matched characters.       |
| `--matches-underline` | `false`     | Underline matched characters.           |
| `--matches-faint`     | `false`     | Faint matched characters.               |
