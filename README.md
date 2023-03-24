# go-fzf

Fuzzy Finder CLI and Library.

- [Usage](#usage)
  - [Using as a CLI](#using-as-a-cli)
  - [Using as a Library](#using-as-a-library)
- [LICENSE](#license)

## Usage

### Using as a CLI

If you want to see what go-fzf can do for you, try the `gofzf` CLI.  

![](/docs/cli/demo.gif)

See [docs/cli/README.md](./docs/cli/README.md) for more information.

### Using as a Library

go-fzf makes it easy to create a highly customizable Fuzzy Finder.  
For example, the following Fuzzy Finder can be created with just this program.

```go
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
```

![](./docs/library/demo.gif)

See [docs/library/README.md](./docs/library/README.md) for more information.

## LICENSE

[MIT](./LICENSE)
