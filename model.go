package fzf

import (
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
	fzf        *FZF
	items      *items
	findOption *findOption

	// state
	abort bool

	cursor         string
	nocursor       string
	cursorPosition int

	promptWidth int

	selectedPrefix   string
	unselectedPrefix string

	matchesStyle           lipgloss.Style
	cursorLineStyle        lipgloss.Style
	cursorLineMatchesStyle lipgloss.Style

	matches matches
	choices []int

	// window
	windowWidth     int
	windowHeight    int
	windowYPosition int

	// components
	input textinput.Model
}

func newModel(fzf *FZF, items *items, opt *findOption) *model {
	input := textinput.New()
	input.Prompt = fzf.option.prompt
	input.Placeholder = fzf.option.inputPlaceholder
	input.Focus()

	if !fzf.multiple() {
		fzf.option.keymap.Toggle.SetEnabled(false)
	}

	var matches matches
	for i := 0; i < items.Len(); i++ {
		matches = append(matches, match{
			Str:   items.String(i),
			Index: i,
		})
	}

	return &model{
		fzf:        fzf,
		items:      items,
		findOption: opt,
		// state
		abort: false,

		cursor:         fzf.option.styles.option.cursor.Render(fzf.option.cursor),
		nocursor:       strings.Repeat(" ", lipgloss.Width(fzf.option.cursor)),
		cursorPosition: 0,

		promptWidth: lipgloss.Width(fzf.option.prompt),

		selectedPrefix:   fzf.option.styles.option.selectedPrefix.Render(fzf.option.selectedPrefix),
		unselectedPrefix: fzf.option.styles.option.unselectedPrefix.Render(fzf.option.unselectedPrefix),

		matchesStyle:           fzf.option.styles.option.matches,
		cursorLineStyle:        fzf.option.styles.option.cursorLine,
		cursorLineMatchesStyle: lipgloss.NewStyle().Inherit(fzf.option.styles.option.matches).Inherit(fzf.option.styles.option.cursorLine),

		matches: matches,
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
	var v strings.Builder

	_, _ = v.WriteString(m.headerView())
	_, _ = v.WriteRune('\n')
	_, _ = v.WriteString(m.itemsView())

	return v.String()
}

func (m *model) headerView() string {
	var v strings.Builder

	// input
	_, _ = v.WriteString(m.input.View())
	// count
	if m.fzf.option.countViewEnabled {
		_, _ = v.WriteRune('\n')
		_, _ = v.WriteString(m.fzf.option.countViewFunc(m.items.Len(), len(m.matches), m.windowWidth))
	}

	return v.String()
}

func (m *model) headerHeight() int {
	return lipgloss.Height(m.headerView())
}

func (m *model) itemsView() string {
	var v strings.Builder

	headerHeight := m.headerHeight()

	for i, match := range m.matches[m.windowYPosition:] {
		cursorLine := m.cursorPosition == i+m.windowYPosition

		// write cursor
		if cursorLine {
			_, _ = v.WriteString(m.cursor)
		} else {
			_, _ = v.WriteString(m.nocursor)
		}

		// write toggle
		if m.fzf.multiple() {
			if intContains(m.choices, match.Index) {
				_, _ = v.WriteString(m.selectedPrefix)
			} else {
				_, _ = v.WriteString(m.unselectedPrefix)
			}
		}

		// write item prefix
		if m.findOption.itemPrefixFunc != nil {
			_, _ = v.WriteString(stringLinesToSpace(m.findOption.itemPrefixFunc(match.Index)))
		}

		// write item
		for ci, c := range match.Str {
			// matches
			if intContains(match.MatchedIndexes, ci) {
				if cursorLine {
					_, _ = v.WriteString(m.cursorLineMatchesStyle.Render(string(c)))
				} else {
					_, _ = v.WriteString(m.matchesStyle.Render(string(c)))
				}
			} else if cursorLine {
				_, _ = v.WriteString(m.cursorLineStyle.Render(string(c)))
			} else {
				_, _ = v.WriteRune(c)
			}
		}

		if i+1 >= m.windowHeight-headerHeight {
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
			m.fixYPosition()
			m.fixCursor()
		case key.Matches(msg, m.fzf.option.keymap.Down):
			// down
			m.cursorDown()
			m.fixYPosition()
			m.fixCursor()
		}
	case tea.WindowSizeMsg:
		// window
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		m.input.Width = m.windowWidth - m.promptWidth
		m.fixYPosition()
		m.fixCursor()
	}

	var cmds []tea.Cmd
	beforeValue := m.input.Value()
	{
		input, cmd := m.input.Update(msg)
		m.input = input
		cmds = append(cmds, cmd)
	}
	if beforeValue != m.input.Value() {
		m.filter()
		m.fixYPosition()
		m.fixCursor()
	}

	return m, tea.Batch(cmds...)
}

func (m *model) choice() {
	if len(m.choices) > 0 {
		return
	}

	if len(m.matches) == 0 {
		return
	}

	m.choices = append(m.choices, m.matches[m.cursorPosition].Index)
}

func (m *model) toggle() {
	if len(m.matches) == 0 {
		return
	}

	match := m.matches[m.cursorPosition]
	if intContains(m.choices, match.Index) {
		m.choices = intFilter(m.choices, func(i int) bool { return i != match.Index })
	} else {
		if m.fzf.option.noLimit || len(m.choices) < m.fzf.option.limit {
			m.choices = append(m.choices, match.Index)
		}
	}
}

func (m *model) cursorUp() {
	if m.cursorPosition > 0 {
		m.cursorPosition--
	}
}

func (m *model) cursorDown() {
	if m.cursorPosition+1 < len(m.matches) {
		m.cursorPosition++
	}
}

func (m *model) filter() {
	s := m.input.Value()
	if s == "" {
		var matches matches
		for i := 0; i < m.items.Len(); i++ {
			matches = append(matches, match{
				Str:   m.items.String(i),
				Index: i,
			})
		}
		m.matches = matches
		return
	}

	m.matches = fuzzySearch(m.items, s)
}

func (m *model) fixCursor() {
	if m.cursorPosition < 0 {
		m.cursorPosition = 0
		return
	}

	if m.cursorPosition+1 > len(m.matches) {
		m.cursorPosition = max(len(m.matches)-1, 0)
		return
	}
}

func (m *model) fixYPosition() {
	headerHeight := m.headerHeight()

	if m.windowHeight-headerHeight > len(m.matches) {
		m.windowYPosition = 0
		return
	}

	if m.cursorPosition < m.windowYPosition {
		m.windowYPosition = m.cursorPosition
		return
	}

	if m.cursorPosition+1 >= (m.windowHeight-headerHeight)+m.windowYPosition {
		m.windowYPosition = max(m.cursorPosition+1-(m.windowHeight-headerHeight), 0)
		return
	}
}
