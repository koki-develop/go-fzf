<h1 align="center">go-fzf</h1>

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
<a href="./README.md">English</a> | 日本語
</p>

## 目次

- [使い方](#使い方)
  - [CLI](#cli-として使用する)
  - [ライブラリ](#ライブラリとして使用する)
- [ライセンス](#ライセンス)

## 使い方

### CLI として使用する

go-fzf で何ができるのかを知りたい場合は `gofzf` CLI を試してください。  
`gofzf` CLI は go-fzf で作られており、 go-fzf のほとんどの機能を利用可能です。

![](/docs/cli/demo.gif)

詳しい情報は[ドキュメント](./docs/cli/README.ja.md)をご参照ください。

### ライブラリとして使用する

go-fzf を使用するとカスタマイズ性の高い Fuzzy Finder を簡単に作ることができます。  
例えば、たったこれだけのコードで以下のような Fuzzy Finder を作ることができます。

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

詳しい情報は[ドキュメント](./docs/library/README.ja.md)をご参照ください。

#### 使用例

[examples](./examples/) には go-fzf の使い方の様々な例が用意されています。

## ライセンス

[MIT](./LICENSE)
