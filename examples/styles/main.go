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
			fzf.WithStylePrompt(fzf.Style{Faint: true}),                                       // Prompt
			fzf.WithStyleInputPlaceholder(fzf.Style{Faint: true, ForegroundColor: "#ff0000"}), // Placeholder for input
			fzf.WithStyleCursor(fzf.Style{Bold: true}),                                        // Cursor
			fzf.WithStyleCursorLine(fzf.Style{Bold: true}),                                    // Cursor line
			fzf.WithStyleMatches(fzf.Style{ForegroundColor: "#ff0000"}),                       // Matched characters
			fzf.WithStyleSelectedPrefix(fzf.Style{ForegroundColor: "#ff0000"}),                // Prefix of selected items
			fzf.WithStyleUnselectedPrefix(fzf.Style{Faint: true}),                             // Prefix of unselected items
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
