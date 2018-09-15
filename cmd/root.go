package cmd

import (
	"fmt"
	"github.com/mrbeskin/tgen/generate"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var templateFile string
var substitutions string
var substitutionFile string
var outputFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tgen",
	Short: "Generate combined output from a template file and associated values",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var g string
		var err error
		if (substitutionFile == "") && (substitutions == "") {
			panic("must provide either substitutions or file to read substitutions from")
		}
		if (substitutionFile != "") && (substitutions != "") {
			panic("found both substitutions and substitution file, please provide only one")
		}
		if substitutionFile != "" {
			g, err = generate.GenerateFromFileWithSubstitutionFile(templateFile, substitutionFile)
			if err != nil {
				panic(err)
			}
		} else {
			s, err := generate.ParseSubstitutions(substitutions)
			if err != nil {
				panic(err)
			}
			g, err = generate.GenerateFromFile(templateFile, s)
		}

		if outputFile == "" {
			fmt.Println(g)
			return
		}
		err = ioutil.WriteFile(outputFile, []byte(g), 0644)
		if err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&templateFile, "template", "t", "file on which template values will be replaced")
	rootCmd.MarkFlagRequired("template")
	rootCmd.Flags().StringVar(&substitutionFile, "file", "f", "file from which to read substitutions")
	rootCmd.Flags().StringVar(&substitutions, "replace", "r", "pass substitutions to replace templated values via the CLI")
	rootCmd.Flags().StringVar(&outputFile, "out", "o", "file where generated output will be written (defaults to STDOUT)")
}
