package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// global vars for command args
var (
	apiDescriptionFile string
	outputFile         string
	Verbose            bool
	validArgs          = []string{"api", "verbose", "output"}
)

var (
	rootCmd = &cobra.Command{
		Use:   "etz",
		Short: "etz root command",
		Long: `Root command for etzqa.
				  Complete documentation is available at https://github.com/nadavbm/etzqa`,
		Run: func(cmd *cobra.Command, args []string) {
			// Empty
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&apiDescriptionFile, "file", "../../files/config.json", "api description file location")
	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "../../files/helpers.json", "test result output file location")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(apiCmd)
}
