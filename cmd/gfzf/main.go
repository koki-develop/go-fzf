package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/koki-develop/go-fzf"
	"github.com/spf13/cobra"
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

		f := fzf.New()
		i, err := f.Find(is)
		if err != nil {
			return err
		}

		fmt.Println(is[i])
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
