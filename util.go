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

func intContains(is []int, i int) bool {
	for _, l := range is {
		if l == i {
			return true
		}
	}
	return false
}
