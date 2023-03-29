package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/koki-develop/go-fzf"
)

func main() {
	var items []string
	var mu sync.Mutex

	go func() {
		i := 0
		for {
			time.Sleep(50 * time.Millisecond)
			mu.Lock()
			items = append(items, strconv.Itoa(i))
			mu.Unlock()
			i++
		}
	}()

	f, err := fzf.New(fzf.WithHotReload(&mu))
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(&items, func(i int) string { return items[i] })
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		fmt.Println(items[i])
	}
}
