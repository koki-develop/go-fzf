package fzf

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

var defaultFindOption = findOption{
	itemPrefixFunc:    nil,
	previewWindowFunc: nil,
	preselectAll:      false,
	preselect:         []int{},
}

// Fuzzy Finder.
type FZF struct {
	model   *model
	program *tea.Program
}

// New returns a new Fuzzy Finder.
func New(opts ...Option) (*FZF, error) {
	o := defaultOption
	for _, opt := range opts {
		opt(&o)
	}

	if err := o.valid(); err != nil {
		return nil, err
	}

	m := newModel(&o)

	return &FZF{
		model:   m,
		program: tea.NewProgram(m, tea.WithOutput(os.Stderr)),
	}, nil
}

// Find launches the Fuzzy Finder and returns a list of indexes of the selected items.
func (fzf *FZF) Find(items interface{}, itemFunc func(i int) string, opts ...FindOption) ([]int, error) {
	findOption := defaultFindOption
	for _, opt := range opts {
		opt(&findOption)
	}

	rv := reflect.ValueOf(items)
	if fzf.model.option.hotReloadLocker == nil {
		if rv.Kind() != reflect.Slice {
			return nil, fmt.Errorf("items must be a slice, but got %T", items)
		}
	} else {
		if !(rv.Kind() == reflect.Ptr && reflect.Indirect(rv).Kind() == reflect.Slice) {
			return nil, fmt.Errorf("items must be a pointer to slice, but got %T", items)
		}
	}

	is, err := newItems(rv, itemFunc)
	if err != nil {
		return nil, err
	}
	fzf.model.loadItems(is)
	fzf.model.setFindOption(&findOption)
	if fzf.model.option.noLimit {
		if findOption.preselectAll {
			fzf.model.choices = make([]int, is.Len())
			for i := 0; i < is.Len(); i++ {
				fzf.model.choices[i] = i
			}
		} else {
			fzf.model.choices = make([]int, len(findOption.preselect))
			for i, idx := range findOption.preselect {
				fzf.model.choices[i] = idx
			}
		}
	}

	if _, err := fzf.program.Run(); err != nil {
		return nil, err
	}

	if fzf.model.abort {
		return nil, ErrAbort
	}

	return fzf.model.choices, nil
}

// ForceReload forces the reload of items.
// HotReload must be enabled.
func (fzf *FZF) ForceReload() error {
	if fzf.model.option.hotReloadLocker == nil {
		return errors.New("hot reload is not enabled")
	}

	fzf.program.Send(forceReloadMsg{})
	return nil
}

// Quit quits the Fuzzy Finder.
func (fzf *FZF) Quit() {
	fzf.program.Quit()
}

// Abort aborts the Fuzzy Finder.
func (fzf *FZF) Abort() {
	fzf.model.abort = true
	fzf.Quit()
}

// Option represents a option for the Find.
type FindOption func(o *findOption)

type findOption struct {
	itemPrefixFunc    func(i int) string
	previewWindowFunc func(i, width, height int) string
	preselectAll      bool
	preselect         []int
}

// WithItemPrefix sets the prefix function of the item.
func WithItemPrefix(f func(i int) string) FindOption {
	return func(o *findOption) {
		o.itemPrefixFunc = f
	}
}

// WithPreviewWindow sets the preview window function of the item.
func WithPreviewWindow(f func(i, width, height int) string) FindOption {
	return func(o *findOption) {
		o.previewWindowFunc = f
	}
}

func WithPreselectAll(preselect bool) FindOption {
	return func(o *findOption) {
		o.preselectAll = preselect
	}
}

// WithPreselect sets the preselected indexes.
func WithPreselect(idxs []int) FindOption {
	return func(o *findOption) {
		o.preselect = idxs
	}
}
