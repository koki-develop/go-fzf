package fzf

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	_ tea.Model = (*model)(nil)
)

type model struct {
	items      *items
	itemsLen   int
	option     *option
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

	matches Matches
	choices []int

	// window
	windowWidth     int
	windowHeight    int
	windowYPosition int

	// components
	input textinput.Model
}

func newModel(opt *option) *model {
	input := textinput.New()
	input.Prompt = opt.prompt
	input.PromptStyle = opt.styles.option.prompt
	input.Placeholder = opt.inputPlaceholder
	input.PlaceholderStyle = opt.styles.option.inputPlaceholder
	input.TextStyle = opt.styles.option.inputText
	input.Focus()

	if !opt.multiple() {
		opt.keymap.Toggle.SetEnabled(false)
	}

	return &model{
		option: opt,
		// state
		abort: false,

		cursor:         opt.styles.option.cursor.Render(opt.cursor),
		nocursor:       strings.Repeat(" ", lipgloss.Width(opt.cursor)),
		cursorPosition: 0,

		promptWidth: lipgloss.Width(opt.prompt),

		selectedPrefix:   opt.styles.option.selectedPrefix.Render(opt.selectedPrefix),
		unselectedPrefix: opt.styles.option.unselectedPrefix.Render(opt.unselectedPrefix),

		matchesStyle:           opt.styles.option.matches,
		cursorLineStyle:        opt.styles.option.cursorLine,
		cursorLineMatchesStyle: lipgloss.NewStyle().Inherit(opt.styles.option.matches).Inherit(opt.styles.option.cursorLine),

		choices: []int{},
		// window
		windowWidth:     0,
		windowHeight:    0,
		windowYPosition: 0,
		// components
		input: input,
	}
}

func (m *model) loadItems(items *items) {
	m.items = items
	m.itemsLen = items.Len()
	m.filter()
}

func (m *model) setFindOption(findOption *findOption) {
	m.findOption = findOption
}

func (m *model) Init() tea.Cmd {
	cmds := []tea.Cmd{
		textinput.Blink,
		tea.EnterAltScreen,
	}
	if m.option.hotReloadLocker != nil {
		cmds = append(cmds, m.watchReload())
	}

	return tea.Batch(
		cmds...,
	)
}

/*
 * view
 */

func (m *model) View() string {
	if m.option.hotReloadLocker != nil {
		m.option.hotReloadLocker.Lock()
		defer m.option.hotReloadLocker.Unlock()
	}

	var v strings.Builder

	var windowStyle lipgloss.Style
	switch m.option.inputPosition {
	case InputPositionTop:
		windowStyle = lipgloss.NewStyle().Height(m.windowHeight).AlignVertical(lipgloss.Top)
		_, _ = v.WriteString(m.inputView())
		_, _ = v.WriteRune('\n')
		_, _ = v.WriteString(m.itemsView())

	case InputPositionBottom:
		windowStyle = lipgloss.NewStyle().Height(m.windowHeight).AlignVertical(lipgloss.Bottom)
		_, _ = v.WriteString(m.itemsView())
		_, _ = v.WriteRune('\n')
		_, _ = v.WriteString(m.inputView())
	}

	return windowStyle.Render(v.String())
}

func (m *model) inputView() string {
	var v strings.Builder

	switch m.option.inputPosition {
	case InputPositionTop:
		// input
		_, _ = v.WriteString(m.input.View())
		// count
		if m.option.countViewEnabled {
			_, _ = v.WriteRune('\n')
			_, _ = v.WriteString(m.option.countViewFunc(CountViewMeta{
				ItemsCount:    m.items.Len(),
				MatchesCount:  len(m.matches),
				SelectedCount: len(m.choices),
				WindowWidth:   m.windowWidth,
				Limit:         m.option.limit,
				NoLimit:       m.option.noLimit,
			}))
		}

	case InputPositionBottom:
		// count
		if m.option.countViewEnabled {
			_, _ = v.WriteString(m.option.countViewFunc(CountViewMeta{
				ItemsCount:    m.items.Len(),
				MatchesCount:  len(m.matches),
				SelectedCount: len(m.choices),
				WindowWidth:   m.windowWidth,
				Limit:         m.option.limit,
				NoLimit:       m.option.noLimit,
			}))
			_, _ = v.WriteRune('\n')
		}
		// input
		_, _ = v.WriteString(m.input.View())
	}

	return v.String()
}

