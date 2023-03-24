# go-fzf

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/koki-develop/go-fzf)](https://github.com/koki-develop/go-fzf/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/koki-develop/go-fzf/ci.yml?logo=github)](https://github.com/koki-develop/go-fzf/actions/workflows/ci.yml)
[![Maintainability](https://img.shields.io/codeclimate/maintainability/koki-develop/go-fzf?style=flat&logo=codeclimate)](https://codeclimate.com/github/koki-develop/go-fzf/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/koki-develop/go-fzf)](https://goreportcard.com/report/github.com/koki-develop/go-fzf)
[![LICENSE](https://img.shields.io/github/license/koki-develop/go-fzf)](./LICENSE)

Fuzzy Finder CLI and Library.

- [Usage](#usage)
  - [CLI](#using-as-a-cli)
  - [Library](#using-as-a-library)
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
