package main

import (
	"fmt"
	"log"
	"os"

	"github.com/koki-develop/go-fzf"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(
		files,
		func(i int) string { return files[i].Name() },
		fzf.WithPreviewWindow(func(i, width, height int) string {
			info, _ := files[i].Info()
			return fmt.Sprintf(
				"Name: %s\nModTime: %s\nSize: %d bytes",
				info.Name(), info.ModTime(), info.Size(),
			)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		fmt.Println(files[i].Name())
	}
}
