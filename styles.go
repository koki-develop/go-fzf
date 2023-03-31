package fzf

import "github.com/charmbracelet/lipgloss"

var (
	defaultColor        = "#00ADD8"
	defaultStylesOption = stylesOption{
		prompt:           lipgloss.NewStyle(),
		cursor:           lipgloss.NewStyle().Foreground(lipgloss.Color(defaultColor)),
		cursorLine:       lipgloss.NewStyle().Bold(true),
		matches:          lipgloss.NewStyle().Foreground(lipgloss.Color(defaultColor)),
		selectedPrefix:   lipgloss.NewStyle().Foreground(lipgloss.Color(defaultColor)),
		unselectedPrefix: lipgloss.NewStyle().Faint(true),
	}
)

// Style represents a style.
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

// Styles is the styles for each component.
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

// Option represents a option for the styles.
type StylesOption func(o *stylesOption)

type stylesOption struct {
	prompt           lipgloss.Style
	cursor           lipgloss.Style
	cursorLine       lipgloss.Style
	selectedPrefix   lipgloss.Style
	unselectedPrefix lipgloss.Style
	matches          lipgloss.Style
}

// NewStyles returns a new styles.
func NewStyles(opts ...StylesOption) *Styles {
	o := defaultStylesOption
	for _, opt := range opts {
		opt(&o)
	}
	return &Styles{option: &o}
}

// WithStylePrompt sets the style of prompt.
func WithStylePrompt(s Style) StylesOption {
	return func(o *stylesOption) {
		o.prompt = s.lipgloss()
	}
}

// WithStyleCursor sets the style of cursor.
func WithStyleCursor(s Style) StylesOption {
	return func(o *stylesOption) {
		o.cursor = s.lipgloss()
	}
}

// WithStyleCursorLine sets the style of cursor line.
func WithStyleCursorLine(s Style) StylesOption {
	return func(o *stylesOption) {
		o.cursorLine = s.lipgloss()
	}
}

// WithStyleSelectedPrefix sets the style of prefix of the selected item.
func WithStyleSelectedPrefix(s Style) StylesOption {
	return func(o *stylesOption) {
		o.selectedPrefix = s.lipgloss()
	}
}

// WithStyleUnselectedPrefix sets the style of prefix of the unselected item.
func WithStyleUnselectedPrefix(s Style) StylesOption {
	return func(o *stylesOption) {
		o.unselectedPrefix = s.lipgloss()
	}
}

// WithStyleMatches sets the style of the matched characters.
func WithStyleMatches(s Style) StylesOption {
	return func(o *stylesOption) {
		o.matches = s.lipgloss()
	}
}
