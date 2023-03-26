package main

import (
	"bufio"
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime/debug"
	"sync"

	"github.com/koki-develop/go-fzf"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const (
	mainColor = "#00ADD8"
)

var (
	version string
)

var (
	flagLimit         int
	flagNoLimit       bool
	flagCaseSensitive bool

	flagPrompt           string
	flagCursor           string
	flagSelectedPrefix   string
	flagUnselectedPrefix string
	flagInputPlaceholder string

	flagCountView bool

	flagCursorFg            string
	flagCursorBg            string
	flagCursorBold          bool
	flagCursorBlink         bool
	flagCursorItalic        bool
	flagCursorStrikethrough bool
	flagCursorUnderline     bool
	flagCursorFaint         bool

	flagCursorLineFg            string
	flagCursorLineBg            string
	flagCursorLineBold          bool
	flagCursorLineBlink         bool
	flagCursorLineItalic        bool
	flagCursorLineStrikethrough bool
	flagCursorLineUnderline     bool
	flagCursorLineFaint         bool

	flagSelectedPrefixFg            string
	flagSelectedPrefixBg            string
	flagSelectedPrefixBold          bool
	flagSelectedPrefixBlink         bool
	flagSelectedPrefixItalic        bool
	flagSelectedPrefixStrikethrough bool
	flagSelectedPrefixUnderline     bool
	flagSelectedPrefixFaint         bool

	flagUnselectedPrefixFg            string
	flagUnselectedPrefixBg            string
	flagUnselectedPrefixBold          bool
	flagUnselectedPrefixBlink         bool
	flagUnselectedPrefixItalic        bool
	flagUnselectedPrefixStrikethrough bool
	flagUnselectedPrefixUnderline     bool
	flagUnselectedPrefixFaint         bool

	flagMatchesFg            string
	flagMatchesBg            string
	flagMatchesBold          bool
	flagMatchesBlink         bool
	flagMatchesItalic        bool
	flagMatchesStrikethrough bool
	flagMatchesUnderline     bool
	flagMatchesFaint         bool
)

var rootCmd = &cobra.Command{
	Use:          "gofzf",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var items []string
		var mu sync.RWMutex

		f := fzf.New(
			fzf.WithNoLimit(flagNoLimit),
			fzf.WithLimit(flagLimit),
			fzf.WithCaseSensitive(flagCaseSensitive),

			fzf.WithHotReload(mu.RLocker()),

			fzf.WithPrompt(flagPrompt),
			fzf.WithCursor(flagCursor),
			fzf.WithSelectedPrefix(flagSelectedPrefix),
			fzf.WithUnselectedPrefix(flagUnselectedPrefix),
			fzf.WithInputPlaceholder(flagInputPlaceholder),

			fzf.WithCountViewEnabled(flagCountView),

			fzf.WithStyles(
				fzf.WithStyleCursor(fzf.Style{
					ForegroundColor: flagCursorFg,
					BackgroundColor: flagCursorBg,
					Bold:            flagCursorBold,
					Blink:           flagCursorBlink,
					Italic:          flagCursorItalic,
					Strikethrough:   flagCursorStrikethrough,
					Underline:       flagCursorUnderline,
					Faint:           flagCursorFaint,
				}),
				fzf.WithStyleCursorLine(fzf.Style{
					ForegroundColor: flagCursorLineFg,
					BackgroundColor: flagCursorLineBg,
					Bold:            flagCursorLineBold,
					Blink:           flagCursorLineBlink,
					Italic:          flagCursorLineItalic,
					Strikethrough:   flagCursorLineStrikethrough,
					Underline:       flagCursorLineUnderline,
					Faint:           flagCursorLineFaint,
				}),
				fzf.WithStyleSelectedPrefix(fzf.Style{
					ForegroundColor: flagSelectedPrefixFg,
					BackgroundColor: flagSelectedPrefixBg,
					Bold:            flagSelectedPrefixBold,
					Blink:           flagSelectedPrefixBlink,
					Italic:          flagSelectedPrefixItalic,
					Strikethrough:   flagSelectedPrefixStrikethrough,
					Underline:       flagSelectedPrefixUnderline,
					Faint:           flagSelectedPrefixFaint,
				}),
				fzf.WithStyleUnselectedPrefix(fzf.Style{
					ForegroundColor: flagUnselectedPrefixFg,
					BackgroundColor: flagUnselectedPrefixBg,
					Bold:            flagUnselectedPrefixBold,
					Blink:           flagUnselectedPrefixBlink,
					Italic:          flagUnselectedPrefixItalic,
					Strikethrough:   flagUnselectedPrefixStrikethrough,
					Underline:       flagUnselectedPrefixUnderline,
					Faint:           flagUnselectedPrefixFaint,
				}),
				fzf.WithStyleMatches(fzf.Style{
					ForegroundColor: flagMatchesFg,
					BackgroundColor: flagMatchesBg,
					Bold:            flagMatchesBold,
					Blink:           flagMatchesBlink,
					Italic:          flagMatchesItalic,
					Strikethrough:   flagMatchesStrikethrough,
					Underline:       flagMatchesUnderline,
					Faint:           flagMatchesFaint,
				}),
			),
		)

		ctx := context.Background()
		g, ctx := errgroup.WithContext(ctx)

		info, err := os.Stdin.Stat()
		if err != nil {
			return err
		}

		if info.Mode()&os.ModeCharDevice == 0 {
			g.Go(func() error {
				sc := bufio.NewScanner(os.Stdin)
				for sc.Scan() {
					mu.Lock()
					items = append(items, sc.Text())
					mu.Unlock()
				}
				return nil
			})
		} else {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			g.Go(func() error {
				err := filepath.WalkDir(wd, func(path string, d fs.DirEntry, err error) error {
					select {
					case <-ctx.Done():
						return ctx.Err()
					default:
					}

					if err != nil {
						return err
					}

					if d.Name()[0] == '.' {
						if d.IsDir() {
							return fs.SkipDir
						}
						return nil
					}

					if !d.IsDir() {
						mu.Lock()
						items = append(items, path)
						mu.Unlock()
					}

					return nil
				})
				if err != nil {
					f.Quit()
					return err
				}
				return nil
			})
		}

		g.Go(func() error {
			choices, err := f.Find(&items, func(i int) string { return items[i] })
			if err != nil {
				return err
			}

			for _, choice := range choices {
				fmt.Println(items[choice])
			}

			return nil
		})

		if err := g.Wait(); err != nil {
			return err
		}

		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// version
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}

	rootCmd.Version = version

	// flags
	rootCmd.Flags().SortFlags = false

	rootCmd.Flags().IntVar(&flagLimit, "limit", 1, "maximum number of items to select")
	rootCmd.Flags().BoolVar(&flagNoLimit, "no-limit", false, "unlimited number of items to select")
	rootCmd.MarkFlagsMutuallyExclusive("limit", "no-limit")

	rootCmd.Flags().BoolVar(&flagCaseSensitive, "case-sensitive", false, "case sensitive search")

	rootCmd.Flags().StringVar(&flagPrompt, "prompt", "> ", "")
	rootCmd.Flags().StringVar(&flagCursor, "cursor", "> ", "")
	rootCmd.Flags().StringVar(&flagSelectedPrefix, "selected-prefix", "● ", "")
	rootCmd.Flags().StringVar(&flagUnselectedPrefix, "unselected-prefix", "◯ ", "")
	rootCmd.Flags().StringVar(&flagInputPlaceholder, "input-placeholder", "Filter...", "")

	rootCmd.Flags().BoolVar(&flagCountView, "count-view", true, "")

	rootCmd.Flags().StringVar(&flagCursorFg, "cursor-fg", mainColor, "")
	rootCmd.Flags().StringVar(&flagCursorBg, "cursor-bg", "", "")
	rootCmd.Flags().BoolVar(&flagCursorBold, "cursor-bold", false, "")
	rootCmd.Flags().BoolVar(&flagCursorBlink, "cursor-blink", false, "")
	rootCmd.Flags().BoolVar(&flagCursorItalic, "cursor-italic", false, "")
	rootCmd.Flags().BoolVar(&flagCursorStrikethrough, "cursor-strike", false, "")
	rootCmd.Flags().BoolVar(&flagCursorUnderline, "cursor-underline", false, "")
	rootCmd.Flags().BoolVar(&flagCursorFaint, "cursor-faint", false, "")

	rootCmd.Flags().StringVar(&flagCursorLineFg, "cursorline-fg", "", "")
	rootCmd.Flags().StringVar(&flagCursorLineBg, "cursorline-bg", "", "")
	rootCmd.Flags().BoolVar(&flagCursorLineBold, "cursorline-bold", true, "")
	rootCmd.Flags().BoolVar(&flagCursorLineBlink, "cursorline-blink", false, "")
	rootCmd.Flags().BoolVar(&flagCursorLineItalic, "cursorline-italic", false, "")
	rootCmd.Flags().BoolVar(&flagCursorLineStrikethrough, "cursorline-strke", false, "")
	rootCmd.Flags().BoolVar(&flagCursorLineUnderline, "cursorline-underline", false, "")
	rootCmd.Flags().BoolVar(&flagCursorLineFaint, "cursorline-faint", false, "")

	rootCmd.Flags().StringVar(&flagSelectedPrefixFg, "selected-prefix-fg", mainColor, "")
	rootCmd.Flags().StringVar(&flagSelectedPrefixBg, "selected-prefix-bg", "", "")
	rootCmd.Flags().BoolVar(&flagSelectedPrefixBold, "selected-prefix-bold", false, "")
	rootCmd.Flags().BoolVar(&flagSelectedPrefixBlink, "selected-prefix-blink", false, "")
	rootCmd.Flags().BoolVar(&flagSelectedPrefixItalic, "selected-prefix-italic", false, "")
	rootCmd.Flags().BoolVar(&flagSelectedPrefixStrikethrough, "selected-prefix-strke", false, "")
	rootCmd.Flags().BoolVar(&flagSelectedPrefixUnderline, "selected-prefix-underline", false, "")
	rootCmd.Flags().BoolVar(&flagSelectedPrefixFaint, "selected-prefix-faint", false, "")

	rootCmd.Flags().StringVar(&flagUnselectedPrefixFg, "unselected-prefix-fg", "", "")
	rootCmd.Flags().StringVar(&flagUnselectedPrefixBg, "unselected-prefix-bg", "", "")
	rootCmd.Flags().BoolVar(&flagUnselectedPrefixBold, "unselected-prefix-bold", false, "")
	rootCmd.Flags().BoolVar(&flagUnselectedPrefixBlink, "unselected-prefix-blink", false, "")
	rootCmd.Flags().BoolVar(&flagUnselectedPrefixItalic, "unselected-prefix-italic", false, "")
	rootCmd.Flags().BoolVar(&flagUnselectedPrefixStrikethrough, "unselected-prefix-strke", false, "")
	rootCmd.Flags().BoolVar(&flagUnselectedPrefixUnderline, "unselected-prefix-underline", false, "")
	rootCmd.Flags().BoolVar(&flagUnselectedPrefixFaint, "unselected-prefix-faint", true, "")

	rootCmd.Flags().StringVar(&flagMatchesFg, "matches-fg", mainColor, "")
	rootCmd.Flags().StringVar(&flagMatchesBg, "matches-bg", "", "")
	rootCmd.Flags().BoolVar(&flagMatchesBold, "matches-bold", false, "")
	rootCmd.Flags().BoolVar(&flagMatchesBlink, "matches-blink", false, "")
	rootCmd.Flags().BoolVar(&flagMatchesItalic, "matches-italic", false, "")
	rootCmd.Flags().BoolVar(&flagMatchesStrikethrough, "matches-strke", false, "")
	rootCmd.Flags().BoolVar(&flagMatchesUnderline, "matches-underline", false, "")
	rootCmd.Flags().BoolVar(&flagMatchesFaint, "matches-faint", false, "")
}
