package fzf

import tea "github.com/charmbracelet/bubbletea"

type FZF struct {
	option *option
}

func New(opts ...Option) *FZF {
	o := defaultOption
	for _, opt := range opts {
		opt(&o)
	}

	return &FZF{
		option: &o,
	}
}

func (fzf *FZF) Find(items Items) ([]int, error) {
	m := newModel(fzf, newItems(items))

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return []int{}, err
	}

	if m.abort {
		return []int{}, ErrAbort
	}
	if len(m.choices) == 0 {
		return []int{}, ErrAbort
	}

	return m.choices, nil
}

func (fzf *FZF) multiple() bool {
	return fzf.option.noLimit || fzf.option.limit > 1
}
