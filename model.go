package fzf

import tea "github.com/charmbracelet/bubbletea"

var (
	_ tea.Model = (*model)(nil)
)

type model struct {
	fzf *FZF
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) View() string {
	return m.fzf.prompt
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}