func (m *model) inputHeight() int {
	return lipgloss.Height(m.inputView())
}

func (m *model) itemsView() string {
	var v strings.Builder

	inputHeight := m.inputHeight()

	switch m.option.inputPosition {
	case InputPositionTop:
		for i, match := range m.matches {
			if i < m.windowYPosition {
				continue
			}

			cursorLine := m.cursorPosition == i
			v.WriteString(m.itemView(match, cursorLine))
			if i+1-m.windowYPosition >= m.windowHeight-inputHeight {
				break
			}
			v.WriteRune('\n')
		}
	case InputPositionBottom:
		for i := len(m.matches) - 1; i >= 0; i-- {
			if len(m.matches)-i+m.windowHeight-inputHeight < m.windowYPosition {
				continue
			}

			cursorLine := m.cursorPosition == i
			v.WriteString(m.itemView(m.matches[i], cursorLine))
			if i-1 < m.windowYPosition {
				break
			}
			v.WriteRune('\n')
		}
	}

	return v.String()
}

func (m *model) itemView(match Match, cursorLine bool) string {
	var v strings.Builder

	// write cursor
	if cursorLine {
		_, _ = v.WriteString(m.cursor)
	} else {
		_, _ = v.WriteString(m.nocursor)
	}

	// write toggle
	if m.option.multiple() {
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

	return v.String()
}

/*
 * update
 */

type watchReloadMsg struct{}
type forceReloadMsg struct{}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.option.hotReloadLocker != nil {
		m.option.hotReloadLocker.Lock()
		defer m.option.hotReloadLocker.Unlock()
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// key
		switch {
		case key.Matches(msg, m.option.keymap.Abort):
			// abort
			m.abort = true
			return m, tea.Quit
		case key.Matches(msg, m.option.keymap.Choose):
			// choose
			m.choice()
			return m, tea.Quit
		case key.Matches(msg, m.option.keymap.Toggle):
			// toggle
			m.toggle()
		case key.Matches(msg, m.option.keymap.Up):
			// up
			switch m.option.inputPosition {
			case InputPositionTop:
				m.cursorUp()
			case InputPositionBottom:
				m.cursorDown()
			}
			m.fixYPosition()
			m.fixCursor()
		case key.Matches(msg, m.option.keymap.Down):
			// down
			switch m.option.inputPosition {
			case InputPositionTop:
				m.cursorDown()
			case InputPositionBottom:
				m.cursorUp()
			}
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
	case watchReloadMsg:
		// watch reload
		return m, m.watchReload()
	case forceReloadMsg:
		// force reload
		m.forceReload()
		return m, nil
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
		if m.option.noLimit || len(m.choices) < m.option.limit {
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
		var matches Matches
		for i := 0; i < m.items.Len(); i++ {
			matches = append(matches, Match{
				Str:   m.items.ItemString(i),
				Index: i,
			})
		}
		m.matches = matches
		return
	}

	m.matches = Search(m.items, s, WithSearchCaseSensitive(m.option.caseSensitive))
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
	inputHeight := m.inputHeight()

	if m.windowHeight-inputHeight > len(m.matches) {
		m.windowYPosition = 0
		return
	}

	if m.cursorPosition < m.windowYPosition {
		m.windowYPosition = m.cursorPosition
		return
	}

	if m.cursorPosition+1 >= (m.windowHeight-inputHeight)+m.windowYPosition {
		m.windowYPosition = max(m.cursorPosition+1-(m.windowHeight-inputHeight), 0)
		return
	}
}

func (m *model) forceReload() {
	m.loadItems(m.items)
}

func (m *model) watchReload() tea.Cmd {
	return tea.Tick(30*time.Millisecond, func(_ time.Time) tea.Msg {
		if m.itemsLen != m.items.Len() {
			m.loadItems(m.items)
		}

		return watchReloadMsg{}
	})
}
