package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/koki-develop/go-fzf"
	"github.com/spf13/cobra"
)

const (
	mainColor = "#00ADD8"
)

var (
	flagLimit   int
	flagNoLimit bool
)

type items []string

var (
	_ fzf.Items = (items)(nil)
)

func (is items) ItemString(i int) string {
	return is[i]
}

func (is items) Len() int {
	return len(is)
}

var rootCmd = &cobra.Command{
	Use:          "gfzf",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		sc := bufio.NewScanner(os.Stdin)

		var is items
		for sc.Scan() {
			is = append(is, sc.Text())
		}

		f := fzf.New(
			fzf.WithNoLimit(flagNoLimit),
			fzf.WithLimit(flagLimit),
			fzf.WithStyles(
				fzf.WithStyleCursor(fzf.Style{ForegroundColor: mainColor}),
				fzf.WithStyleCursorLine(fzf.Style{Bold: true}),
				fzf.WithStyleMatches(fzf.Style{ForegroundColor: mainColor}),
				fzf.WithStyleUnselectedPrefix(fzf.Style{Faint: true}),
			),
		)
		choices, err := f.Find(is)
		if err != nil {
			return err
		}

		for _, choice := range choices {
			fmt.Println(is[choice])
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
	rootCmd.Flags().IntVarP(&flagLimit, "limit", "l", 1, "maximum number of items to select")
	rootCmd.Flags().BoolVar(&flagNoLimit, "no-limit", false, "unlimited number of items to select")
}
