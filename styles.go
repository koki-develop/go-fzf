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
	Blink           bool
	Italic          bool
	Strikethrough   bool
	Underline       bool
	Faint           bool
}

type Styles struct {
	option *stylesOption
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
	if s.Blink {
		style = style.Blink(true)
	}
	if s.Italic {
		style = style.Italic(true)
	}
	if s.Strikethrough {
		style = style.Strikethrough(true)
	}
	if s.Underline {
		style = style.Underline(true)
	}
	if s.Faint {
		style = style.Faint(true)
	}

	return style
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
