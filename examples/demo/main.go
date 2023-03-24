package main

import (
	"fmt"
	"log"

	"github.com/koki-develop/go-fzf"
)

type Items []string

func (items Items) ItemString(i int) string {
	return items[i]
}

func (items Items) Len() int {
	return len(items)
}

func main() {
	items := Items{"hello", "world", "foo", "bar"}

	fzf := fzf.New(
		fzf.WithStyles(
			fzf.WithStyleCursor(fzf.Style{ForegroundColor: "#ff0000"}),
			fzf.WithStyleMatches(fzf.Style{ForegroundColor: "#00ff00"}),
		),
	)
	idxs, err := fzf.Find(items)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		fmt.Println(items[i])
	}
}
