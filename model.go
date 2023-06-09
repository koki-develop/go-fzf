package fzf

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
	"github.com/muesli/termenv"
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
	ellipsisStyle          lipgloss.Style

	matches Matches
	choices []int

	// window
	windowWidth     int
	windowHeight    int
	windowYPosition int

	mainViewWidth      int
	previewWindowWidth int

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
	lipgloss.SetColorProfile(termenv.TrueColor)

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
		ellipsisStyle:          lipgloss.NewStyle().Faint(true),

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
	runewidth.DefaultCondition.EastAsianWidth = false

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

	views := []string{m.mainView()}
	if m.findOption.previewWindowFunc != nil {
		views = append(views, m.previewWindowView())
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, views...)
}

func (m *model) mainView() string {
	rows := make([]string, 2)

	windowStyle := lipgloss.NewStyle().Height(m.windowHeight).Width(m.mainViewWidth)
	switch m.option.inputPosition {
	case InputPositionTop:
		windowStyle = windowStyle.AlignVertical(lipgloss.Top)
		rows[0] = m.inputView()
		rows[1] = m.itemsView()
	case InputPositionBottom:
		windowStyle = windowStyle.AlignVertical(lipgloss.Bottom)
		rows[0] = m.itemsView()
		rows[1] = m.inputView()
	}

	return windowStyle.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
}

func (m *model) inputView() string {
	rows := []string{}

	countView := ""
	countViewEnabled := m.option.countViewEnabled && m.option.countViewFunc != nil
	if countViewEnabled {
		countView = m.option.countViewFunc(CountViewMeta{
			ItemsCount:    m.items.Len(),
			MatchesCount:  len(m.matches),
			SelectedCount: len(m.choices),
			Width:         m.mainViewWidth,
			Limit:         m.option.limit,
			NoLimit:       m.option.noLimit,
		})
	}

	switch m.option.inputPosition {
	case InputPositionTop:
		rows = append(rows, m.input.View())
		if countViewEnabled {
			rows = append(rows, countView)
		}
	case InputPositionBottom:
		if countViewEnabled {
			rows = append(rows, countView)
		}
		rows = append(rows, m.input.View())
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m *model) inputHeight() int {
	return lipgloss.Height(m.inputView())
}

func (m *model) itemsHeight() int {
	return min(m.windowHeight-m.inputHeight(), len(m.matches))
}

func (m *model) itemsView() string {
	itemsHeight := m.itemsHeight()
	if itemsHeight < 1 {
		return ""
	}
	matches := m.matches[m.windowYPosition : itemsHeight+m.windowYPosition]
	rows := make([]string, len(matches))
	switch m.option.inputPosition {
	case InputPositionTop:
		for i, match := range matches {
			cursorLine := m.cursorPosition == (i + m.windowYPosition)
			rows[i] = m.itemView(match, cursorLine)
		}
	case InputPositionBottom:
		for i, match := range matches {
			cursorLine := m.cursorPosition == (i + m.windowYPosition)
			rows[len(matches)-1-i] = m.itemView(match, cursorLine)
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
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

	// truncate string
	// TODO: refactor
	maxItemWidth := m.mainViewWidth - lipgloss.Width(v.String())
	from := 0
	over := 5
	ellipsis := ".."
	if maxItemWidth <= lipgloss.Width(ellipsis)*2 {
		ellipsis = "."
	}
	if maxItemWidth < lipgloss.Width(ellipsis)*2 {
		return v.String()
	}
	leftDots, rightDots := false, false
	str := match.Str
	runes := []rune(match.Str)

	// if string width exceeds maxWidth, truncate
	if lipgloss.Width(match.Str) > maxItemWidth {
		if len(match.MatchedIndexes) == 0 {
			// if not filtered, truncate the right
			rightDots = true
			for lipgloss.Width(str)+2 > maxItemWidth {
				runes := []rune(str)
				str = string(runes[:len(runes)-1])
			}
		} else {
			// if filtered

			lastMatchedIndex := match.MatchedIndexes[len(match.MatchedIndexes)-1]
			if lipgloss.Width(string(runes[:min(lastMatchedIndex+1+over, len(runes))])+ellipsis) <= maxItemWidth {
				// if width from the beginning to index within maxWidth, truncate only the right
				rightDots = true
				for lipgloss.Width(str+ellipsis) > maxItemWidth {
					runes := []rune(str)
					str = string(runes[:len(runes)-1])
				}
			} else {
				// if width from the beginning to index not within maxWidth, truncate the left
				leftDots = true
				if lipgloss.Width(string(runes[min(lastMatchedIndex+1+over, len(runes)-1):])) > lipgloss.Width(ellipsis) {
					// if the right also not within, truncate
					rightDots = true
					for lipgloss.Width(string([]rune(str)[lastMatchedIndex+1+over:])+ellipsis) > lipgloss.Width(ellipsis) {
						runes := []rune(str)
						str = string(runes[:len(runes)-1])
					}

					// truncate the left
					for lipgloss.Width(str+ellipsis+ellipsis) > maxItemWidth {
						runes := []rune(str)
						str = string(runes[1:])
						from++
					}
				} else {
					// truncate the left
					for lipgloss.Width(str+ellipsis) > maxItemWidth {
						runes := []rune(str)
						str = string(runes[1:])
						from++
					}
				}
			}
		}
	}

	if leftDots {
		_, _ = v.WriteString(m.ellipsisStyle.Render(ellipsis))
	}

	// write item
	for ci, c := range []rune(str) {
		var s string
		// matches
		if intContains(match.MatchedIndexes, ci+from) {
			if cursorLine {
				s = m.cursorLineMatchesStyle.Render(string(c))
			} else {
				s = m.matchesStyle.Render(string(c))
			}
		} else if cursorLine {
			s = m.cursorLineStyle.Render(string(c))
		} else {
			s = string(c)
		}

		_, _ = v.WriteString(s)
	}

	if rightDots {
		_, _ = v.WriteString(m.ellipsisStyle.Render(ellipsis))
	}

	return v.String()
}

func (m *model) previewWindowView() string {
	v := ""
	if len(m.matches) > 0 {
		v = m.findOption.previewWindowFunc(m.matches[m.cursorPosition].Index, m.previewWindowWidth, m.windowHeight)
	}

	return lipgloss.NewStyle().
		Width(m.previewWindowWidth).
		Height(m.windowHeight).
		BorderStyle(lipgloss.NormalBorder()).
		BorderLeft(true).
		Render(v)
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
			m.fixYPosition()
			m.fixCursor()
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
		m.fixYPosition()
		m.fixCursor()
		m.fixWidth()
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
		if len(m.matches) > 0 || !strings.HasPrefix(m.input.Value(), beforeValue) {
			m.filter()
			m.fixYPosition()
			m.fixCursor()
		}
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

	switch m.option.inputPosition {
	case InputPositionTop:
		m.cursorDown()
	case InputPositionBottom:
		m.cursorUp()
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

func (m *model) fixWidth() {
	m.mainViewWidth = m.windowWidth
	if m.findOption.previewWindowFunc != nil {
		m.mainViewWidth /= 2
		m.previewWindowWidth = m.windowWidth - m.mainViewWidth
	}
	m.input.Width = m.mainViewWidth - m.promptWidth - 1
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
