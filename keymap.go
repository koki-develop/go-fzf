package fzf

import "github.com/charmbracelet/bubbles/key"

// KeyMap is the key mapping for the Fuzzy Finder.
type KeyMap struct {
	Up     []string
	Down   []string
	Toggle []string
	Choose []string
	Abort  []string
}

type keymap struct {
	Up     key.Binding
	Down   key.Binding
	Toggle key.Binding
	Choose key.Binding
	Abort  key.Binding
}
