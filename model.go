package fzf

import (
	"fmt"

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
	abort  bool
	cursor int

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
	return tea.Batch(
		textinput.Blink,
		tea.EnterAltScreen,
	)
}

/*
 * view
 */

func (m *model) View() string {
	return fmt.Sprintf("%s\n%s", m.headerView(), m.itemsView())
}

func (m *model) headerView() string {
	return m.input.View()
}

func (m *model) itemsView() string {
	return fmt.Sprintf("cursor: %d\n%s", m.cursor, m.items)
}

/*
 * update
 */

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.fzf.option.keymap.Abort):
			// abort
			m.abort = true
			return m, tea.Quit
		case key.Matches(msg, m.fzf.option.keymap.Choose):
			// choose
			return m, tea.Quit
		case key.Matches(msg, m.fzf.option.keymap.Up):
			// up
			m.cursorUp()
		case key.Matches(msg, m.fzf.option.keymap.Down):
			// down
			m.cursorDown()
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

func (m *model) cursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *model) cursorDown() {
	if m.cursor < m.items.Len() {
		m.cursor++
	}
}
