package fzf

import "github.com/charmbracelet/lipgloss"

var (
	defaultStylesOption = &stylesOption{
		cursor:           lipgloss.NewStyle(),
		cursorLine:       lipgloss.NewStyle(),
		matches:          lipgloss.NewStyle(),
		selectedPrefix:   lipgloss.NewStyle(),
		unselectedPrefix: lipgloss.NewStyle(),
	}
)

type Style struct {
	ForegroundColor string
	BackgroundColor string
	Bold            bool
}

type Styles struct {
	option *stylesOption
}

type StylesOption func(o *stylesOption)

type stylesOption struct {
	cursor           lipgloss.Style
	cursorLine       lipgloss.Style
	selectedPrefix   lipgloss.Style
	unselectedPrefix lipgloss.Style
	matches          lipgloss.Style
}

func NewStyles(opts ...StylesOption) *Styles {
	o := defaultStylesOption
	for _, opt := range opts {
		opt(o)
	}
	return &Styles{option: o}
}

func WithStyleCursor(s Style) StylesOption {
	return func(o *stylesOption) {
		o.cursor = s.lipgloss()
	}
}

func WithStyleCursorLine(s Style) StylesOption {
	return func(o *stylesOption) {
		o.cursorLine = s.lipgloss()
	}
}

func WithStyleSelectedPrefix(s Style) StylesOption {
	return func(o *stylesOption) {
		o.selectedPrefix = s.lipgloss()
	}
}

func WithStyleUnselectedPrefix(s Style) StylesOption {
	return func(o *stylesOption) {
		o.unselectedPrefix = s.lipgloss()
	}
}

func WithStyleMatches(s Style) StylesOption {
	return func(o *stylesOption) {
		o.matches = s.lipgloss()
	}
}

func (s *Style) lipgloss() lipgloss.Style {
	style := lipgloss.NewStyle()

	if s.ForegroundColor != "" {
		style = style.Foreground(lipgloss.Color(s.ForegroundColor))
	}
	if s.BackgroundColor != "" {
		style = style.Background(lipgloss.Color(s.BackgroundColor))
	}
	if s.Bold {
		style = style.Bold(true)
	}

	return style
}
