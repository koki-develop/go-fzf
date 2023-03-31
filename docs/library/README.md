# Using as a Library

## Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

First, install the latest version with `go get`.

```console
$ go get -u github.com/koki-develop/go-fzf
```

Next, import go-fzf.

```go
import "github.com/koki-develop/go-fzf"
```

## Usage

- [Basic](#basic)
- [Select Multiple](#select-multiple)
- [Case Sensitive/Insensitive](#case-sensitiveinsensitive)
- [Key Mapping](#key-mapping)
- [Hot Reload](#hot-reload)
- [Customize UI](#customize-ui)

### Basic

First, initialize Fuzzy Finder with `fzf.New()`.

```go
f, err := fzf.New()
if err != nil {
  // ...
}
```

Next, `Find()` launches Fuzzy Finder.  
The first argument is the slice of the item to be searched.  
The second argument is a function that returns the text of the i-th item.

```go
idxs, err := f.Find(items, func(i int) string { return items[i] })
if err != nil {
  // ...
}
```

[Example](/examples/basic/)

### Select Multiple

`fzf.WithLimit()` can be used to set the number of items that can be selected.

```go
f, err := fzf.New(fzf.WithLimit(4))
if err != nil {
  // ...
}
```

`fzf.WithNoLimit()` allows unlimited item selection.

```go
f, err := fzf.New(fzf.WithNoLimit(true)) // no limit
if err != nil {
  // ...
}
```

[Example](/examples/multiple/)

### Case Sensitive/Insensitive

`fzf.WithCaseSensitive()` can be used to make fuzzy searches case-sensitive.

```go
f, err := fzf.New(fzf.WithCaseSensitive(true))
if err != nil {
  // ...
}
```

[Example](/examples/case-sensitive/)

### Key Mapping

`fzf.WithKeyMap()` can be used to customize the key mapping.  
The key mapping that can be customized are as follows

- `Up` - Move cursor up.
- `Down` - Move cursor down.
- `Toggle` - Select or unselect items.
- `Choose` - Complete search.
- `Abort` - Abort search.

```go
f, err := fzf.New(
  fzf.WithNoLimit(true),
  fzf.WithKeyMap(fzf.KeyMap{
    Up:     []string{"up", "ctrl+b"},   // ↑, Ctrl+b
    Down:   []string{"down", "ctrl+f"}, // ↓, Ctrl+f
    Toggle: []string{"tab"},            // tab
    Choose: []string{"enter"},          // Enter
    Abort:  []string{"esc"},            // esc
  }),
)
if err != nil {
  // ...
}
```

[Example](/examples/keymap/)

### Hot Reload

Hot reloading can be enabled using `fzf.WithHotReload()`.

```go
var items []string
var mu sync.Mutex

go func() {
  i := 0
  for {
    time.Sleep(50 * time.Millisecond)
    mu.Lock()
    items = append(items, strconv.Itoa(i))
    mu.Unlock()
    i++
  }
}()

// Pass locker
f, err := fzf.New(fzf.WithHotReload(&mu))
if err != nil {
  // ...
}

// Note that a pointer to slice must be passed.
idxs, err := f.Find(&items, func(i int) string { return items[i] })
if err != nil {
  // ...
}
```

[Example](/examples/hotreload/)

### Customize UI

- [Prompt](#prompt)
- [Cursor](#cursor)
- [Prefix of selected/unselected items](#prefix-of-selectedunselected-items)
- [Placeholder for input](#placeholder-for-input)
- [Count View](#count-view)
- [Styles](#styles)

#### Prompt

`fzf.WithPrompt()` can be used to set the prompt string.

```go
f, err := fzf.New(fzf.WithPrompt("=> "))
if err != nil {
  // ...
}
```

[Example](/examples/prompt/)

#### Cursor

`fzf.WithCursor()` can be used to set the cursor string.

```go
f, err := fzf.New(fzf.WithCursor("=> "))
if err != nil {
  // ...
}
```

[Example](/examples/cursor/)

#### Prefix of selected/unselected items

`fzf.WithSelectedPrefix()` can be used to set the prefix of selected items.  
Similarly, `fzf.WithUnselectedPrefix()` can be used to set the prefix of unselected items.

```go
f, err := fzf.New(
  fzf.WithNoLimit(true),
  fzf.WithSelectedPrefix("[x] "),
  fzf.WithUnselectedPrefix("[ ] "),
)
if err != nil {
  // ...
}
```

[Example](/examples/prefix/)

#### Placeholder for input

`fzf.WithCursor()` can be used to set the placeholder for input.

```go
f, err := fzf.New(fzf.WithInputPlaceholder("Search..."))
if err != nil {
  // ...
}
```

[Example](/examples/placeholder/)

#### Count View

`WithCountViewEnabled()` can be used to enable/disable the count view (enabled by default).  
`WithCountView()` can be used to set the function that renders the count view.  
The argument of the function is a `fzf.CountViewMeta` structure containing the necessary information for the count view.

```go
f, err := fzf.New(
  fzf.WithNoLimit(true),
  fzf.WithCountViewEnabled(true),
  fzf.WithCountView(func(meta fzf.CountViewMeta) string {
    return fmt.Sprintf("items: %d, selected: %d", meta.ItemsCount, meta.SelectedCount)
  }),
)
if err != nil {
  // ...
}
```

[Example](/examples/countview/)

#### Styles

`WithStyles()` can be used to set the style of each component.  
See [reference](https://pkg.go.dev/github.com/koki-develop/go-fzf#Style) for available styles.

```go
f, err := fzf.New(
  fzf.WithNoLimit(true),
  fzf.WithStyles(
    fzf.WithStylePrompt(fzf.Style{Faint: true}),                                       // Prompt
    fzf.WithStyleInputPlaceholder(fzf.Style{Faint: true, ForegroundColor: "#ff0000"}), // Placeholder for input
    fzf.WithStyleInputText(fzf.Style{Italic: true}),                                   // Input text
    fzf.WithStyleCursor(fzf.Style{Bold: true}),                                        // Cursor
    fzf.WithStyleCursorLine(fzf.Style{Bold: true}),                                    // Cursor line
    fzf.WithStyleMatches(fzf.Style{ForegroundColor: "#ff0000"}),                       // Matched characters
    fzf.WithStyleSelectedPrefix(fzf.Style{ForegroundColor: "#ff0000"}),                // Prefix of selected items
    fzf.WithStyleUnselectedPrefix(fzf.Style{Faint: true}),                             // Prefix of unselected items
  ),
)
if err != nil {
  // ...
}
```

[Example](/examples/styles/)
