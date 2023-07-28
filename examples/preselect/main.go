package main

import (
	"fmt"
	"log"

	"github.com/koki-develop/go-fzf"
)

func main() {
	items := []string{"hello", "world", "foo", "bar"}

	f, err := fzf.New(fzf.WithNoLimit(true))
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(
		items,
		func(i int) string { return items[i] },
		fzf.WithPreselect([]int{0, 2}),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		fmt.Println(items[i])
	}
}
