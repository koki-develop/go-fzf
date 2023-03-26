package fzf

import (
	"fmt"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

var defaultFindOption = findOption{
	itemPrefixFunc: nil,
}

// Fuzzy Finder.
type FZF struct {
	option *option
	model  *model
}

// New returns a new Fuzzy Finder.
func New(opts ...Option) *FZF {
	o := defaultOption
	for _, opt := range opts {
		opt(&o)
	}

	m := newModel(&o)

	return &FZF{
		option: &o,
		model:  m,
	}
}

// Find launches the Fuzzy Finder and returns a list of indexes of the selected items.
func (fzf *FZF) Find(items interface{}, itemFunc func(i int) string, opts ...FindOption) ([]int, error) {
	findOption := defaultFindOption
	for _, opt := range opts {
		opt(&findOption)
	}

	rv := reflect.ValueOf(items)
	switch {
	case rv.Kind() == reflect.Slice:
	case rv.Kind() == reflect.Ptr && reflect.Indirect(rv).Kind() == reflect.Slice:
	default:
		return nil, fmt.Errorf("items must be a slice, but got %T", items)
	}

	is, err := newItems(rv, itemFunc)
	if err != nil {
		return nil, err
	}
	fzf.model.setItems(is)
	fzf.model.setFindOption(&findOption)

	p := tea.NewProgram(fzf.model)
	if _, err := p.Run(); err != nil {
		return nil, err
	}

	if fzf.model.abort {
		return nil, ErrAbort
	}

	return fzf.model.choices, nil
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
