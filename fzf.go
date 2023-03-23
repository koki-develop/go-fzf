package fzf

import "fmt"

type FZF struct {
	prompt string
}

func New(opts ...Option) *FZF {
	o := defaultOption

	for _, opt := range opts {
		opt(&o)
	}

	return &FZF{
		prompt: o.prompt,
	}
}

func (fzf *FZF) Find() (int, error) {
	fmt.Printf("%#v\n", fzf)
	return 0, nil
}
