package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/xybor/xyauth/internal/utils"
)

var output string
var rootCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate file from template",
	Long:  `A helper tool for generate file from template`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		src := args[0]
		t, err := template.ParseFiles(src)
		if err != nil {
			panic(err)
		}

		fin, err := os.Stat(src)
		if err != nil {
			panic(err)
		}

		if output == "" {
			if strings.HasSuffix(src, ".template") {
				output = src[:len(src)-9]
			} else {
				output = src + ".out"
			}
		}

		fout, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, fin.Mode())
		if err != nil {
			panic(err)
		}

		if err := t.Execute(fout, utils.GetConfig().ToMap()); err != nil {
			panic(err)
		}

		fmt.Println("Generate file successfully")
	},
}

func main() {
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "The output path")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
