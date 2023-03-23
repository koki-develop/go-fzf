package fzf

import (
	"regexp"
	"unicode/utf8"
)

var (
	ansiEscapeRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)
)

func stringLen(s string) int {
	return utf8.RuneCountInString(ansiEscapeRegex.ReplaceAllString(s, ""))
}
