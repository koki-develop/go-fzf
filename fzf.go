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

func (fzf *FZF) Find(items Items) (int, error) {
	m := newModel(fzf, items)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return 0, err
	}

	if m.abort {
		return 0, ErrAbort
	}
	if len(m.choices) == 0 {
		return 0, ErrAbort
	}

	return m.choices[0], nil
}
