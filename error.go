package fzf

import "errors"

var (
	// ErrAbort is returned when a Fuzzy Finder is aborted.
	ErrAbort = errors.New("abort")
)
