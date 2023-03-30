package main

import (
	"fmt"
	"log"

	"github.com/koki-develop/go-fzf"
)

func main() {
	items := []string{"HELLO", "Hello", "hello"}

	f, err := fzf.New(fzf.WithCaseSensitive(true))
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
