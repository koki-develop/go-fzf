package fzf

import (
	"github.com/charmbracelet/bubbles/key"
)

var defaultOption = option{
	prompt:           "> ",
	cursor:           "> ",
	inputPlaceholder: "Filter...",
	styles:           NewStyles(),
	keymap: &keymap{
		Up:     key.NewBinding(key.WithKeys("up", "ctrl+p")),
		Down:   key.NewBinding(key.WithKeys("down", "ctrl+n")),
		Toggle: key.NewBinding(key.WithKeys("tab")),
		Choose: key.NewBinding(key.WithKeys("enter")),
		Abort:  key.NewBinding(key.WithKeys("ctrl+c", "esc")),
	},
}

type option struct {
	prompt           string
	cursor           string
	styles           *Styles
	inputPlaceholder string
	keymap           *keymap
}

type Option func(o *option)

func WithPrompt(p string) Option {
	return func(o *option) {
		o.prompt = p
	}
}

func WithCursor(c string) Option {
	return func(o *option) {
		o.cursor = c
	}
}

func WithStyles(opts ...StylesOption) Option {
	return func(o *option) {
		o.styles = NewStyles(opts...)
	}
}

func WithKeyMap(km *KeyMap) Option {
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
			o.keymap.Abort = key.NewBinding(key.WithKeys(km.Abort...))
		}
	}
}

func WithInputPlaceholder(p string) Option {
	return func(o *option) {
		o.inputPlaceholder = p
	}
}
