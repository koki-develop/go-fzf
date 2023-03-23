package fzf

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	_ tea.Model = (*model)(nil)
)

type model struct {
	fzf   *FZF
	items Items

	// state
	abort   bool
	cursor  int
	choices []int

	// window
	windowWidth     int
	windowHeight    int
	windowYPosition int

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
		// state
		abort:   false,
		cursor:  0,
		choices: []int{},
		// window
		windowWidth:     0,
		windowHeight:    0,
		windowYPosition: 0,
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
	var v strings.Builder

	headerHeight := lipgloss.Height(m.headerView())
	cursorLen := stringLen(m.fzf.option.cursor)

	for i := 0; i < m.items.Len(); i++ {
		if i < m.windowYPosition {
			continue
		}

		// write cursor
		cursor := strings.Repeat(" ", cursorLen)
		if m.cursor == i {
			cursor = m.fzf.option.cursor
		}
		_, _ = v.WriteString(cursor)

		// write item
		itemstring := m.items.ItemString(i)
		var itemv strings.Builder
		for _, c := range itemstring {
			// cursor line
			if i == m.cursor {
				_, _ = itemv.WriteString(m.fzf.option.styles.CursorLine.Render(string(c)))
			} else {
				itemv.WriteString(string(c))
			}
		}
		_, _ = v.WriteString(itemv.String())

		if i+1 == m.windowYPosition+(m.windowHeight-(headerHeight)) {
			break
		}
		v.WriteString("\n")
	}

	return v.String()
}

/*
 * update
 */

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// key
		switch {
		case key.Matches(msg, m.fzf.option.keymap.Abort):
			// abort
			m.abort = true
			return m, tea.Quit
		case key.Matches(msg, m.fzf.option.keymap.Choose):
			// choose
			m.choice()
			return m, tea.Quit
		case key.Matches(msg, m.fzf.option.keymap.Up):
			// up
			m.cursorUp()
		case key.Matches(msg, m.fzf.option.keymap.Down):
			// down
			m.cursorDown()
		}
	case tea.WindowSizeMsg:
		// window
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		m.input.Width = m.windowWidth - 3
	}

	var cmds []tea.Cmd
	{
		input, cmd := m.input.Update(msg)
		m.input = input
		cmds = append(cmds, cmd)
	}

	m.fixYPosition()
	m.fixCursor()

	return m, tea.Batch(cmds...)
}

func (m *model) choice() {
	if m.items.Len() == 0 {
		m.abort = true
		return
	}

	if len(m.choices) == 0 && m.cursor >= 0 {
		m.choices = append(m.choices, m.cursor)
	}
}

func (m *model) cursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *model) cursorDown() {
	if m.cursor+1 < m.items.Len() {
		m.cursor++
	}
}

func (m *model) fixCursor() {
	if m.cursor < 0 && m.items.Len() > 0 {
		m.cursor = 0
		return
	}

	if m.cursor+1 > m.items.Len() {
		m.cursor = m.items.Len() - 1
		return
	}
}

func (m *model) fixYPosition() {
	headerHeight := lipgloss.Height(m.headerView())

	if m.windowHeight-headerHeight > m.items.Len() {
		m.windowYPosition = 0
		return
	}

	if m.cursor < m.windowYPosition {
		m.windowYPosition = m.cursor
		return
	}

	if m.cursor+1 >= (m.windowHeight-headerHeight)+m.windowYPosition {
		m.windowYPosition = m.cursor + 1 - (m.windowHeight - headerHeight)
		return
	}
}
