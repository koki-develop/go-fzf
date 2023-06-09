# Using as a CLI

## Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Basic](#basic)
  - [Select Multiple](#select-multiple)
  - [Customize UI](#customize-ui)

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

### Customize UI

The `gofzf` CLI allows for various visual customizations using flags.

- [Prompt](#prompt)
- [Position of input](#position-of-input)
- [Placeholder for input](#placeholder-for-input)
- [Input text](#input-text)
- [Cursor](#cursor)
- [Cursor Line](#cursor-line)
- [Prefix of selected/unselected items](#prefix-of-selectedunselected-items)
- [Matched characters](#matched-characters)

#### Prompt

| Flag                 | Default | Description                 |
| -------------------- | ------- | --------------------------- |
| `--prompt`           | `"> "`  | Prompt string.              |
| `--prompt-fg`        | N/A     | Foreground color of prompt. |
| `--prompt-bg`        | N/A     | Background color of prompt. |
| `--prompt-bold`      | `false` | Bold prompt.                |
| `--prompt-blink`     | `false` | Blink prompt.               |
| `--prompt-italic`    | `false` | Italicize prompt.           |
| `--prompt-strike`    | `false` | Strkethrough prompt.        |
| `--prompt-underline` | `false` | Underline prompt.           |
| `--prompt-faint`     | `false` | Faint prompt.               |

#### Position of input

| Flag               | Default | Description                                           |
| ------------------ | ------- | ----------------------------------------------------- |
| `--input-position` | `"top"` | Position of input. Either `top` or `bottom` is valid. |

#### Placeholder for input

| Flag                            | Default       | Description                                |
| ------------------------------- | ------------- | ------------------------------------------ |
| `--input-placeholder`           | `"Filter..."` | Placeholder for input.                     |
| `--input-placeholder-fg`        | N/A           | Foreground color of placeholder for input. |
| `--input-placeholder-bg`        | N/A           | Background color of placeholder for input. |
| `--input-placeholder-bold`      | `false`       | Bold placeholder for input.                |
| `--input-placeholder-blink`     | `false`       | Blink placeholder for input.               |
| `--input-placeholder-italic`    | `false`       | Italicize placeholder for input.           |
| `--input-placeholder-strike`    | `false`       | Strkethrough placeholder for input.        |
| `--input-placeholder-underline` | `false`       | Underline placeholder for input.           |
| `--input-placeholder-faint`     | `true`        | Faint placeholder for input.               |

#### Input text

| Flag                     | Default | Description                     |
| ------------------------ | ------- | ------------------------------- |
| `--input-text-fg`        | N/A     | Foreground color of input text. |
| `--input-text-bg`        | N/A     | Background color of input text. |
| `--input-text-bold`      | `false` | Bold input text.                |
| `--input-text-blink`     | `false` | Blink input text.               |
| `--input-text-italic`    | `false` | Italicize input text.           |
| `--input-text-strike`    | `false` | Strkethrough input text.        |
| `--input-text-underline` | `false` | Underline input text.           |
| `--input-text-faint`     | `false` | Faint input text.               |

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
