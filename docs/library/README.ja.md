# ライブラリとして使用する

## 目次

- [インストール](#インストール)
- [使い方](#使い方)

## インストール

まず `go get` で最新バージョンをインストールしてください。

```console
$ go get -u github.com/koki-develop/go-fzf
```

次に go-fzf を import してください。

```go
import "github.com/koki-develop/go-fzf"
```

## 使い方

- [基本](#基本)
- [複数選択](#複数選択)
- [大文字 / 小文字を区別する](#大文字--小文字を区別する)
- [キーマップ](#キーマップ)
- [ホットリロード](#ホットリロード)
- [見た目をカスタマイズする](#見た目をカスタマイズする)

### 基本

まず `fzf.New()` で Fuzzy Finder を初期化します。

```go
f, err := fzf.New()
if err != nil {
  // ...
}
```

次に `Find()` を実行すると Fuzzy Finder が起動します。  
第一引数には検索対象のアイテムの slice を渡します。  
第二引数には i 番目のアイテムのテキストを返す関数を渡します。

```go
idxs, err := f.Find(items, func(i int) string { return items[i] })
if err != nil {
  // ...
}
```

[Example](/examples/basic/)

### 複数選択

`fzf.WithLimit()` を使用すると選択できるアイテムの数を設定できます。

```go
f, err := fzf.New(fzf.WithLimit(4))
if err != nil {
  // ...
}
```

`fzf.WithNoLimit()` を使用するとアイテムを無制限に選択できるようになります。

```go
f, err := fzf.New(fzf.WithNoLimit(true)) // no limit
if err != nil {
  // ...
}
```

[Example](/examples/multiple/)

### 大文字 / 小文字を区別する

`fzf.WithCaseSensitive()` を使用するとあいまい検索で大文字/小文字を区別するようにできます。

```go
f, err := fzf.New(fzf.WithCaseSensitive(true))
if err != nil {
  // ...
}
```

[Example](/examples/case-sensitive/)

### キーマップ

`fzf.WithKeyMap()` を使用するとキーマップをカスタマイズできます。  
カスタマイズできるキーマップは次の通りです。

- `Up` - カーソルを上に移動する
- `Down` - カーソルを下に移動する
- `Toggle` - アイテムを選択もしくは選択解除する
- `Choose` - 検索を完了する
- `Abort` - 検索を中止する。

```go
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
  // ...
}
```

[Example](/examples/keymap/)

### ホットリロード

`fzf.WithHotReload()` を使用するとホットリロードを有効にできます。

```go
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
  // ...
}

// NOTE: slice のポインタを渡す必要があることに注意してください。
idxs, err := f.Find(&items, func(i int) string { return items[i] })
if err != nil {
  // ...
}
```

[Example](/examples/hotreload/)

### 見た目をカスタマイズする

- [プロンプト](#プロンプト)
- [カーソル](#カーソル)
- [選択中 / 未選択アイテムの接頭辞](#選択中--未選択アイテムの接頭辞)
- [インプットのプレースホルダ](#インプットのプレースホルダ)
- [カウントビュー](#カウントビュー)
- [スタイル](#スタイル)

#### プロンプト

`fzf.WithPrompt()` を使用するとプロンプトの文字列を設定できます。

```go
f, err := fzf.New(fzf.WithPrompt("=> "))
if err != nil {
  // ...
}
```

[Example](/examples/prompt/)

#### カーソル

`fzf.WithCursor()` を使用するとカーソルの文字列を設定できます。

```go
f, err := fzf.New(fzf.WithCursor("=> "))
if err != nil {
  // ...
}
```

[Example](/examples/cursor/)

#### 選択中 / 未選択アイテムの接頭辞

`fzf.WithSelectedPrefix()` を使用すると選択中アイテムの接頭辞を設定できます。  
同じ様に `fzf.WithUnselectedPrefix()` を使用すると未選択アイテムの接頭辞を設定できます。

```go
f, err := fzf.New(
  fzf.WithSelectedPrefix("[x] "),
  fzf.WithUnselectedPrefix("[ ] "),
)
if err != nil {
  // ...
}
```

[Example](/examples/prefix/)

#### インプットのプレースホルダ

`fzf.WithInputPlaceholder()` を使用するとインプットのプレースホルダを設定できます。

```go
f, err := fzf.New(fzf.WithInputPlaceholder("Search..."))
if err != nil {
  // ...
}
```

[Example](/examples/placeholder/)

#### カウントビュー

`fzf.WithCountViewEnabled()` を使用するとカウントビューの有効/無効を切り替えることができます ( デフォルトは有効 ) 。  
`fzf.WithCountView()` を使用するとカウントビューをレンダリングする関数を設定することができます。  
関数の引数にはカウントビューに必要な情報が含まれる `fzf.CountViewMeta` 構造体が渡されます。

```go
f, err := fzf.New(
  fzf.WithNoLimit(true),
  fzf.WithCountViewEnabled(true),
  fzf.WithCountView(func(meta fzf.CountViewMeta) string {
    return fmt.Sprintf("items: %d, selected: %d", meta.ItemsCount, meta.SelectedCount)
  }),
)
if err != nil {
  // ...
}
```

[Example](/examples/countview/)

#### スタイル

`fzf.WithStyles()` を使用すると各コンポーネントのスタイルを設定することができます。  
使用できるスタイルについては[リファレンス](https://pkg.go.dev/github.com/koki-develop/go-fzf#Style)をご参照ください。

```go
f, err := fzf.New(
  fzf.WithNoLimit(true),
  fzf.WithStyles(
    fzf.WithStylePrompt(fzf.Style{Faint: true}),                                       // プロンプト
    fzf.WithStyleInputPlaceholder(fzf.Style{Faint: true, ForegroundColor: "#ff0000"}), // インプットのプレースホルダ
    fzf.WithStyleInputText(fzf.Style{Italic: true}),                                   // インプットのテキスト
    fzf.WithStyleCursor(fzf.Style{Bold: true}),                                        // カーソル
    fzf.WithStyleCursorLine(fzf.Style{Bold: true}),                                    // カーソル行
    fzf.WithStyleMatches(fzf.Style{ForegroundColor: "#ff0000"}),                       // 一致文字
    fzf.WithStyleSelectedPrefix(fzf.Style{ForegroundColor: "#ff0000"}),                // 選択中アイテムの接頭辞
    fzf.WithStyleUnselectedPrefix(fzf.Style{Faint: true}),                             // 未選択のアイテムの接頭辞
  ),
)
if err != nil {
  // ...
}
```

[Example](/examples/styles/)
