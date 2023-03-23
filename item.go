package fzf

type Items interface {
	ItemString(int) string
	Len() int
}
