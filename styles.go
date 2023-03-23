package fzf

import "github.com/charmbracelet/lipgloss"

type Style struct {
	ForegroundColor string
	BackgroundColor string
	Bold            bool
}

type Styles struct {
	Matches    *Style
	CursorLine *Style
}

func (s *Style) style() lipgloss.Style {
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
