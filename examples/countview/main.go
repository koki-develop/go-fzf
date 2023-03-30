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
		fzf.WithCountViewEnabled(true),
		fzf.WithCountView(func(meta fzf.CountViewMeta) string {
			return fmt.Sprintf("items: %d, selected: %d", meta.ItemsCount, meta.SelectedCount)
		}),
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
