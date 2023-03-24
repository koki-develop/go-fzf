package fzf

import tea "github.com/charmbracelet/bubbletea"

var defaultFindOption = findOption{
	itemPrefixFunc: nil,
}

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
func (fzf *FZF) Find(items interface{}, itemFunc func(i int) string, opts ...FindOption) ([]int, error) {
	o := defaultFindOption
	for _, opt := range opts {
		opt(&o)
	}

	is, err := newItems(items, itemFunc, o.itemPrefixFunc)
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

// Option represents a option for the Find.
type FindOption func(o *findOption)

type findOption struct {
	itemPrefixFunc func(i int) string
}

// WithItemPrefix sets the prefix function of the item.
func WithItemPrefix(f func(i int) string) FindOption {
	return func(o *findOption) {
		o.itemPrefixFunc = f
	}
}
