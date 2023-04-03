<h1 align="center">🔍 go-fzf</h1>

<p align="center">
<a href="https://pkg.go.dev/github.com/koki-develop/go-fzf"><img src="https://pkg.go.dev/badge/github.com/koki-develop/go-fzf.svg" alt="Go Reference"></a>
<a href="https://github.com/koki-develop/go-fzf/releases/latest"><img src="https://img.shields.io/github/v/release/koki-develop/go-fzf?style=flat-square" alt="GitHub release (latest by date)"></a>
<a href="https://github.com/koki-develop/go-fzf/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/koki-develop/go-fzf/ci.yml?logo=github&amp;style=flat-square" alt="GitHub Workflow Status"></a>
<a href="https://codeclimate.com/github/koki-develop/go-fzf/maintainability"><img src="https://img.shields.io/codeclimate/maintainability/koki-develop/go-fzf?style=flat-square&amp;logo=codeclimate" alt="Maintainability"></a>
<a href="https://goreportcard.com/report/github.com/koki-develop/go-fzf"><img src="https://goreportcard.com/badge/github.com/koki-develop/go-fzf?style=flat-square" alt="Go Report Card"></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/go-fzf?style=flat-square" alt="LICENSE"></a>
</p>

<p align="center">
Fuzzy Finder CLI and Library.
</p>

<p align="center">
English | <a href="./README.ja.md">日本語</a>
</p>

## Contents

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

![](./examples/basic/demo.gif)

For more information, please refer to the [documentation](./docs/library/README.md)

#### Examples

Various examples of how to use go-fzf are available in the [examples](./examples/) directory.

## LICENSE

[MIT](./LICENSE)
