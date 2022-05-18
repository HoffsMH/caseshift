package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hoffsmh/caseshift"
	"github.com/spf13/cobra"
)

var removeExtensions bool
var baseName bool
var keepOriginal bool

var rootCmd = &cobra.Command{
	Use: "caseshift",
	Run: func(cmd *cobra.Command, args []string) {

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if baseName {
				text = filepath.Base(text)
			}

			if removeExtensions {
				text = strings.TrimSuffix(text, path.Ext(text))
			}

			fmt.Println(text)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	},
}

var subCmd = &cobra.Command{
  Use:   "pascal",
  Short: "PascalCase",
  Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if baseName {
				text = filepath.Base(text)
			}

			if removeExtensions {
				text = strings.TrimSuffix(text, path.Ext(text))
			}

			if keepOriginal {
				fmt.Println(text, " ", caseshift.PascalCase(text))
				continue
			}
			fmt.Println(caseshift.PascalCase(text))
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
  },
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(subCmd)
	rootCmd.PersistentFlags().BoolVarP(&removeExtensions, "remove-extensions", "e", false, "before processing")
	rootCmd.PersistentFlags().BoolVarP(&baseName, "basename", "b", false, "before processing")
	rootCmd.PersistentFlags().BoolVarP(&keepOriginal, "keep", "k", false, "before processing")
}
