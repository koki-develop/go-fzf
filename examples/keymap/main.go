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
		fzf.WithKeyMap(fzf.KeyMap{
			Up:     []string{"up", "ctrl+b"},   // ↑, Ctrl+b
			Down:   []string{"down", "ctrl+f"}, // ↓, Ctrl+f
			Toggle: []string{"tab"},            // tab
			Choose: []string{"enter"},          // Enter
			Abort:  []string{"esc"},            // esc
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
