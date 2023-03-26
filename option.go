package fzf

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
)

var defaultOption = option{
	limit:   1,
	noLimit: false,

	prompt:           "> ",
	cursor:           "> ",
	selectedPrefix:   "● ",
	unselectedPrefix: "◯ ",
	inputPlaceholder: "Filter...",
	styles:           NewStyles(),

	keymap: &keymap{
		Up:     key.NewBinding(key.WithKeys("up", "ctrl+p")),
		Down:   key.NewBinding(key.WithKeys("down", "ctrl+n")),
		Toggle: key.NewBinding(key.WithKeys("tab")),
		Choose: key.NewBinding(key.WithKeys("enter")),
		Abort:  key.NewBinding(key.WithKeys("ctrl+c", "esc")),
	},

	countViewEnabled: true,
	countViewFunc: func(itemsCount, matchesCount, windowWidth int) string {
		var v strings.Builder
		_, _ = v.WriteString(strconv.Itoa(matchesCount))
		_, _ = v.WriteRune('/')
		_, _ = v.WriteString(strconv.Itoa(itemsCount))
		_, _ = v.WriteRune(' ')
		_, _ = v.WriteString(strings.Repeat("─", max(windowWidth-v.Len(), 0)))
		return v.String()
	},
}

type option struct {
	limit   int
	noLimit bool

	prompt           string
	cursor           string
	selectedPrefix   string
	unselectedPrefix string
	inputPlaceholder string
	styles           *Styles

	keymap *keymap

	countViewEnabled bool
	countViewFunc    func(itemsCount, matchesCount, windowWidth int) string
}

// Option represents a option for the Fuzzy Finder.
type Option func(o *option)

// WithLimit sets the number of items that can be selected.
func WithLimit(l int) Option {
	return func(o *option) {
		o.limit = l
	}
}

// WithNoLimit can be set to `true` to allow unlimited item selection.
func WithNoLimit(n bool) Option {
	return func(o *option) {
		o.noLimit = n
	}
}

// WithPrompt sets the prompt text.
func WithPrompt(p string) Option {
	return func(o *option) {
		o.prompt = p
	}
}

// WithCursor sets the cursor.
func WithCursor(c string) Option {
	return func(o *option) {
		o.cursor = c
	}
}

// WithSelectedPrefix sets the prefix of the selected item.
func WithSelectedPrefix(p string) Option {
	return func(o *option) {
		o.selectedPrefix = p
	}
}

// WithUnselectedPrefix sets the prefix of the unselected item.
func WithUnselectedPrefix(p string) Option {
	return func(o *option) {
		o.unselectedPrefix = p
	}
}

// WithStyles sets the various styles.
func WithStyles(opts ...StylesOption) Option {
	return func(o *option) {
		o.styles = NewStyles(opts...)
	}
}

// WithKeyMap sets the key mapping.
func WithKeyMap(km KeyMap) Option {
	return func(o *option) {
		if len(km.Up) > 0 {
			o.keymap.Up = key.NewBinding(key.WithKeys(km.Up...))
		}
		if len(km.Down) > 0 {
			o.keymap.Down = key.NewBinding(key.WithKeys(km.Down...))
		}
		if len(km.Toggle) > 0 {
			o.keymap.Toggle = key.NewBinding(key.WithKeys(km.Toggle...))
		}
		if len(km.Choose) > 0 {
			o.keymap.Choose = key.NewBinding(key.WithKeys(km.Choose...))
		}
		if len(km.Abort) > 0 {
			o.keymap.Abort = key.NewBinding(key.WithKeys(append(km.Abort, "ctrl+c")...))
		}
	}
}

// WithInputPlaceholder sets the placeholder for input.
func WithInputPlaceholder(p string) Option {
	return func(o *option) {
		o.inputPlaceholder = p
	}
}

// WithCountViewEnabled enables or disables count view.
func WithCountViewEnabled(b bool) Option {
	return func(o *option) {
		o.countViewEnabled = b
	}
}

// WithCountView sets the function to create the count view.
func WithCountView(f func(itemsCount, matchesCount, windowWidth int) string) Option {
	return func(o *option) {
		o.countViewFunc = f
	}
}
