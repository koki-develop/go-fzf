# go-fzf

[![Go Reference](https://pkg.go.dev/badge/github.com/koki-develop/go-fzf.svg)](https://pkg.go.dev/github.com/koki-develop/go-fzf)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/koki-develop/go-fzf?style=flat-square)](https://github.com/koki-develop/go-fzf/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/koki-develop/go-fzf/ci.yml?logo=github&style=flat-square)](https://github.com/koki-develop/go-fzf/actions/workflows/ci.yml)
[![Maintainability](https://img.shields.io/codeclimate/maintainability/koki-develop/go-fzf?style=flat-square&logo=codeclimate)](https://codeclimate.com/github/koki-develop/go-fzf/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/koki-develop/go-fzf?style=flat-square)](https://goreportcard.com/report/github.com/koki-develop/go-fzf)
[![LICENSE](https://img.shields.io/github/license/koki-develop/go-fzf?style=flat-square)](./LICENSE)

Fuzzy Finder CLI and Library.

- [Usage](#usage)
  - [CLI](#using-as-a-cli)
  - [Library](#using-as-a-library)
- [LICENSE](#license)

## Usage

### Using as a CLI

If you want to know what you can do with go-fzf, try the `gofzf` CLI.  
The `gofzf` CLI is built with go-fzf and can utilize most of go-fzf's features.

![](/docs/cli/demo.gif)

For more information, please refer to the [documentation](./docs/cli/README.md).

### Using as a library

With go-fzf, you can easily create a highly customizable Fuzzy Finder.  
For example, you can create a Fuzzy Finder like the one below with just this simple code:

```go
package main

import (
	"fmt"
	"log"

	"github.com/koki-develop/go-fzf"
)

func main() {
	items := []string{"hello", "world", "foo", "bar"}

	f, _ := fzf.New()
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
```

![](./docs/library/demo.gif)

For more information, please refer to the [documentation](./docs/library/README.md)  
Additionally, various samples of how to use go-fzf are available in the [examples](./examples/) directory.

## LICENSE

[MIT](./LICENSE)
