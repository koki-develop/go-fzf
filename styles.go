package fzf

import "github.com/charmbracelet/lipgloss"

type Style struct {
	ForegroundColor string
	BackgroundColor string
	Bold            bool
}

type Styles struct {
	Cursor     *Style
	CursorLine *Style
	Matches    *Style
}

type styles struct {
	Cursor     lipgloss.Style
	CursorLine lipgloss.Style
	Matches    lipgloss.Style
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
