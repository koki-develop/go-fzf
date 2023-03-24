package fzf

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sahilm/fuzzy"
)

var (
	_ tea.Model = (*model)(nil)
)

type model struct {
	fzf   *FZF
	items *items

	// state
	abort   bool
	cursor  int
	matches fuzzy.Matches
	choices []int

	// window
	windowWidth     int
	windowHeight    int
	windowYPosition int

	// components
	input textinput.Model
}

func newModel(fzf *FZF, items *items) *model {
	input := textinput.New()
	input.Prompt = fzf.option.prompt
	input.Placeholder = fzf.option.inputPlaceholder
	input.Focus()

	if !fzf.multiple() {
		fzf.option.keymap.Toggle.SetEnabled(false)
	}

	return &model{
		fzf:   fzf,
		items: items,
		// state
		abort:   false,
		cursor:  0,
		matches: fuzzy.Matches{},
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
	cursorLen := lipgloss.Width(m.fzf.option.cursor)

	for i, match := range m.matches {
		if i < m.windowYPosition {
			continue
		}

		// write cursor
		cursor := strings.Repeat(" ", cursorLen)
		if m.cursor == i {
			cursor = m.fzf.option.styles.option.cursor.Render(m.fzf.option.cursor)
		}
		_, _ = v.WriteString(cursor)

		// write toggle
		if m.fzf.multiple() {
			var togglev strings.Builder
			if intContains(m.choices, match.Index) {
				_, _ = togglev.WriteString(m.fzf.option.styles.option.selectedPrefix.Render(m.fzf.option.selectedPrefix))
			} else {
				_, _ = togglev.WriteString(m.fzf.option.styles.option.unselectedPrefix.Render(m.fzf.option.unselectedPrefix))
			}
			_, _ = v.WriteString(togglev.String())
		}

		// write item
		var itemv strings.Builder
		for ci, c := range match.Str {
			// matches
			style := lipgloss.NewStyle()
			if intContains(match.MatchedIndexes, ci) {
				style = style.Inherit(m.fzf.option.styles.option.matches)
			}
			if i == m.cursor {
				style = style.Inherit(m.fzf.option.styles.option.cursorLine)
			}
			_, _ = itemv.WriteString(style.Render(string(c)))
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
		case key.Matches(msg, m.fzf.option.keymap.Toggle):
			// toggle
			m.toggle()
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

	m.filter()
	m.fixYPosition()
	m.fixCursor()

	return m, tea.Batch(cmds...)
}

func (m *model) choice() {
	if len(m.choices) == 0 && m.cursor >= 0 {
		m.choices = append(m.choices, m.matches[m.cursor].Index)
	}
}

func (m *model) toggle() {
	if m.matches.Len() == 0 {
		return
	}

	match := m.matches[m.cursor]
	if intContains(m.choices, match.Index) {
		m.choices = intFilter(m.choices, func(i int) bool { return i != match.Index })
	} else {
		if m.fzf.option.noLimit || len(m.choices) < m.fzf.option.limit {
			m.choices = append(m.choices, match.Index)
		}
	}
}

func (m *model) cursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *model) cursorDown() {
	if m.cursor+1 < len(m.matches) {
		m.cursor++
	}
}

func (m *model) filter() {
	var matches fuzzy.Matches

	s := m.input.Value()
	if s == "" {
		for i := 0; i < m.items.Len(); i++ {
			matches = append(matches, fuzzy.Match{
				Str:   m.items.String(i),
				Index: i,
			})
		}
		m.matches = matches
		return
	}

	m.matches = fuzzy.FindFrom(s, m.items)
}

func (m *model) fixCursor() {
	if m.cursor < 0 && len(m.matches) > 0 {
		m.cursor = 0
		return
	}

	if m.cursor+1 > len(m.matches) {
		m.cursor = len(m.matches) - 1
		return
	}
}

func (m *model) fixYPosition() {
	headerHeight := lipgloss.Height(m.headerView())

	if m.windowHeight-headerHeight > len(m.matches) {
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
