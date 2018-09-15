package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var templateFile string
var substitutions string
var outputFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tgen",
	Short: "Generate combined output from a template file and associated values",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

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
	rootCmd.Flags().StringVar(&substitutions, "replace", "r", "substitutions to replace templates with")
	rootCmd.Flags().StringVar(&outputFile, "out", "o", "file where generated output will be written (defaults to STDOUT)")
}
