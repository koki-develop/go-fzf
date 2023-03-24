package fzf

import tea "github.com/charmbracelet/bubbletea"

// Fuzzy Finder.
type FZF struct {
	option *option
}

// New returns a new Fuzzy Finder.
func New(opts ...Option) *FZF {
	o := defaultOption
	for _, opt := range opts {
		opt(&o)
	}

	return &FZF{
		option: &o,
	}
}

// Find launches the Fuzzy Finder and returns a list of indexes of the selected items.
func (fzf *FZF) Find(items interface{}, itemFunc func(i int) string) ([]int, error) {
	is, err := newItems(items, itemFunc)
	if err != nil {
		return nil, err
	}
	m := newModel(fzf, is)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return nil, err
	}

	if m.abort {
		return nil, ErrAbort
	}

	return m.choices, nil
}

func (fzf *FZF) multiple() bool {
	return fzf.option.noLimit || fzf.option.limit > 1
}
