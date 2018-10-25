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
	Short: "Interpolate values into a templated file.",
	Long:  `This tool was developed to make it simple to replace template strings in config files with the proper values.
	It uses a key/value system, where you pass in key/value pairs that match with templated keys in the file being interpolated. 

Am example file, called "template": 

	Hello,
	The secret is {{ SECRET }}.

An example substitution file called "substitutions": 

	SECRET=I am a Cylon

Would work this way:
 
	$ tgen -t /path/to/template -f /path/to/substitutions
	Hello,
	The secret is I am a Cylon.`,
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
