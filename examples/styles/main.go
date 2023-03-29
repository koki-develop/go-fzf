package main

import (
	"fmt"
	"log"

	"github.com/koki-develop/go-fzf"
)

func main() {
	items := []string{"hello", "world", "foo", "bar"}

	f, err := fzf.New(
		fzf.WithNoLimit(true),
		fzf.WithStyles(
			fzf.WithStyleCursor(fzf.Style{Bold: true}),
			fzf.WithStyleCursorLine(fzf.Style{Bold: true}),
			fzf.WithStyleMatches(fzf.Style{ForegroundColor: "#ff0000"}),
			fzf.WithStyleSelectedPrefix(fzf.Style{ForegroundColor: "#ff0000"}),
			fzf.WithStyleUnselectedPrefix(fzf.Style{Faint: true}),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(items, func(i int) string { return items[i] })
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		fmt.Println(items[i])
	}
}
