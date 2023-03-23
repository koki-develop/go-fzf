package fzf

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	_ tea.Model = (*model)(nil)
)

type model struct {
	fzf   *FZF
	items Items

	// state
	aborted bool

	// components
	input textinput.Model
}

func newModel(fzf *FZF, items Items) *model {
	input := textinput.New()
	input.Prompt = fzf.option.prompt
	input.Placeholder = fzf.option.inputPlaceholder
	input.Focus()

	return &model{
		fzf:   fzf,
		items: items,
		// components
		input: input,
	}
}

func (m *model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m *model) View() string {
	return m.input.View()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.fzf.option.keymap.Abort):
			m.aborted = true
			return m, tea.Quit
		case key.Matches(msg, m.fzf.option.keymap.Choose):
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd
	{
		input, cmd := m.input.Update(msg)
		m.input = input
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
