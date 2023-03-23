package fzf

var defaultOption = option{
	prompt: "> ",
}

type option struct {
	prompt string
}

type Option func(o *option)

func WithPrompt(p string) Option {
	return func(o *option) {
		o.prompt = p
	}
}
